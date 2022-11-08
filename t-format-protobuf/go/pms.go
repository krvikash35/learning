package main

import (
	pb "go-proto/proto.person"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var file_name = "address_boook_proto"

func add_address_book() *pb.AddressBook {
	log.Println("adding address book")
	p1 := pb.Person{
		Id:    1,
		Name:  "vikash",
		Email: "v@gmail",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123", Type: pb.Person_HOME},
		},
		LastUpdated: timestamppb.Now(),
	}
	p2 := pb.Person{
		Id:    2,
		Name:  "vikash",
		Email: "v@gmail",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123", Type: pb.Person_HOME},
		},
	}

	ab := pb.AddressBook{
		People: []*pb.Person{
			&p1,
			&p2,
		},
	}
	log.Printf("adding address book done %+v", ab)
	return &ab
}

func write_to_file() {
	log.Println("writing to file")
	ab := add_address_book()

	bytes, err := proto.Marshal(ab)
	if err != nil {
		log.Fatalln("error while marshling addressbook", err)
	}

	err = os.WriteFile(file_name, bytes, 0666)
	if err != nil {
		log.Fatalln("error while writing to file", file_name, err)
	}
	log.Println("writing to file done")
}

func read_from_file() {
	log.Println("reading from file")
	ab := &pb.AddressBook{}

	bytes, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatalln("error while reading from file", file_name, err)
	}

	err = proto.Unmarshal(bytes, ab)
	if err != nil {
		log.Fatalf("error while unmarshing address book %s", err)
	}
	log.Printf("reading from file DONE %+v", ab)
}

func main() {
	add_address_book()
	write_to_file()
	read_from_file()
}
