package file_test

import (
	"testing"

	"edmgen/pkg/edm"
	"edmgen/pkg/filters/file"
	"github.com/stretchr/testify/assert"
)

func Test_Filename(t *testing.T) {

	tests := []struct {
		name        string
		filter      file.Filter
		expectPaths []string
		expectError error
	}{
		{
			name:   "ignore-filename",
			filter: file.IgnoreFilename([]string{"hello.py"}),
			expectPaths: []string{
				"../../../test/fixtures/static/hello.java",
				"../../../test/fixtures/static/foo/bar.txt",
				"../../../test/fixtures/static/hello.c",
			},
			expectError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			edmc, err := edm.New(edm.Opts{})
			assert.NoError(t, err)

			err = edmc.SelectFiles("../../../test/fixtures/static", tt.filter)
			assert.ElementsMatch(t, tt.expectPaths, edmc.Paths)
			assert.Equal(t, tt.expectError, err)
		})
	}
}
