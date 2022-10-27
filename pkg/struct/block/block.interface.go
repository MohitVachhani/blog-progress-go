package blockInterface

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateBlockInput struct {
	Type     string `json:"type,omitempty"`
	Url      string `json:"url,omitempty"`
	Text     string `json:"text,omitempty"`
	ParentId string `json:"parentId"`
	Duration int    `json:"duration,omitempty"`
}

type UpdateBlockInput struct {
	BlockId     string   `json:"blockId,omitempty"`
	Type        string   `json:"type,omitempty"`
	Url         string   `json:"url,omitempty"`
	Text        string   `json:"text,omitempty"`
	ParentId    string   `json:"parentId"`
	Duration    int      `json:"duration,omitempty"`
	ChildrenIds []string `json:"childrenIds,omitempty"`
}

type GetBlockById struct {
	BlockId string
}

type BlockSchema struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Type        string               `json:"type,omitempty" bson:"type,omitempty"`
	CreatedAt   time.Time            `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time            `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Url         string               `json:"url,omitempty" bson:"url,omitempty"`
	Text        string               `json:"text,omitempty" bson:"text,omitempty"`
	ParentId    primitive.ObjectID   `json:"parentId" bson:"parentId,omitempty"`
	Duration    int                  `json:"duration,omitempty" bson:"duration,omitempty"`
	ChildrenIds []primitive.ObjectID `json:"childrenIds,omitempty" bson:"childrenIds,omitempty"`
}
