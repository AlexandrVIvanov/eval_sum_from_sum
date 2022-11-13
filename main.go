package main

import "fmt"

type myresulttype struct {
	sumelements    int
	index_elements []int
}

// fuction to check given string is in array or not
func implContains(sl []int, val int) bool {
	// iterate over the array and compare given string to each element
	for _, value := range sl {
		if value == val {
			return true
		}
	}
	return false
}

func evalsum(sums []int, maxsum int, curresult myresulttype) myresulttype {

	maxeval := myresulttype{
		sumelements:    curresult.sumelements,
		index_elements: curresult.index_elements,
	}
	//	fmt.Println(sums)
	//fmt.Println(curresult)
	for i := 0; i < len(sums); i++ {

		curelement := sums[i]

		if implContains(curresult.index_elements, i) {
			continue
		}

		if (curelement + curresult.sumelements) > maxsum {

			continue

		}
		eval := myresulttype{
			sumelements:    curresult.sumelements + curelement,
			index_elements: append(curresult.index_elements, i),
		}
		res := evalsum(sums, maxsum, eval)
		if res.sumelements > maxeval.sumelements {
			maxeval.sumelements = res.sumelements
			maxeval.index_elements = res.index_elements
		}
	}
	//	fmt.Print("return:")
	//	fmt.Println(maxeval)
	return maxeval
}

func main() {
	var ie []int
	sums := []int{100, 100, 154, 195, 202, 340, 365, 486, 550, 611}

	curresult := myresulttype{
		sumelements:    0,
		index_elements: ie,
	}

	fmt.Println(sums)
	index_elements := evalsum(sums, 643, curresult)

	fmt.Println(index_elements)
	fmt.Println(index_elements.sumelements)
	for i := 0; i < len(index_elements.index_elements); i++ {
		fmt.Println(sums[index_elements.index_elements[i]])
	}
}
