package main

var definitions = map[string]Rule{
	"lowercase": new(RuleLowercase).Init(),
	"regex":     new(RuleRegex).Init(),

	"camelcase":          new(RuleCamelCase).Init(),
	"pascalcase":         new(RulePascalCase).Init(),
	"snakecase":          new(RuleSnakeCase).Init(),
	"screamingsnakecase": new(RuleScreamingSnakeCase).Init(),
	"kebabcase":          new(RuleKebabCase).Init(),
	"pointcase":          new(RulePointCase).Init(),
}

var rules = map[string]Rule{
	"lowercase": definitions["lowercase"],
	"regex":     definitions["regex"],

	"camelcase": definitions["camelcase"],
	"camelCase": definitions["camelcase"],

	"pascalcase": definitions["pascalcase"],
	"PascalCase": definitions["pascalcase"],

	"snakecase":  definitions["snakecase"],
	"snake_case": definitions["snakecase"],

	"screamingsnakecase":   definitions["screamingsnakecase"],
	"SCREAMING_SNAKE_CASE": definitions["screamingsnakecase"],

	"kebabcase":  definitions["kebabcase"],
	"kebab-case": definitions["kebabcase"],

	"pointcase":  definitions["pointcase"],
	"point.case": definitions["pointcase"],
}

type Rule interface {
	Init() Rule
	GetName() string
	SetParameters(params []string) error
	Validate(value string) (bool, error)
	GetErrorMessage() string
}
