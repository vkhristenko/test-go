package main

import (
    "fmt"
    "flag"
    "net"
    "log"
    "context"

    "google.golang.org/grpc"
    cg "github.com/vkhristenko/test-go/test-grpc/quickstart/codegen"
)

var (
    port = flag.Uint("port", 12345, "server port")
)

type server struct {
    cg.UnimplementedTestInterfaceServer
}

func (s *server) Test(
        ctx context.Context, 
        request *cg.TestRequest) (*cg.TestReply, error) {
    log.Printf("a new request")
    return &cg.TestReply{}, nil
}

func startServer(port uint16) {
    fmt.Printf("starting a server on port %d\n", port)

    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("failed to listen on %v", err)
    }

    s := grpc.NewServer()
    cg.RegisterTestInterfaceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func main() {
    flag.Parse()
    startServer(uint16(*port))
}
