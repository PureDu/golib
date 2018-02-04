package data

import "fmt"

type GUID struct {
	high uint64
	low  uint64
}

func NewGUID(high, low uint64) GUID {
	return GUID{
		high: high,
		low:  low,
	}
}

func NewNULLGUID() GUID {
	return NewGUID(0, 0)
}

func (g *GUID) GetHigh() uint64 {
	return g.high
}

func (g *GUID) GetLow() uint64 {
	return g.low
}

func (g *GUID) IsNULL() bool {
	return (g.high == 0) && (g.low == 0)
}

func (g *GUID) String() string {
	return fmt.Sprintf("%d-%d", g.high, g.low)
}
