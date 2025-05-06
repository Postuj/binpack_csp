package entities

import (
	"github.com/gnboorse/centipede"
	"github.com/postuj/binpack_csp/libs/csp/varnames"
)

type ItemType uint

const (
	FRUIT ItemType = iota
	VEGETABLE
	MEAT
	SEAFOOD
)

type Item struct {
	id             int
	itemType       ItemType
	name           string
	size           int
	possibleBinIds []int
}

func (i *Item) GetID() int {
	return i.id
}

func (i *Item) GetType() ItemType {
	return i.itemType
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetPossibleBinIds() []int {
	return i.possibleBinIds
}

func (i *Item) GetAllowedBinType() BinType {
	switch i.itemType {
	case FRUIT, VEGETABLE:
		return REGULAR
	case MEAT, SEAFOOD:
		return COOLED
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
