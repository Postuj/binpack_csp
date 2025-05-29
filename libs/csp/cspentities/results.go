package cspentities

import (
	"time"

	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/core/entities"
)

func NewAlocationResult(
	items []*Item,
	bins []*entities.Bin,
	solvedState centipede.CSPState[Placement],
	elapsedTime time.Duration,
) *entities.AllocationResult {
	allocatedBins := make([]*entities.AllocatedBin, len(bins))
	for i, bin := range bins {
		allocatedBins[i] = &entities.AllocatedBin{
			Id:       bin.GetID(),
			Name:     bin.GetName(),
			Capacity: bin.GetCapacity(),
			Type:     entities.BinType(bin.GetType()),
			Items:    make([]*entities.AllocatedItem, 0),
		}
	}

	for _, item := range items {
		placement := solvedState.Vars.Find(item.GetPlacementVarName()).Value

		bin := allocatedBins[placement.BinId]
		bin.Items = append(bin.Items, &entities.AllocatedItem{
			Id:     item.GetID(),
			Name:   item.GetName(),
			Size:   item.GetSize(),
			Offset: placement.Offset,
		})
	}

	return &entities.AllocationResult{
		Success:       true,
		AllocatedBins: allocatedBins,
		Time:          elapsedTime,
	}
}
