package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTraining_distance(t1 *testing.T) {
	type fields struct {
		TrainingType string
		Action       int
		LenStep      float64
		Duration     time.Duration
		Weight       float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Distance Swimming",
			fields: fields{
				TrainingType: "Плавание",
				Action:       2000,
				LenStep:      SwimmingLenStep,
				Duration:     90 * time.Minute,
				Weight:       85,
			},
			want: 2.76,
		},
		{
			name: "Test Distance Walking",
			fields: fields{
				TrainingType: "Ходьба",
				Action:       20000,
				LenStep:      LenStep,
				Duration:     3*time.Hour + 45*time.Minute,
				Weight:       85,
			},
			want: 13,
		},
		{
			name: "Test Distance Running",
			fields: fields{
				TrainingType: "Бег",
				Action:       5000,
				LenStep:      LenStep,
				Duration:     30 * time.Minute,
				Weight:       85,
			},
			want: 3.25,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Training{
				TrainingType: tt.fields.TrainingType,
				Action:       tt.fields.Action,
				LenStep:      tt.fields.LenStep,
				Duration:     tt.fields.Duration,
				Weight:       tt.fields.Weight,
			}
			if got := t.distance(); got != tt.want {
				t1.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTraining_meanSpeed(t1 *testing.T) {
	type fields struct {
		TrainingType string
		Action       int
		LenStep      float64
		Duration     time.Duration
		Weight       float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Mean Speed Swimming",
			fields: fields{
				TrainingType: "Плавание",
				Action:       2000,
				LenStep:      SwimmingLenStep,
				Duration:     90 * time.Minute,
				Weight:       85,
			},
			want: 1.8399999999999999,
		},
		{
			name: "Test Mean Speed Walking",
			fields: fields{
				TrainingType: "Ходьба",
				Action:       20000,
				LenStep:      LenStep,
				Duration:     3*time.Hour + 45*time.Minute,
				Weight:       85,
			},
			want: 3.466666666666667,
		},
		{
			name: "Test Mean Speed Running",
			fields: fields{
				TrainingType: "Бег",
				Action:       5000,
				LenStep:      LenStep,
				Duration:     30 * time.Minute,
				Weight:       85,
			},
			want: 6.5,
		},
		{
			name: "Test Division by 0",
			fields: fields{
				TrainingType: "Плавание",
				Action:       2000,
				LenStep:      SwimmingLenStep,
				Duration:     0,
				Weight:       85,
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Training{
				TrainingType: tt.fields.TrainingType,
				Action:       tt.fields.Action,
				LenStep:      tt.fields.LenStep,
				Duration:     tt.fields.Duration,
				Weight:       tt.fields.Weight,
			}
			if got := t.meanSpeed(); got != tt.want {
				t1.Errorf("meanSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTraining_Calories(t1 *testing.T) {
	type fields struct {
		TrainingType string
		Action       int
		LenStep      float64
		Duration     time.Duration
		Weight       float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Null Data",
			fields: fields{
				TrainingType: "",
				Action:       0,
				LenStep:      0,
				Duration:     0,
				Weight:       0,
			},
			want: 0,
		},
		{
			name: "Test No Null Data",
			fields: fields{
				TrainingType: "Плавание",
				Action:       2000,
				LenStep:      SwimmingLenStep,
				Duration:     0,
				Weight:       85,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Training{
				TrainingType: tt.fields.TrainingType,
				Action:       tt.fields.Action,
				LenStep:      tt.fields.LenStep,
				Duration:     tt.fields.Duration,
				Weight:       tt.fields.Weight,
			}
			if got := t.Calories(); got != tt.want {
				t1.Errorf("Calories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTraining_TrainingInfo(t1 *testing.T) {
	type fields struct {
		TrainingType string
		Action       int
		LenStep      float64
		Duration     time.Duration
		Weight       float64
	}
	tests := []struct {
		name   string
		fields fields
		want   InfoMessage
	}{
		{
			name: "Test Training Info Running",
			fields: fields{
				TrainingType: "Бег",
				Action:       5000,
				LenStep:      LenStep,
				Duration:     30 * time.Minute,
				Weight:       85,
			},
			want: InfoMessage{
				TrainingType: "Бег",
				Duration:     30 * time.Minute,
				Distance:     3.25,
				Speed:        6.50,
				Calories:     0,
			},
		},
		{
			name: "Test Training Info No Data",
			fields: fields{
				TrainingType: "",
				Action:       0,
				LenStep:      0,
				Duration:     0,
				Weight:       0,
			},
			want: InfoMessage{
				TrainingType: "",
				Duration:     0,
				Distance:     0,
				Speed:        0,
				Calories:     0,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Training{
				TrainingType: tt.fields.TrainingType,
				Action:       tt.fields.Action,
				LenStep:      tt.fields.LenStep,
				Duration:     tt.fields.Duration,
				Weight:       tt.fields.Weight,
			}
			if got := t.TrainingInfo(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("TrainingInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
