package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pluralizer_pluralize(t *testing.T) {
	tests := map[string]struct {
		recv pluralizer
		arg  string
		want string
	}{
		"ok: common case": {
			recv: pluralizer{},
			arg:  "User",
			want: "Users",
		},
	}
	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.recv.pluralize(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
