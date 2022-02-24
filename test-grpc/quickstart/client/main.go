package main

import (
    "fmt"
    "flag"
    "log"
    "context"
    "time"

    "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
    cg "github.com/vkhristenko/test-go/test-grpc/quickstart/codegen"
)

func startClient(port uint16) {
    fmt.Printf("connecting to the server on port %d\n", port)

    // connec to the server
    full_address := fmt.Sprintf("localhost:%d", port)
    conn, err := grpc.Dial(
        full_address, 
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect to %v", err)
    }

    // create a new client
    defer conn.Close()
    client := cg.NewTestInterfaceClient(conn)

    // do one rpc call
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := client.Test(ctx, &cg.TestRequest{})
    if err != nil {
        log.Fatalf("rpc failed %v", err)
    }
    log.Printf("all good! %s", r.String())
}

var (
    port = flag.Uint("port", 12345, "server port")
)

func main() {
    flag.Parse()
    startClient(uint16(*port))
}
