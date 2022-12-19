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
func (t Timecode) ToFrames(framerate int) int {
	return t.Hours*3600*framerate + t.Minutes*60*framerate + t.Seconds*framerate + t.Frames
}

// TimecodeFromFrames returns a timecode from the given number of frames.
func TimecodeFromFrames(frames int, framerate int) Timecode {
	hours := frames / (3600 * framerate)
	frames -= hours * 3600 * framerate
	minutes := frames / (60 * framerate)
	frames -= minutes * 60 * framerate
	seconds := frames / framerate
	frames -= seconds * framerate
	return Timecode{hours, minutes, seconds, frames}
}

// ToDuration returns timecode time.Duration representation.
func (t Timecode) ToDuration(framerate int) time.Duration {
	return time.Duration(t.ToFrames(framerate)) * time.Second / time.Duration(framerate)
}

// TimecodeFromDuration returns a timecode from the given time.Duration.
func TimecodeFromDuration(duration time.Duration, framerate int) Timecode {
	return TimecodeFromFrames(int(duration*time.Duration(framerate)/time.Second), framerate)
}

// Correct corrects the timecode to make sure that the values are within the
// valid ranges. For example, if the timecode is 00:00:00:30 with a framerate
// of 25, the timecode will be corrected to 00:00:01:05.
func (t *Timecode) Correct(framerate int) {
	frames := t.ToFrames(framerate)
	t.Hours = frames / (3600 * framerate)
	frames -= t.Hours * 3600 * framerate
	t.Minutes = frames / (60 * framerate)
	frames -= t.Minutes * 60 * framerate
	t.Seconds = frames / framerate
	frames -= t.Seconds * framerate
	t.Frames = frames
}
