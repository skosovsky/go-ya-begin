package ftracker

import (
	"testing"
)

func Test_distance(t *testing.T) {
	type args struct {
		action int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Success test",
			args: args{
				action: 2000,
			},
			want: 1.3,
		},
		{
			name: "Null action",
			args: args{
				action: 0,
			},
			want: 0.0,
		},
		{
			name: "One action",
			args: args{
				action: 1,
			},
			want: 0.00065,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.action); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_meanSpeed(t *testing.T) {
	type args struct {
		action   int
		duration float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Successful test",
			args: args{
				action:   2000,
				duration: 2.0,
			},
			want: 0.65,
		},
		{
			name: "Null duration",
			args: args{
				action:   2000,
				duration: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := meanSpeed(tt.args.action, tt.args.duration); got != tt.want {
				t.Errorf("meanSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowTrainingInfo(t *testing.T) {
	type args struct {
		action       int
		trainingType string
		duration     float64
		weight       float64
		height       float64
		lengthPool   int
		countPool    int
	}

	runTest := "Тип тренировки: Бег\nДлительность: 0.15 ч.\nДистанция: 2.60 км.\nСкорость: 17.33 км/ч\nСожгли калорий: 427.24\n"

	walkingTest := "Тип тренировки: Ходьба\nДлительность: 1.00 ч.\nДистанция: 2.60 км.\nСкорость: 2.60 км/ч\nСожгли калорий: 220.27\n"

	swimmingTest := "Тип тренировки: Плавание\nДлительность: 0.25 ч.\nДистанция: 0.65 км.\nСкорость: 1.60 км/ч\nСожгли калорий: 114.75\n"

	unknownTest := "неизвестный тип тренировки"

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "run test",
			args: args{
				action:       4000,
				trainingType: "Бег",
				duration:     0.15,
				weight:       85,
				height:       185,
				lengthPool:   50,
				countPool:    2,
			},
			want: runTest,
		},

		{
			name: "walking test",
			args: args{
				action:       4000,
				trainingType: "Ходьба",
				duration:     1,
				weight:       85,
				height:       185,
				lengthPool:   50,
				countPool:    2,
			},
			want: walkingTest,
		},

		{
			name: "swimming test",
			args: args{
				action:       1000,
				trainingType: "Плавание",
				duration:     0.25,
				weight:       85,
				height:       185,
				lengthPool:   100,
				countPool:    4,
			},
			want: swimmingTest,
		},

		{
			name: "unknown test",
			args: args{
				action:       1000,
				trainingType: "Керлинг",
				duration:     5,
				weight:       85,
				height:       185,
				lengthPool:   50,
				countPool:    2,
			},
			want: unknownTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShowTrainingInfo(tt.args.action, tt.args.trainingType, tt.args.duration, tt.args.weight, tt.args.height, tt.args.lengthPool, tt.args.countPool); got != tt.want {
				t.Errorf("ShowTrainingInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
