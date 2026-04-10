package usecase

import (
	"fmt"

	"github.com/pastequo/experiments.unit-test/internal/domain/repo"
	"github.com/pastequo/experiments.unit-test/internal/utils/observability"
)

// Orchestration logic
// Should call both Act and Do, the order doesn't matter

type ActAndDo struct {
	actor repo.Actor
	doer  repo.Doer

	metrics observability.Metrics
}

func NewActAndDo(actor repo.Actor, doer repo.Doer, metrics observability.Metrics) ActAndDo {
	return ActAndDo{
		actor: actor,
		doer:  doer,

		metrics: metrics,
	}
}

// Sequential implementation, Act first
func (a ActAndDo) Run() error {
	a.metrics.Increment("usecase.Run")

	err := a.actor.Act()
	if err != nil {
		return fmt.Errorf("failed to Act: %w", err)
	}

	err = a.doer.Do()
	if err != nil {
		return fmt.Errorf("failed to Do: %w", err)
	}

	return nil
}

// // Sequential implementation, Do first
// func (a ActAndDo) Run() error {
//	a.metrics.Increment("usecase.Run")

// 	err := a.doer.Do()
// 	if err != nil {
// 		return fmt.Errorf("failed to Do: %w", err)
// 	}

// 	err = a.actor.Act()
// 	if err != nil {
// 		return fmt.Errorf("failed to Act: %w", err)
// 	}

// 	return nil
// }

// // Concurrent implementation, Act and Do in parallel
// func (a ActAndDo) Run() error {
//	a.metrics.Increment("usecase.Run")

// 	g := errgroup.Group{}

// 	g.Go(func() error {
// 		err := a.actor.Act()
// 		if err != nil {
// 			return fmt.Errorf("failed to Act: %w", err)
// 		}

// 		return nil
// 	})

// 	g.Go(func() error {
// 		err := a.doer.Do()
// 		if err != nil {
// 			return fmt.Errorf("failed to Do: %w", err)
// 		}

// 		return nil
// 	})

// 	err := g.Wait()
// 	if err != nil {
// 		return fmt.Errorf("failed to run Act and Do: %w", err)
// 	}

// 	return nil
// }
