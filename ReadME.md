#SDS Project 4
##### Paxos implementation with Key Value store

To run this you need to install go, protobuf and grpc in your machine. 

Please install go, in Mac by default it should be available. So I am assuming it will be available

To Install Protobuf please run the following. 

`go get github.com/golang/protobuf/proto`

To Install GRPC in your machine, please run the following:

 `go get -u google.golang.org/grpc`

Now please download the zip. So far I have seen, platform independent go executable needs external library. If I find a better solution, I will mail you guys. Sorry for the trouble.


To run the Server:

```
cd DSProject4/server/main/

go run main.go -port=5000 -peerports=5001,5002,5003,5004
```

To run the client

```
cd DSProject4/client

go run main.go localhost 5000
```

Note: 

Acceptor will randomly fail as mentioned in the requirement documents and printed in the document.

If the majority is not acheived, then paxos will keep on trying till consensus is achieved. 