package homework

import "testing"

func Test_average(t *testing.T) {
	type args struct {
		input [15]float32
	}
	tests := []struct {
		name       string
		args       args
		wantResult float32
	}{
		{name: "t1", args: args{[15]float32{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3}}, wantResult: 3.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := average(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("average() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
