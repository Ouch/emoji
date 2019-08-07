package react

import (
	"testing"
)

func TestGetEmoji(t *testing.T) {
	emoji := GetEmoji(Sad, 0)
	if emoji.GetString() != "☹️" {
		t.Fail()
	}
}

func TestGetRandomEmoji(t *testing.T) {
	emoji := GetRandomEmoji(Happy)
	if emoji.GetEmotion() != "happy" {
		t.Fail()
	}
}
