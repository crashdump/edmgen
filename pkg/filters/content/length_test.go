package content_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/crashdump/edmgen/pkg/filters/content"
)

func Test_LineLength(t *testing.T) {
	type args struct {
		min      int
		max      int
		ignoreWs bool
	}
	tests := []struct {
		name    string
		args    args
		content []string
		want    []string
	}{
		{
			name: "min-35-max-40-includes-ws",
			content: []string{
				"Donec a dui et dui fringilla consectetur id nec massa.",          // 54 chars
				"    Nulla aliquet porttitor venenatis.",                          // 38 chars
				"Nam tristique maximus ante hendrerit aliquet.",                   // 45 chars
				"Integer dignissim posuere lobortis.",                             // 35 chars
				"          Aenean ultrices erat ut augue ultrices",                // 38 chars
				"Suspendisse lacinia ante nunc, pulvinar blandit nisl ornare ut.", // 63 chars
			},
			args: args{
				min:      35,
				max:      40,
				ignoreWs: false,
			},
			want: []string{
				"    Nulla aliquet porttitor venenatis.",
				"Integer dignissim posuere lobortis.",
			},
		},
		{
			name: "min-35-max-40-ignore-ws",
			content: []string{
				"Donec a dui et dui fringilla consectetur id nec massa.",          // 54 chars
				"Nulla aliquet porttitor venenatis.",                              // 34 chars
				"Nam tristique maximus ante hendrerit aliquet.",                   // 45 chars
				"Integer dignissim posuere lobortis.",                             // 35 chars
				"          Aenean ultrices erat ut augue ultrices",                // 48 chars
				"Suspendisse lacinia ante nunc, pulvinar blandit nisl ornare ut.", // 63 chars
			},
			args: args{
				min:      35,
				max:      40,
				ignoreWs: true,
			},
			want: []string{
				"Nam tristique maximus ante hendrerit aliquet.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := content.LineLength(tt.args.min, tt.args.max, tt.args.ignoreWs)(tt.content)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
