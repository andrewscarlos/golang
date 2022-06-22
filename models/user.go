package models

type Address struct {
	State   string `json:"state" bjon:"state"`
	City    string `json:"city" bjon:"city"`
	Pincode int    `json:"pincode" bjon:"pincode"`
}

type User struct {
	Name    string  `json:"name" bson:"name"`
	Age     int     `json:"age" bson:"age"`
	Address Address `json:"address"bson:"address"`
}
