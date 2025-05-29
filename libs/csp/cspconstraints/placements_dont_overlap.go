package cspconstraints

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/cspentities"
)

func PlacementsInBinDontOverlap(
	item1, item2 *cspentities.Item,
) centipede.Constraint[cspentities.Placement] {
	item1PlacementVarName := item1.GetPlacementVarName()
	item2PlacementVarName := item2.GetPlacementVarName()
	size1 := item1.GetSize()
	size2 := item2.GetSize()

	return centipede.Constraint[cspentities.Placement]{
		Vars: centipede.VariableNames{item1PlacementVarName, item2PlacementVarName},
		ConstraintFunction: func(variables *centipede.Variables[cspentities.Placement]) bool {
			placement1 := variables.Find(item1PlacementVarName)
			placement2 := variables.Find(item2PlacementVarName)

			if placement1.Empty || placement2.Empty {
				return true
			}

			if placement1.Value.BinId != placement2.Value.BinId {
				return true
			}

			offset1 := placement1.Value.Offset
			offset2 := placement2.Value.Offset

			return offset1+size1 <= offset2 || offset2+size2 <= offset1
		},
	}
}

func AddPlacementsDontOverlapConstraints(
	items []*cspentities.Item,
) []centipede.Constraint[cspentities.Placement] {
	constraints := make([]centipede.Constraint[cspentities.Placement], 0)

	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i].GetAllowedBinType() == items[j].GetAllowedBinType() &&
				items[i].GetID() != items[j].GetID() {
				constraints = append(constraints, PlacementsInBinDontOverlap(items[i], items[j]))
			}

		}
	}

	return constraints
}
