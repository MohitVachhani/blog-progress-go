package blockService

import (
	"log"
	"strings"

	blockRepo "github.com/MohitVachhani/blog-progress/cmd/repo/block"
	blockInterface "github.com/MohitVachhani/blog-progress/pkg/struct/block"
)

func getDurationOfBlock(input blockInterface.CreateBlockInput) int {
	if strings.Compare(input.Type, "paragraph") == 0 {
		log.Printf(input.Type)

		textWordsLength := len(strings.Split(input.Text, " "))
		log.Print(textWordsLength)

		return (textWordsLength / 30) * 60
	}

	if strings.Compare(input.Type, "image") == 0 {
		return 7
	}

	return input.Duration
}

func CreateBlock(input blockInterface.CreateBlockInput) blockInterface.BlockSchema {
	blockDuration := getDurationOfBlock(input)
	log.Print(blockDuration)

	createBlockInput := blockInterface.CreateBlockInput{
		Type:     input.Type,
		Url:      input.Url,
		Text:     input.Text,
		ParentId: input.ParentId,
		Duration: blockDuration,
	}

	createdBlock := blockRepo.CreateBlock(createBlockInput)

	return createdBlock
}

func UpdateBlock(input blockInterface.UpdateBlockInput) blockInterface.BlockSchema {

	return blockRepo.UpdateBlock(input)
}
