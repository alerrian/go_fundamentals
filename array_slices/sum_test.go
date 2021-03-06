package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectSum := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()

		if got != want {
			t.Errorf("\nGot: %d\nWant: %d\nGiven: %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
			t.Errorf("\nGot: %d\nWant: %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nGot: %d\nWant: %v", got, want)
		}
	}
	
	t.Run("can get sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSum(t, got, want)
	})

	t.Run("can sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSum(t, got, want)
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)

	fmt.Println(sum)
	//Output: 15
}

func ExampleSumAll() {
	sumAll := SumAll([]int{1,2}, []int{0,9})

	fmt.Println(sumAll)
	//Output: [3 9]
}

func ExampleSumAllTails() {
	sumAllTails := SumAllTails([]int{1, 2}, []int{0, 9})

	fmt.Println(sumAllTails)
	//Output: [2 9]
}
