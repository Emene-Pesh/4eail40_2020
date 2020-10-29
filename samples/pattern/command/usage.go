package main

func main() {
	// Receiver
	lamp := new(Light)

	// ConcreteCommand
	switchOn := &SwitchOnCommand{Light: lamp}
	switchOff := &SwitchOffCommand{Light: lamp}

	// Invoker
	mySwitch := new(Switch)
	mySwitch.StoreAndExecute(switchOn)
	mySwitch.StoreAndExecute(switchOff)
}
