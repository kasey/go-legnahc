package main

import "testing"

func TestVersionLine(t *testing.T) {
	// Test the version line regex
	if !versionRE.MatchString("# v1.0.0") {
		t.Error("versionRE failed to match")
	}
}
