package waitGroup

import (
	"sync"
	"testing"
)

func TestWaitGroup_Add(t *testing.T) {
	type fields struct {
		count int
		sem   chan struct{}
		mu    sync.Mutex
	}
	type args struct {
		delta int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      int
		wantPanic bool
	}{
		{
			name: "test 1",
			fields: fields{
				count: 0,
				sem:   nil,
				mu:    sync.Mutex{},
			},
			args:      args{delta: 3},
			want:      3,
			wantPanic: false,
		},
		{
			name: "test 2",
			fields: fields{
				count: 1,
				sem:   nil,
				mu:    sync.Mutex{},
			},
			args:      args{delta: -1},
			want:      0,
			wantPanic: false,
		},
		{
			name: "test 3",
			fields: fields{
				count: 0,
				sem:   nil,
				mu:    sync.Mutex{},
			},
			args:      args{delta: -1},
			want:      -1,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &WaitGroup{
				count: tt.fields.count,
				sem:   tt.fields.sem,
				mu:    tt.fields.mu,
			}

			defer func() {
				if r := recover(); (r == nil) == tt.wantPanic {
					t.Errorf("The code did not panic")
				}

			}()

			wg.Add(tt.args.delta)

			if wg.count != tt.want {
				t.Errorf("add() = %v, want %v", wg.count, tt.want)
			}

		})
	}
}

func TestWaitGroup_Done(t *testing.T) {
	type fields struct {
		count int
		sem   chan struct{}
		mu    sync.Mutex
	}
	tests := []struct {
		name      string
		fields    fields
		want      int
		wantPanic bool
	}{
		{
			name: "test 1",
			fields: fields{
				count: 3,
				sem:   nil,
				mu:    sync.Mutex{},
			},
			want:      2,
			wantPanic: false,
		},
		{
			name: "test 2",
			fields: fields{
				count: 0,
				sem:   nil,
				mu:    sync.Mutex{},
			},
			want:      -1,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &WaitGroup{
				count: tt.fields.count,
				sem:   tt.fields.sem,
				mu:    tt.fields.mu,
			}

			defer func() {
				if r := recover(); (r == nil) == tt.wantPanic {
					t.Errorf("The code did not panic")
				}

			}()
			wg.Done()

			if wg.count != tt.want {
				t.Errorf("add() = %v, want %v", wg.count, tt.want)
			}
		})
	}
}
