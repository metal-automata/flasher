package worker

import (
	"context"
	"sync"
	"time"

	"go.hollow.sh/toolbox/events"
	"go.hollow.sh/toolbox/events/registry"

	"github.com/nats-io/nats.go"
)

var (
	once           sync.Once
	checkinCadence = 30 * time.Second
)

// This starts a go-routine to peridocally check in with the NATS kv
func (w *Worker) startWorkerLivenessCheckin(ctx context.Context) {
	natsJS, ok := w.stream.(*events.NatsJetstream)
	if !ok {
		w.logger.Error("Non-NATS stores are not supported for worker-liveness")
		return
	}

	once.Do(func() {
		if err := registry.InitializeActiveControllerRegistry(natsJS); err != nil {
			w.logger.WithError(err).Error("unable to initialize active worker registry")
			return
		}

		go w.checkinRoutine(ctx)
	})
}

func (w *Worker) checkinRoutine(ctx context.Context) {
	me := registry.GetID(w.id)
	if err := registry.RegisterController(me); err != nil {
		w.logger.WithError(err).Warn("unable to do initial worker liveness registration")
	}

	tick := time.NewTicker(checkinCadence)
	defer tick.Stop()

	var stop bool
	for !stop {
		select {
		case <-tick.C:
			err := registry.ControllerCheckin(me)
			switch err {
			case nil:
			case nats.ErrKeyNotFound: // generally means NATS reaped our entry on TTL
				if err = registry.RegisterController(me); err != nil {
					w.logger.WithError(err).
						WithField("id", me.String()).
						Warn("unable to re-register worker")
				}
			default:
				w.logger.WithError(err).
					WithField("id", me.String()).
					Warn("worker checkin failed")
			}
		case <-ctx.Done():
			w.logger.Info("liveness check-in stopping on done context")
			stop = true
		}
	}
}