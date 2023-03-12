package chatgpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"open-ai-feishu/g"
)

func Ask(question string) string {
	client := openai.NewClient(g.GetConfig().ChatGpt.AuthToken)
	//client.CreateEditImage()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return "出错了，请重试"
	}
	return resp.Choices[0].Message.Content
}
