package victor

import (
	"testing"
)

func TestRouting(t *testing.T) {
	robot := &robot{name: "ralph"}
	dispatch := NewDispatch(robot)

	called := 0

	dispatch.HandleFunc(robot.Direct("howdy"), func(s *State) {
		called++
	})
	dispatch.HandleFunc(robot.Direct("tell (him|me)"), func(s *State) {
		called++
	})
	dispatch.HandleFunc("alot", func(s *State) {
		called++
	})

	// Should trigger
	dispatch.ProcessMessage(&msg{text: "ralph howdy"})
	dispatch.ProcessMessage(&msg{text: "ralph tell him"})
	dispatch.ProcessMessage(&msg{text: "ralph tell me"})
	dispatch.ProcessMessage(&msg{text: "/tell me"})
	dispatch.ProcessMessage(&msg{text: "I heard alot of them."})

	if called != 5 {
		t.Errorf("One or more register actions weren't triggered")
	}
}

func TestParams(t *testing.T) {
	robot := &robot{name: "ralph"}
	dispatch := NewDispatch(robot)

	called := 0

	dispatch.HandleFunc(robot.Direct("yodel (it)"), func(s *State) {
		called++
		params := s.Params()
		if len(params) == 0 || params[0] != "it" {
			t.Errorf("Incorrect message params expected=%v got=%v", []string{"it"}, params)
		}
	})

	dispatch.ProcessMessage(&msg{text: "ralph yodel it"})

	if called != 1 {
		t.Error("Registered action was never triggered")
	}
}

func TestNonFiringRoutes(t *testing.T) {
	robot := &robot{name: "ralph"}
	dispatch := NewDispatch(robot)

	called := 0

	dispatch.HandleFunc(robot.Direct("howdy"), func(s *State) {
		called++
	})

	dispatch.ProcessMessage(&msg{text: "Tell ralph howdy."})

	if called > 0 {
		t.Error("Registered action was triggered when it shouldn't have been")
	}
}

type msg struct {
	userID      string
	userName    string
	channelID   string
	channelName string
	text        string
}

func (m *msg) UserID() string {
	return m.userID
}

func (m *msg) UserName() string {
	return m.userName
}

func (m *msg) ChannelID() string {
	return m.channelID
}

func (m *msg) ChannelName() string {
	return m.channelName
}

func (m *msg) Text() string {
	return m.text
}
