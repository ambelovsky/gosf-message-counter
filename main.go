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
}

// AppMethods is a struct optionally exposed to clients after the plugin has been registered
type AppMethods struct{}

// Tick adds 1 to the message count
func (a AppMethods) Tick() {
	Tick()
}

// Plugin is the aspect oriented element required by the modular plugin framework
type Plugin struct{}

// Activate is an aspect-oriented modular plugin requirement
func (p Plugin) Activate(app *map[string]interface{}) {
	active = true
	(*app)["message-counter"] = new(AppMethods)
	go process()
}

// Deactivate is an aspect-oriented modular plugin requirement
func (p Plugin) Deactivate(app *map[string]interface{}) {
	active = false
}

// PreReceive is an aspect-oriented modular plugin requirement
func (p Plugin) PreReceive(clientMessage *f.Message) {

}

// PostReceive is an aspect-oriented modular plugin requirement
func (p Plugin) PostReceive(clientMessage *f.Message) {

}

// PreRespond is an aspect-oriented modular plugin requirement
func (p Plugin) PreRespond(clientMessage *f.Message, serverMessage *f.Message) {

}

// PostRespond is an aspect-oriented modular plugin requirement
func (p Plugin) PostRespond(clientMessage *f.Message, serverMessage *f.Message) {
	Tick()
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
