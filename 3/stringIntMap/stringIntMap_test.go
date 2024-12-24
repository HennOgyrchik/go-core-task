package stringIntMap

import (
	"reflect"
	"testing"
)

func TestStringIntMap_Add(t *testing.T) {
	type fields struct {
		x map[string]int
	}
	type args struct {
		key   string
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int
	}{

		{
			name:   "1",
			fields: fields{x: map[string]int{}},
			args: args{
				key:   "test",
				value: 150156,
			},
			want: map[string]int{"test": 150156},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{
				x: tt.fields.x,
			}
			s.Add(tt.args.key, tt.args.value)

			if !reflect.DeepEqual(s.x, tt.want) {
				t.Errorf("Add() = %v, want %v", s.x, tt.want)
			}
		})
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	type fields struct {
		x map[string]int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int
	}{
		// TODO: Add test cases.
		{
			name:   "test",
			fields: fields{x: map[string]int{"hello": 484681, "world": 1}},
			args:   args{key: "hello"},
			want:   map[string]int{"world": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{
				x: tt.fields.x,
			}
			s.Remove(tt.args.key)

			if !reflect.DeepEqual(s.x, tt.want) {
				t.Errorf("Add() = %v, want %v", s.x, tt.want)
			}
		})
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	type fields struct {
		x map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]int
	}{
		{
			name:   "test",
			fields: fields{x: map[string]int{"hello": 2, "world": 1}},
			want:   map[string]int{"hello": 2, "world": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{
				x: tt.fields.x,
			}
			if got := s.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringIntMap_Exists(t *testing.T) {
	type fields struct {
		x map[string]int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "test1",
			fields: fields{x: map[string]int{"hello": 2, "world": 1}},
			args:   args{key: "hello"},
			want:   true,
		},
		{
			name:   "test2",
			fields: fields{x: map[string]int{"hello": 2, "world": 1}},
			args:   args{key: "cat"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{
				x: tt.fields.x,
			}
			if got := s.Exists(tt.args.key); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringIntMap_Get(t *testing.T) {
	type fields struct {
		x map[string]int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		{
			name:   "test1",
			fields: fields{x: map[string]int{"hello": 2, "world": 1}},
			args:   args{key: "hello"},
			want:   2,
			want1:  true,
		},
		{
			name:   "test2",
			fields: fields{x: map[string]int{"hello": 2, "world": 1}},
			args:   args{key: "cat"},
			want:   0,
			want1:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{
				x: tt.fields.x,
			}
			got, got1 := s.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
