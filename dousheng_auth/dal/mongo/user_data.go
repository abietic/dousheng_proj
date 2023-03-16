package mongo

import "context"

type UserData struct {
	UserID          uint64 `bson:"user_id"`
	Avatar          string `bson:"avatar"`
	BackgroundImage string `bson:"background_image"`
	Signature       string `bson:"signature"`
	Username        string `bson:"username"`
}

func CreateUserData(ctx context.Context, userData *UserData) error {
	coll := MongoDB.Database("dousheng").Collection("user_data")
	_, err := coll.InsertOne(ctx, *userData)
	return err
}
