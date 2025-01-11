package main

import (
	"fmt"
	"log"
	"net"

	"cloudbees/dao/inmem"
	cloudbeespb "cloudbees/proto"
	"cloudbees/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	ticketDatastore := inmem.NewTickets()
	seatDataStore := inmem.NewSeat()
	userDatastore := inmem.NewUserStore()
	ticketService := service.NewTicketService(ticketDatastore)
	seatService := service.NewSeatService(seatDataStore)
	userService := service.NewUserService(userDatastore)
	cloudbeespb.RegisterTicketingServer(server, ticketService)
	cloudbeespb.RegisterSeatingServer(server, seatService)
	cloudbeespb.RegisterUserServiceServer(server, userService)
	server.Serve(lis)
}
