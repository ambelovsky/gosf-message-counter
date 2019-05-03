package gosfmessagecounter

import (
	"log"
	"strconv"
	"time"

	f "github.com/ambelovsky/gosf"
)

var appStartTime int64
var totalMessageCount int64

var periodMessageCount int
var periodStartTime int64

var currTime int64
var messagesPerSecond float64

var active bool

// Plugin is the aspect oriented element required by the modular plugin framework
type Plugin struct{}

// Activate is an aspect-oriented modular plugin requirement
func (p Plugin) Activate(app *f.AppSettings) {
	active = true
	go process()
}

// Deactivate is an aspect-oriented modular plugin requirement
func (p Plugin) Deactivate(app *f.AppSettings) {
	active = false
}

// GetMessagesPerSecond returns the last recorded number of messages per second
func GetMessagesPerSecond() float64 {
	return messagesPerSecond
}

// Tick records that a single message has been processed
func Tick() {
	totalMessageCount++
	periodMessageCount++
}

func init() {
	totalMessageCount = 0
	periodMessageCount = 0

	appStartTime = time.Now().UnixNano()
	periodStartTime = time.Now().UnixNano()

	// Register hooks
	f.OnAfterResponse(func(client *f.Client, request *f.Request, response *f.Message) {
		Tick()
	})
}

// Process is a long-running process that is kicked off when this plugin is activated
func process() {
	for active {
		time.Sleep(10 * time.Second)

		currTime = time.Now().UnixNano()
		messagesPerSecond = float64(periodMessageCount) / (float64((currTime - periodStartTime)) / 1000000000)

		ConsoleClear()
		log.Println("Messages Per Second: " + strconv.FormatFloat(messagesPerSecond, 'f', 2, 64))

		periodStartTime = time.Now().UnixNano()
		periodMessageCount = 0
	}
}
