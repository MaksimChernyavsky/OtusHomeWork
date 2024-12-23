package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		fromPath string
		toPath   string
		expPath  string
		offset   int64
		limit    int64
		name     string
		err      error
	}{
		{
			fromPath: "testdata/input.txt",
			toPath:   "/tmp/success.txt",
			expPath:  "testdata/out_offset100_limit1000.txt",
			offset:   100,
			limit:    1000,
			name:     "success",
			err:      nil,
		},
		{
			fromPath: "/dev/urandom",
			toPath:   "",
			expPath:  "",
			offset:   0,
			limit:    0,
			name:     "unsupported_file",
			err:      ErrUnsupportedFile,
		},
		{
			fromPath: "testdata/input.txt",
			toPath:   "",
			expPath:  "",
			offset:   1000000000000,
			limit:    1000,
			name:     "offset_exceeds_file_size",
			err:      ErrOffsetExceedsFileSize,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case %s", tt.name), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Copy(tt.fromPath, tt.toPath, tt.offset, tt.limit)
			if tt.err == nil {
				require.Empty(t, err, "error after copy must be empty")
				res, err := os.ReadFile("/tmp/" + tt.name + ".txt")
				require.Empty(t, err, "error after read all result must be empty")
				exp, err := os.ReadFile(tt.expPath)
				require.Empty(t, err, "error after read all expected must be empty")
				require.Equal(t, exp, res, "incorrect result after copy")
			} else {
				require.ErrorIs(t, err, tt.err, "error after copy must be correct")
			}
		})
	}
}
