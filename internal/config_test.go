package config

import (
	"math"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	os.Setenv("GOGEN_HOME", "..")
	c := NewConfig()
	c2 := NewConfig()
	assert.Equal(t, c, c2)
}

func TestGlobal(t *testing.T) {
	os.Setenv("GOGEN_HOME", "..")
	c := NewConfig()
	global := Global{Debug: false, Verbose: false, GeneratorWorkers: 1, OutputWorkers: 1, ROTInterval: 1}
	assert.Equal(t, c.Global, global)
}

func TestFlatten(t *testing.T) {
	// Setup environment
	os.Setenv("GOGEN_HOME", "..")
	os.Setenv("GOGEN_ALWAYS_REFRESH", "1")
	home := ".."
	rand.Seed(0)

	var s *Sample
	s = FindSampleInFile(home, "flatten")
	assert.Equal(t, false, s.Disabled)
	assert.Equal(t, "sample", s.Generator)
	assert.Equal(t, "stdout", s.Outputter)
	assert.Equal(t, "raw", s.OutputTemplate)
	assert.Equal(t, "config", s.Rater)
	assert.Equal(t, 0, s.Interval)
	assert.Equal(t, 0, s.Count)
	assert.Equal(t, "now", s.Earliest)
	assert.Equal(t, "now", s.Latest)
	if diff := math.Abs(float64(0.2 - s.RandomizeCount)); diff > 0.000001 {
		t.Fatalf("RandomizeCount not equal")
	}
	assert.Equal(t, true, s.RandomizeEvents)
	assert.Equal(t, 1, s.EndIntervals)
}

func TestValidate(t *testing.T) {
	// Setup environment
	os.Setenv("GOGEN_HOME", "..")
	os.Setenv("GOGEN_ALWAYS_REFRESH", "1")
	home := ".."
	rand.Seed(0)
	// loc, _ := time.LoadLocation("Local")
	// n := time.Date(2001, 10, 20, 12, 0, 0, 100000, loc)
	// now := func() time.Time {
	// 	return n
	// }

	var s *Sample
	s = FindSampleInFile(home, "validate-lower-upper")
	assert.Equal(t, true, s.Disabled)

	s = FindSampleInFile(home, "validate-upper")
	assert.Equal(t, true, s.Disabled)

	s = FindSampleInFile(home, "validate-string-length")
	assert.Equal(t, true, s.Disabled)

	s = FindSampleInFile(home, "validate-choice-items")
	assert.Equal(t, true, s.Disabled)

	s = FindSampleInFile(home, "validate-weightedchoice-items")
	assert.Equal(t, true, s.Disabled)

	s = FindSampleInFile(home, "validate-fieldchoice-items")
	assert.Equal(t, true, s.Disabled)

	s = FindSampleInFile(home, "validate-fieldchoice-badfield")
	assert.Equal(t, true, s.Disabled)

	s = FindSampleInFile(home, "validate-badrandom")
	assert.Equal(t, true, s.Disabled)

	s = FindSampleInFile(home, "validate-earliest-latest")
	assert.Equal(t, true, s.Disabled)
}
