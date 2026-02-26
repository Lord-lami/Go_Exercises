package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

// Greeter interface type is the type that must be made to allow 
// the robot greet the airport customers in their language
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

type Italian struct{}

func (italian Italian) LanguageName() string {return "Italian"}

func (italian Italian) Greet(name string) string {return "Ciao " + name + "!"}

type Portuguese struct{}

func (portuguese Portuguese) LanguageName() string {return "Portuguese"}

func (portuguese Portuguese) Greet(name string) string {return "Olá " + name + "!"}
