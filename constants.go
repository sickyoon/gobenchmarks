package main

const (
	chello  = "hello"
	cspace  = " "
	cworld  = "world"
	canswer = 42
)

var (
	hello  string
	space  string
	world  string
	answer int
)

// variables behave differently than constants
func init() {
	hello = "hello"
	space = " "
	world = "world"
	answer = 42
}
