package csp

import (
	"context"
	"fmt"
	"time"

	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/core/entities"
	"github.com/postuj/binpack_csp/libs/csp/cspconstraints"
	"github.com/postuj/binpack_csp/libs/csp/cspentities"
	"github.com/postuj/binpack_csp/libs/csp/propagations"
)

type CspSolver struct{}

func (s *CspSolver) Solve(bins []*entities.Bin, items []*entities.Item) *entities.AllocationResult {
	itemFactory := cspentities.NewItemFactory(bins)

	for _, item := range items {
		itemFactory.AddItem(
			item.GetName(),
			item.GetSize(),
			item.GetType(),
		)
	}

	nonMixableItemTypes := []cspconstraints.NonMixableItemTypes{
		{entities.FRUIT, entities.VEGETABLE},
		{entities.MEAT, entities.SEAFOOD},
	}

	cspItems := itemFactory.GetItems()
	constraints := cspconstraints.MakeConstraints(cspItems, nonMixableItemTypes)
	propagators := propagations.MakePropagations(cspItems)

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	vars := itemFactory.GetAllVariables()
	solver := centipede.NewBackTrackingCSPSolverWithPropagation(vars, constraints, propagators)
	begin := time.Now()

	success, err := solver.Solve(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return &entities.AllocationResult{
			Success:       false,
			AllocatedBins: nil,
			Time:          0,
		}
	}

	elapsed := time.Since(begin)
	if !success {
		return &entities.AllocationResult{
			Success:       false,
			AllocatedBins: nil,
			Time:          elapsed,
		}
	}

	return cspentities.NewAlocationResult(cspItems, bins, solver.State, elapsed)
}
