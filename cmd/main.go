package main

import (
	"fmt"
	"log"
	"net"

	"cloudbees/dao/inmem"
	cloudbeespb "cloudbees/proto"
	"cloudbees/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	log.Println("Init datastores..")
	ticketDatastore := inmem.NewTickets()
	seatDataStore := inmem.NewSeat()
	userDatastore := inmem.NewUserStore()

	ticketService := service.NewTicketService(ticketDatastore, userDatastore, seatDataStore)
	seatService := service.NewSeatService(seatDataStore, ticketDatastore, userDatastore)
	userService := service.NewUserService(userDatastore, ticketDatastore, seatDataStore)

	cloudbeespb.RegisterTicketingServer(server, ticketService)
	cloudbeespb.RegisterSeatingServer(server, seatService)
	cloudbeespb.RegisterUserServiceServer(server, userService)
	reflection.Register(server)
	log.Println("starting server... listening on", fmt.Sprintf("localhost:%d", 8080))
	err = server.Serve(lis)
	if err != nil {
		return
	}
}
