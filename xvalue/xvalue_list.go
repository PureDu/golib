package xvalue

type XValueList struct {
	noCopy
}

func (vl *XValueList) Concat(src *XValueList) bool  {
	return true
}

func (vl *XValueList) AppendValue(value *XValue) bool  {
	return true
}

func (vl *XValueList) AppendValueList(src *XValueList, start, count int) bool  {
	return true
}

func (vl *XValueList) Clear()  {
	
}

func (vl *XValueList) Empty() bool  {
	return true
}

func (vl *XValueList) Count() int  {
	return 0
}

func (vl *XValueList) Kind(index int) Kind  {
	return Invalid
}

func (vl *XValueList) AddBool(x bool) bool  {
	return true
}