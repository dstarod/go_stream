# go_stream

Golang int slices actions, inspired Java8 stream

You can create type IntArray and call methods:

- Append(int): add element
- Extend([]int): add all elements
- Filter(func(int) bool): pass elements, returns true for func(elem)
- Map(func(int) int): replace every element
- All(): returns true, if all element != 0
- Any(): returns true, if any element != 0
- Max(): returns maximal value 
- Min(): returns minimal value
- Sort(): sorting elements in natural order
- SortReverse(): sorting elements in reversed order
- First(): returns first element
- Last(): returns last element
- Count(): returns collection length
- Distinct(): drop duplicates
- SortDistinct(): more efficient, than Distinct(), but original order will be lost

For example:

    package main
    
    import (
        "fmt"
        "github.com/dstarodubtsev/go_stream"
    )
    
    func main()  {
        ia := go_stream.IntArray{[]int{4, 2, 4, 3}}
    
        fmt.Println(ia) // []int{4, 2, 1, 3}
    
        ia.Distinct()
        fmt.Println(ia) // []int{4, 2, 3}
    
        fmt.Println(ia.First()) // 4
        fmt.Println(ia.Last()) // 3
        fmt.Println(ia.Max()) // 4 == ia.Sort().Last()
        fmt.Println(ia.Min()) // 2 == ia.Sort().First()
    
        ia.
            Extend([]int{45,56}).Filter(func(x int) bool { return x%3 == 0}).
            Map(func(x int) int { return x<<2 }).SortReverse().Append(12)
        fmt.Println(ia) // []int{180, 12, 12}
    
        ia.SortDistinct()
        fmt.Println(ia) // []int{12, 180}
    }
