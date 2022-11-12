package breflect

import "testing"

type tInterface interface {
	test()
}

type tStruct struct {
	t any
}

func (t tStruct) test() {
}

func TestIsNil(t *testing.T) {
	var (
		baseInt    int     = 1
		baseString string  = "1"
		baseFloat  float64 = 1.1
	)
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "base",
			args: args{
				value: baseInt,
			},
			want: false,
		},
		{
			name: "base",
			args: args{
				value: baseString,
			},
			want: false,
		},
		{
			name: "base",
			args: args{
				value: baseFloat,
			},
			want: false,
		},
		{
			name: "point",
			args: args{
				value: &baseInt,
			},
			want: false,
		},
		{
			name: "point",
			args: args{
				value: &baseString,
			},
			want: false,
		},
		{
			name: "point",
			args: args{
				value: &baseFloat,
			},
			want: false,
		},
		{
			name: "point",
			args: args{
				value: (*int)(nil),
			},
			want: true,
		},
		{
			name: "point",
			args: args{
				value: (*string)(nil),
			},
			want: true,
		},
		{
			name: "point",
			args: args{
				value: (*float64)(nil),
			},
			want: true,
		},
		{
			name: "point",
			args: args{
				value: (*float64)(nil),
			},
			want: true,
		},
		{
			name: "struct",
			args: args{
				value: tStruct{t: "1"},
			},
			want: false,
		},
		{
			name: "point struct",
			args: args{
				value: &tStruct{t: "1"},
			},
			want: false,
		},
		{
			name: "point struct",
			args: args{
				value: (*tStruct)(nil),
			},
			want: true,
		},
		{
			name: "interface",
			args: args{
				value: (tInterface)(tStruct{t: "1"}),
			},
			want: false,
		},
		{
			name: "interface",
			args: args{
				value: (tInterface)(nil),
			},
			want: true,
		},
		{
			name: "point interface",
			args: args{
				value: (*tInterface)(nil),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.value); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
