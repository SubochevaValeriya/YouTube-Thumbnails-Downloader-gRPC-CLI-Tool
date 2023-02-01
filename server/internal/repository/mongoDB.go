package repository

import (
	"github.com/SubochevaValeriya/grpcYoutubeThumbnails/server/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

type MongoConfig struct {
	Host            string
	Port            string
	DefaultDatabase string
	Collection      string
}

var collection *mongo.Collection

func MongoConnection(mongoConfig MongoConfig) (*mongo.Client, error) {
	address := "mongodb://" + mongoConfig.Host + ":" + mongoConfig.Port
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	err = client.Connect(ctx)
	if err != nil {
		return client, err
	}

	collection = client.Database(mongoConfig.DefaultDatabase).Collection(mongoConfig.Collection)

	return client, nil
}

func FindVideoByID(ctx context.Context, videoID string) (*internal.VideoItem, error) {
	data := &internal.VideoItem{}
	filter := bson.M{"video_id": videoID}

	res := collection.FindOne(ctx, filter)

	return data, res.Decode(data)
}

func CreateVideoItem(ctx context.Context, data *internal.VideoItem) error {
	_, err := collection.InsertOne(ctx, data)

	return err
}

func DropCollection(ctx context.Context) error {
	err := collection.Drop(ctx)
	return err
}
