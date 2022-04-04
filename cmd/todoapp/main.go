package main

import (
	_ "github.com/OksidGen/todoapp/internal/middleware"
	"github.com/OksidGen/todoapp/internal/router"
)

func main() {
	router.InitializeRouter()
}
