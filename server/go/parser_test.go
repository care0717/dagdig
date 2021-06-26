package openapi

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"io"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		input io.Reader
		want Job
	}{
		{
			name: "normal",
			input:  bytes.NewBufferString("[JobID]\n1\n\n[Created]\n00:00:05\n\n[Priority]\nLow\n\n[Tasks]\n4\n3\n"),
			want: Job{
				Id:       1,
				Created:  "00:00:05",
				Priority: "Low",
				Tasks:    []int32{4,3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, tt.want) {
				t.Fatalf("Parse (-got +want): %s", cmp.Diff(got, tt.want))
			}
		})
	}

	errorTests := []struct {
		name string
		input io.Reader
	}{
		{
			name: "Job id is not int",
			input:  bytes.NewBufferString("[JobID]\nhoge\n\n[Created]\n00:00:05\n\n[Priority]\nLow\n\n[Tasks]\n4\n3\n"),
		},
		{
			name: "Tasks is not int",
			input:  bytes.NewBufferString("[JobID]\n1\n\n[Created]\n00:00:05\n\n[Priority]\nLow\n\n[Tasks]\n4\nfuga\n3\n"),
		},
	}
	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.input)
			if err == nil {
				t.Errorf("input invalid case but succeeded")
			}
		})
	}
}
