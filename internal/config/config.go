package config

import (
	"github.com/spf13/viper"
)

// Config 定义了配置文件中的各个字段
type Config struct {
	AppConfig struct {
		Mock bool `mapstructure:"mock"`
	} `mapstructure:"app_config"`
	Hotkey string `mapstructure:"hotkey"`
	OCR    struct {
		Language string `mapstructure:"language"`
	} `mapstructure:"ocr"`
	History struct {
		Enabled bool `mapstructure:"enabled"`
	} `mapstructure:"history"`
	Search struct {
		PromptTemplates []string `mapstructure:"prompt_templates"`
	} `mapstructure:"search"`
	// 可以根据需要添加更多的配置项
}

// LoadConfig 加载配置文件并返回配置结构体
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")    // 设置配置文件名（不包括扩展名）
	viper.AddConfigPath("./configs") // 设置配置文件路径
	viper.SetConfigType("yaml")      // 设置配置文件类型
	viper.AutomaticEnv()             // 读取环境变量

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type OpenAI struct {
	APIKey     string `json:"api_key"`
	BaseUrl    string `json:"base_url"`
	ApiVersion string `json:"api_version"`
	Model      string `json:"model"`
}

type GuiConfig struct {
	OpenAI OpenAI `json:"openai"`
}

func NewGuiConfig() *GuiConfig {
	return &GuiConfig{}
}

func (g *GuiConfig) IsOpenAIEmpty() bool {
	return g.OpenAI.APIKey == "" && g.OpenAI.BaseUrl == "" && g.OpenAI.ApiVersion == ""
}
