package openapi

import (
	"testing"
	"time"
)

func TestToCreatedFormat(t *testing.T) {
	tests := []struct {
		name string
		d time.Duration
		want string
	}{
		{
			name: "0s",
			d: 0*time.Second,
			want: "00:00:00",
		},
		{
			name: "59s",
			d: 59*time.Second,
			want: "00:00:59",
		},
		{
			name: "61s",
			d: 61*time.Second,
			want: "00:01:01",
		},
		{
			name: "3599s",
			d: 3599*time.Second,
			want: "00:59:59",
		},
		{
			name: "3661s",
			d: 3661*time.Second,
			want: "01:01:01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCreatedFormat(tt.d); got != tt.want {
				t.Errorf("ToCreatedFormat = %v, want %v", got, tt.want)
			}
		})
	}
}

