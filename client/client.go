package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/SubochevaValeriya/grpcYoutubeThumbnails/client/internal"
	grpcYoutubeThumbnails "github.com/SubochevaValeriya/grpcYoutubeThumbnails/proto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"sync"
)

func main() {

	cc, client := startingServer()
	defer cc.Close()

	async, args, err := flagsAndCommands()

	if err != nil {
		return
	}

	internal.CreateFolder()
	// downloading video thumbnail

	wg := &sync.WaitGroup{}
	downloadingThumbnailRequest(client, async, args, wg)
	wg.Wait()
}

// initialization of config
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// starting the server
func startingServer() (*grpc.ClientConn, grpcYoutubeThumbnails.YoutubeThumbnailsServiceClient) {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial(fmt.Sprintf("localhost:%s", viper.GetString("port")), opts)
	if err != nil {
		logrus.Fatalf("could not connect: %v", err)
	}

	return cc, grpcYoutubeThumbnails.NewYoutubeThumbnailsServiceClient(cc)
}

// parsing flags, commands and parameters
func flagsAndCommands() (*bool, []string, error) {
	async := flag.Bool("async", false, "flag for the program to run asynchronously")
	help := flag.Bool("help", false, "help message")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usageMessage)
	}

	flag.Parse()

	if len(os.Args) == 1 {
		write("%s", helpMessage)
		return async, os.Args, errors.New("parameter not found")
	}

	if *help {
		write("%s", helpMessage)
		return async, os.Args, errors.New("help message only")
	}

	return async, os.Args, nil
}

// making request
func downloadingThumbnailRequest(client grpcYoutubeThumbnails.YoutubeThumbnailsServiceClient, async *bool, args []string, wg *sync.WaitGroup) {
	var argI = 1
	if *async {
		argI = 2
		write("%s\n", "-- Working in Async mode --")
	}

	// URLs from file (command "file")

	if args[argI] == "file" {
		file, err := os.Open(args[argI+1])
		if err != nil {
			log.Fatal("Can't open file")
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			YouTubeLink := scanner.Text()
			asyncWork(client, async, YouTubeLink, wg)
		}
		return
	} else { // URLs from parameters
		for i := argI; i < len(os.Args); i++ {
			YouTubeLink := args[i]
			asyncWork(client, async, YouTubeLink, wg)
		}
	}
}

// writing message to console or other destination
func write(format string, message string) {
	dst := os.Stdout // can be changed also using config variables
	fmt.Fprintf(dst, format, message)
}

// determination in which mode to start downloading: async or not
func asyncWork(client grpcYoutubeThumbnails.YoutubeThumbnailsServiceClient, async *bool, YouTubeLink string, wg *sync.WaitGroup) {
	if *async {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendRequest(client, YouTubeLink)
		}()
	} else {
		sendRequest(client, YouTubeLink)
	}
}

// sending request to the server
func sendRequest(c grpcYoutubeThumbnails.YoutubeThumbnailsServiceClient, YouTubeLink string) {
	downloadThumbnailRes, err := c.DownloadThumbnail(context.Background(), &grpcYoutubeThumbnails.DownloadThumbnailLinkRequest{URL: YouTubeLink})

	write("%s: ", YouTubeLink)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		//testChanErr <- err
		return
	}

	err = internal.SaveThumbnail(downloadThumbnailRes.Response.Name, downloadThumbnailRes.Response.Image)
	if err != nil {
		fmt.Printf("%v", err)
	}
	write("%s\n", "Downloaded Successfully")
}

const usageMessage = `Please try to input YouTubeLink as a parameter:
go run client/client.go [https://www.youtube.com/yourVideoID]
or use flag --help`

const helpMessage = `
YOUTUBE THUMBNAILS DOWNLOADER 

Description:
Youtube Thumbnails Downloader is a CLI tool to download YouTube thumbnails by video URLs.
You can input several URLs divided by backspaces or load links from file.

  Usage: 
client/client.go [flags] URLs 
client/client.go [flags] file name.ext

  Usage Examples:
go run client/client.go https://www.youtube.com/yourVideoID
go run client/client.go file urls.txt
go run client/client.go --async file urls.txt

  Commands:
	
file name.ext

  Flags:
    --help     Show this help message
    --async    Flag for the program to run asynchronously
`
