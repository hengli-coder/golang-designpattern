package adapter

import "fmt"

// client this client can only insert Lightning connector into computer
type client struct{}

func (c *client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
