package main

import (
	"fmt"
	"log"
	"os"

	grpcYoutubeThumbnails "github.com/SubochevaValeriya/grpcYoutubeThumbnails/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const defaultPort = "4041"

func main() {

	fmt.Println("Load Thumbnail Client")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	opts := grpc.WithInsecure() // change

	cc, err := grpc.Dial("localhost:4041", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() //

	c := grpcYoutubeThumbnails.NewYoutubeThumbnailsServiceClient(cc)

	// loading video thumbnail

	fmt.Printf("| Please input video link:")
	var link string
	fmt.Scanf("%s", &link)

	fmt.Printf("Loading video thumbnail...")

	video := &grpcYoutubeThumbnails.Video{
		Id:            "",
		Link:          link,
		ThumbnailLink: "",
	}

}
