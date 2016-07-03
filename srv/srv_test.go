package srv

import "testing"

// TIME IS THE BITCH HERE...

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
         
        "temperature_sensor": 0.0
    }`

func TestSerializer(t *testing.T) {
	rawJSON := []byte(incomingData)
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
	}

	location, err := SerializeLocation(rawJSON)
	if err != nil {
		t.Error(err.Error())
	}

	if locationStruct != location {
		t.Error("Expected - ", locationStruct)
		t.Error("Received - ", location)
	}
}
