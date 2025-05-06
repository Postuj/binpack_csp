package entities

import "github.com/gnboorse/centipede"

type itemFactory struct {
	bins  []*Bin
	items []*Item
	vars  []centipede.Variable[Placement]
}

func NewItemFactory(bins []*Bin) *itemFactory {
	return &itemFactory{
		bins:  bins,
		items: make([]*Item, 0),
		vars:  make([]centipede.Variable[Placement], 0),
	}
}

func (f *itemFactory) AddItem(name string, size int, itemType ItemType) *Item {
	itemId := len(f.items)
	item := &Item{
		id:       itemId,
		itemType: itemType,
		name:     name,
		size:     size,
	}
	possibleBinIds := getPossibleBinIds(item.GetAllowedBinType(), f.bins)
	item.possibleBinIds = possibleBinIds
	f.items = append(f.items, item)

	f.vars = append(f.vars, centipede.NewVariable(
		item.GetPlacementVarName(),
		makePlacementDomainForBins(item, f.bins),
	))

	return item
}

func (f *itemFactory) GetItems() []*Item {
	return f.items
}

func (f *itemFactory) GetAllVariables() []centipede.Variable[Placement] {
	return f.vars
}

func getPossibleBinIds(binType BinType, bins []*Bin) []int {
	possibleBinIds := make([]int, 0)
	for _, bin := range bins {
		if binType == bin.GetType() {
			possibleBinIds = append(possibleBinIds, bin.GetID())
		}
	}
	return possibleBinIds
}
