package cspentities

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/core/entities"
	"github.com/postuj/binpack_csp/libs/csp/varnames"
)

type Item struct {
	id             int
	itemType       entities.ItemType
	name           string
	size           int
	possibleBinIds []int
}

func (i *Item) GetID() int {
	return i.id
}

func (i *Item) GetType() entities.ItemType {
	return i.itemType
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetPossibleBinIds() []int {
	return i.possibleBinIds
}

func (i *Item) GetAllowedBinType() entities.BinType {
	switch i.itemType {
	case entities.FRUIT, entities.VEGETABLE:
		return entities.REGULAR
	case entities.MEAT, entities.SEAFOOD:
		return entities.COOLED
	default:
		panic("unknown item type")
	}
}

func (i *Item) GetSize() int {
	return i.size
}

func (i *Item) GetPlacementVarName() centipede.VariableName {
	return varnames.Placement(i.id)
}
