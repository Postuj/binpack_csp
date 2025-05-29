package brute

import (
	"time"

	"github.com/postuj/binpack_csp/libs/core/entities"
)

type BruteSolver struct{}

type binState struct {
	bin   *entities.Bin
	used  []bool
	items []*entities.AllocatedItem
	types map[entities.ItemType]bool
}

func (s *BruteSolver) Solve(
	bins []*entities.Bin,
	items []*entities.Item,
) *entities.AllocationResult {
	start := time.Now()
	n := len(items)
	m := len(bins)
	if n == 0 || m == 0 {
		return &entities.AllocationResult{
			AllocatedBins: nil,
			Success:       false,
			Time:          time.Since(start),
		}
	}

	binStates := prepareBinStates(bins)

	var result []*entities.AllocatedBin
	found := bruteForceSolve(items, binStates, &result)

	return &entities.AllocationResult{
		AllocatedBins: result,
		Success:       found,
		Time:          time.Since(start),
	}
}

func prepareBinStates(bins []*entities.Bin) []*binState {
	binStates := make([]*binState, len(bins))
	for i, b := range bins {
		binStates[i] = &binState{
			bin:   b,
			used:  make([]bool, b.GetCapacity()),
			items: []*entities.AllocatedItem{},
			types: map[entities.ItemType]bool{},
		}
	}
	return binStates
}

func canPlace(item *entities.Item, bs *binState, offset int) bool {
	if offset < 0 || offset+item.GetSize() > bs.bin.GetCapacity() {
		return false
	}
	for i := 0; i < item.GetSize(); i++ {
		if bs.used[offset+i] {
			return false
		}
	}
	if bs.bin.GetType() != item.GetAllowedBinType() {
		return false
	}
	itype := item.GetType()
	if (itype == entities.FRUIT && bs.types[entities.VEGETABLE]) ||
		(itype == entities.VEGETABLE && bs.types[entities.FRUIT]) {
		return false
	}
	if (itype == entities.SEAFOOD && bs.types[entities.MEAT]) ||
		(itype == entities.MEAT && bs.types[entities.SEAFOOD]) {
		return false
	}
	return true
}

func place(item *entities.Item, bs *binState, offset int) {
	for i := 0; i < item.GetSize(); i++ {
		bs.used[offset+i] = true
	}
	bs.items = append(bs.items, &entities.AllocatedItem{
		Id:     item.GetID(),
		Name:   item.GetName(),
		Size:   item.GetSize(),
		Offset: offset,
	})
	bs.types[item.GetType()] = true
}

func unplace(item *entities.Item, bs *binState, offset int, items []*entities.Item) {
	for i := 0; i < item.GetSize(); i++ {
		bs.used[offset+i] = false
	}
	bs.items = bs.items[:len(bs.items)-1]
	bs.types = map[entities.ItemType]bool{}
	for _, ai := range bs.items {
		for _, it := range items {
			if it.GetID() == ai.Id {
				bs.types[it.GetType()] = true
				break
			}
		}
	}
}

func buildResult(binStates []*binState) []*entities.AllocatedBin {
	var result []*entities.AllocatedBin
	for _, bs := range binStates {
		if len(bs.items) > 0 {
			result = append(result, &entities.AllocatedBin{
				Id:       bs.bin.GetID(),
				Name:     bs.bin.GetName(),
				Capacity: bs.bin.GetCapacity(),
				Type:     bs.bin.GetType(),
				Items:    append([]*entities.AllocatedItem{}, bs.items...),
			})
		}
	}
	return result
}

func bruteForceSolve(
	items []*entities.Item,
	binStates []*binState,
	result *[]*entities.AllocatedBin,
) bool {
	n := len(items)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == n {
			*result = buildResult(binStates)
			return true
		}
		item := items[idx]
		for _, bs := range binStates {
			for offset := 0; offset <= bs.bin.GetCapacity()-item.GetSize(); offset++ {
				if canPlace(item, bs, offset) {
					place(item, bs, offset)
					if dfs(idx + 1) {
						return true
					}
					unplace(item, bs, offset, items)
				}
			}
		}
		return false
	}
	return dfs(0)
}
