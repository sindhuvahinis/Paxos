package wrapper

import (
	"../../paxos"
	"../../proto"
	responseCode "../../util"
	"../KeyValueStore"
	"context"
	"fmt"
	"log"
	"time"
)

type Server struct {
	CurrentProcessId int
	Paxos            *paxos.Paxos
	KeyValueStore    *KeyValueStore.KeyValue
}

const (
	Put int = iota + 1
	Get
	Delete
)

func (s *Server) Initialize(p *paxos.Paxos) {
	s.Paxos = p
	s.CurrentProcessId = 0
	s.KeyValueStore = &KeyValueStore.KeyValue{}
	s.KeyValueStore.Initialize()
}

func (s *Server) synchronizeWithPeers(maxProcessId int) {

	log.Printf("Synchronize with peers is called... Current process id %d MaxProcessID %d ", s.CurrentProcessId, maxProcessId)

	for s.CurrentProcessId <= maxProcessId {
		status, value := s.Paxos.Status(s.CurrentProcessId)

		if status == paxos.Empty {
			log.Printf("Paxos is Empty.. ")
			s.startPaxosProcess(s.CurrentProcessId, paxos.OperationValue{})
			status, value = s.Paxos.Status(s.CurrentProcessId)
		}

		for {
			if status == paxos.Decided {
				if value.OperationId == Put {
					s.KeyValueStore.PUT(value.Key, value.Value)
				} else if value.OperationId == Delete {
					s.KeyValueStore.DELETE(value.Key)
				}
				break
			}

			time.Sleep(10 * time.Millisecond)
			status, value = s.Paxos.Status(s.CurrentProcessId)
		}

		s.CurrentProcessId += 1
	}

	log.Printf("Synchronize is over... ")
}

func (s *Server) PUT(ctx context.Context, request *proto.PutRequest) (*proto.Response, error) {
	log.Printf("PUT call is received from the client")

	log.Printf("Peers in PUT are %v ", s.Paxos.PeerClients)
	value := paxos.OperationValue{
		Key:         request.Key,
		Value:       request.Value,
		OperationId: Put,
	}

	s.connectWithPeersForAgreement(value)
	return &proto.Response{
		ResponseCode: responseCode.SUCCESS,
		Message:      fmt.Sprintf("PUT is successful. Key-Value pair %v - %v is updated successfully", request.Key, request.Value),
	}, nil
}

func (s *Server) GET(ctx context.Context, request *proto.GetAndDeleteRequest) (*proto.Response, error) {

	key := request.Key

	if key == "" {
		return &proto.Response{
			ResponseCode: responseCode.INVALID_INPUT,
			Message:      "Key cannot be null or empty. GET is not initiated",
		}, nil
	}

	valueToBeSent := paxos.OperationValue{
		Key:         key,
		Value:       "",
		OperationId: Get,
	}

	s.connectWithPeersForAgreement(valueToBeSent)
	value, isKeyExist := s.KeyValueStore.GET(key)

	if !isKeyExist {
		return &proto.Response{
			ResponseCode: responseCode.INVALID_INPUT,
			Message:      "Value is not found for the given key",
		}, nil
	}

	return &proto.Response{
		ResponseCode: responseCode.SUCCESS,
		Message:      fmt.Sprintf("GET has been successful. Value for the key %s is %s", key, value),
	}, nil

}

func (s *Server) DELETE(ctx context.Context, request *proto.GetAndDeleteRequest) (*proto.Response, error) {
	log.Printf("DELETE call is received from the client")

	valueToBeSent := paxos.OperationValue{
		Key:         request.Key,
		Value:       "",
		OperationId: Delete,
	}

	s.connectWithPeersForAgreement(valueToBeSent)
	return &proto.Response{
		ResponseCode: responseCode.SUCCESS,
		Message:      fmt.Sprintf("DELETE is successful. Key-Value pair %v is deleted", request.Key),
	}, nil
}

/************************ PRIVATE METHODS ************************/

func (s *Server) startPaxosProcess(processId int, value paxos.OperationValue) paxos.OperationValue {
	s.Paxos.InitiatePaxos(processId, value)
	var decidedValue paxos.OperationValue

	isDecided := false

	for !isDecided {
		status, value := s.Paxos.Status(processId)

		log.Printf("Paxos status right now.. %v", status)
		if status == paxos.Decided {
			decidedValue = value
			isDecided = true

			fmt.Println()
			log.Printf("PAXOS IS SUCCESSFULLY COMPLETED FOR THE PROCESS ID %d...!", processId)
		}
		time.Sleep(10 * time.Millisecond)

		log.Printf("PAXOS IS NOT DECIDED FOR THE PROCESS ID %d . TRYING AGAIN...! ", processId)
	}

	log.Printf("Start Paxos is over...")
	return decidedValue
}

func (s *Server) connectWithPeersForAgreement(value paxos.OperationValue) {
	processId := s.Paxos.GetMaxProcessId() + 1
	s.synchronizeWithPeers(processId - 1)
	agreementValue := s.startPaxosProcess(processId, value)

	log.Printf("Values Key %v Value %v Operation %v ", value.Key, value.Value, value.OperationId)
	log.Printf("Agreement Key %v Value %v Operation %v ", agreementValue.Key, agreementValue.Value, agreementValue.OperationId)
	if agreementValue.Key == value.Key && agreementValue.Value == value.Value && agreementValue.OperationId == value.OperationId {
		s.synchronizeWithPeers(processId)
	}
}
