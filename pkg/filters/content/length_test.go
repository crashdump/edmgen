package content_test

import (
	"testing"

	"edmgen/pkg/filters/content"
	"github.com/stretchr/testify/assert"
)

func Test_LineLength(t *testing.T) {
	tests := []struct {
		name    string
		content []string
		want    []string
	}{
		{
			name: "",
			content: []string{
				"Donec a dui et dui fringilla consectetur id nec massa.",          // 54 chars
				"Nulla aliquet porttitor venenatis.",                              // 34 chars
				"Nam tristique maximus ante hendrerit aliquet.",                   // 45 chars
				"Integer dignissim posuere lobortis.",                             // 35 chars
				"Aenean ultrices erat ut augue ultrices",                          // 38 chars
				"Suspendisse lacinia ante nunc, pulvinar blandit nisl ornare ut.", // 63 chars
			},
			want: []string{
				"Integer dignissim posuere lobortis.",
				"Aenean ultrices erat ut augue ultrices",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := content.LineLength(35, 40)(tt.content)
			assert.Equal(t, tt.want, got)
		})
	}
}
