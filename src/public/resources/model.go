package resources

// Model ...
type Model struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func NewModelResource(id, name, cat string) *Model {
	return &Model{
		ID:       id,
		Name:     name,
		Category: cat,
	}
}
