package main

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
        fmt.Println()
        return
    }
	m := len(slice) / size
	rem := len(slice) % size
	if rem >= 1  {
		m++
	}

	maped := make([][]int, m)
	
	n := 0
	count:=0	

	for i:=0; i<len(slice); i++ {
		if count == size {
			n++
			count=0
			i--
		} else {
			maped[n] =append(maped[n], slice[i])
			count++
		}
	}
	println(size, ": ", maped)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
