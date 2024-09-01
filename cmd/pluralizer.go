package cmd

type pluralizer struct{}

func (p pluralizer) pluralize(str string) string {
	if str[len(str)-1] == 's' {
		return str + "es"
	}
	return str + "s"
}
