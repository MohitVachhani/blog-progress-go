package blockrepo

import (
	"context"
	"log"
	"time"

	blockInterface "github.com/MohitVachhani/blog-progress/pkg/struct/block"
	mongoUtils "github.com/MohitVachhani/blog-progress/pkg/utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBlockById(input blockInterface.GetBlockById) blockInterface.BlockSchema {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// convert string to objectId
	blockObjectId := mongoUtils.ConvertStringToPrimitiveObjectId(input.BlockId)

	blockCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "blocks")

	var Block blockInterface.BlockSchema

	blockFindOneError := blockCollection.FindOne(ctx, bson.M{"_id": blockObjectId}).Decode(&Block)

	if blockFindOneError != nil {
		log.Fatal(blockFindOneError)
	}

	return Block
}

func CreateBlock(input blockInterface.CreateBlockInput) blockInterface.BlockSchema {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	blocksCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "blocks")

	var createBlockInput blockInterface.BlockSchema = blockInterface.BlockSchema{
		Type:      input.Type,
		Text:      input.Text,
		CreatedAt: time.Now().UTC(),
		Url:       input.Url,
		ParentId:  mongoUtils.ConvertStringToPrimitiveObjectId(input.ParentId),
		Duration:  input.Duration,
	}

	insertOneBlockResult, err := blocksCollection.InsertOne(ctx, createBlockInput)

	if err != nil {
		log.Fatal("Error occurred while creating a new block in mongo", err)
	}

	// convert interface{} to primitive objectId
	var insertedBlockId = insertOneBlockResult.InsertedID.(primitive.ObjectID)
	log.Println("Block Created with Id:", insertedBlockId)

	var getBlockByIdInput blockInterface.GetBlockById = blockInterface.GetBlockById{
		BlockId: insertedBlockId.Hex(),
	}

	block := GetBlockById(getBlockByIdInput)

	return block
}

func UpdateBlock(input blockInterface.UpdateBlockInput) blockInterface.BlockSchema {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	blocksCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "blocks")
	var blockChildrenIds []primitive.ObjectID

	for inputChildrenIdsIterator := 0; inputChildrenIdsIterator < len(input.ChildrenIds); inputChildrenIdsIterator++ {
		blockChildrenIds = append(blockChildrenIds, mongoUtils.ConvertStringToPrimitiveObjectId(input.ChildrenIds[inputChildrenIdsIterator]))
	}

	updateBlockInput := blockInterface.BlockSchema{
		Type:        input.Type,
		Url:         input.Url,
		Text:        input.Text,
		ParentId:    mongoUtils.ConvertStringToPrimitiveObjectId(input.ParentId),
		Duration:    input.Duration,
		UpdatedAt:   time.Now().UTC(),
		ChildrenIds: blockChildrenIds,
	}

	_, err := blocksCollection.UpdateOne(ctx, bson.M{"_id": mongoUtils.ConvertStringToPrimitiveObjectId(input.BlockId)}, bson.M{"$set": updateBlockInput})

	if err != nil {
		log.Fatal("Error occurred while updating a block in mongo", err)
	}

	var getBlockByIdInput blockInterface.GetBlockById = blockInterface.GetBlockById{
		BlockId: input.BlockId,
	}

	block := GetBlockById(getBlockByIdInput)

	return block
}
