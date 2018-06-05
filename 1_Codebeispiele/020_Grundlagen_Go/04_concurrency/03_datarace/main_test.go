package main

import "testing"

func Test_datarace(t *testing.T) {
	type args struct {
		nr int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"6",
			args{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datarace(tt.args.nr)
		})
	}
}
