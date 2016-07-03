package srv

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// NetHost is the default port to be used in the net connection
const NetHost = "localhost"

// NetPort is the default port to be used in the net connection
const NetPort = 2800

// DatabaseName is the name of the database used
const DatabaseName = "vts"

// MongoHost is the default host for MongoDB
const MongoHost = "127.0.0.1"

// SetLogger sets the logger to both file and stdin
func SetLogger() {
	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:   "./logs/server.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}, os.Stdout))
}
