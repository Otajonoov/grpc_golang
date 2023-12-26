package main

import (
	"google.golang.org/grpc"
	pb "grpc_golang/server_streaming/protofiles/data_streaming"
	"log"
	"math/rand"
	"net"
)

type server struct {
	pb.UnimplementedStreamingServiceServer
}

// *DataRequest, StreamingService_GetDataStreamingServer) error
func (s server) GetDataStreaming(req *pb.DataRequest, srv pb.StreamingService_GetDataStreamingServer) error {
	log.Println("Fetch data streaming")

	for i := 0; i < 10; i++ {
		value := randStringBytes(50)

		resp := pb.DataResponse{
			Part:   int32(i),
			Buffer: value,
		}

		if err := srv.Send(&resp); err != nil {
			log.Println("error generating response")
			return err
		}
	}

	return nil
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func main() {
	// create listener
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic("error building server: " + err.Error())
	}

	// create gRPC server
	s := grpc.NewServer()

	// registering the server implementation with the gRPC server
	pb.RegisterStreamingServiceServer(s, server{})

	log.Println("start server")

	if err := s.Serve(listener); err != nil {
		panic("error building server: " + err.Error())
	}
}
