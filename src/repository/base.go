package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	DB          *mongo.Database
	MongoClient *mongo.Client
)

const BannerCollection = "banners"
const UserCollection = "user"

type UserDoc struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Created  time.Time          `bson:"created"`
	Updated  time.Time          `bson:"updated"`
}

type BannerDoc struct {
	ID             primitive.ObjectID `json:"_id"`
	Title          string             `json:"title"`
	Summary        string             `json:"summary"`
	Status         string             `json:"status"`
	Tags           string             `json:"tags"`
	MediaType      string             `json:"mediaType"`
	MediaReference string             `json:"mediaReference"`
	ExternalLink   string             `json:"externalLink"`
	Priority       string             `json:"priority"`
	Created        time.Time          `json:"created"`
	Updated        time.Time          `json:"updated"`
}
