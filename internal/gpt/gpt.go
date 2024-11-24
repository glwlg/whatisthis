package gpt

import (
	"context"
	"github.com/glwlg/whatisthis/internal/config"
	"github.com/glwlg/whatisthis/internal/utils"
	openai "github.com/sashabaranov/go-openai"
)

type Gpt struct {
	client         *openai.Client
	model          string
	promptTemplate string
	mock           bool
}

func NewClient(cfg config.Config, gCfg *config.GuiConfig) *Gpt {

	//clientConfig := openai.DefaultAzureConfig(gCfg.OpenAI.APIKey, gCfg.OpenAI.BaseUrl)
	clientConfig := openai.DefaultConfig(gCfg.OpenAI.APIKey)
	clientConfig.BaseURL = gCfg.OpenAI.BaseUrl
	//clientConfig.APIVersion = gCfg.OpenAI.ApiVersion

	client := openai.NewClientWithConfig(clientConfig)

	return &Gpt{
		client:         client,
		model:          gCfg.OpenAI.Model,
		promptTemplate: cfg.Search.PromptTemplates[0],
		mock:           cfg.AppConfig.Mock,
	}
}

func (g *Gpt) Search(msg string) (string, error) {
	if g.mock {
		return msg, nil
	}
	resp, err := g.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: g.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: g.promptTemplate,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)
	if err != nil {
		utils.LogError("ChatCompletion error: %v\n", err)
		return "", err
	}

	utils.LogInfo(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}
