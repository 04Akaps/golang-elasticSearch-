package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type MigrationFile struct {
	ObjectId primitive.ObjectID `json:"_id"`
	Owner    string             `json:"owner"`
	Tid      interface{}        `json:"tid"`
	KeyWord  string             `json:"keyWord"`
}

type Migration struct {
	PK      string `json:"pk"`
	Owner   string `json:"owner"`
	Tid     int    `json:"tid"`
	KeyWord string `json:"keyWord"`
}
