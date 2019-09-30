package models

// IP model
type IP struct {
	IP string `json:"ip" bson:"ip"`
}

// Counter model
type Counter struct {
	ID      string `json:"id" bson:"id"`
	Counter int32  `json:"counter" bson:"counter"`
}
