package main

import (
	"go-ff/pkg/def/packet/packettype"
	"go-ff/pkg/service/dotenv"
	"go-ff/pkg/service/external"
	"go-ff/cmd/connectionserver/service/connectionmanager"
	"log"
)

func onConnectedHandler(ch <-chan *external.Client) {
	for {
		c := <-ch
		if c == nil {
			continue
		}

		connectionmanager.Add(c)
		c.SendGreetings()
	}
}

func onDisconnectedHandler(ch <-chan *external.Client) {
	for {
		c := <-ch
		if c == nil {
			continue
		}

		connectionmanager.Remove(c)
	}
}

func onMessageHandler(ch <-chan *external.PacketHandler) {
	for {
		p := <-ch
		if p == nil {
			continue
		}

		c := connectionmanager.Get(p.ClientID)
		if c == nil {
			continue
		}

		id := p.Packet.ReadUInt32()
		if id == packettype.Certify {
			sendServerList(c)
		}
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	dotenv.Initialize()

	// External ----
	onConnected := make(chan *external.Client)
	onDisconnected := make(chan *external.Client)
	onMessage := make(chan *external.PacketHandler)

	go onConnectedHandler(onConnected)
	go onDisconnectedHandler(onDisconnected)
	go onMessageHandler(onMessage)

	server := external.Create("0.0.0.0:23000")
	server.OnConnected(onConnected)
	server.OnDisconnected(onDisconnected)
	server.OnMessage(onMessage)
	server.Start()
}