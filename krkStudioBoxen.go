package main

import "strings"

type krkStudioBoxen struct {
	statusLoop

	joeryzenDefaultSink string

	previous string
}

func (l *krkStudioBoxen) ProcessEvent(ev MQTTEvent) []MQTTPublish {
	// var lease struct {
	//	Expiration time.Time `json:"expiration"`
	// }
	switch ev.Topic {
	case "regelwerk/ticker/1s":
	// current time influences our state

	case "github.com/stapelberg/defaultsink2mqtt/default_sink":
		l.joeryzenDefaultSink = string(ev.Payload.([]byte))

	default:
		return nil // event did not influence our state
	}

	l.statusf("joeryzenDefaultSink=%s", l.joeryzenDefaultSink)
	payload := "OFF"

	if strings.Contains(l.joeryzenDefaultSink, "UR22C") {
		payload = "ON"
	}
	if l.previous == payload {
		return nil // skip, no change
	}
	l.previous = payload
	return []MQTTPublish{
		{
			Topic:    "cmnd/delock/Power",
			Payload:  payload,
			Retained: true,
		},
	}
}
