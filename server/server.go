package main

import (
	"context"
	"fmt"
	grpcYoutubeThumbnails "github.com/SubochevaValeriya/grpcYoutubeThumbnails/proto"
	"github.com/SubochevaValeriya/grpcYoutubeThumbnails/server/internal"
	"github.com/SubochevaValeriya/grpcYoutubeThumbnails/server/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.Println("Reading configs")

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	dbConfig := repository.MongoConfig{
		Host:            viper.GetString("db.host"),
		Port:            viper.GetString("db.port"),
		DefaultDatabase: viper.GetString("db.default_database"),
		Collection:      viper.GetString("db.collection"),
	}

	logrus.Println("Connecting to MongoDB")

	client, err := repository.MongoConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	logrus.Println("Starting Service...")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	grpcYoutubeThumbnails.RegisterYoutubeThumbnailsServiceServer(s, &server{})
	reflection.Register(s)

	logrus.Println("Service started.")
	go func() {
		logrus.Println("Service is waiting for requests...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)

	<-ch
	// Closing connection with DB and stopping the server
	fmt.Println()
	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on closing connection with MongoDB : %v", err)
	}
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}

type server struct {
	grpcYoutubeThumbnails.YoutubeThumbnailsServiceServer
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (s *server) DownloadThumbnail(ctx context.Context, req *grpcYoutubeThumbnails.DownloadThumbnailLinkRequest) (*grpcYoutubeThumbnails.DownloadThumbnailLinkResponse, error) {
	video := internal.VideoItem{}
	err := video.FindVideoID(req.URL)
	if err != nil {
		return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: "Please try to use different URL"}, err
	}

	data, err := repository.FindVideoByID(ctx, video.VideoID)
	if err == nil {
		logrus.Printf("Found in the cash")
	} else {
		logrus.Printf("Not found in the cash")
		data = &video
		data.FindTitle(req.URL)

		if data.FindThumbnailLink() != nil {
			return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: fmt.Sprintf("Can't find thumbnail link")}, err
		}

		err = repository.CreateVideoItem(ctx, data)

		if err != nil {
			logrus.Printf("Can't add data to DB: %v", err)
		}
	}

	log.Println("Download thumbnail")

	if internal.CreateFolder() != nil {
		return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: fmt.Sprintf("Can't create directory for thumbnails")}, err
	}

	response, err := data.GetImage()

	if err != nil {
		return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: fmt.Sprintf("Can't get image by URL: %s", data.ThumbnailLink)}, err
	}

	if internal.SaveThumbnail(data.Name, response) != nil {
		return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: fmt.Sprintf("Can't download image: %s", data.ThumbnailLink)}, err
	}

	return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: "Downloaded successfully"}, nil
}
