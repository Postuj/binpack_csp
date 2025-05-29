package propagations

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/cspentities"
)

func AllocatedSlotPropagation(item *cspentities.Item) centipede.Propagation[cspentities.Placement] {
	placementVarName := item.GetPlacementVarName()
	return centipede.Propagation[cspentities.Placement]{
		Vars: centipede.VariableNames{placementVarName},
		PropagationFunction: func(
			asassignment centipede.VariableAssignment[cspentities.Placement],
			variables *centipede.Variables[cspentities.Placement],
		) []centipede.DomainRemoval[cspentities.Placement] {
			binId := asassignment.Value.BinId
			offset := asassignment.Value.Offset
			removals := make([]centipede.DomainRemoval[cspentities.Placement], 0)

			for _, variable := range *variables {
				if variable.Name == placementVarName {
					continue
				}

				for _, dom := range variable.Domain {
					if dom.BinId != binId {
						continue
					}

					if dom.Offset >= offset && dom.Offset < offset+item.GetSize() {
						removals = append(removals, centipede.DomainRemoval[cspentities.Placement]{
							VariableName: variable.Name,
							Value:        dom,
						})
					}
				}
			}
			return removals
		},
	}
}
