package main

import (
	"fmt"

	"github.com/anirudhganwal06/visitor-counter-with-golang/server/app"
	"github.com/anirudhganwal06/visitor-counter-with-golang/server/config"
)

func main() {
	fmt.Println("Starting application...")
	MongoURI := config.MongoURI
	app := &app.App{}
	app.Initialize(MongoURI)
	var port string = "8080"
	fmt.Println(`Server running @` + port)
	app.Run(":" + port)
}
