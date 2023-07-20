package edm_test

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"edmgen/pkg/edm"
	"edmgen/pkg/filters/content"
	"edmgen/pkg/filters/file"
	"github.com/stretchr/testify/assert"
)

var fixturesDir = "../../test/fixtures"

func TestWalk_New(t *testing.T) {
	_, err := edm.New(edm.Opts{})
	assert.NoError(t, err)
}

func Test_SelectFiles(t *testing.T) {
	tests := []struct {
		name          string
		filters       []file.Filter
		expectNbFiles int
		expectError   bool
	}{
		{
			name: "simple-filename-filter",
			filters: []file.Filter{
				// Expect to match only files with a ".c" extension
				func(path fs.DirEntry) (bool, error) {
					if strings.HasSuffix(path.Name(), ".c") {
						return true, nil
					}
					return false, nil
				},
			},
			expectNbFiles: 1,
			expectError:   false,
		},
		{
			name: "simple-directory-name-filter",
			filters: []file.Filter{
				// Expect to ignore the "foo" directory and the files within
				func(path fs.DirEntry) (bool, error) {
					if path.IsDir() && path.Name() == "foo" {
						return false, filepath.SkipDir
					}
					return true, nil
				},
			},
			expectNbFiles: 3,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			edmc, err := edm.New(edm.Opts{})
			assert.NoError(t, err)

			path := fmt.Sprintf("%s/static", fixturesDir)
			err = edmc.SelectFiles(path, tt.filters...)
			assert.NoError(t, err)
			assert.Len(t, edmc.Paths, tt.expectNbFiles, "should match only X files")
		})
	}
}

func Test_ExamineFiles(t *testing.T) {
	tests := []struct {
		name          string
		filepaths     []string
		filters       []content.Filter
		expectNbLines int
		expectError   bool
	}{
		{
			name: "simple-line-length-filter",
			filepaths: []string{
				fmt.Sprintf("%s/static/hello.java", fixturesDir),
			},
			filters: []content.Filter{
				func(content []string) []string {
					var result []string
					for _, line := range content {
						if len(line) >= 40 {
							result = append(result, line)
						}
					}
					return result
				},
			},
			expectNbLines: 2,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			edmc, err := edm.New(edm.Opts{})
			assert.NoError(t, err)

			edmc.Paths = tt.filepaths
			err = edmc.ExamineFiles(tt.filters...)
			assert.NoError(t, err)
			assert.Len(t, edmc.Content, tt.expectNbLines, "should match only X lines")
		})
	}
}

func Test_SampleContent(t *testing.T) {
	tests := []struct {
		name          string
		content       []string
		filters       []content.Filter
		expectContent []string
	}{
		{
			name: "simple-line-length-filter",
			content: []string{
				"foo",
				"bar",
				"foo",
				"baz",
				"foo",
			},
			filters: []content.Filter{
				content.Uniq,
			},
			expectContent: []string{
				"foo",
				"bar",
				"baz",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			edmc, err := edm.New(edm.Opts{})
			assert.NoError(t, err)

			edmc.Content = tt.content
			res := edmc.SampleContent(tt.filters...)
			assert.EqualValues(t, tt.expectContent, res)
		})
	}
}

func Test_ApplyFilter(t *testing.T) {
	tests := []struct {
		name        string
		lines       []string
		filters     []content.Filter
		expectLines []string
	}{
		{
			name:        "no-filters",
			lines:       []string{"foo"},
			filters:     []content.Filter{},
			expectLines: []string{"foo"},
		},
		{
			name: "one-filter",
			lines: []string{
				"foobar",
				"baz",
			},
			filters: []content.Filter{
				content.LongestLine,
			},
			expectLines: []string{"foobar"},
		},
		{
			name: "two-filters",
			lines: []string{
				"foobar",
				"foobar",
				"baz",
				"foobar",
				"baz",
			},
			filters: []content.Filter{
				content.LongestLine,
				content.Uniq,
			},
			expectLines: []string{"foobar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := edm.ApplyFilters(tt.lines, tt.filters...)
			assert.Equal(t, tt.expectLines, result)
		})
	}
}
