package srv

import (
	"testing"
	"time"
)

// SingleLineJSON is important for testing in telnet.
const SingleLineJSON string = `
{"status":"A","bearing":136.67,"door4_sensor":false,"door2_sensor":false,"number":"XMVT863071010819875","engine_sensor":false,"door3_sensor":false,"active":true,"door1_sensor":false,"ac_sensor":false,"speed":0.0,"loc":{"type":"Point","coordinates":[90.36850666666666,23.74985]},"time":"2016-07-03T12:09:54.035402+00:00","temperature_sensor":0.0}
`

const incomingData string = `
   {
        "status": "A", 
        "bearing": 136.67, 
        "door4_sensor": false, 
        "door2_sensor": false, 
        "number": "XMVT863071010819875", 
        "engine_sensor": false, 
        "door3_sensor": false, 
        "active": true, 
        "door1_sensor": false, 
        "ac_sensor": false, 
        "speed": 0.0, 
        "loc": {
            "type": "Point", 
            "coordinates": [90.36850666666666, 23.74985]
        }, 
        "time": "2016-07-03T12:09:54.035402+00:00", 
        "temperature_sensor": 0.0
    }
`

func convertTime(timeStr string) (time.Time, error) {
	t, e := time.Parse(time.RFC3339, timeStr)
	if e != nil {
		return time.Now(), e
	}

	return t, nil
}

func TestSerializer(t *testing.T) {
	rawJSON := []byte(incomingData)
	goTime, e := convertTime("2016-07-03T12:09:54.035402+00:00")
	if e != nil {
		t.Fatal(e.Error())
	}

	locationStruct := Location{
		Status:       "A",
		Bearing:      136.67,
		Door1Sensor:  false,
		Door2Sensor:  false,
		Door3Sensor:  false,
		Door4Sensor:  false,
		EngineSensor: false,
		Active:       true,
		AcSensor:     false,
		Number:       "XMVT863071010819875",
		Speed:        0.0,
		Locate: Loc{
			Coordinates: [2]float64{90.36850666666666, 23.74985},
			Type:        "Point",
		},
		TemperatureSensor: 0.0,
		Time:              goTime,
	}

	location, err := SerializeLocation(rawJSON)
	if err != nil {
		t.Fatal(err.Error())
	}

	if locationStruct.Time.UnixNano() != location.Time.UnixNano() {
		t.Fatal("Times do not match")
	} else {
		now := time.Now().UTC()
		location.Time = now
		locationStruct.Time = now

		if locationStruct != location {
			t.Error("Expected - ", locationStruct)
			t.Error("Got - ", location)
		}
	}
}
