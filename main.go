package main

import (
	"embed"
	"go-base/cmd"
)

//go:embed cmd/*
var EmbedFs embed.FS

// @title           Go Gin Base
// @description     Go-Gin-Base quickly build and develop web applications. restful API, microservice...
// @version         v0.0.1

// @author.name     Nguyen Quang Truong
// @author.github   https://github.com/truongbo17
// @author.email    truongnq017@gmail.com

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				    Type "Bearer" followed by a space and JWT token.
func main() {
	cmd.Execute(EmbedFs)
}
