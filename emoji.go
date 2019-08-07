package emoji

import (
	"math/rand"
	"time"
)

// Emotion - An emotion in string value.
type Emotion string

const (
	Happy   Emotion = "happy"
	Neutral Emotion = "neutral"
	Sad     Emotion = "sad"
)

// Emoji - An emoji with emotion type, sentiment score and bytes.
type Emoji struct {
	emotion Emotion
	score   int
	bytes   []byte
}

// GetEmotion - REturn the emotion of this emoji.
func (e *Emoji) GetEmotion() Emotion {
	return e.emotion
}

// GetScore - Return the sentiment score of this emoji.
func (e *Emoji) GetScore() int {
	return e.score
}

// GetBytes - Return the byte representation of this emoji.
func (e *Emoji) GetBytes() []byte {
	return e.bytes
}

// GetString - Return a string representation of this emoji.
func (e *Emoji) GetString() string {
	return string(e.bytes)
}

var (
	emoji = []*Emoji{
		&Emoji{
			Happy,
			1,
			[]byte("😄"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😃"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😀"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😄"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😁"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😆"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😅"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("🤣"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("🤠"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("🙂"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😊"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😇"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("🥰"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😍"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("🤩"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😘"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😜"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("🤪"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😝"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("😎"),
		},
		&Emoji{
			Happy,
			1,
			[]byte("🥳"),
		},
		&Emoji{
			Neutral,
			0,
			[]byte("😐"),
		},
		&Emoji{
			Neutral,
			0,
			[]byte("😑"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("☹️"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😐"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😑"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😬"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😔"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("🤒"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("🤕"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("🤮"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("🥴"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😵"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("🤯"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("🧐"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😕"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("🥺"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😭"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😖"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😞"),
		},
		&Emoji{
			Sad,
			-1,
			[]byte("😫"),
		},
	}
)

// GetEmoji - Get an emoji based on emotion and sentiment score.
func GetEmoji(emotion Emotion, score int) *Emoji {
	// TODO: add support for sentiment score
	for _, emoji := range emoji {
		if emoji.emotion == emotion {
			return emoji
		}
	}
	panic("That emotion and/or score are not yet supported.")
}

// GetRandomEmoji - Get a random emoji based on emotion.
func GetRandomEmoji(emotion Emotion) *Emoji {
	possible := []*Emoji{}
	for _, emoji := range emoji {
		if emoji.emotion == emotion {
			possible = append(possible, emoji)
		}
	}
	if len(possible) == 0 {
		panic("That emotion and/or score are not yet supported.")
	}
	rand.Seed(time.Now().Unix())
	return possible[rand.Intn(len(possible))]
}

// GetScore - Return a sentiment score from a given string.
func GetScore(input string) int {
	return 0
}
