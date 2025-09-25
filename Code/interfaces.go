package main

type gasEngineNew struct {
	mpg     uint
	gallons uint
}

func (g gasEngineNew) start() {
	println("Gas engine started")
}

func (g gasEngineNew) stop() {
	println("Gas engine stopped")
}

type engine interface {
	start()
	stop()
}

func drive(e engine) {
	e.start()
	println("Driving...")
	e.stop()
}

// testDrive demonstrates the use of the engine interface with a gasEngine
func test_interface() {
	g := gasEngineNew{mpg: 25, gallons: 10}
	drive(g)
}
