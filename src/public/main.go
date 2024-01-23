package main

import (
	"github.com/Braly-Ltd/voice-changer-api-public/bootstrap"
	"go.uber.org/fx"
)

// @title Voice Changer API Public
// @version 1.0.0
func main() {
	fx.New(bootstrap.All()).Run()
}
