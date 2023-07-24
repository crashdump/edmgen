package content_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/crashdump/edmgen/pkg/filters/content"
)

func Test_Uniq(t *testing.T) {
	tests := []struct {
		name       string
		content    []string
		want       []string
		wantLength int
	}{
		{
			name: "",
			content: []string{
				"Nulla aliquet porttitor venenatis.",
				"Donec a dui et dui fringilla consectetur id nec massa.",
				"Nulla aliquet porttitor venenatis.",
				"Nam tristique maximus ante hendrerit aliquet.",
				"Nulla aliquet porttitor venenatis.",
			},
			want: []string{
				"Donec a dui et dui fringilla consectetur id nec massa.",
				"Nulla aliquet porttitor venenatis.",
				"Nam tristique maximus ante hendrerit aliquet.",
			},
			wantLength: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := content.Uniq(tt.content)
			assert.Len(t, got, tt.wantLength)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
