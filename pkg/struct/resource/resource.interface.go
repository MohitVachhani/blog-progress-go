package resourceinterface

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateResourceInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateResourceRepoInput struct {
	ResourceId    string
	Name          string
	Status        string
	Description   string
	ParentBlockId string
}

type GetResourceById struct {
	ResourceId string
}

type ResourceSchema struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	ParentBlockId primitive.ObjectID `json:"parentBlockId" bson:"parentBlockId,omitempty"`
}
