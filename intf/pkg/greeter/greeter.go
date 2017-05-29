package greeter

import "fmt"

// Greeter implements Greet()
type Greeter interface {
	Greet() string
}

type lobby struct {
	text string
}

func (l lobby) Greet() string {
	return fmt.Sprintf("Greeting with %q", l.text)
}

// New returns a new Greeter
func New(s string) Greeter {
	if s == "" {
		s = "You did not specify me :("
	}
	return &lobby{
		text: s,
	}
}
