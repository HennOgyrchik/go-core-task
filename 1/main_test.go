package main

import (
	"reflect"
	"testing"
)

func Test_printType(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{a: 5.4},
		},
		{
			name: "1",
			args: args{a: 035},
		},
		{
			name: "1",
			args: args{a: "agasgasdg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printType(tt.args.a)
		})
	}
}

func Test_numbers_String(t *testing.T) {
	type fields struct {
		d int
		o int
		h int
		f float64
		s string
		b bool
		c complex64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				d: 42,
				o: 052,
				h: 0x2A,
				f: 3.14,
				s: "Golang",
				b: true,
				c: 1 + 2i,
			},
			want: "42, 052, 0x2A, 3.14, Golang, true, (1+2i)",
		},
		{
			name: "2",
			fields: fields{
				d: 654,
				o: 025,
				h: 0xAA,
				f: 5.6483,
				s: "test",
				b: false,
				c: 2 - 3i,
			},
			want: "654, 025, 0xAA, 5.65, test, false, (2-3i)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := numbers{
				d: tt.fields.d,
				o: tt.fields.o,
				h: tt.fields.h,
				f: tt.fields.f,
				s: tt.fields.s,
				b: tt.fields.b,
				c: tt.fields.c,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toSliceRune(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "1",
			args: args{s: "hello"},
			want: []rune{104, 101, 108, 108, 111},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSliceRune(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toSliceRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printHash(t *testing.T) {
	type args struct {
		r    []rune
		salt string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				r:    []rune{104, 101, 108, 108, 111},
				salt: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printHash(tt.args.r, tt.args.salt)
		})
	}
}
