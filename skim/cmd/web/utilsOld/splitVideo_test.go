package utilsOld

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestSplitVideo(t *testing.T) {
	tests := []struct {
		name        string
		expectedErr error
	}{
		{
			name:        "initial",
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SplitVideo()
			if tt.expectedErr != nil {
				t.Fatal("have not yet implement errors ")
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %q\n", err)
			}

			_, err = os.Stat("../../../uploads/splitVids")
			if os.IsNotExist(err) {
				t.Fatal("splitVids dir does not exist.")
			}

			clips := []string{}

			filepath.Walk(
				"../../../uploads/splitVids",
				func(path string, info fs.FileInfo, err error) error {
					if info.IsDir() {
						return nil
					}
					clips = append(clips, path)
					return nil
				},
			)
			if len(clips) < 1 {
				t.Fatal("the clips aren't splitVids dir")
			}
		})
	}
}
