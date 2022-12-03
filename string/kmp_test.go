package string

import "testing"

func TestKMP(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "测试KMP算法",
			args: args{
				a: "United-States",
				b: "ted",
			},
			want: 3,
		},
		{
			name: "测试KMP算法2",
			args: args{
				a: "Uni-ed-States",
				b: "-ed",
			},
			want: 3,
		},
		{
			name: "测试KMP算法2",
			args: args{
				a: "United-States",
				b: "-",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMP(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}
