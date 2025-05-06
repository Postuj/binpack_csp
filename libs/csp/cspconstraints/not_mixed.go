package cspconstraints

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/entities"
)

func NotMixed(item1, item2 *entities.Item) centipede.Constraint[entities.Placement] {
	item1PlacementVarName := item1.GetPlacementVarName()
	item2PlacementVarName := item2.GetPlacementVarName()

	return centipede.Constraint[entities.Placement]{
		Vars: centipede.VariableNames{item1PlacementVarName, item2PlacementVarName},
		ConstraintFunction: func(variables *centipede.Variables[entities.Placement]) bool {
			placement1 := variables.Find(item1PlacementVarName)
			placement2 := variables.Find(item2PlacementVarName)

			if placement1.Empty || placement2.Empty {
				return true
			}

			return placement1.Value.BinId != placement2.Value.BinId
		},
	}
}

type NonMixableItemTypes [2]entities.ItemType

func findItemsOfType(
	itemsMap map[entities.ItemType][]*entities.Item,
	itemType entities.ItemType,
) []*entities.Item {
	if items, ok := itemsMap[itemType]; ok {
		return items
	}
	return nil
}

func AddNonMixableItemTypesConstraints(
	items []*entities.Item,
	nonMixableItemTypes []NonMixableItemTypes,
) []centipede.Constraint[entities.Placement] {
	constraints := make([]centipede.Constraint[entities.Placement], 0)
	itemsMap := make(map[entities.ItemType][]*entities.Item)
	for _, item := range items {
		itemType := item.GetType()
		if _, ok := itemsMap[itemType]; !ok {
			itemsMap[itemType] = make([]*entities.Item, 0)
		}
		itemsMap[itemType] = append(itemsMap[itemType], item)
	}

	for _, itemTypePair := range nonMixableItemTypes {
		itemsOfType1 := findItemsOfType(itemsMap, itemTypePair[0])
		itemsOfType2 := findItemsOfType(itemsMap, itemTypePair[1])

		if itemsOfType1 == nil || itemsOfType2 == nil {
			continue
		}

		for _, item1 := range itemsOfType1 {
			for _, item2 := range itemsOfType2 {
				if item1.GetID() != item2.GetID() {
					constraints = append(constraints, NotMixed(item1, item2))
				}
			}
		}
	}

	return constraints
}
