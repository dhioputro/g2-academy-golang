package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	en "task/entity"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

var park string
var parkTimes = make([]int64, 0)
var parkList = make([]string, 0)

func main() {
	var input string

	for input != "99" {
		fmt.Println("1. Parking Entry")
		fmt.Println("2. Parking Out")
		fmt.Println("99. Done")
		fmt.Println("What to choose?")
		fmt.Scanln(&input)

		switch input {
		case "1":
			parkIn()

		case "2":
			var vType string
			var vPlate string
			var parkID string
			fmt.Println("Vehicle Type? (Car/Bike)")
			fmt.Scanln(&vType)
			fmt.Println("Vehicle plate?")
			fmt.Scanln(&vPlate)
			fmt.Println("Parking ID : ")
			fmt.Scanln(&parkID)
			parkOut(vType, vPlate, parkID)

		case "99":
			fmt.Println("Thank You")

		default:
			fmt.Println("You enter the wrong input")

		}
	}
}

func parkIn() {
	conn := ConnectServer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getPark, err := conn.GetParkID(ctx, new(empty.Empty))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Your Park Id: %s", getPark.GetId())
}

func parkOut(vType, vPlate, parkID string) {
	conn := ConnectServer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	parked := en.Parked{Id: parkID, Vehicle: vType, Plate: vPlate}

	getBill, err := conn.GetParkBill(ctx, &parked)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	Print(getBill)
	log.Printf(Print(getBill))
}

// connection setup
func ConnectServer() en.ConnectClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	return en.NewConnectClient(conn)
}

func Print(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
