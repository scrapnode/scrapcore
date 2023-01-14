package nats

import (
	"context"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/samber/lo"
	"os"
	"syscall"
	"time"
)

func (natsbus *Nats) Connect(ctx context.Context) error {
	natsbus.mu.Lock()
	defer natsbus.mu.Unlock()

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	opts := []nats.Option{
		nats.Name(hostname),
		nats.ReconnectWait(2 * time.Second),
		nats.Timeout(2 * time.Second),
		nats.MaxReconnects(7),
		nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
			natsbus.Logger.Error(fmt.Sprintf("msgbus.nats: got disconnected with reason: %q", err))
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			natsbus.Logger.Error(fmt.Sprintf("msgbus.nats: got reconnected to %v", conn.ConnectedUrl()))
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			natsbus.Logger.Error(fmt.Sprintf("msgbus.nats: connection is closed with reason: %q", nc.LastError()))

			// terminal presses then let another gorouting handle other component disconnection
			if err := syscall.Kill(syscall.Getpid(), syscall.SIGINT); err != nil {
				natsbus.Logger.Error(fmt.Sprintf("msgbus.nats: could not kill process: %q", err))
			}
		}),
		nats.ErrorHandler(func(c *nats.Conn, s *nats.Subscription, err error) {
			natsbus.Logger.Errorw(fmt.Sprintf("msgbus.nats: got error: %q", err), "subject", s.Subject, "queue", s.Queue)

			// terminal presses then let another gorouting handle other component disconnection
			if err := syscall.Kill(syscall.Getpid(), syscall.SIGINT); err != nil {
				natsbus.Logger.Error(fmt.Sprintf("msgbus.nats: could not kill process: %q", err))
			}
		}),
	}
	conn, err := nats.Connect(natsbus.Configs.Dsn, opts...)
	if err != nil {
		return err
	}

	natsbus.conn = conn
	natsbus.Logger.Debug("connected")
	return natsbus.setStream(ctx)
}

func (natsbus *Nats) setStream(ctx context.Context) error {
	jsc, err := natsbus.conn.JetStream()
	if err != nil {
		return err
	}

	name := NewStreamName(natsbus.Configs)
	stream, err := jsc.StreamInfo(name)

	// we only accept 2 case: no error & ErrStreamNotFound
	if err != nil && !errors.Is(err, nats.ErrStreamNotFound) {
		return err
	}

	jscfg := ParseJetStreamConfigs(ctx, natsbus.Configs)
	// if there is no stream was created, create a new one
	if err != nil {
		if _, err = jsc.AddStream(jscfg); err != nil {
			return err
		}

		natsbus.Logger.Debugw("create new stream", "stream_name", name, "subjects", jscfg.Subjects)
	} else {
		stream.Config.Subjects = lo.Uniq(append(stream.Config.Subjects, jscfg.Subjects...))
		if _, err = jsc.UpdateStream(&stream.Config); err != nil {
			return err
		}
		natsbus.Logger.Debugw("found stream", "stream_name", stream.Config.Name, "stream_cfg", stream.Config)
	}

	natsbus.jsc = jsc
	return nil
}
