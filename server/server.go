package server

import (
	"context"
	"fmt"
	grpcYoutubeThumbnails "github.com/SubochevaValeriya/grpcYoutubeThumbnails/proto"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {

	// configs
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing congigs: %s", err.Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoURl := os.Getenv("MONGODB_URL")

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURl))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pokemon Service Started")
	collection = client.Database("videos-db").Collection("videos")

	lis, err := net.Listen("tcp", "0.0.0.0:4041")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	grpcYoutubeThumbnails.RegisterYoutubeThumbnailsServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:
	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}

const defaultPort = "4041"

var collection *mongo.Collection

type server struct {
	grpcYoutubeThumbnails.YoutubeThumbnailsServiceServer
}

type videoItem struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	VideoID       string             `bson:"video_id"`
	Name          string             `bson:"name"`
	ThumbnailLink string             `bson:"thumbnail_link"`
}

func getVideoData(data *videoItem) *grpcYoutubeThumbnails.Video {
	return &grpcYoutubeThumbnails.Video{
		Id:            data.ID.Hex(),
		Link:          data.ThumbnailLink,
		ThumbnailLink: data.ThumbnailLink,
	}
}

func (*server) LoadThumbnailLink(ctx context.Context, req *grpcYoutubeThumbnails.LoadThumbnailLinkRequest) (*grpcYoutubeThumbnails.LoadThumbnailLinkResponse, error) {
	log.Println("Load thumbnail")
	pokemon := req.GetVideo()
	data := videoItem{
		VideoID:       pokemon.GetId(),
		Name:          pokemon.GetLink(),
		ThumbnailLink: pokemon.GetThumbnailLink(),
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}
	return &grpcYoutubeThumbnails.LoadThumbnailLinkResponse{
		Video: &grpcYoutubeThumbnails.Video{
			Id: oid.Hex(),
			//Pid:         pokemon.GetPid(),
			Link:          pokemon.GetLink(),
			ThumbnailLink: pokemon.GetThumbnailLink(),
		},
	}, nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
