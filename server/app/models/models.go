package models

// IP model
type IP struct {
	IP string `json:"ip" bson:"ip"`
}

// Counter model
type Counter struct {
	Counter int64 `json:"counter" bson:"counter"`
}
