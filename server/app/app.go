package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/anirudhganwal06/visitor-counter-with-golang/server/app/controllers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App has Router instance
type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

// Initialize the app
func (app *App) Initialize(MongoURI string) {
	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	db := client.Database("visitor-counter-with-golang")
	app.DB = db
	app.Router = mux.NewRouter()
	app.setRouters()

}

func (app *App) setRouters() {
	app.Get("/api/counter", app.handleRequest(controllers.GetHomepage))
}

// Get wraps the router for GET method
func (app *App) Get(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// RequestHandlerFunction handles the requests
type RequestHandlerFunction func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(resw http.ResponseWriter, req *http.Request) {
		handler(app.DB, resw, req)
	}
}

// Run listens to the port
func (app *App) Run(port string) {
	http.ListenAndServe(port, app.Router)
}
