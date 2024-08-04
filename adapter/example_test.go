package adapter

func Example() {
	client := &client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
	// Output:
	// Client inserts Lightning connector into computer.
	// Lightning connector is plugged into mac machine.
	// Client inserts Lightning connector into computer.
	// Adapter converts Lightning signal to USB.
	// USB connector is plugged into windows machine.
}
