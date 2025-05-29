package cspconstraints

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/core/entities"
	"github.com/postuj/binpack_csp/libs/csp/cspentities"
)

func NotMixed(item1, item2 *cspentities.Item) centipede.Constraint[cspentities.Placement] {
	item1PlacementVarName := item1.GetPlacementVarName()
	item2PlacementVarName := item2.GetPlacementVarName()

	return centipede.Constraint[cspentities.Placement]{
		Vars: centipede.VariableNames{item1PlacementVarName, item2PlacementVarName},
		ConstraintFunction: func(variables *centipede.Variables[cspentities.Placement]) bool {
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
	itemsMap map[entities.ItemType][]*cspentities.Item,
	itemType entities.ItemType,
) []*cspentities.Item {
	if items, ok := itemsMap[itemType]; ok {
		return items
	}
	return nil
}

func AddNonMixableItemTypesConstraints(
	items []*cspentities.Item,
	nonMixableItemTypes []NonMixableItemTypes,
) []centipede.Constraint[cspentities.Placement] {
	constraints := make([]centipede.Constraint[cspentities.Placement], 0)
	itemsMap := make(map[entities.ItemType][]*cspentities.Item)
	for _, item := range items {
		itemType := item.GetType()
		if _, ok := itemsMap[itemType]; !ok {
			itemsMap[itemType] = make([]*cspentities.Item, 0)
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
