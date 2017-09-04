// Package jterm is a binding of gopherjs to jquery.terminal. This lets you use jquery.terminal without writing
// any JavaScript. There is the examples directory to show some simple use cases.
// In some cases jqeury.terminal has properties and methods that have the same name. Go doesn't
// allow this to happen. In these cases the method name may have been changed slightly in this package.
package jterm

import (
	"strconv"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

// CallTerminal will call the jQuery.terminal method, which is the entry point for the library.
func CallTerminal(j, t Terminal, post func(c, t string) func() string) Terminal {
	j.jq.Call("terminal", post, t)
	return t
}

// NewTerminal will return a new Terminal struct. It is a "superclass" of jquery.JQuery
func NewTerminal(args string) Terminal {
	return Terminal{jq: jquery.NewJQuery(args)}
}

// DefaultTerminal returns a Terminal struct with all of the default values pre-populated.
func DefaultTerminal() Terminal {
	obj := js.Global.Get("$").Get("terminal").Get("defaults")
	t := &Terminal{o: obj}
	return *t
}

// Terminal is a struct with fields that correspond to jquery.terminal options.
type Terminal struct {
	o                  *js.Object
	jq                 jquery.JQuery
	Greetings          string      `js:"greetings"`
	History            bool        `js:"history"`
	Exit               bool        `js:"exit"`
	Prompt             interface{} `js:"prompt"`
	Name               string      `js:"name"`
	ScrollBottomOffset int         `js:"scollBottomOffse"`
	ScrollOnEcho       bool        `js:"scrollOnEcho"`
	EchoCommand        bool        `js:"echoCommand"`
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

// Error displays the input in red.
func (t Terminal) Error(e interface{}) {
	t.jq.Call("error", e)
}

// Disable disables the terminal.
func (t Terminal) Disable() {
	t.jq.Call("disable")
}

// Flush will echo all stored content to the terminal at once, if the `flush`` property is set to `false`.
func (t Terminal) flush() {
	t.jq.Call("flush")
}

// GetCommand will return the text for the current command.
func (t Terminal) GetCommand() string {
	return t.jq.Call("get_command").Val()
}

// GetPrompt willl return the current prompt.
func (t Terminal) GetPrompt() string {
	return t.jq.Call("get_prompt").Val()
}

// Insert will insert text intot he current cursor position.
func (t Terminal) Insert(s string) {
	t.jq.Call("insert", s)
}

// IsBottom returns true if the terminal is scrolled to the bottom.
func (t Terminal) IsBottom() bool {
	return t.jq.Call("is_bottom").Is()
}

// Level will return how deeply nested in the interpreters you currently are.
func (t Terminal) Level() int {
	i, _ := strconv.Atoi(t.jq.Call("level").Val())
	return i
}

// LoginName will reutrn the login name which was used in authetntication.
func (t Terminal) LoginName() string {
	return t.jq.Call("login_name").Val()
}

// Logout will log you out from the current terminal.
func (t Terminal) Logout() {
	t.jq.Call("logout")
}

// GetName will return the name of the current terminal. Renamed from Name() to GetName().
func (t Terminal) GetName() string {
	return t.jq.Call("name").Val()
}

// Paused will return true if the terminal is paused.
func (t Terminal) Paused() bool {
	return t.jq.Call("paused").Is()
}

// Pop will remove the current interpreter from the stack and run ht eprevious one.
func (t Terminal) Pop() {
	t.jq.Call("pop")
}

// Purge will remove all local storage left by the terminal.
func (t Terminal) Purge() {
	t.jq.Call("purge")
}

// Reset will reset the terminal
func (t Terminal) Reset() {
	// TODO this func does not fully impliment the jquery.terminal echo function
	t.jq.Call("reset")
}

// SetCommand will set the current command.
func (t Terminal) SetCommand(s string) {
	t.jq.Call("set_command", s)
}

// ScrollToBottom will scroll to the bottom of the terminal.
func (t Terminal) ScrollToBottom() {
	t.jq.Call("scroll_to_bottom")
}
