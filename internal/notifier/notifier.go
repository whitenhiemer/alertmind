package notifier

import (
	"context"

	"github.com/whitenhiemer/alertmind/internal/alert"
)

// Notifier sends a triage summary for an alert group.
type Notifier interface {
	Notify(ctx context.Context, payload *alert.AlertmanagerPayload, triage *alert.Triage) error
}

// Multi sends to all configured notifiers, collecting errors.
type Multi struct {
	notifiers []Notifier
}

func NewMulti(notifiers ...Notifier) *Multi {
	return &Multi{notifiers: notifiers}
}

func (m *Multi) Notify(ctx context.Context, payload *alert.AlertmanagerPayload, triage *alert.Triage) error {
	var last error
	for _, n := range m.notifiers {
		if err := n.Notify(ctx, payload, triage); err != nil {
			last = err
		}
	}
	return last
}
