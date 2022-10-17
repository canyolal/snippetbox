package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	// Initialize a new time.Time obj. and pass it to humanDate() func.
	tm := time.Date(2022, 10, 17, 19, 45, 0, 0, time.UTC)
	hd := humanDate(tm)

	// Check that the output from the humanDate function is in the format we
	// expect. If it isn't what we expect, use the t.Errorf() function to
	// indicate that the test has failed and log the expected and actual
	// values.
	if hd != "17 Oct 2022 at 19:45" {
		t.Errorf("got %q; want %q", hd, "17 Oct 2022 at 19:45")
	}
}
