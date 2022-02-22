/*
Дан целочисленный массив nums, выведите на экран true, если любое значение встречается в массиве как минимум дважды, и false, если каждый элемент различен.
*/

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 8, 55, 44, 4}
	var flag bool

	for i := 0; i < len(a); i++ {
		//fmt.Println(i)
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				flag = true
			}
		}
	}
	fmt.Println(flag)
}
