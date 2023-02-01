package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCaseFindIndex struct {
	URL           string
	Expected      string
	ExpectedError bool
}

func TestFindIndex(t *testing.T) {
	video := VideoItem{
		ID:            "",
		Name:          "",
		VideoID:       "",
		ThumbnailLink: "",
	}

	testCases := []TestCaseFindIndex{
		{URL: "https://www.youtube.com/watch?v=6o1m7ofjCCY", Expected: "6o1m7ofjCCY", ExpectedError: false},
		{URL: "http://youtube.com/watch?vi=dQw4w9WgXcQ&feature=youtube_gdata_player", Expected: "dQw4w9WgXcQ", ExpectedError: false},
		{URL: "help", Expected: "", ExpectedError: true},
		{URL: "itsNotCorre", Expected: "", ExpectedError: true},
	}

	for _, cse := range testCases {
		video.VideoID = ""
		cse := cse
		t.Run(cse.URL, func(t *testing.T) {
			err := video.FindVideoID(cse.URL)
			assert.Equalf(t, cse.Expected, video.VideoID, "for %d expected %t, got %t", cse.URL, cse.Expected, video.VideoID)
			if cse.ExpectedError {
				if err == nil {
					t.Errorf("should be an error")
				}
			} else {
				if err != nil {
					t.Errorf("should be no error")
				}
			}
		})
	}
}
