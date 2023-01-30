package server

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type TestCaseFindIndex struct {
	URL      string
	Expected string
}

func TestFindIndex(t *testing.T) {
	video := videoItem{
		ID:            primitive.ObjectID{},
		Name:          "",
		VideoID:       "",
		ThumbnailLink: "",
	}

	testCases := []TestCaseFindIndex{
		{URL: "https://www.youtube.com/watch?v=6o1m7ofjCCY", Expected: "6o1m7ofjCCY"},
		{URL: "http://youtube.com/watch?vi=dQw4w9WgXcQ&feature=youtube_gdata_player", Expected: "dQw4w9WgXcQ"},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.URL, func(t *testing.T) {
			video.findIndex(cse.URL)
			assert.Equalf(t, cse.Expected, video.VideoID, "for %d expected %t, got %t", cse.URL, cse.Expected, video.VideoID)
		})
	}
}
