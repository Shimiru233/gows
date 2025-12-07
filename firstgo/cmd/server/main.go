package main

import (
	"firstgo/internal/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
