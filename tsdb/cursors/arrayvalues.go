package cursors

import "sort"

func (a *FloatArray) Size() int {
	// size of timestamps + values
	return len(a.Timestamps)*8 + len(a.Values)*8
}

func (a *IntegerArray) Size() int {
	// size of timestamps + values
	return len(a.Timestamps)*8 + len(a.Values)*8
}

func (a *UnsignedArray) Size() int {
	// size of timestamps + values
	return len(a.Timestamps)*8 + len(a.Values)*8
}

func (a *StringArray) Size() int {
	sz := len(a.Timestamps) * 8
	for _, s := range a.Values {
		sz += len(s)
	}
	return sz
}

func (a *BooleanArray) Size() int {
	// size of timestamps + values
	return len(a.Timestamps)*8 + len(a.Values)
}

var _ sort.Interface = (*TimestampArray)(nil)

func (a *TimestampArray) Less(i, j int) bool {
	return a.Timestamps[i] < a.Timestamps[j]
}

func (a *TimestampArray) Swap(i, j int) {
	a.Timestamps[i], a.Timestamps[j] = a.Timestamps[j], a.Timestamps[i]
}
