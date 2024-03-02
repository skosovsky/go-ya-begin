package main

import (
	"reflect"
	"testing"
)

func Test_sortSlice(t *testing.T) {
	type args struct {
		slc []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{[]int{5, 6, 9, -10, 0}},
			want: []int{9, 6, 5, 0, -10},
		},
		{
			name: "Example 2",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "Example 3",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "Example 4",
			args: args{[]int{1, 1, 1, 1, 2, 2}},
			want: []int{2, 2, 1, 1, 1, 1},
		},
		{
			name: "Example 5",
			args: args{[]int{-15, -1, -10}},
			want: []int{-1, -10, -15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSlice(tt.args.slc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortSliceLib(t *testing.T) {
	type args struct {
		slc []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{[]int{5, 6, 9, -10, 0}},
			want: []int{9, 6, 5, 0, -10},
		},
		{
			name: "Example 2",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "Example 3",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "Example 4",
			args: args{[]int{1, 1, 1, 1, 2, 2}},
			want: []int{2, 2, 1, 1, 1, 1},
		},
		{
			name: "Example 5",
			args: args{[]int{-15, -1, -10}},
			want: []int{-1, -10, -15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSliceLib(tt.args.slc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortSliceLib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortSliceLibNew(t *testing.T) {
	type args struct {
		slc []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{[]int{5, 6, 9, -10, 0}},
			want: []int{9, 6, 5, 0, -10},
		},
		{
			name: "Example 2",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "Example 3",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "Example 4",
			args: args{[]int{1, 1, 1, 1, 2, 2}},
			want: []int{2, 2, 1, 1, 1, 1},
		},
		{
			name: "Example 5",
			args: args{[]int{-15, -1, -10}},
			want: []int{-1, -10, -15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSliceLibNew(tt.args.slc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortSliceLibNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findX(t *testing.T) {
	type args struct {
		slc []int
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{[]int{10, 9, 8, 7, 6, 5}, 11},
			want: false,
		},
		{
			name: "Example 2",
			args: args{[]int{10, 9, 8, 7, 6, 5}, 7},
			want: true,
		},
		{
			name: "Example 3",
			args: args{[]int{}, 7},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findX(tt.args.slc, tt.args.num); got != tt.want {
				t.Errorf("findX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findXLib(t *testing.T) {
	type args struct {
		slc []int
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{[]int{10, 9, 8, 7, 6, 5}, 11},
			want: false,
		},
		{
			name: "Example 2",
			args: args{[]int{10, 9, 8, 7, 6, 5}, 7},
			want: true,
		},
		{
			name: "Example 3",
			args: args{[]int{}, 7},
			want: false,
		},
		{
			name: "Example 4",
			args: args{[]int{5, 6, 7, 8, 9, 10}, 7},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findXLib(tt.args.slc, tt.args.num); got != tt.want {
				t.Errorf("findXLib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nameSort(t *testing.T) {
	type args struct {
		slc []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{[]string{"Иван Фролов", "Ян Косовский", "Сергей Косовский", "Марк Косовский"}},
			want: []string{"Марк Косовский", "Сергей Косовский", "Ян Косовский", "Иван Фролов"},
		},
		{
			name: "Example 2",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "Example 3",
			args: args{[]string{"Иван", "Ян Косовский"}},
			want: []string{"Ян Косовский"},
		},
		{
			name: "Example 4",
			args: args{[]string{"Иван"}},
			want: []string{},
		},
		{
			name: "Example 5",
			args: args{[]string{"Mark Kosovsky", "Ivan Frolov"}},
			want: []string{"Ivan Frolov", "Mark Kosovsky"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nameSort(tt.args.slc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nameSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
