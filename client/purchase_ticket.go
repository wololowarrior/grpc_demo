package main

import (
	"context"
	"log"
	"time"

	cloudbeespb "cloudbees/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	ticketingClient := cloudbeespb.NewTicketingClient(conn)
	seatingC := cloudbeespb.NewSeatingClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := ticketingClient.Purchase(ctx, &cloudbeespb.TicketRequest{
		Ticket: &cloudbeespb.TicketDetails{
			Source:      "London",
			Destination: "Paris",
			Price:       20,
		},
		User: &cloudbeespb.User{
			Email:     "abc",
			FirstName: "Harshil",
			LastName:  "Gupta",
		},
	})
	if err != nil {
		log.Fatalf("could not purchase ticket: %v", err)
	}
	log.Printf("Ticket purchased: %s", r)
	section := "A"
	r1, err := seatingC.Allocate(ctx, &cloudbeespb.AllocationRequest{
		TicketId: r.GetTicket().GetId(),
		Section:  &section,
	})
	if err != nil {
		log.Fatalf("could not allocate ticket: %v", err)
	}
	log.Printf("Seat allocated: %s", r1)

	r2, err := seatingC.List(ctx, &cloudbeespb.ListSeatsRequest{Section: section})
	if err != nil {
		log.Fatalf("could not get ticket: %v", err)
	}
	log.Printf("List of seats: %s", r2)

	section = "B"
	r3, err := seatingC.Modify(ctx, &cloudbeespb.ModifySeatRequest{
		Section:  &section,
		TicketId: r.GetTicket().GetId(),
	})
	if err != nil {
		log.Fatalf("could not modify seat: %v", err)
	}
	log.Printf("New seat: %s", r3)

	r4, err := ticketingClient.GetReceipt(ctx,
		&cloudbeespb.User{Email: "abc"})

	if err != nil {
		log.Fatalf("could not get receipt: %v", err)
	}
	log.Printf("Receipt: %s", r4)

	r2, err = seatingC.List(ctx, &cloudbeespb.ListSeatsRequest{Section: section})
	if err != nil {
		log.Fatalf("could not get ticket: %v", err)
	}
	log.Printf("List of seats: %s", r2)
}
