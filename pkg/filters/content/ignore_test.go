package content_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/crashdump/edmgen/pkg/filters/content"
)

func Test_IgnoreLine(t *testing.T) {
	type args struct {
		string string
	}
	tests := []struct {
		name    string
		args    args
		content []string
		want    []string
	}{
		{
			name: "empty-string",
			content: []string{
				"Donec a dui et dui fringilla consectetur id nec massa.",
				"    Nulla aliquet porttitor venenatis.",
				"Nam tristique maximus ante hendrerit aliquet.",
				"Integer dignissim posuere lobortis.",
				"          Aenean ultrices erat ut augue ultrices",
				"Suspendisse lacinia ante nunc, pulvinar blandit nisl ornare ut.",
			},
			args: args{
				string: "",
			},
			want: []string(nil),
		},
		{
			name: "ignore-maximus",
			content: []string{
				"Donec a dui et dui fringilla consectetur id nec massa.",
				"Nulla aliquet porttitor venenatis.",
				"Nam tristique maximus ante hendrerit aliquet.",
				"Integer dignissim posuere lobortis.",
				"          Aenean ultrices erat ut augue ultrices",
				"Suspendisse lacinia ante nunc, pulvinar blandit nisl ornare ut.",
			},
			args: args{
				string: "maximus",
			},
			want: []string{
				"Donec a dui et dui fringilla consectetur id nec massa.",
				"Nulla aliquet porttitor venenatis.",
				"Integer dignissim posuere lobortis.",
				"          Aenean ultrices erat ut augue ultrices",
				"Suspendisse lacinia ante nunc, pulvinar blandit nisl ornare ut.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := content.IgnoreLine(tt.args.string)(tt.content)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
