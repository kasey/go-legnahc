package changelog

import "testing"

func TestVersionLine(t *testing.T) {
	// Test the version line regex
	if !versionRE.MatchString("# v1.0.0") {
		t.Error("versionRE failed to match")
	}
}

func TestParseBulletOverride(t *testing.T) {
	line := "- added override [[PR]](https://github.com/prysmaticlabs/prysm/pull/1)"
	res := parseBullet(line, "NOPE")
	if line != res {
		t.Error("parseBullet did not recognize the override")
	}
}
