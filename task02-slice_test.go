package homework

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		input []int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int64
	}{
		{name: "t1", args: args{input: []int64{1, 2, 3, 4, 5}}, wantResult: []int64{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := reverse(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("reverse() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
