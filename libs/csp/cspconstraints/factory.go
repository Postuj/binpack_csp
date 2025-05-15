package cspconstraints

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/entities"
)

func MakeConstraints(
	items []*entities.Item,
	nonMixableItemTypes []NonMixableItemTypes,
) centipede.Constraints[entities.Placement] {
	constraints := centipede.Constraints[entities.Placement]{}

	constraints = append(constraints, AddPlacementsDontOverlapConstraints(items)...)
	constraints = append(
		constraints,
		AddNonMixableItemTypesConstraints(items, nonMixableItemTypes)...)

	return constraints
}
