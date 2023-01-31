package main

import (
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

	YouTubeLink := os.Args[1]
	//flag.String("YouTubeLink", "", "URL which thumbnail you want to download")
	async := flag.Bool("async", false, "flags for the program to run asynchronously")
	help := flag.Bool("help", false, "help")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usageMessage)
	}
	flag.Parse()
	if *help {
		fmt.Printf(helpMessage)
	}
	fmt.Println(async)
	fmt.Println(YouTubeLink)

	downloadThumbnailRes, err := c.DownloadThumbnail(context.Background(), &grpcYoutubeThumbnails.DownloadThumbnailLinkRequest{URL: YouTubeLink})

	//resp, err := c.DownloadThumbnail(context.Background(), &grpcYoutubeThumbnails.DownloadThumbnailLinkRequest{URL: link})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println("Downloading video thumbnail...")
	fmt.Println(downloadThumbnailRes.Response)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

const usageMessage = `Please try to input YouTubeLink as a parameter:
go run client/client.go [https://www.youtube.com/yourVideoID]
or use flag -h`

const helpMessage = `Youtube Thumbnails Downloader is a CLI tool to download YouTube thumbnails by video URLs.

usage: client/client.go [flags] URLs
  options:
    -h, --help     Does something
    -a, --async    Does something with "required"
`
