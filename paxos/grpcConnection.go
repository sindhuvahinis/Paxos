package paxos

import (
	"../proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func ConnectToPeers(peerPorts []int, myPort int) *Paxos {
	paxos := &Paxos{}
	paxos.peerPorts = peerPorts
	noOfPeers := len(peerPorts)
	paxos.MyPort = myPort
	paxos.PeerClients = make(map[int]proto.PaxosServiceClient)

	for _, peer := range peerPorts {
		connection, err := grpc.Dial("localhost"+fmt.Sprintf(":%d", peer), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Error occurred while establishing the connection with the Peer: %v ", err)
		}

		client := proto.NewPaxosServiceClient(connection)
		paxos.PeerClients[peer] = client
	}

	paxos.Processes = make(map[int]*ProcessState)
	paxos.MajoritySize = noOfPeers/2 + 1

	return paxos
}
