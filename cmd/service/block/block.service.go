package blockService

import (
	"strings"

	blockRepo "github.com/MohitVachhani/blog-progress/cmd/repo/block"
	blockInterface "github.com/MohitVachhani/blog-progress/pkg/struct/block"
)

func getDurationOfBlock(input blockInterface.CreateBlockInput) int {
	if strings.Compare(input.Type, "paragraph") == 0 {
		textWordsLength := len(strings.Split(input.Text, " "))
		return (textWordsLength / 250) * 60
	}

	if strings.Compare(input.Type, "image") == 0 {
		return 7
	}

	return input.Duration
}

func CreateBlock(input blockInterface.CreateBlockInput) blockInterface.BlockSchema {
	blockDuration := getDurationOfBlock(input)

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
