package gpt

import (
	"context"
	"errors"
	"github.com/glwlg/whatisthis/internal/config"
	"github.com/glwlg/whatisthis/internal/utils"
	openai "github.com/sashabaranov/go-openai"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
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

func (g *Gpt) SearchStream(ctx context.Context, msg string) {
	if g.mock {
		runtime.EventsEmit(ctx, "onSearchResult", msg)
		return
	}
	stream, err := g.client.CreateChatCompletionStream(
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
			Stream: true,
		},
	)
	if err != nil {
		utils.LogError("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		var response openai.ChatCompletionStreamResponse
		response, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			utils.LogInfo("\nStream finished")
			runtime.EventsEmit(ctx, "onSearchResultStreamEnd", "")
			return
		}

		if err != nil {
			utils.LogError("\nStream error: %v\n", err)
			runtime.EventsEmit(ctx, "onSearchResultStreamEnd", "error")
			return
		}

		runtime.EventsEmit(ctx, "onSearchResultStream", response.Choices[0].Delta.Content)

		//fmt.Println(response.Choices[0].Delta.Content)
	}
}
