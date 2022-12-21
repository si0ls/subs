package stl

import (
	"fmt"
	"time"
)

// Timecode represents a temporal position.
type Timecode struct {
	Hours   int
	Minutes int
	Seconds int
	Frames  int
}

// String returns a string representation of the timecode.
func (t Timecode) String() string {
	return fmt.Sprintf("%02d:%02d:%02d:%02d", t.Hours, t.Minutes, t.Seconds, t.Frames)
}

// ToFrames returns the total number of frames.
func (t Timecode) ToFrames(framerate uint) int {
	return t.Hours*3600*int(framerate) + t.Minutes*60*int(framerate) + t.Seconds*int(framerate) + t.Frames
}

// TimecodeFromFrames returns a timecode from the given number of frames.
func TimecodeFromFrames(frames int, framerate uint) Timecode {
	hours := frames / (3600 * int(framerate))
	frames -= hours * 3600 * int(framerate)
	minutes := frames / (60 * int(framerate))
	frames -= minutes * 60 * int(framerate)
	seconds := frames / int(framerate)
	frames -= seconds * int(framerate)
	return Timecode{hours, minutes, seconds, frames}
}

// ToDuration returns timecode time.Duration representation.
func (t Timecode) ToDuration(framerate uint) time.Duration {
	return time.Duration(t.ToFrames(framerate)) * time.Second / time.Duration(framerate)
}

// TimecodeFromDuration returns a timecode from the given time.Duration.
func TimecodeFromDuration(duration time.Duration, framerate uint) Timecode {
	return TimecodeFromFrames(int(duration*time.Duration(framerate)/time.Second), framerate)
}

// Correct corrects the timecode to make sure that the values are within the
// valid ranges. For example, if the timecode is 00:00:00:30 with a framerate
// of 25, the timecode will be corrected to 00:00:01:05.
func (t *Timecode) Correct(framerate uint) {
	frames := t.ToFrames(framerate)
	t.Hours = frames / (3600 * int(framerate))
	frames -= t.Hours * 3600 * int(framerate)
	t.Minutes = frames / (60 * int(framerate))
	frames -= t.Minutes * 60 * int(framerate)
	t.Seconds = frames / int(framerate)
	frames -= t.Seconds * int(framerate)
	t.Frames = frames
}
