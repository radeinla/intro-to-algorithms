package main

import "fmt"

type binary struct {
	bits []int
}

func (number *binary) getValue() int {
	var acc = 0
	var pow2 = 1

	for i := len(number.bits)-1; i >= 0; i-- {
		if number.bits[i] == 1 {
			acc += pow2
		}

		pow2 *= 2
	}

	return acc
}

func (number *binary) add(other *binary) binary {
	var addedBits = make([]int, len(number.bits)+1)

	var carry = 0

	for i := len(number.bits)-1; i >= 0; i-- {
		if carry == 1 {
			if number.bits[i] == 1 {
				carry = 1
			} else {
				carry = 0
			}
		} else {
			carry = number.bits[i]
		}
		if carry == 1 {
			if other.bits[i] == 1 {
				addedBits[i+1] = 0
				carry = 1
			} else {
				addedBits[i+1] = 1
				carry = 0
			}
		} else {
			addedBits[i+1] = other.bits[i]
		}
	}

	addedBits[0] = carry

	return binary {addedBits }
}

func main() {
	var four = binary { []int {1, 0, 0} }

	fmt.Printf("%d\n", four.getValue())

	var six = binary { []int {1, 1, 0} }

	ten := four.add(&six)

	fmt.Printf("%d\n", ten.getValue())

	var one = binary {[]int{0, 0, 1} }

	var five = four.add(&one)

	fmt.Printf("%d\n", five.getValue())

	var two = binary {[]int{0, 0, 1, 0} }

	var seven = five.add(&two)

	fmt.Printf("%d\n", seven.getValue())
}