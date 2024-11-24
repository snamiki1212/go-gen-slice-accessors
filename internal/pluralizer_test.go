package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pluralizer_Pluralize(t *testing.T) {
	tests := map[string]struct {
		recv Pluralizer
		arg  string
		want string
	}{
		"ok: common: 1":  {recv: NewPluralizer(), arg: "User", want: "Users"},
		"ok: common: 2":  {recv: NewPluralizer(), arg: "UserID", want: "UserIDs"},
		"ok: common: 3":  {recv: NewPluralizer(), arg: "column1", want: "column1s"},
		"ok: f or fe: 1": {recv: NewPluralizer(), arg: "knife", want: "knives"},
		"ok: f or fe: 2": {recv: NewPluralizer(), arg: "leaf", want: "leaves"},
		"ok: x + y: 1":   {recv: NewPluralizer(), arg: "city", want: "cities"},
		"ok: x + y: 2":   {recv: NewPluralizer(), arg: "History", want: "Histories"},
		"ok: x + y: 3":   {recv: NewPluralizer(), arg: "HistorY", want: "HistorYs"},
		"ok: x + o: 1":   {recv: NewPluralizer(), arg: "hero", want: "heroes"},
		"ok: es: 1":      {recv: NewPluralizer(), arg: "kiss", want: "kisses"},
		"ok: es: 2":      {recv: NewPluralizer(), arg: "bus", want: "buses"},
		"ok: es: 3":      {recv: NewPluralizer(), arg: "box", want: "boxes"},
	}
	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			got := tt.recv.Pluralize(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
