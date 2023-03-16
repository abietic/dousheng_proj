package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type VideoData struct {
	VideoID     uint64    `bson:"video_id"`
	AuthorID    uint64    `bson:"author_id"`
	Title       string    `bson:"title"`
	PlayUrl     string    `bson:"play_url"`
	CoverUrl    string    `bson:"cover_url"`
	PublishTime time.Time `bson:"publish_time"`
}

func CreateVideoData(ctx context.Context, videoData *VideoData) error {
	coll := MongoDB.Database("dousheng").Collection("video_data")
	_, err := coll.InsertOne(ctx, *videoData)
	return err
}

func QueryVideoDataList(ctx context.Context, authorID uint64) ([]VideoData, error) {
	filter := bson.D{{"author_id", authorID}}
	// sorter := bson.D{{"publish_time", -1}}
	// opts := options.Find().SetSort(sorter)
	coll := MongoDB.Database("dousheng").Collection("video_data")
	// cursor, err := coll.Find(ctx, filter, opts)
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var res []VideoData
	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
