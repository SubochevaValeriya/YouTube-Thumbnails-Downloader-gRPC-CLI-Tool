package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestSaveThumbnail(t *testing.T) {
	resp, err := http.Get("https://i.ytimg.com/vi/6o1m7ofjCCY/hqdefault.jpg")
	if err != nil {
		t.Fatal()
	}

	image, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = SaveThumbnail("one", image)
	if err != nil {
		t.Errorf("file not found")
	}

	if _, err := os.Stat("one.jpg"); os.IsNotExist(err) {
		fmt.Println(err)
		t.Errorf("file not found")
	}

	os.RemoveAll("one.jpg")
}
