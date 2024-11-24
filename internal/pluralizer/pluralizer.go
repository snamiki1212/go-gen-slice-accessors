package pluralizer

import "regexp"

type Pluralizer struct {
	rules pluralizeRules
}

// REF: https://github.com/gertd/go-pluralize/blob/master/pluralize.go#L317
type (
	pluralizeRule struct {
		expression  *regexp.Regexp
		replacement string
	}
	pluralizeRules []pluralizeRule
	strRule        struct {
		expr string
		rep  string
	}
	strRules []strRule
)

func (sc strRules) toPluralizeRules() []pluralizeRule {
	rules := make([]pluralizeRule, 0, len(sc))
	for _, v := range sc {
		rules = append(rules, v.toPluralizeRule())
	}
	return rules
}

func (s strRule) toPluralizeRule() pluralizeRule {
	return pluralizeRule{newRegexpRule(s.expr), s.rep}
}

func (p Pluralizer) replace(word string, rules []pluralizeRule) string {
	// reverse order
	for i := len(rules) - 1; i >= 0; i-- {
		if rules[i].expression.MatchString(word) {
			return p.doReplace(word, rules[i])
		}
	}
	return word
}

func (p Pluralizer) doReplace(word string, rule pluralizeRule) string {
	return rule.expression.ReplaceAllString(word, rule.replacement)
}

// NOTE: (?i) is case insentive flag(https://stackoverflow.com/questions/15326421/how-do-i-do-a-case-insensitive-regular-expression-in-go)
// NOTE: https://github.com/gertd/go-pluralize/blob/master/pluralize.go#L317
func newDefaultStrRules() strRules {
	return strRules{
		{`(?i).*`, `${0}s`},
		{`(?:fe|f)$`, `${1}ves`},
		{`([^aiueo])y$`, `${1}ies`},
		{`([^aiueo])o$`, `${1}oes`},
		{`([s|z|sh|ch|x])$`, `${1}es`},
	}
}

func isExpr(s string) bool {
	return s[:1] == `(`
}

func newRegexpRule(rule string) *regexp.Regexp {
	rl := func() string {
		if isExpr(rule) {
			return rule
		}
		return `(?i)^` + rule + `$`
	}()
	return regexp.MustCompile(rl)
}

func NewPluralizer(rulesList ...strRules) Pluralizer {
	rules := make(strRules, 0)
	if len(rulesList) > 0 {
		rs := rulesList[0] // only get the first elements
		rules = append(rules, rs...)
	}
	rules = append(rules, newDefaultStrRules()...)
	return Pluralizer{rules.toPluralizeRules()}
}

func (p Pluralizer) Pluralize(str string) string {
	return p.replace(str, p.rules)
}
