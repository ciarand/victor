package victor

import (
    "math/rand"
)

type Message interface {
    Id() int
    Type() string
    UserId() int
    RoomId() int
    Body() string
}

type Context struct {
    message Message

    Send  func(text string)
    Reply func(text string)
    Paste func(text string)
    Sound func(text string)

    matches []string
}

func (c *Context) SetMessage(msg Message) *Context {
    c.message = msg

    return c
}

func (c *Context) Message() Message {
    return c.message
}

func (c *Context) SetMatches(matches []string) *Context {
    c.matches = matches

    return c
}

func (c *Context) Matches() []string {
    return c.matches
}

func (c *Context) RandomString(strings []string) string {
    return strings[rand.Intn(len(strings))]
}