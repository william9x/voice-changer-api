package resources

// Model ...
type Model struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	LogoURL  string `json:"logo_url,omitempty"`
}

func NewModelResource(id, name, cat, logo string) *Model {
	return &Model{
		ID:       id,
		Name:     name,
		Category: cat,
		LogoURL:  logo,
	}
}
