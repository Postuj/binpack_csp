package cspentities

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/core/entities"
)

type Placement struct {
	BinId  int
	Offset int
}

func makePlacementDomainForBins(item *Item, bins []*entities.Bin) centipede.Domain[Placement] {
	placements := make(centipede.Domain[Placement], 0)
	for _, bin := range bins {
		if item.GetAllowedBinType() == bin.GetType() {
			for slot := range bin.GetCapacity() - item.GetSize() + 1 {
				if slot%2 != 0 {
					continue
				}
				placements = append(placements, Placement{
					BinId:  bin.GetID(),
					Offset: slot,
				})
			}
		}
	}
	return placements
}
