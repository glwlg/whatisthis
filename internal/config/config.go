package config

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
)

type AppConfig struct {
	Mock bool `mapstructure:"mock"`
}

type Search struct {
	PromptTemplates []string `mapstructure:"prompt_templates"`
}

// Config 定义了配置文件中的各个字段
type Config struct {
	AppConfig AppConfig `mapstructure:"app_config"`
	Hotkey    string    `mapstructure:"hotkey"`
	OCR       struct {
		Language string `mapstructure:"language"`
	} `mapstructure:"ocr"`
	History struct {
		Enabled bool `mapstructure:"enabled"`
	} `mapstructure:"history"`
	Search Search `mapstructure:"search"`
	// 可以根据需要添加更多的配置项
}

// InitConfig 初始化配置
func InitConfig() (*Config, error) {
	filePath := "./configs/config.yaml"
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// 如果文件不存在，创建一个新的文件并写入默认配置
		defaultConfig := &Config{
			AppConfig: AppConfig{
				Mock: false,
			},
			Search: Search{
				PromptTemplates: []string{"# Role：百科机器人\n       ## Background：\n       用户希望你作为一个百科机器人，能够向用户说明他们提供的内容是什么，而不试图去理解问题本身。这可能是因为用户希望获得对某个特定内容的解释或定义，而不需要进一步的分析或推理。\n       ## Attention：\n       用户对任务的渴求是非常高的，他们需要一个准确、简洁的解释，而不是复杂的分析。提供清晰明确的信息将让用户感到满意。\n       ## Profile：\n       - Author: abai\n       - Version: 1.0\n       - Language: 中文\n       - Description: 你是一名百科机器人，擅长解释和定义用户提供的内容。你擅长提供简洁、准确的解释，而不会尝试对问题本身进行分析或推理。\n       ### Skills:\n       - 解释能力：能够准确解释用户提供的内容。\n       - 简洁表述：能够用简洁的语言表达复杂的内容。\n       - 中立客观：在解释内容时保持中立，不加入个人观点。\n       - 快速响应：能够在短时间内提供高质量的解释。\n       - 多领域知识：具备广泛的知识储备，能够解释各种类型的内容。\n       ## Goals:\n       - 解释用户提供的内容。\n       - 不进行问题分析或推理。\n       - 提供简洁、准确的解释。\n       - 保持中立客观。\n       - 快速响应用户需求。\n       ## Constrains:\n       - 不进行问题的分析或推理。\n       - 提供的信息必须准确且简洁。\n       - 保持中立和客观。\n       - 避免使用复杂的术语，确保解释易于理解。\n       - 遵守用户提供的上下文和要求。\n       - 每次解释不超过100字。\n       - 可以附上链接和参考文献以提供更多信息。\n       ## Workflow:\n       1. 接收用户提供的内容。\n       2. 对内容进行初步理解，不进行深入分析。\n       3. 提供简洁、准确的解释，每次不超过100字。\n       4. 确保解释中不包含个人观点或复杂分析。\n       5. 根据需要附上链接和参考文献。\n       6. 检查解释内容，确保其准确性和简洁性。\n       ## OutputFormat:\n       - 提供简洁的解释，每次不超过100字。\n       - 避免复杂的术语。\n       - 保持中立和客观。\n       - 不进行问题的分析或推理。\n       - 根据需要附上链接和参考文献。\n       ## Suggestions:\n       - 提供准确的解释，确保用户能够理解。\n       - 避免复杂的分析或推理，保持解释的简洁性。\n       - 确保信息的中立和客观，不加入个人观点。\n       - 使用简单的语言，确保解释易于理解。\n       - 根据需要附上链接和参考文献，以提供更多信息。"},
			},
		}
		//先创建文件夹
		err = os.MkdirAll("./configs", os.ModePerm)
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		encoder := yaml.NewEncoder(file)
		err = encoder.Encode(defaultConfig)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		// 如果发生其他错误，返回错误
		return nil, err
	}

	// 读取并解析配置文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
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
	Delay  string `json:"delay"`
}

func NewGuiConfig() *GuiConfig {
	return &GuiConfig{}
}

func (g *GuiConfig) IsOpenAIEmpty() bool {
	return g.OpenAI.APIKey == "" && g.OpenAI.BaseUrl == "" && g.OpenAI.ApiVersion == ""
}
