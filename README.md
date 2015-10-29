# go_stream

Golang int slices actions, inspired Java8 stream

You can create types IntArray or FloatArray and call methods:

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
    
        // IntArray example
        
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
        
        // FloatArray - it's the same
        
        fa := go_stream.FloatArray{[]float64{3.6, 1.4, 7.5}}
        fmt.Println(fa) // []float64{3.6, 1.4, 7.5}
    
        fa.Append(4.5).Extend([]float64{7.56, 93.234}).Map(func(f float64) float64 { return f*2 }).SortDistinct()
        fmt.Println(fa) // []float64{2.8, 7.2, 9, 15, 15.12, 186.468}
    
        fa.Filter(func(x float64) bool { return x > 3  })
        fmt.Println(fa) // []float64{7.2, 9, 15, 15.12, 186.468}
    
        p := fa.Max()
        fmt.Printf("%T, %v\n", p, p) // float64, 15
    }

