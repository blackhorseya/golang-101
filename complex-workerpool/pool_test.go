package main

import (
	"testing"
)

func TestNewWorkerPool(t *testing.T) {
	type args struct {
		workerNums int
	}
	tests := []struct {
		name    string
		args    args
		want    Pool
		wantErr bool
	}{
		{
			name:    "no workers then error",
			args:    args{workerNums: -1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "create worker pool with 5 workers",
			args:    args{workerNums: 5},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWorkerPool(tt.args.workerNums)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWorkerPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got == nil {
				t.Errorf("NewWorkerPool() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestNewWorker_MultipleStartAndStop(t *testing.T) {
	pool, err := NewWorkerPool(5)
	if err != nil {
		t.Fatal(err)
	}

	pool.Start()
	pool.Start()

	pool.Stop()
	pool.Stop()
}
