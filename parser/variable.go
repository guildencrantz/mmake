package parser

import "regexp"

var (
	assignmentRe = regexp.MustCompile("(.+?)[:?+]?=(.+?)")
)

// Variable node.
type Variable struct {
	Name  string
	Value string
}
