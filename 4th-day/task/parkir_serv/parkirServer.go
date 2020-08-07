package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	en "task/entity"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type Parked struct {
	ParkID  string `json:"id"`
	Vehicle string `json:"vehicle"`
	Plate   string `json:"plate"`
}

type ParkBill struct {
	Vehicle  string `json:"vehicle"`
	Plate    string `json:"plate"`
	ParkTime string `json:"entry"`
	ExitTime string `json:"exit"`
	Bill     int32  `json:"bill"`
}

type server struct {
	en.UnimplementedConnectServer
}

func (s *server) ConnectToServer(ctx context.Context, void *empty.Empty) (*en.ServerReply, error) {
	return &en.ServerReply{Message: "Client is connected to Server"}, nil
}

func (s *server) GetParkID(ctx context.Context, void *empty.Empty) (*en.ParkID, error) {
	generateID := parkIn()
	return &en.ParkID{Id: generateID}, nil
}
func (s *server) GetParkFee(ctx context.Context, parkExit *en.Parked) (*en.ParkBill, error) {
	entryTime, exitTime, parkingFee := parkOut(parkExit.GetId(), parkExit.GetVehicle(), parkExit.GetPlate())
	result := en.ParkBill{Vehicle: parkExit.GetVehicle(), Plate: parkExit.GetPlate(), Entry: entryTime, Exit: exitTime, Bill: parkingFee}

	return &result, nil
}

var park string
var parkTimes = make([]int64, 0)
var parkList = make([]string, 0)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	en.RegisterConnectServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func parkIn() string {
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 30000
	park := strconv.Itoa(rand.Intn(max-min) + min)
	parkTimes = append(parkTimes, time.Now().Unix())
	parkList = append(parkList, park)
	return park
}

func parkOut(vType, vPlate, parkID string) (string, string, int32) {
	var parkTime int64
	var currentTime int64
	var duration int64
	var totalFee int32 = 0
	var bool = false

	for key, val := range parkList {
		if parkID == val {
			bool = true
			parkTime = parkTimes[key]
			// fmt.Println("parking", parkTimes)
			removeParkTimes(&parkTimes, key)
			removeParkID(&parkList, key)

			currentTime = time.Now().Unix()
			duration = currentTime - parkTime
			switch vType {
			case "car":

				for i := 1; i <= int(duration); i++ {
					if i == 1 {
						totalFee += 5000
						continue
					}
					totalFee += 3000
				}
			case "bike":

				for i := 1; i <= int(duration); i++ {
					if i == 1 {
						totalFee += 3000
						continue
					}
					totalFee += 2000
				}
			default:

				return "", "", 0
			}
		}
	}

	if !bool {
		return "", "", 0
	}

	entry := timeParsing(strconv.Itoa(int(parkTime))).String()
	exit := timeParsing(strconv.Itoa(int(currentTime))).String()

	return entry, exit, totalFee
}

func timeParsing(s string) time.Time {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm
}

func removeParkTimes(arr *[]int64, index int) {
	newArr := *arr
	*arr = append(newArr[:index], newArr[index+1:]...)
}
func removeParkID(arr *[]string, index int) {
	newArr := *arr
	*arr = append(newArr[:index], newArr[index+1:]...)
}
