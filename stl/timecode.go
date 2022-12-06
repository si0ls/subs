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

// ToTime returns timecode time.Time representation.
func (t Timecode) ToTime(framerate int) time.Time {
	return time.Unix(0, t.ToDuration(framerate).Nanoseconds())
}

// TimecodeFromTime returns a timecode from the given time.Time.
func TimecodeFromTime(t time.Time, framerate int) Timecode {
	return TimecodeFromDuration(t.Sub(time.Unix(0, 0)), framerate)
}
