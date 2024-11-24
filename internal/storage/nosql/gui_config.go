package nosql

import (
	"fmt"
	"github.com/glwlg/whatisthis/internal/config"
)

func (s *Storage) SaveOpenai(openai config.OpenAI) error {
	cfg, err := s.GetGuiConfig()
	if err != nil {
		return fmt.Errorf(`failed to read config: %w`, err)
	}
	cfg.OpenAI = openai
	err = s.writeConfig(cfg)
	if err != nil {
		return fmt.Errorf(`failed to save locale: %w`, err)
	}
	return nil
}

func (s *Storage) SaveConfig(config *config.GuiConfig) error {
	cfg, err := s.GetGuiConfig()
	if err != nil {
		return fmt.Errorf(`failed to read config: %w`, err)
	}
	cfg.Delay = config.Delay
	err = s.writeConfig(cfg)
	if err != nil {
		return fmt.Errorf(`failed to save locale: %w`, err)
	}
	return nil
}
