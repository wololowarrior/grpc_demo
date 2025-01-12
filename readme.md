## Summary

This project demonstrates a rudimentary ticket booking backend build on GRPc.

## How to run
Built binaries for grpc server and a sample client are kept in the `bin/` folder.

1. Run the server by
```shell
chmod u+x ./bin/grpc_server
./bin/grpc_server
```
2. Run the sample client from another term to see the flow
```shell
chmod u+x ./bin/grpc_client
./bin/grpc_client
```

You can use the sample client for exploring the project. 
Or use `grpcurl` to run the rpc. Below are the rpc listed

To install grpcurl see [this](https://github.com/fullstorydev/grpcurl?tab=readme-ov-file#installation). 
You can just download the binary or use `go install` to build from source.

## RPCs
Server reflection is enabled on the Grpc server, so you can always explore the list of rpc available. `grpcurl  -plaintext "localhost:8080" list cloudbees.Seating`

1. cloudbees.Ticketing/Purchase: This rpc will purchase a ticket for the user. Sample command is given below
```shell
grpcurl -plaintext -d '{"ticket":{"source":"a","destination":"b","price":123},"user":{"email":"abc","firstName":"harshil","lastName":"gupta"}}' 
localhost:8080 cloudbees.Ticketing/Purchase
```
2. cloudbees.Seating/Allocate: For a purchased ticket, we can call `cloudbees.Seating/Allocate` to allocate a seat for 
the passed ticket. Optionally you can also pass `section: "A"|"B"` to get a seat in a particular section
```shell
grpcurl -plaintext -d '{"ticket_id":0, "section":"B"}' localhost:8080 cloudbees.Seating/Allocate
```
3. cloudbees.Ticketing/GetReceipt: You can get a receipt which contains the summary of the itinerary
```shell
grpcurl -plaintext -d '{"ticket_id":0}' localhost:8080 cloudbees.Ticketing/GetReceipt
```
4. cloudbees.Seating/List: Get a list of allocated seat & user details per section
```shell
grpcurl -plaintext -d '{"section": "B"}' localhost:8080 cloudbees.Seating/List
```
5. cloudbees.Seating/Modify: User can modify ticket to move to other section. Pass the new section in the payload
```shell
grpcurl -plaintext -d '{"section": "B","ticket_id":0}' localhost:8080 cloudbees.Seating/Modify
```
6. cloudbees.UserService/DeleteUser: We can delete user from the train. This will deallocated seat, delete ticket and the
```shell
 grpcurl -plaintext -d '{"email":"abc"}' localhost:8080 cloudbees.UserService/DeleteUser 
```
