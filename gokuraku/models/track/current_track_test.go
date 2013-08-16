package track

import (
	"github.com/ToQoz/Gokuraku/gokuraku/test_helpers"
	"strconv"
	"testing"
	"time"
)

func TestCurrentTrackValidate(t *testing.T) {
	var err error
	now := strconv.FormatInt(time.Now().Unix(), 10)

	currentTrack := CurrentTrack{
		TrackId:   "",
		StartedAt: now,
	}

	err = currentTrack.Validate()
	if err == nil {
		t.Error("CurrentTrack#Validate() should return error if CurrentTrack#TrackId is empty")
		return
	}

	test_helpers.AssertEqual(t, "Current track should have track ID", err.Error())
}
