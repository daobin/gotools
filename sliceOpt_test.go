package gotools

import (
	"reflect"
	"testing"
)

func Test_sliceOpt_DiffInt(t *testing.T) {
	type args struct {
		data1 []int
		data2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "前空后有",
			args: args{
				data1: []int{},
				data2: []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			name: "前有后空",
			args: args{
				data1: []int{1, 2, 3},
				data2: []int{},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "前有后有",
			args: args{
				data1: []int{1, 2, 3},
				data2: []int{2, 3, 4},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := sliceOpt{}
			if got := opt.DiffInt(tt.args.data1, tt.args.data2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_sliceOpt_DiffString(t *testing.T) {
//	type args struct {
//		data1 []string
//		data2 []string
//	}
//	tests := []struct {
//		name string
//		args args
//		want []string
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			opt := sliceOpt{}
//			if got := opt.DiffString(tt.args.data1, tt.args.data2); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("DiffString() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_sliceOpt_IndexOfInt(t *testing.T) {
//	type args struct {
//		data    []int
//		element int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			opt := sliceOpt{}
//			if got := opt.IndexOfInt(tt.args.data, tt.args.element); got != tt.want {
//				t.Errorf("IndexOfInt() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_sliceOpt_IndexOfString(t *testing.T) {
//	type args struct {
//		data    []string
//		element string
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			opt := sliceOpt{}
//			if got := opt.IndexOfString(tt.args.data, tt.args.element); got != tt.want {
//				t.Errorf("IndexOfString() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_sliceOpt_IntersectInt(t *testing.T) {
//	type args struct {
//		data1 []int
//		data2 []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want []int
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			opt := sliceOpt{}
//			if got := opt.IntersectInt(tt.args.data1, tt.args.data2); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("IntersectInt() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_sliceOpt_IntersectString(t *testing.T) {
//	type args struct {
//		data1 []string
//		data2 []string
//	}
//	tests := []struct {
//		name string
//		args args
//		want []string
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			opt := sliceOpt{}
//			if got := opt.IntersectString(tt.args.data1, tt.args.data2); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("IntersectString() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_sliceOpt_MergeInt(t *testing.T) {
//	type args struct {
//		data1 []int
//		data2 []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want []int
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			opt := sliceOpt{}
//			if got := opt.MergeInt(tt.args.data1, tt.args.data2); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MergeInt() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_sliceOpt_MergeString(t *testing.T) {
//	type args struct {
//		data1 []string
//		data2 []string
//	}
//	tests := []struct {
//		name string
//		args args
//		want []string
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			opt := sliceOpt{}
//			if got := opt.MergeString(tt.args.data1, tt.args.data2); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MergeString() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
