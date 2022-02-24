#!/bin/bash

grpcurl \
    -plaintext \
    -proto /Users/vkhristenko/soft/test-go/test-grpc/quickstart/codegen/test.proto \
    -import-path ../ \
    localhost:12345 \
    quickstart.codegen.TestInterface.Test
