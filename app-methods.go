package gosfmessagecounter

// AppMethods is a struct optionally exposed to clients to use after the plugin has been registered
type AppMethods struct{}

// Tick adds 1 to the message count
func (a AppMethods) Tick() {
	Tick()
}
