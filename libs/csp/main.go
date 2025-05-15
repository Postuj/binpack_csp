package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/cspconstraints"
	"github.com/postuj/binpack_csp/libs/csp/entities"
	"github.com/postuj/binpack_csp/libs/csp/propagations"
)

func main() {
	bins := []*entities.Bin{
		entities.NewBin(0, "Crate", entities.REGULAR, 16),
		entities.NewBin(1, "Box", entities.REGULAR, 16),
		entities.NewBin(2, "Crate with ice", entities.COOLED, 24),
		entities.NewBin(3, "Cooled box", entities.COOLED, 32),
		entities.NewBin(4, "Fridge", entities.COOLED, 32),
	}

	itemFactory := entities.NewItemFactory(bins)

	itemFactory.AddItem("Potato", 8, entities.VEGETABLE)
	itemFactory.AddItem("Apple", 4, entities.FRUIT)
	itemFactory.AddItem("Tomato", 4, entities.VEGETABLE)
	itemFactory.AddItem("Banana", 8, entities.FRUIT)
	itemFactory.AddItem("Orange", 4, entities.FRUIT)
	itemFactory.AddItem("Onion", 4, entities.VEGETABLE)
	itemFactory.AddItem("Salmon", 16, entities.SEAFOOD)
	itemFactory.AddItem("Tuna", 24, entities.SEAFOOD)
	itemFactory.AddItem("Chicken", 8, entities.MEAT)
	itemFactory.AddItem("Shrimp", 8, entities.SEAFOOD)
	itemFactory.AddItem("Crab", 8, entities.SEAFOOD)
	itemFactory.AddItem("Pork", 16, entities.MEAT)
	itemFactory.AddItem("Beef", 8, entities.MEAT)

	items := itemFactory.GetItems()

	nonMixableItemTypes := []cspconstraints.NonMixableItemTypes{
		{entities.FRUIT, entities.VEGETABLE},
		{entities.MEAT, entities.SEAFOOD},
	}

	constraints := cspconstraints.MakeConstraints(items, nonMixableItemTypes)
	propagators := propagations.MakePropagations(items)

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	vars := itemFactory.GetAllVariables()
	solver := centipede.NewBackTrackingCSPSolverWithPropagation(vars, constraints, propagators)
	begin := time.Now()

	success, err := solver.Solve(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	elapsed := time.Since(begin)
	if !success {
		fmt.Printf("Could not find solution in %s\n", elapsed)
		return
	}

	fmt.Printf("Found solution in %s\n", elapsed)
	result := entities.NewAlocationResult(items, bins, solver.State)
	fmt.Print(result.String())
}
