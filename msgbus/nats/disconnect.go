package nats

import "context"

func (natsbus *Nats) Disconnect(ctx context.Context) error {
	natsbus.mu.Lock()
	defer natsbus.mu.Unlock()

	if natsbus.conn == nil {
		return nil
	}

	if err := natsbus.conn.Drain(); err != nil {
		return err
	}

	natsbus.conn = nil
	natsbus.jsc = nil
	natsbus.Logger.Debug("disconnected")
	return nil
}
