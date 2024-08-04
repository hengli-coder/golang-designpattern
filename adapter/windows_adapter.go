package adapter

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

// InsertIntoLightningPort is a method that allows the WindowsAdapter to insert into the Lightning port
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
