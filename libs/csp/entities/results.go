package entities

import (
	"fmt"

	"github.com/gnboorse/centipede"
)

type AllocatedItem struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Offset int    `json:"offset"`
}

func (a *AllocatedItem) String() string {
	return fmt.Sprintf(
		"Item %s[%d] at %d, %d\n",
		a.Name,
		a.Id,
		a.Offset,
		a.Offset+a.Size-1,
	)
}

type AllocatedBin struct {
	Id       int              `json:"id"`
	Name     string           `json:"name"`
	Capacity int              `json:"capacity"`
	Type     BinType          `json:"type"`
	Items    []*AllocatedItem `json:"items"`
}

func (a *AllocatedBin) String() string {
	out := fmt.Sprintf(
		"Bin %s[%d], capacity %d, has %d items\n",
		a.Name,
		a.Id,
		a.Capacity,
		len(a.Items),
	)
	for _, item := range a.Items {
		out += fmt.Sprintf("  %s", item.String())
	}
	return out
}

type AllocationResult struct {
	AllocatedBins []*AllocatedBin `json:"allocatedBins"`
}

func (a *AllocationResult) String() string {
	out := "Allocation result:\n"
	for _, bin := range a.AllocatedBins {
		out += bin.String()
	}
	return out
}

func NewAlocationResult(
	items []*Item,
	bins []*Bin,
	solvedState centipede.CSPState[Placement],
) *AllocationResult {
	allocatedBins := make([]*AllocatedBin, len(bins))
	for i, bin := range bins {
		allocatedBins[i] = &AllocatedBin{
			Id:       bin.GetID(),
			Name:     bin.GetName(),
			Capacity: bin.GetCapacity(),
			Type:     bin.GetType(),
			Items:    make([]*AllocatedItem, 0),
		}
	}

	for _, item := range items {
		placement := solvedState.Vars.Find(item.GetPlacementVarName()).Value

		bin := allocatedBins[placement.BinId]
		bin.Items = append(bin.Items, &AllocatedItem{
			Id:     item.GetID(),
			Name:   item.GetName(),
			Size:   item.GetSize(),
			Offset: placement.Offset,
		})
	}

	return &AllocationResult{
		AllocatedBins: allocatedBins,
	}
}
