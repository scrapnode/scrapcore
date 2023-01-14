package msgbus

import (
	"github.com/gosimple/slug"
	"github.com/samber/lo"
	"math"
	"strings"
	"time"
)

func NatsStreamName(cfg *Configs) string {
	return strings.ReplaceAll(slug.Make(cfg.Name), "-", "_")
}

func NatsSubject(cfg *Configs, sample *Event) string {
	segments := []string{cfg.Region, cfg.Name}
	if sample == nil {
		return strings.Join(append(segments, ">"), ".")
	}

	if sample.Workspace != "" {
		segments = append(segments, sample.Workspace)
	} else {
		segments = append(segments, "*")
	}

	if sample.App != "" {
		segments = append(segments, sample.App)
	} else {
		segments = append(segments, "*")
	}

	if sample.Type != "" {
		segments = append(segments, sample.Type)
	} else {
		segments = append(segments, "*")
	}

	var dots []string
	count := len(segments)
	for i := count - 1; i >= 0; i-- {
		if segments[i] != "*" {
			dots = append(dots, lo.Reverse(segments[:i+1])...)
			break
		}

		if segments[i] == "*" && segments[i-1] == "*" {
			continue
		}
		dots = append(dots, segments[i])
	}
	newdots := lo.Reverse(dots)
	if newdots[len(dots)-1] == "*" {
		newdots[len(dots)-1] = ">"
	}
	return strings.Join(newdots, ".")
}

func NewBackoff(count int) []time.Duration {
	var backoff []time.Duration
	for i := 0; i < count; i++ {
		backoff = append(backoff, time.Second*time.Duration(3+math.Pow(2, float64(i))))
	}
	return backoff
}
