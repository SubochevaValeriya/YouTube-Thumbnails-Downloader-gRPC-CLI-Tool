package main

import (
	"context"
	"fmt"
	grpcYoutubeThumbnails "github.com/SubochevaValeriya/grpcYoutubeThumbnails/proto"
	"github.com/SubochevaValeriya/grpcYoutubeThumbnails/server/internal"
	"github.com/SubochevaValeriya/grpcYoutubeThumbnails/server/internal/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
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

	// environmental variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	dbConfig := repository.MongoConfig{
		Host:            os.Getenv("host"),
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

	// removing data from cache
	if viper.GetString("drop_cache") == "yes" {
		err := repository.DropCollection(context.Background())
		if err != nil {
			logrus.Errorf("cache data not deleted: %s", err)
		} else {
			logrus.Printf("Сache cleared successfully.")
		}
	}

	logrus.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on closing connection with MongoDB : %v", err)
	}
	logrus.Println("Stopping the server")
	s.Stop()
	logrus.Println("End of Program")
}

type server struct {
	grpcYoutubeThumbnails.YoutubeThumbnailsServiceServer
}

// initialization configs for app
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// DownloadThumbnail function downloads thumbnails in the specific directory. If video data already is in the cache than uses thumbnail link from cache.
func (s *server) DownloadThumbnail(ctx context.Context, req *grpcYoutubeThumbnails.DownloadThumbnailLinkRequest) (*grpcYoutubeThumbnails.DownloadThumbnailLinkResponse, error) {
	video := internal.VideoItem{}
	err := video.FindVideoID(req.URL)

	// incorrect URL
	if err != nil {
		return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: nil}, err
	}

	data, err := repository.FindVideoByID(ctx, video.VideoID)
	if err == nil {
		logrus.Printf("Found in the cache: %s", req.URL)
	} else {
		logrus.Printf("Not found in the cache: %s", req.URL)
		data = &video
		data.FindTitle(req.URL)

		// Can't find thumbnail link
		if data.FindThumbnailLink() != nil {
			return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: nil}, err
		}

		err = repository.CreateVideoItem(ctx, data)

		if err != nil {
			logrus.Printf("Can't add data to DB: %v", err)
		}
	}
	logrus.Printf("Downloaded thumbnail for URL: %s", req.URL)

	res, err := data.GetImage()

	// Can't get image by URL
	if err != nil {
		return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: nil}, err
	}

	// Can't download image
	image, err := io.ReadAll(res.Body)
	if err != nil {
		return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: nil}, err
	}
	// Downloaded successfully

	response := grpcYoutubeThumbnails.Response{
		Name:  fmt.Sprintf("%s ID%s", data.Name, data.VideoID),
		Image: image,
	}

	return &grpcYoutubeThumbnails.DownloadThumbnailLinkResponse{Response: &response}, nil
}
