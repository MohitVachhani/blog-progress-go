package resourceService

import (
	resourceRepo "github.com/MohitVachhani/blog-progress/cmd/repo/resource"
	blockService "github.com/MohitVachhani/blog-progress/cmd/service/block"
	blockInterface "github.com/MohitVachhani/blog-progress/pkg/struct/block"
	resourceInterface "github.com/MohitVachhani/blog-progress/pkg/struct/resource"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateResource(input resourceInterface.CreateResourceInput) resourceInterface.ResourceSchema {

	resourceId := primitive.NewObjectID().Hex()

	// create root block of blog input
	createBlockInput := blockInterface.CreateBlockInput{
		Type:     "blog",
		ParentId: resourceId,
	}

	// create root block of blog
	createdResourceParentBlock := blockService.CreateBlock(createBlockInput)

	// create resource input
	var createResourceInput resourceInterface.CreateResourceRepoInput = resourceInterface.CreateResourceRepoInput{
		ResourceId:    resourceId,
		Name:          input.Name,
		Description:   input.Description,
		Status:        "unPublished",
		ParentBlockId: createdResourceParentBlock.ID.Hex(),
	}

	// resource document created
	createdResource := resourceRepo.CreateResource(createResourceInput)

	return createdResource
}
