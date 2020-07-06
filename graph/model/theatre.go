
package model

type Theater struct {
	TheaterID int      `json:"theaterId"`
	Location  Location `bson:"location" json:"location"`
}
type Address struct {
	Street1 string `json:"street1"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
}
type Geo struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type Location struct {
	Address Address `bson:"address" json:"address"`
	Geo     Geo     `bson:"geo" json:"geo"`
}