package xvalue

import "fmt"

type XGUID struct {
	high uint64
	low uint64
}

func (g XGUID) High() uint64  {
	return g.high
}

func (g XGUID) Low() uint64  {
	return g.low
}

func (g XGUID) IsZero() bool  {
	return (g.low == 0) && (g.high == 0)
}

func (g XGUID) String() string  {
	return fmt.Sprintf("%d-%d", g.high, g.low)
}

func NewXGUID(h, l uint64) XGUID  {
	return XGUID{high:h, low:l}
}
