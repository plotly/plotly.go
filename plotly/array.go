package plotly

import "time"

type Array []interface{}

func (a *Array) AsTimes() []time.Time {
	// Loop through interface{}, decode and return the time.Time objects
	// if inconsistencies, panic outright ?
	return nil
}

func (a *Array) SetTimes(times []time.Time) {
	a = &Array{times}
}

func (a *Array) AsInts() []int64 {
	return nil
}

func (a *Array) SetInts(ints []int64) {
	a = &Array{ints}
}

func (a *Array) AsStrings() []int64 {
	return nil
}

func (a *Array) SetStrings(strs []string) {
	a = &Array{strs}
}

func (a *Array) Holds() ArrayType {
	if len(*a) == 0 {
		return EmptyArray
	}
	first := (*a)[0]
	switch first.(type) {
	case int64:
		return IntsArray
	case string:
		return StringsArray
	case float64:
		return FloatsArray
	case time.Time:
		return TimesArray
	default:
		return TypeUnknown
	}
}

type ArrayType int

const (
	TypeUnknown ArrayType = iota
	EmptyArray
	TimesArray
	IntsArray
	FloatsArray
	StringsArray
)
