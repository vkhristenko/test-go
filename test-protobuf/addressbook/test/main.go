package main

import (
    "fmt"
    cg "github.com/vkhristenko/test-go/test-protobuf/addressbook/codegen"
    timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
    fmt.Println("hello")

    person := cg.Person{
        Id: 10,
        Name: "viktor",
        Email: "something@google.com",
        Phones: []*cg.Person_PhoneNumber{
            {
                Number: "012345", 
                Type: cg.Person_MOBILE,
            },
        },
        LastUpdated: &timestamppb.Timestamp{},
    }

    fmt.Println(person.String())
}
