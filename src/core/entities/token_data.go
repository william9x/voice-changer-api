package entities

type TokenData struct {
	Issuer   string
	Expires  int64
	IssuedAt int64
	Subject  string
	UserID   string
	Claims   map[string]interface{}
}
