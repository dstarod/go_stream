package go_stream

import (
	"fmt"
	"sort"
	"strings"

)

type FloatArray struct  {
	Data []float64
}

// Value for fmt module functions
func (ia FloatArray) String() string {
	sa := []string{}
	for _, value := range ia.Data {
		sa = append(sa, fmt.Sprint(value))
	}
	result := fmt.Sprintf("[]float64{%s}", strings.Join(sa, ", "))
	return result
}

// Append one element
func (ia *FloatArray) Append(elem float64)  *FloatArray {
	ia.Data = append(ia.Data, elem)
	return ia
}

// Append all element from array
func (ia *FloatArray) Extend(elements []float64)  *FloatArray {
	for _, elem := range elements {
		ia.Data = append(ia.Data, elem)
	}
	return ia
}

// Filtering elements
func (ia *FloatArray) Filter(ff func(float64) bool) *FloatArray {
	new_data := []float64{}
	for _, value := range ia.Data {
		if ff(value){
			new_data = append(new_data, value)
		}
	}
	ia.Data = new_data
	return ia
}

// Make some action with avery elements in array
func (ia *FloatArray) Map(ff func(float64) float64) *FloatArray {
	for n, value := range ia.Data {
		ia.Data[n] = ff(value)
	}
	return ia
}

// Are all values != 0
func (ia *FloatArray) All() bool {
	for _, value := range ia.Data {
		if value == 0{
			return false
		}
	}
	return true
}

// Is array contains any value != 0
func (ia *FloatArray) Any() bool {
	for _, value := range ia.Data {
		if value != 0{
			return true
		}
	}
	return false
}

// Maximal value
func (ia *FloatArray) Max() float64 {
	if len(ia.Data) == 0{
		fmt.Println("func Max() : Array is empty")
		return 0
	}
	max := ia.Data[0]
	for _, value := range ia.Data {
		if value > max{
			max = value
		}
	}
	return max
}

// Minimal value
func (ia *FloatArray) Min() float64 {
	if len(ia.Data) == 0{
		fmt.Println("func Min() : Array is empty")
		return 0
	}
	min := ia.Data[0]
	for _, value := range ia.Data {
		if value < min{
			min = value
		}
	}
	return min
}

// Sorting
func (ia *FloatArray) Sort() *FloatArray {
	sort.Sort(sort.Float64Slice(ia.Data))
	return ia
}

// Reversed sorting
func (ia *FloatArray) SortReverse() *FloatArray {
	sort.Sort(sort.Reverse(sort.Float64Slice(ia.Data)))
	return ia
}

// First element
func (ia *FloatArray) First() float64 {
	if len(ia.Data) == 0{
		fmt.Println("func First() : Array is empty")
		return 0
	}
	return ia.Data[0]
}

// Last element
func (ia *FloatArray) Last() float64 {
	if len(ia.Data) == 0{
		fmt.Println("func Last() : Array is empty")
		return 0
	}
	return ia.Data[len(ia.Data) - 1]
}

// Array length
func (ia *FloatArray) Count() int {
	return len(ia.Data)
}

// Not efficient, SortDistinct recommended
func (ia *FloatArray) Distinct() *FloatArray {
	new_data := []float64{}
	seen := map[float64]float64{}
	for _, value := range ia.Data {
		if _, ok := seen[value]; !ok{
			new_data = append(new_data, value)
			seen[value] = value
		}
	}
	ia.Data = new_data
	return ia
}

// Sort and drop duplicates
func (ia *FloatArray) SortDistinct() *FloatArray {
	ia.Sort()
	new_data := []float64{}
	p := 0.0
	for _, value := range ia.Data {
		if p != value{
			new_data = append(new_data, value)
			p = value
		}
	}
	ia.Data = new_data
	return ia
}