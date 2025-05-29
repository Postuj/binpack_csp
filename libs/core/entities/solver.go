package entities

type Solver interface {
	Solve(bins []*Bin, items []*Item) *AllocationResult
}
