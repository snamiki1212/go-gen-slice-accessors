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
		"ok: common: 1":  {recv: newPluralizer(), arg: "User", want: "Users"},
		"ok: common: 2":  {recv: newPluralizer(), arg: "UserID", want: "UserIDs"},
		"ok: common: 3":  {recv: newPluralizer(), arg: "column1", want: "column1s"},
		"ok: f or fe: 1": {recv: newPluralizer(), arg: "knife", want: "knives"},
		"ok: f or fe: 2": {recv: newPluralizer(), arg: "leaf", want: "leaves"},
		"ok: x + y: 1":   {recv: newPluralizer(), arg: "city", want: "cities"},
		"ok: x + y: 2":   {recv: newPluralizer(), arg: "History", want: "Histories"},
		"ok: x + o: 1":   {recv: newPluralizer(), arg: "hero", want: "heroes"},
		"ok: es: 1":      {recv: newPluralizer(), arg: "kiss", want: "kisses"},
		"ok: es: 2":      {recv: newPluralizer(), arg: "bus", want: "buses"},
		"ok: es: 3":      {recv: newPluralizer(), arg: "box", want: "boxes"},
	}
	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.recv.pluralize(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
