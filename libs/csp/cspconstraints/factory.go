package cspconstraints

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/cspentities"
)

func MakeConstraints(
	items []*cspentities.Item,
	nonMixableItemTypes []NonMixableItemTypes,
) centipede.Constraints[cspentities.Placement] {
	constraints := centipede.Constraints[cspentities.Placement]{}

	constraints = append(constraints, AddPlacementsDontOverlapConstraints(items)...)
	constraints = append(
		constraints,
		AddNonMixableItemTypesConstraints(items, nonMixableItemTypes)...)

	return constraints
}
