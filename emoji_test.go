package emoji

import (
	"testing"
)

func TestGetEmoji(t *testing.T) {
	emoji := GetEmoji(Sad, 0)
	if emoji.String() != "☹️" {
		t.Fail()
	}
}

func TestGetRandomEmoji(t *testing.T) {
	emoji := GetRandomEmoji(Happy)
	if emoji.Emotion() != "happy" {
		t.Fail()
	}
}
