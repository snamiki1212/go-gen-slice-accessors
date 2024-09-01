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
		"ok: 1": {recv: newPluralizer(), arg: "User", want: "Users"},
		"ok: 2": {recv: newPluralizer(), arg: "UserID", want: "UserIDs"},
		"ok: 3": {recv: newPluralizer(), arg: "column1", want: "column1s"},
		// "ok: ": {recv: newPluralizer(), arg: "History", want: "Histories"},
	}
	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.recv.pluralize(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
