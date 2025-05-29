package main

import (
	"fmt"

	"github.com/postuj/binpack_csp/libs/brute"
	"github.com/postuj/binpack_csp/libs/core/entities"
	"github.com/postuj/binpack_csp/libs/csp"
)

func main() {
	bins := []*entities.Bin{
		entities.NewBin(0, "Crate", entities.REGULAR, 16),
		entities.NewBin(1, "Box", entities.REGULAR, 16),
		entities.NewBin(2, "Crate with ice", entities.COOLED, 24),
		entities.NewBin(3, "Cooled box", entities.COOLED, 32),
		entities.NewBin(4, "Fridge", entities.COOLED, 32),
	}

	items := []*entities.Item{
		entities.NewItem(0, "Potato", 8, entities.VEGETABLE),
		entities.NewItem(1, "Apple", 4, entities.FRUIT),
		entities.NewItem(2, "Tomato", 4, entities.VEGETABLE),
		entities.NewItem(3, "Banana", 8, entities.FRUIT),
		entities.NewItem(4, "Orange", 4, entities.FRUIT),
		entities.NewItem(5, "Onion", 4, entities.VEGETABLE),
		entities.NewItem(6, "Salmon", 16, entities.SEAFOOD),
		entities.NewItem(7, "Tuna", 24, entities.SEAFOOD),
		entities.NewItem(8, "Chicken", 8, entities.MEAT),
		entities.NewItem(9, "Shrimp", 8, entities.SEAFOOD),
		entities.NewItem(10, "Crab", 8, entities.SEAFOOD),
		entities.NewItem(11, "Pork", 16, entities.MEAT),
		entities.NewItem(12, "Beef", 8, entities.MEAT),
	}

	fmt.Println("Solving bin packing problem with CSP...")
	solveProblem(&csp.CspSolver{}, bins, items)

	fmt.Println("Solving bin packing problem with Brute Force...")
	solveProblem(&brute.BruteSolver{}, bins, items)
}

func solveProblem(
	solver entities.Solver,
	bins []*entities.Bin,
	items []*entities.Item,
) *entities.AllocationResult {
	result := solver.Solve(bins, items)
	if result.Success {
		fmt.Printf("Solution found in %s\n%s\n", result.Time, result)
	} else {
		fmt.Println("No solution found.")
	}
	return result
}
