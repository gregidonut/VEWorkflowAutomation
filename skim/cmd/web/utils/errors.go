package utils

import "errors"

var (
	generateTTSErr = errors.New("error generating tts")
	noFSVidsErr    = errors.New("there are no fsVids yet")
)
