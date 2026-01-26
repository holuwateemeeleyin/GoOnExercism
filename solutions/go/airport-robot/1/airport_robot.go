package airportrobot

// Greeter defines the contract for all language greeters
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// SayHello generates a greeting using any Greeter implementation
func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

// --- Language implementations ---

type German struct{}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(name string) string {
	return "Hallo " + name + "!"
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}
