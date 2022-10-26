package resourcerepo

import (
	"context"
	"log"
	"time"

	resourceInterface "github.com/MohitVachhani/blog-progress/pkg/struct/resource"
	mongoUtils "github.com/MohitVachhani/blog-progress/pkg/utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetResourceById(input resourceInterface.GetResourceById) resourceInterface.ResourceSchema {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// convert string to objectId
	resourceObjectId := mongoUtils.ConvertStringToPrimitiveObjectId(input.ResourceId)

	resourceCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "resources")

	var resource resourceInterface.ResourceSchema

	resourceFindOneError := resourceCollection.FindOne(ctx, bson.M{"_id": resourceObjectId}).Decode(&resource)

	if resourceFindOneError != nil {
		log.Fatal(resourceFindOneError)
	}

	return resource
}

func CreateResource(input resourceInterface.CreateResourceRepoInput) resourceInterface.ResourceSchema {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	resourceCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "resources")

	var createResourceInput resourceInterface.ResourceSchema = resourceInterface.ResourceSchema{
		ID:            mongoUtils.ConvertStringToPrimitiveObjectId(input.ResourceId),
		Name:          input.Name,
		Description:   input.Description,
		CreatedAt:     time.Now().UTC(),
		Status:        input.Status,
		ParentBlockId: mongoUtils.ConvertStringToPrimitiveObjectId(input.ParentBlockId),
	}

	insertOneResourceResult, err := resourceCollection.InsertOne(ctx, createResourceInput)

	if err != nil {
		log.Fatal("Error occurred while creating a new user in mongo", err)
	}

	// convert interface{} to primitive objectId
	var insertedResourceId = insertOneResourceResult.InsertedID.(primitive.ObjectID)

	var getResourceByIdInput resourceInterface.GetResourceById = resourceInterface.GetResourceById{
		ResourceId: insertedResourceId.Hex(),
	}

	resource := GetResourceById(getResourceByIdInput)

	return resource
}
