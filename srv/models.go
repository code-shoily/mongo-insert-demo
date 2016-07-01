package srv

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Loc depicts the latitude and longitude of an asset
type Loc struct {
	Coordinates [2]float64 `bson:"coordinates"`
	Type        string     `bson:"type"`
}

// Location is an abstraction of the location string to be inserted into the database
type Location struct {
	Number            string        `bson:"number"`
	ID                bson.ObjectId `bson:"_id,omitempty"`
	AcSensor          bool          `bson:"ac_sensor"`
	Active            bool          `bson:"active"`
	Bearing           float64       `bson:"bearing"`
	Door1Sensor       bool          `bson:"door1_sensor"`
	Door2Sensor       bool          `bson:"door2_sensor"`
	Door3Sensor       bool          `bson:"door3_sensor"`
	Door4Sensor       bool          `bson:"door4_sensor"`
	EngineSensor      bool          `bson:"engine_sensor"`
	Locate            Loc           `bson:"loc"`
	Speed             float64       `bson:"speed"`
	Status            string        `bson:"status"`
	TemperatureSensor bool          `bson:"temperature_sensor"`
	Time              time.Time     `bson:"time"`
}

// Save saves the struct as mongo document
func (data *Location) Save(session *mgo.Session) {
	collectionName := "asset_" + data.Number
	c := session.DB(DatabaseName).C(collectionName)
	err := c.Insert(data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

// SerializeLocation serializes from string into bson data
func SerializeLocation(line []byte) (Location, error) {
	data := Location{}
	err := json.Unmarshal(line, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//<-- TODO SNIP ME - This is to be used for smaller data -->//

// SerializeTodo serilizes the JSON string into native golang struct
func SerializeTodo(line []byte) (TodoList, error) {
	data := TodoList{}
	err := json.Unmarshal(line, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Todo shorter test model
type Todo struct {
	Priority    int    `json:"priority" bson:"priority"`
	Description string `json:"description" bson:"description"`
	Completed   bool   `json:"completed" bson:"completed"`
}

// TodoList shorter container model
type TodoList struct {
	Number string        `json:"number" bson:"-"`
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Title  string        `json:"title" bson:"title"`
	List   Todo          `json:"list" bson:"list"`
}

// Save saves the struct as mongo document
func (data *TodoList) Save(session *mgo.Session) {
	collectionName := "asset_" + data.Number
	c := session.DB("vts").C(collectionName)
	err := c.Insert(data)
	if err != nil {
		fmt.Print(err.Error())
	}
}
