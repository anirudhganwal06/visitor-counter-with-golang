package controllers

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/anirudhganwal06/visitor-counter-with-golang/app/models"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetHomepage is for getting homepage
func GetHomepage(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	IP := models.IP{}
	filter := bson.M{"ip": ip}
	decoder := db.Collection("IPs").FindOne(context.TODO(), filter)
	decoder.Decode(&IP)
	fmt.Println(IP)
	if (IP == models.IP{}) {
		IP.IP = ip
		_, err := db.Collection("IPs").InsertOne(context.TODO(), IP)
		if err != nil {
			log.Fatal(err)
		}
	}
	counter, err := db.Collection("IPs").EstimatedDocumentCount(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counter)
	response := models.Counter{}
	response.Counter = counter
	fmt.Print("response: ")
	fmt.Println(response)
	respondJSON(res, http.StatusOK, response)
}
