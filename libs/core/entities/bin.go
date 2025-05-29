package entities

type BinType uint

const (
	REGULAR BinType = iota
	COOLED
)

type Bin struct {
	id       int
	name     string
	binType  BinType
	capacity int
}

func NewBin(id int, name string, binType BinType, capacity int) *Bin {
	return &Bin{
		id:       id,
		name:     name,
		binType:  binType,
		capacity: capacity,
	}
}

func (b *Bin) GetID() int {
	return b.id
}

func (b *Bin) GetName() string {
	return b.name
}

func (b *Bin) GetType() BinType {
	return b.binType
}

func (b *Bin) GetCapacity() int {
	return b.capacity
}
