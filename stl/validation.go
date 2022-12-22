package stl

import (
	"fmt"
	"time"
)

func validateRange(value, min, max int, err error, fatal bool) error {
	if value < min || value > max {
		return validateErr(fmt.Errorf("%w: must be in range [%d;%d]", err, min, max), value, fatal)
	}
	return nil
}

func validateNotInRange(value, min, max int, err error, fatal bool) error {
	if value >= min && value <= max {
		return validateErr(fmt.Errorf("%w: must not be in range [%d;%d]", err, min, max), value, fatal)
	}
	return nil
}

func validateList[T comparable](value T, list []T, err error, fatal bool) error {
	for _, v := range list {
		if value == v {
			return nil
		}
	}
	return validateErr(fmt.Errorf("%w: must be one of %v", err, list), value, fatal)
}

func validateTimecode(tc Timecode, framerate uint, err error, fatal bool) error {
	if tcErr := tc.Validate(framerate); tcErr != nil {
		return validateErr(fmt.Errorf("%w: %s", err, tcErr), tc, fatal)
	}
	return nil
}

func validateTimecodeOrder(tc1, tc2 Timecode, framerate uint, err error, fatal bool) error {
	if tc1.ToDuration(framerate) > tc2.ToDuration(framerate) {
		return validateErr(fmt.Errorf("%w: %s > %s", err, tc1, tc2), tc1, fatal)
	}
	return nil
}

func validateTimecodeOrderStrict(tc1, tc2 Timecode, framerate uint, err error, fatal bool) error {
	if tc1.ToDuration(framerate) >= tc2.ToDuration(framerate) {
		return validateErr(fmt.Errorf("%w: %s > %s", err, tc1, tc2), tc1, fatal)
	}
	return nil
}

func validateNonEmptyString(s string, err error, fatal bool) error {
	if s == "" {
		return validateErr(err, s, fatal)
	}
	return nil
}

func validateDate(date time.Time, err error, fatal bool) error {
	if date.IsZero() {
		return validateErr(err, date, fatal)
	}
	return nil
}

func validateDateOrder(date1, date2 time.Time, err error, fatal bool) error {
	if date1.After(date2) {
		return validateErr(fmt.Errorf("%w: %s > %s", err, date1, date2), date1, fatal)
	}
	return nil
}
