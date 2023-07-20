package content_test

import (
	"testing"

	"edmgen/pkg/filters/content"
	"github.com/stretchr/testify/assert"
)

func Test_LongestLine(t *testing.T) {
	tests := []struct {
		name    string
		content []string
		want    []string
	}{
		{
			name: "",
			content: []string{
				"Donec a dui et dui fringilla consectetur id nec massa.",          // > 40 chars
				"Nulla aliquet porttitor venenatis.",                              // < 40 chars
				"Nam tristique maximus ante hendrerit aliquet.",                   // > 40 chars
				"Suspendisse lacinia ante nunc, pulvinar blandit nisl ornare ut.", // > 40 chars
				"Integer dignissim posuere lobortis. ",                            // < 40 chars
				"Aenean ultrices erat ut augue ultrices",                          // < 40 chars

			},
			want: []string{
				"Suspendisse lacinia ante nunc, pulvinar blandit nisl ornare ut.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := content.LongestLine(tt.content)
			assert.Len(t, got, 1)
			assert.Equal(t, tt.want, got)
		})
	}
}
