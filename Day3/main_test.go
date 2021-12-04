package main

import "testing"

func Test_calculateOxygenAndCo2Rating(t *testing.T) {
	type args struct {
		diagnostics [][]int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		want1   int64
		wantErr bool
	}{
		{
			"test1",
			args{
				[][]int{
					{0, 0, 1, 0, 0},
					{1, 1, 1, 1, 0},
					{1, 0, 1, 1, 0},
					{1, 0, 1, 1, 1},
					{1, 0, 1, 0, 1},
					{0, 1, 1, 1, 1},
					{0, 0, 1, 1, 1},
					{1, 1, 1, 0, 0},
					{1, 0, 0, 0, 0},
					{1, 1, 0, 0, 1},
					{0, 0, 0, 1, 0},
					{0, 1, 0, 1, 0},
				},
			},
			23,
			10,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := calculateOxygenAndCo2Rating(tt.args.diagnostics)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateOxygenAndCo2Rating() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateOxygenAndCo2Rating() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculateOxygenAndCo2Rating() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
