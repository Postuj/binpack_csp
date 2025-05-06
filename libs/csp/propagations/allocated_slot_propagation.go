package propagations

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/entities"
)

func AllocatedSlotPropagation(item *entities.Item) centipede.Propagation[entities.Placement] {
	placementVarName := item.GetPlacementVarName()
	return centipede.Propagation[entities.Placement]{
		Vars: centipede.VariableNames{placementVarName},
		PropagationFunction: func(
			asassignment centipede.VariableAssignment[entities.Placement],
			variables *centipede.Variables[entities.Placement],
		) []centipede.DomainRemoval[entities.Placement] {
			binId := asassignment.Value.BinId
			offset := asassignment.Value.Offset
			removals := make([]centipede.DomainRemoval[entities.Placement], 0)

			for _, variable := range *variables {
				if variable.Name == placementVarName {
					continue
				}

				for _, dom := range variable.Domain {
					if dom.BinId != binId {
						continue
					}

					if dom.Offset >= offset && dom.Offset < offset+item.GetSize() {
						removals = append(removals, centipede.DomainRemoval[entities.Placement]{
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
