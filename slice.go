package urandom

import (
	"reflect"
)

// SliceRand assign a random value from a string to the "r" arg
func SliceRand(s interface{}, r interface{}) {
	sl := reflect.ValueOf(s).Elem()

	// Handle types or exit if not supported type
	switch s.(type) {
	case *[]string:
		sli := sl.Interface().([]string)
		*r.(*string) = sli[seed.Intn(len(sli))]
	case *[]int:
		sli := sl.Interface().([]int)
		*r.(*int) = sli[seed.Intn(len(sli))]
	case *[]float64:
		sli := sl.Interface().([]float64)
		*r.(*float64) = sli[seed.Intn(len(sli))]
	default:
		return
	}
}

// SliceRandString returns a random value from a string slice
func SliceRandString(s []string) string {
	return s[seed.Intn(len(s))]
}

// SliceRandS is an alias of SliceRandString
func SliceRandS(s []string) string {
	return SliceRandString(s)
}

// SliceRandInt returns a random value from an int slice
func SliceRandInt(s []int) int {
	return s[seed.Intn(len(s))]
}

// SliceRandI is an alias of SliceRandInt
func SliceRandI(s []int) int {
	return SliceRandInt(s)
}

// SliceRandFloat returns a random value from a float64 slice
func SliceRandFloat(s []float64) float64 {
	return s[seed.Intn(len(s))]
}

// SliceRandF is an alias of SliceRandFloat
func SliceRandF(s []float64) float64 {
	return SliceRandFloat(s)
}
