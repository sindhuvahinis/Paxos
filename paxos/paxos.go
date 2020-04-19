package paxos

import (
	"../proto"
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type OperationValue struct {
	Key         string
	Value       string
	OperationId int
}

type PaxosStatus int

const (
	Decided PaxosStatus = iota + 1
	Pending
	Empty
)

type ProcessState struct {
	status                 PaxosStatus
	promisedProposedNumber string
	acceptedNumber         string
	acceptedValue          string
	acceptedKey            string
	operationId            int
}

type Paxos struct {
	peerPorts    []int
	PeerClients  map[int]proto.PaxosServiceClient
	maxProcessId int
	mutex        sync.Mutex
	Processes    map[int]*ProcessState
	MajoritySize int
	MyPort       int
}

func (p *Paxos) Prepare(ctx context.Context, request *proto.PrepareRequest) (*proto.Promise, error) {

	canPromise := false

	// Checking if process exist or not
	processId := int(request.ProcessId)
	if _, isExist := p.Processes[processId]; isExist {
		if p.Processes[processId].promisedProposedNumber < request.ProposeNumber {
			canPromise = true
		}
	} else {
		p.Processes[processId] = &ProcessState{
			status:                 Empty,
			promisedProposedNumber: "",
			acceptedNumber:         "",
			acceptedValue:          "",
		}
		canPromise = true
	}

	log.Printf("Can Promise value %v ", canPromise)
	if canPromise {
		p.Processes[processId].promisedProposedNumber = request.ProposeNumber
		return &proto.Promise{
			IsPromised:            true,
			ProposeNumber:         request.ProposeNumber,
			AcceptedProposeNumber: p.Processes[processId].acceptedNumber,
			AcceptedKey:           p.Processes[processId].acceptedKey,
			AcceptedValue:         p.Processes[processId].acceptedValue,
			OperationId:           int64(p.Processes[processId].operationId),
		}, nil
	}

	return &proto.Promise{
		IsPromised:            false,
		ProposeNumber:         request.ProposeNumber,
		AcceptedProposeNumber: p.Processes[processId].acceptedNumber,
		AcceptedKey:           p.Processes[processId].acceptedKey,
		AcceptedValue:         p.Processes[processId].acceptedValue,
		OperationId:           int64(p.Processes[processId].operationId),
	}, nil
}

func (p *Paxos) Accept(ctx context.Context, request *proto.AcceptRequest) (*proto.Accepted, error) {

	processId := int(request.ProcessId)
	p.updateMaximumProcessId(processId)
	_, isExist := p.Processes[processId]

	if !isExist {
		p.Processes[processId] = &ProcessState{
			status:                 0,
			promisedProposedNumber: "",
			acceptedNumber:         "",
			acceptedValue:          "",
			acceptedKey:            "",
			operationId:            0,
		}
	}

	isAccepted := false
	proposedNumber := request.ProposeNumber
	if proposedNumber >= p.Processes[processId].promisedProposedNumber {
		p.Processes[processId].promisedProposedNumber = request.ProposeNumber
		p.Processes[processId].acceptedNumber = request.ProposeNumber
		p.Processes[processId].acceptedKey = request.MaxKey
		p.Processes[processId].acceptedValue = request.MaxValue
		p.Processes[processId].status = Pending
		isAccepted = true
	}

	return &proto.Accepted{
		ProposeNumber:         p.Processes[processId].promisedProposedNumber,
		AcceptedProposeNumber: p.Processes[processId].acceptedNumber,
		AcceptedKey:           p.Processes[processId].acceptedKey,
		AcceptedValue:         p.Processes[processId].acceptedValue,
		IsAccepted:            isAccepted,
	}, nil
}

func (p *Paxos) MarkDecided(ctx context.Context, request *proto.DecidedRequest) (*proto.DecidedResponse, error) {

	processId := int(request.ProcessId)

	_, exist := p.Processes[processId]
	if exist {
		p.Processes[processId].operationId = int(request.OperationId)
		p.Processes[processId].acceptedNumber = request.ProposeNumber
		p.Processes[processId].acceptedKey = request.MaxKey
		p.Processes[processId].acceptedValue = request.MaxValue
		p.Processes[processId].status = Decided
	} else {
		p.Processes[processId] = &ProcessState{
			status:                 Decided,
			promisedProposedNumber: request.ProposeNumber,
			acceptedNumber:         request.ProposeNumber,
			acceptedValue:          request.MaxValue,
			acceptedKey:            request.MaxKey,
		}
	}
	if processId > p.maxProcessId {
		p.maxProcessId = processId
	}

	return &proto.DecidedResponse{
		IsMarkedDecided: true,
	}, nil
}

func (p *Paxos) GetMaxProcessId() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.maxProcessId
}

func (p *Paxos) Status(processId int) (PaxosStatus, OperationValue) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if process, isExist := p.Processes[processId]; isExist {
		return process.status, OperationValue{
			Key:         p.Processes[processId].acceptedKey,
			Value:       p.Processes[processId].acceptedValue,
			OperationId: p.Processes[processId].operationId,
		}
	} else {
		return Empty, OperationValue{}
	}
}

func (p *Paxos) InitiatePaxos(processId int, value OperationValue) {

	fmt.Println()
	log.Printf(" PAXOS HAS BEEN INITIATED FOR THE PROCESS ID %d ", processId)
	isPaxosDecided := false
	for !isPaxosDecided {
		proposalNumber := p.GetProposalNumber()

		log.Printf("Peers are %v ", p.PeerClients)

		noOfServerPrepareReceived, maximumValue := p.sendPrepare(processId, proposalNumber, value)
		log.Printf("No of Prepared servers are %d ", noOfServerPrepareReceived)

		if noOfServerPrepareReceived >= p.MajoritySize {
			noOfAcceptedServers := p.sendAccept(processId, proposalNumber, maximumValue)

			log.Printf("No of Accepted servers are %d ", noOfAcceptedServers)
			if noOfAcceptedServers >= p.MajoritySize {
				p.sendDecided(processId, proposalNumber, maximumValue)
				//log.Printf("breaking the forever loop")
				isPaxosDecided = true
			} else {
				time.Sleep(30 * time.Millisecond)
			}
		} else {
			time.Sleep(30 * time.Millisecond)
		}
	}

}

func (p *Paxos) GetProposalNumber() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func (p *Paxos) sendPrepare(processId int, proposeNumber string, value OperationValue) (int, OperationValue) {

	noOfServerPrepareReceived := 0
	maximumProposeNumber := ""
	maximumValue := value
	log.Printf("Prepare method is visited... ")
	for _, peer := range p.PeerClients {
		promise, err := peer.Prepare(context.Background(), &proto.PrepareRequest{ProcessId: int64(processId), ProposeNumber: proposeNumber})

		if err != nil {
			log.Printf("Error happened while connecting to peers in Prepare phase %v ", err)
			panic("")
		}

		if promise.IsPromised {
			if promise.AcceptedProposeNumber > maximumProposeNumber {
				maximumProposeNumber = promise.AcceptedProposeNumber
				maximumValue = OperationValue{
					Key:         promise.AcceptedKey,
					Value:       promise.AcceptedValue,
					OperationId: int(promise.OperationId),
				}
			}
			noOfServerPrepareReceived += 1
		}
	}
	return noOfServerPrepareReceived, maximumValue
}

func (p *Paxos) sendAccept(processId int, proposeNumber string, maximumValue OperationValue) int {
	noOfAcceptedServers := 0

	randomPeerPortToFail := p.getRandomPorts()
	for port, peer := range p.PeerClients {

		if randomPeerPortToFail[port] {
			log.Printf("Acceptor is not active. Hence accept request is not sent to this server for port number %d", port)
			continue
		}

		accepted, _ := peer.Accept(context.Background(), &proto.AcceptRequest{
			ProcessId:     int64(processId),
			ProposeNumber: proposeNumber,
			MaxKey:        maximumValue.Key,
			MaxValue:      maximumValue.Value,
		})

		if accepted.IsAccepted {
			log.Printf("Accept request has been accepted by the server with port number %d", port)
			noOfAcceptedServers += 1
		} else {
			log.Printf("Proposal has been accepted by the server with port number %d", port)
		}
	}

	return noOfAcceptedServers
}

func (p *Paxos) sendDecided(processId int, proposeNumber string, maximumValue OperationValue) {
	for _, peer := range p.PeerClients {
		decidedResponse, _ := peer.MarkDecided(context.Background(), &proto.DecidedRequest{
			ProcessId:     int64(processId),
			ProposeNumber: proposeNumber,
			MaxKey:        maximumValue.Key,
			MaxValue:      maximumValue.Value,
			OperationId:   int64(maximumValue.OperationId),
		})

		if decidedResponse.IsMarkedDecided {
			log.Printf("Decided for the propose number %d", processId)
		}
	}
}

func (p *Paxos) updateMaximumProcessId(processId int) {
	if processId > p.maxProcessId {
		p.maxProcessId = processId
	}
}

func (p *Paxos) getRandomPorts() map[int]bool {
	randomIndex := rand.Int() % 3

	//fmt.Printf("Random Index = %v\n", randomIndex)
	peerRandomSet := make(map[int]bool)
	for i := 0; i < randomIndex; i++ {
		peerPort := p.peerPorts[i]
		if peerPort != p.MyPort {
			//fmt.Printf("Random Port Number %d is %d ", i, peerPort)
			peerRandomSet[peerPort] = true
		} else {
			i--
		}
	}

	return peerRandomSet
}
