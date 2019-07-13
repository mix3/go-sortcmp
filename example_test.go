package sortcmp_test

import (
	"fmt"
	"sort"

	"github.com/mix3/go-sortcmp"
)

func ExampleCompare() {
	data := []struct {
		A int
		B int
		C int
	}{
		{3, 2, 5},
		{9, 5, 5},
		{6, 7, 4},
		{1, 2, 3},
		{2, 10, 3},
		{10, 6, 3},
		{10, 1, 7},
		{1, 6, 5},
		{5, 1, 8},
		{4, 4, 4},
		{10, 4, 4},
		{4, 9, 8},
		{9, 8, 1},
		{2, 8, 10},
		{7, 8, 4},
		{9, 1, 5},
		{9, 3, 9},
		{3, 7, 2},
		{1, 9, 4},
		{4, 2, 2},
		{5, 10, 3},
		{2, 8, 8},
		{1, 3, 2},
		{10, 8, 5},
		{8, 9, 10},
		{9, 1, 10},
		{10, 6, 2},
		{5, 1, 8},
		{6, 5, 3},
		{10, 9, 2},
	}

	sort.Slice(data, func(i, j int) bool {
		return sortcmp.Compare(
			sortcmp.CompareInt(data[i].A, data[j].A),
			sortcmp.CompareInt(data[i].B, data[j].B),
			sortcmp.CompareInt(data[i].C, data[j].C),
		)
	})
	for _, v := range data {
		fmt.Println(v)
	}

	// Output:
	// {1 2 3}
	// {1 3 2}
	// {1 6 5}
	// {1 9 4}
	// {2 8 8}
	// {2 8 10}
	// {2 10 3}
	// {3 2 5}
	// {3 7 2}
	// {4 2 2}
	// {4 4 4}
	// {4 9 8}
	// {5 1 8}
	// {5 1 8}
	// {5 10 3}
	// {6 5 3}
	// {6 7 4}
	// {7 8 4}
	// {8 9 10}
	// {9 1 5}
	// {9 1 10}
	// {9 3 9}
	// {9 5 5}
	// {9 8 1}
	// {10 1 7}
	// {10 4 4}
	// {10 6 2}
	// {10 6 3}
	// {10 8 5}
	// {10 9 2}
}
