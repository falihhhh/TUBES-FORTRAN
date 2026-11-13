package main

import "fmt"

type arrInt [100]int

func digitPuluhan(x int) int {
	return (x / 10) % 10
}
func isiArray(A *arrInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}
func tampilArray(A arrInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(A[i])
	}
}
func insertionSort(A *arrInt, n int) {
	var pass, i, temp int

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && (digitPuluhan(temp) < digitPuluhan(A[i-1]) ||
			(digitPuluhan(temp) == digitPuluhan(A[i-1]) && temp < A[i-1])) {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}
func main() {
	var A arrInt
	var n int
	fmt.Scan(&n)
	isiArray(&A, n)
	fmt.Println("Data sebelum sorting:")
	tampilArray(A, n)
	insertionSort(&A, n)
	fmt.Println()
	fmt.Println("Data setelah sorting:")
	tampilArray(A, n)
}
