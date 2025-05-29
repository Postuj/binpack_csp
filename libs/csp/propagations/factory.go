package propagations

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/cspentities"
)

func MakePropagations(
	items []*cspentities.Item,
) centipede.Propagations[cspentities.Placement] {
	propagators := centipede.Propagations[cspentities.Placement]{}
	for _, item := range items {
		propagators = append(propagators, AllocatedSlotPropagation(item))
	}
	return propagators
}
