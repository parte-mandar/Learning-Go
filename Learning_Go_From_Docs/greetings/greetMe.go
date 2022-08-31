// Created another file to test the impact

package greetings

func Namaskaram(name string) string {
    return "Namaskaram, " + name + "!!"
}

func AddingNamaskaram(name string, lastname string) string {
	return "namaskaram() - " + namaskaram(name) + " " + lastname
	// This is accessing namaskaram() which is in greet.go file in same module
}