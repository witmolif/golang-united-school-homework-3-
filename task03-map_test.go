package homework

import (
	"reflect"
	"testing"
)

func Test_sortMapValues(t *testing.T) {
	type args struct {
		input map[int]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{"t1", args{map[int]string{2: "a", 0: "b", 1: "c"}}, []string{"b", "c", "a"}},
		{"t2", args{map[int]string{10: "aa", 0: "bb", 500: "cc"}}, []string{"bb", "aa", "cc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := sortMapValues(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("sortMapValues() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
