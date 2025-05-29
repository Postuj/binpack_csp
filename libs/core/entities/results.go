package entities

import (
	"fmt"
	"time"
)

type AllocatedItem struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Offset int    `json:"offset"`
}

func (a *AllocatedItem) String() string {
	return fmt.Sprintf(
		"Item %s[%d] at %d, %d\n",
		a.Name,
		a.Id,
		a.Offset,
		a.Offset+a.Size-1,
	)
}

type AllocatedBin struct {
	Id       int              `json:"id"`
	Name     string           `json:"name"`
	Capacity int              `json:"capacity"`
	Type     BinType          `json:"type"`
	Items    []*AllocatedItem `json:"items"`
}

func (a *AllocatedBin) String() string {
	out := fmt.Sprintf(
		"Bin %s[%d], capacity %d, has %d items\n",
		a.Name,
		a.Id,
		a.Capacity,
		len(a.Items),
	)
	for _, item := range a.Items {
		out += fmt.Sprintf("  %s", item.String())
	}
	return out
}

type AllocationResult struct {
	AllocatedBins []*AllocatedBin `json:"allocatedBins"`
	Success       bool            `json:"success"`
	Time          time.Duration   `json:"time"`
}

func (a *AllocationResult) String() string {
	out := "Allocation result:\n"
	for _, bin := range a.AllocatedBins {
		out += bin.String()
	}
	return out
}
