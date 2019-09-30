package main

import (
	"fmt"

	"github.com/anirudh06/visitor-counter-with-golang/app"
	"github.com/anirudh06/visitor-counter-with-golang/config"
)

func main() {
	fmt.Println("Starting application...")
	MongoURI := config.MongoURI
	app := &app.App{}
	app.Initialize(MongoURI)
	var port string = "5000"
	fmt.Println(`Server running @` + port)
	app.Run(":" + port)
}
