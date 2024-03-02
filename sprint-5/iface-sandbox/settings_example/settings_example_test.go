package settings_example

import "testing"

func TestGetSum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1. Positive numbers",
			args: args{a: []int{1, 2, 3}},
			want: 6,
		},
		{
			name: "2. Negative numbers",
			args: args{a: []int{-1, -2, -3}},
			want: -6,
		},
		{
			name: "3. 1 element slice",
			args: args{a: []int{10}},
			want: 10,
		},
		{
			name: "4. Empty slice",
			args: args{a: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetSum(tt.args.a); got != tt.want {
				t.Errorf("GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
