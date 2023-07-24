package file_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/crashdump/edmgen/pkg/edm"
	"github.com/crashdump/edmgen/pkg/filters/file"
)

func Test_Extension(t *testing.T) {

	tests := []struct {
		name        string
		filter      file.Filter
		expectPaths []string
		expectError error
	}{
		{
			name:   "require-extension",
			filter: file.RequireExtensions([]string{".java"}),

			expectPaths: []string{"../../../test/fixtures/static/hello.java"},
			expectError: nil,
		},
		{
			name:   "ignore-extension",
			filter: file.IgnoreExtensions([]string{".c"}),
			expectPaths: []string{
				"../../../test/fixtures/static/hello.java",
				"../../../test/fixtures/static/hello.py",
				"../../../test/fixtures/static/foo/bar.txt",
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
