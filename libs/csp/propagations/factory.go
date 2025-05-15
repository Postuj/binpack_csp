package propagations

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/entities"
)

func MakePropagations(
	items []*entities.Item,
) centipede.Propagations[entities.Placement] {
	propagators := centipede.Propagations[entities.Placement]{}
	for _, item := range items {
		propagators = append(propagators, AllocatedSlotPropagation(item))
	}
	return propagators
}
