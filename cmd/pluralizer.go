package cmd

type pluralizer struct{}

func newPluralizer() pluralizer {
	return pluralizer{}
}

func (p pluralizer) pluralize(str string) string {
	if str[len(str)-1] == 's' {
		return str + "es"
	}
	return str + "s"
}
