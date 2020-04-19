package main

import (
	"../../paxos"
	"../../proto"
	sw "../wrapper"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	// create a listener on TCP port 7777
	portNumber := flag.Int("port", 0, "Server's port number")
	peerPortNumbers := getPeerPorts()

	if *portNumber == 0 {
		log.Fatal("Please enter a valid port number")
	}

	fmt.Println(*portNumber)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *portNumber))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server is up and running in the portnumber %d", *portNumber)

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	//creating server instance
	s := sw.Server{}
	p := paxos.ConnectToPeers(append(peerPortNumbers, *portNumber), *portNumber)

	s.Initialize(p)
	log.Printf("Peers in server start are %v", p.PeerClients)

	proto.RegisterKeyValueStoreServiceServer(grpcServer, &s)
	proto.RegisterPaxosServiceServer(grpcServer, p)
	reflection.Register(grpcServer)

	//start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func getPeerPorts() []int {
	var peerPortNumbersStr string
	flag.StringVar(&peerPortNumbersStr, "peerports", "", "List of port number")
	flag.Parse()

	if peerPortNumbersStr == "" {
		log.Fatal("Please enter the valid peer port numbers")
	}

	peerPortNumbersArr := strings.Split(peerPortNumbersStr, ",")

	peerPortNumbers := make([]int, 0, len(peerPortNumbersArr))

	for _, port := range peerPortNumbersArr {
		portInt, _ := strconv.Atoi(port)
		peerPortNumbers = append(peerPortNumbers, portInt)
	}
	return peerPortNumbers
}
