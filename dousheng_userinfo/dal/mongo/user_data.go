package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserData struct {
	UserID          uint64 `bson:"user_id"`
	Avatar          string `bson:"avatar"`
	BackgroundImage string `bson:"background_image"`
	Signature       string `bson:"signature"`
	Username        string `bson:"username"`
}

func QueryUserData(ctx context.Context, userID uint64) (*UserData, error) {
	filter := bson.D{{"user_id", userID}}
	coll := MongoDB.Database("dousheng").Collection("user_data")
	var res UserData
	err := coll.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
