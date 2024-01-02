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
