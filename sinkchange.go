package main

import (
	"time"
)

type sinkchange struct {
	statusLoop

	joeryzenDefaultSink    string
	michaelPhoneExpiration time.Time
	leaPhoneExpiration     time.Time

	previous string
}

func (l *sinkchange) ProcessEvent(ev MQTTEvent) []MQTTPublish {
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

	now := ev.Timestamp
	weekday, hour, minute := now.Weekday(), now.Hour(), now.Minute()
	_, _, _ = weekday, hour, minute // TODO
	phoneHome := l.michaelPhoneExpiration.After(now) ||
		l.leaPhoneExpiration.After(now)
	_ = phoneHome
	l.statusf("joeryzenDefaultSink=%s", l.joeryzenDefaultSink)
	payload := "OFF"
	if l.joeryzenDefaultSink == "alsa_output.usb-Yamaha_Corporation_Steinberg_UR22C-00.analog-stereo" {
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
