package go_stream

import (
	"fmt"
	"sort"
	"strings"

)

type IntArray struct {
	Data []int
}

// Value for fmt module functions
func (ia IntArray) String() string {
	sa := []string{}
	for _, value := range ia.Data {
		sa = append(sa, fmt.Sprint(value))
	}
	result := fmt.Sprintf("[]int{%s}", strings.Join(sa, ", "))
	return result
}

// Append one element
func (ia *IntArray) Append(elem int)  *IntArray {
	ia.Data = append(ia.Data, elem)
	return ia
}

// Append all element from array
func (ia *IntArray) Extend(elements []int)  *IntArray {
	for _, elem := range elements {
		ia.Data = append(ia.Data, elem)
	}
	return ia
}

// Filtering elements
func (ia *IntArray) Filter(ff func(int) bool) *IntArray {
	new_data := []int{}
	for _, value := range ia.Data {
		if ff(value){
			new_data = append(new_data, value)
		}
	}
	ia.Data = new_data
	return ia
}

// Make some action with avery elements in array
func (ia *IntArray) Map(ff func(int) int) *IntArray {
	for n, value := range ia.Data {
		ia.Data[n] = ff(value)
	}
	return ia
}

// Are all values != 0
func (ia *IntArray) All() bool {
	for _, value := range ia.Data {
		if value == 0{
			return false
		}
	}
	return true
}

// Is array contains any value != 0
func (ia *IntArray) Any() bool {
	for _, value := range ia.Data {
		if value != 0{
			return true
		}
	}
	return false
}

// Maximal value
func (ia *IntArray) Max() int {
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
func (ia *IntArray) Min() int {
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
func (ia *IntArray) Sort() *IntArray {
	sort.Sort(sort.IntSlice(ia.Data))
	return ia
}

// Reversed sorting
func (ia *IntArray) SortReverse() *IntArray {
	sort.Sort(sort.Reverse(sort.IntSlice(ia.Data)))
	return ia
}

// First element
func (ia *IntArray) First() int {
	if len(ia.Data) == 0{
		fmt.Println("func First() : Array is empty")
		return 0
	}
	return ia.Data[0]
}

// Last element
func (ia *IntArray) Last() int {
	if len(ia.Data) == 0{
		fmt.Println("func Last() : Array is empty")
		return 0
	}
	return ia.Data[len(ia.Data) - 1]
}

// Array length
func (ia *IntArray) Count() int {
	return len(ia.Data)
}

// Not efficient, SortDistinct recommended
func (ia *IntArray) Distinct() *IntArray {
	new_data := []int{}
	seen := map[int]int{}
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
func (ia *IntArray) SortDistinct() *IntArray {
	ia.Sort()
	new_data := []int{}
	p := 0
	for _, value := range ia.Data {
		if p != value{
			new_data = append(new_data, value)
			p = value
		}
	}
	ia.Data = new_data
	return ia
}