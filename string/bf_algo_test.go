package string

import "testing"

func TestBF(t *testing.T) {
	type args struct {
		mainString  string
		modelString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// BF算法
		{
			"没有成功匹配",
			args{"abc123", "b"},
			1,
		},
		{
			"成功匹配",
			args{"abc123", "a"},
			0,
		},
		{
			"成功匹配2",
			args{"abc123", "3"},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BF(tt.args.mainString, tt.args.modelString); got != tt.want {
				t.Errorf("BF() = %v, want %v", got, tt.want)
			}
		})
	}
}
