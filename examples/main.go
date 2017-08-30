package main

import (
	"fmt"
	"time"

	"github.com/JamesHovious/jterm"
)

func main() {
	defaultTerminal()
	customTerminal()
}

func postCallbackFunction(command, term string) func() string {
	print("[-] Hello, I print out to the console. ")
	print("[-] I get called everytime a new command is entered. ")
	print("[-] The command was: " + command)
	return func() string {
		return "" // TODO is this right? returning something just to make the funciton shut up.
	}
}

// An minimal example
func defaultTerminal() {
	// Create an terminal with default properties.
	jt := jterm.DefaultTerminal()
	// Define the callback function
	post := postCallbackFunction
	// Locate the div with the id of "terminal"
	JQueryTerminal := jterm.NewTerminal("#terminal")
	// Finally put it all together. Create the terminal.
	jterm.CallTerminal(JQueryTerminal, jt, post)
}

// An example terminal with some custom properties.
func customTerminal() {
	// Building an example custom prompt
	t := time.Now().UTC()
	// Using jquery.terminal formatting as defined at http://terminal.jcubic.pl/api_reference.php#echo
	p := fmt.Sprintf("[[b;#FFFF00;#000][%v\\]:] ", t.Format("2006-01-02 15:04"))

	// Create an terminal with default properties.
	jt := jterm.DefaultTerminal()
	// Populate the terminal with custom properties.
	jt.Prompt = p
	jt.Name = "MyCustomTerminal"
	jt.Greetings = "Custom Gopherjs Terminal banner!"

	// Define the callback function
	post := postCallbackFunction

	// Locate the div with the id of "terminal"
	JQueryTerminal := jterm.NewTerminal("#custom-terminal")

	// Finally put it all together. Create the terminal.
	jterm.CallTerminal(JQueryTerminal, jt, post)

	// Call methods on the same terminal, as needed.
	JQueryTerminal.Echo("Example echo output")

	// Note, right now methods that return values don't actually return a value...
	tPrompt := JQueryTerminal.GetName()
	JQueryTerminal.Echo("The current prompt is : " + tPrompt)

}
