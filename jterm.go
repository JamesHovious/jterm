package jterm

import (
	"reflect"

	"unicode"
	"unicode/utf8"

	"github.com/gopherjs/jquery"
)

// thanks https://gist.github.com/derlin/0be53d0d7f38db181198aada024269b8
func structToLowerFirstMap(in Terminal) map[string]interface{} {
	v := reflect.ValueOf(in)
	vType := v.Type()

	result := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		name := vType.Field(i).Name
		if name != "jq" {
			result[lowerFirst(name)] = v.Field(i).Interface()
		}
	}
	return result
}

// thanks https://gist.github.com/derlin/0be53d0d7f38db181198aada024269b8
func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// NewTerminal will return a new Terminal struct. It is a "superclass" of jquery.JQuery
func NewTerminal(args string) Terminal {
	return Terminal{jq: jquery.NewJQuery(args)}
}

// CallTerminal will call the jQuery.terminal method, which is the entry point for the library.
func CallTerminal(j, t Terminal, post func(c, t string) func() string) Terminal {
	m := structToLowerFirstMap(t)
	j.jq.Call("terminal", post, m)
	return t
}

// Terminal is a struct with fields that correspond to jquery.terminal options.
type Terminal struct {
	jq        jquery.JQuery
	Greetings string
	Prompt    interface{}
	Name      string
	ScrollBottomOffset int
	ScrollOnEcho       bool
	EchoCommand        bool
}

// DefaultTerminal contains all of the default values for the Terminal
var DefaultTerminal = Terminal {
	Greetings: "",
	Prompt: "",
	Name: "jTerm",
	ScrollBottomOffset: 0,
	ScrollOnEcho: true,
	EchoCommand: true,
}

// Clear will clear the terminal.
func (t Terminal) Clear() {
	t.jq.Call("clear")
}

// ClearHistoryState will clear the history from the localstorage.
func (t Terminal) ClearHistoryState() {
	t.jq.Call("clear_history_state")
}

// Destroy removes everything created by the terminal. It will not destroy the
// content in the local storage. If you want to destroy the localstorage as well
// use Purge()
func (t Terminal) Destroy() {
	t.jq.Call("destroy")
}

// Echo will insert a string on a new line in the terminal.
func (t Terminal) Echo(s ...interface{}) { // TODO fully document and test.
	t.jq.Call("echo", s)
}

// Enable enables the terminal.
func (t Terminal) Enable() {
	t.jq.Call("enable")
}

// Disable disables the terminal.
func (t Terminal) Disable() {
	t.jq.Call("disable")
}

// IsBottom returns true if the terminal is scrolled to the bottom.
func (t Terminal) IsBottom() bool {
	return t.jq.Call("is_bottom").Is()
}

// Reset will reset the terminal
func (t Terminal) Reset() {
	// TODO this func does not fully impliment the jquery.terminal echo function
	t.jq.Call("reset")
}

// ScrollToBottom will scroll to the bottom of the terminal.
func (t Terminal) ScrollToBottom() {
	t.jq.Call("scroll_to_bottom")
}
