package main

import "fmt"

func main() {

	sliceMain := []int{43, 12, 0, -3, 7, 6, 24, 99, 2}

	sliceEnd := sortMerg(sliceMain)

	fmt.Print("sliceEnd: ", sliceEnd)
}

func sortMerg(sliceMain []int) []int {

	lenSlice := len(sliceMain)
	sliceTemp := make([]int, lenSlice)

	_ = copy(sliceTemp, sliceMain)

	if lenSlice == 2 {

		if sliceTemp[0] > sliceTemp[1] {
			sliceTemp[0], sliceTemp[1] = sliceTemp[1], sliceTemp[0]
		}

	} else if lenSlice > 2 {

		slice1 := sortMerg(sliceTemp[:lenSlice/2])
		slice2 := sortMerg(sliceTemp[lenSlice/2:])

		lenSlice1 := len(slice1)
		lenSlice2 := len(slice2)

		var indSlice1, indSlice2 int

		sliceTemp = nil

		for i := 0; i < lenSlice; i++ {

			if indSlice1 == lenSlice1 && indSlice2 == lenSlice2 {

				break

			} else if indSlice1 == lenSlice1 || (indSlice2 != lenSlice2 && slice1[indSlice1] > slice2[indSlice2]) {

				sliceTemp = append(sliceTemp, slice2[indSlice2])
				indSlice2++

			} else if indSlice2 == lenSlice2 || (indSlice1 != lenSlice1 && slice1[indSlice1] < slice2[indSlice2]) {

				sliceTemp = append(sliceTemp, slice1[indSlice1])
				indSlice1++

			} else {

				sliceTemp = append(sliceTemp, slice1[indSlice1])
				sliceTemp = append(sliceTemp, slice2[indSlice2])
				indSlice1++
				indSlice2++
			}
		}
	}
	return sliceTemp
}
