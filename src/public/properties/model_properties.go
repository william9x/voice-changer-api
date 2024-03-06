package properties

import (
	"github.com/golibs-starter/golib/config"
)

func NewModelProperties(loader config.Loader) (*ModelProperties, error) {
	props := ModelProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type ModelProperties struct {
	Data    []*ModelData
	DataMap map[string]*ModelData
}

type ModelData struct {
	ID       string     `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Category string     `json:"category,omitempty"`
	LogoURL  string     `json:"logo_url,omitempty"`
	Prompts  PromptData `json:"prompts"`
}

type PromptData struct {
	ShortSpeech []string `json:"short_speech,omitempty"`
	FunnyJokes  []string `json:"funny_jokes,omitempty"`
	ShortVerse  []string `json:"short_verse,omitempty"`
}

func (t *ModelProperties) Prefix() string {
	return "app.models"
}

func (t *ModelProperties) PostBinding() error {
	t.DataMap = make(map[string]*ModelData)
	for _, model := range t.Data {
		t.DataMap[model.ID] = model
	}
	return nil
}
