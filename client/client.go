package main

import (
	"bufio"
	"flag"
	"fmt"
	grpcYoutubeThumbnails "github.com/SubochevaValeriya/grpcYoutubeThumbnails/proto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial(fmt.Sprintf("localhost:%s", viper.GetString("port")), opts)
	if err != nil {
		logrus.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() //

	c := grpcYoutubeThumbnails.NewYoutubeThumbnailsServiceClient(cc)

	// loading video thumbnail

	if len(os.Args) == 1 {
		fmt.Printf(helpMessage)
		return
	}

	async := flag.Bool("async", false, "flag for the program to run asynchronously")
	help := flag.Bool("help", false, "help")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usageMessage)
	}
	flag.Parse()

	if *help {
		fmt.Printf(helpMessage)
	}

	if os.Args[1] == "file" {
		file, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal("Can't open file")
		}
		//defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			sendRequest(c, scanner.Text())
		}
		return
	}

	for i := 1; i < len(os.Args); i++ {
		YouTubeLink := os.Args[i]
		sendRequest(c, YouTubeLink)

	}

	fmt.Println(async)

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func sendRequest(c grpcYoutubeThumbnails.YoutubeThumbnailsServiceClient, YouTubeLink string) {
	downloadThumbnailRes, err := c.DownloadThumbnail(context.Background(), &grpcYoutubeThumbnails.DownloadThumbnailLinkRequest{URL: YouTubeLink})

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println("Downloading video thumbnail...")
	fmt.Println(downloadThumbnailRes.Response)
}

const usageMessage = `Please try to input YouTubeLink as a parameter:
go run client/client.go [https://www.youtube.com/yourVideoID]
or use flag -help`

const helpMessage = `Youtube Thumbnails Downloader is a CLI tool to download YouTube thumbnails by video URLs.
You can input several URLs divided by backspaces.

usage: client/client.go [flags] URLs OR client/client.go file name.ext

usageExamples:
go run client/client.go [https://www.youtube.com/yourVideoID]
go run client/client.go file urls.txt

  commands:
	
file name.ext

  flags:
    --help     Show this help message
    --async    Flag for the program to run asynchronously
`
