package srv

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

// Loc depicts the latitude and longitude of an asset
type Loc struct {
	Coordinates [2]float64 `bson:"coordinates"`
	Type        string     `bson:"type"`
}

// Location is an abstraction of the location string to be inserted into the database
type Location struct {
	Number string `json:"number" bson:"-"`
	// ID                bson.ObjectId `bson:"_id,omitempty"`
	AcSensor          bool      `json:"ac_sensor" bson:"ac_sensor"`
	Active            bool      `json:"active" bson:"active"`
	Bearing           float32   `json:"bearing" bson:"bearing"`
	Door1Sensor       bool      `json:"door1_sensor" bson:"door1_sensor"`
	Door2Sensor       bool      `json:"door2_sensor" bson:"door2_sensor"`
	Door3Sensor       bool      `json:"door3_sensor" bson:"door3_sensor"`
	Door4Sensor       bool      `json:"door4_sensor" bson:"door4_sensor"`
	EngineSensor      bool      `json:"engine_sensor" bson:"engine_sensor"`
	Locate            Loc       `json:"loc" bson:"loc"`
	Speed             float32   `json:"speed" bson:"speed"`
	Status            string    `json:"status" bson:"status"`
	TemperatureSensor float32   `json:"temperature_sensor" bson:"temperature_sensor"`
	Time              time.Time `json:"time" bson:"time"`
}

// Save saves the struct as mongo document
func (data *Location) Save(session *mgo.Session) {
	collectionName := "asset_" + data.Number
	c := session.DB(DatabaseName).C(collectionName)
	err := c.Insert(data)
	if err != nil {
		log.Print("[SAVE ERROR] - " + err.Error())
	}
}

// SerializeLocation serializes from string into bson data
// 2601-0000-3690-072 FIXME PASSPORT NUMBER
func SerializeLocation(line []byte) (Location, error) {
	data := Location{}
	err := json.Unmarshal(line, &data)
	if err != nil {
		return data, err
	}

	if data.Locate.Type == "" {
		return data, errors.New("Invalid Data Format")
	}

	return data, nil
}
