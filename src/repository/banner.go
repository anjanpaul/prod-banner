package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"mongo-fiber-api/dto"
	"time"
)

type Banner struct {
	Ctx context.Context
}

func (b *Banner) CreateUserIndex() (err error) {
	col := DB.Collection(BannerCollection)
	_, err = col.Indexes().CreateOne(b.Ctx, mongo.IndexModel{
		Options: options.Index().SetUnique(true),
		Keys:    bsonx.MDoc{"title": bsonx.Int32(-1)},
	})
	return
}

func (b *Banner) CreateBanner(input *dto.CreateBannerInput) (ID string, err error) {
	col := DB.Collection(BannerCollection)
	res, err := col.InsertOne(b.Ctx, BannerDoc{
		ID:             primitive.NewObjectID(),
		Title:          input.Title,
		Summary:        input.Summary,
		Status:         input.Status,
		MediaType:      input.MediaType,
		MediaReference: input.MediaReference,
		Tags:           input.Tags,
		ExternalLink:   input.ExternalLink,
		Priority:       input.Priority,
		Created:        time.Now(),
		Updated:        time.Now(),
	})

	if err != nil {
		return
	}
	ID = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (b *Banner) GetAllBanners() (cursor *mongo.Cursor, err error) {
	col := DB.Collection(BannerCollection)
	cursor, err = col.Find(context.TODO(), bson.D{})
	if err != nil {
		return
	}
	return
}
func (b *Banner) GetTaggedBanner(tag string) (cursor *mongo.Cursor, err error) {
	col := DB.Collection(BannerCollection)
	cursor, err = col.Find(context.TODO(), bson.M{"tags": tag})
	if err != nil {
		return
	}
	return
}
func (b *Banner) GetTaggedStatusBanner(tag string, status string) (cursor *mongo.Cursor, err error) {
	col := DB.Collection(BannerCollection)
	cursor, err = col.Find(context.TODO(), bson.M{"tags": tag, "status": status})
	if err != nil {
		return
	}
	return
}

func (b *Banner) GetStatusBanner(status string) (cursor *mongo.Cursor, err error) {
	col := DB.Collection(BannerCollection)
	cursor, err = col.Find(context.TODO(), bson.M{"status": status})
	if err != nil {
		return
	}
	return
}

func (b *Banner) GetABanner(objId primitive.ObjectID) (res BannerDoc, err error) {
	col := DB.Collection(BannerCollection)
	err = col.FindOne(b.Ctx, bson.M{"id": objId}).Decode(&res)
	if err != nil {
		return
	}
	return
}
func (b *Banner) DeleteBanner(objId primitive.ObjectID) (res *mongo.DeleteResult, err error) {
	col := DB.Collection(BannerCollection)
	res, err = col.DeleteOne(b.Ctx, bson.M{"id": objId})
	if err != nil {
		return
	}
	if res.DeletedCount == 0 {
		err = errors.New("no Data Deleted")
		return
	}
	return
}

func (b *Banner) EditBanner(ID primitive.ObjectID, input dto.EditBannerInput) (err error) {
	col := DB.Collection(BannerCollection)
	_, err = col.UpdateOne(
		b.Ctx,
		bson.M{"id": bson.M{"$eq": ID}},
		bson.M{"$set": bson.M{
			"title":          input.Title,
			"summary":        input.Summary,
			"status":         input.Status,
			"mediatype":      input.MediaType,
			"mediareference": input.MediaReference,
			"tags":           input.Tags,
			"externallink":   input.ExternalLink,
			"priority":       input.Priority,
		}})
	if err != nil {
		return
	}
	return
}

////GetUserByEmai get user  by email
func (b *Banner) GetUserByEmail(email string) (user *UserDoc, err error) {
	col := DB.Collection(UserCollection)
	err = col.FindOne(b.Ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)
	return
}

func (b *Banner) CreateUser(email string, password string) (ID string, err error) {
	col := DB.Collection(UserCollection)
	res, err := col.InsertOne(b.Ctx, UserDoc{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: password,
		Created:  time.Now(),
		Updated:  time.Now(),
	})
	if err != nil {
		return
	}
	ID = res.InsertedID.(primitive.ObjectID).Hex()
	return
}
