package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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
	Bill     int    `json:"bill"`
}

var park string
var parkTimes = make([]int64, 0)
var parkList = make([]string, 0)

func main() {

	http.HandleFunc("/entrypark", parkingID)
	http.HandleFunc("/exitpark", parkingOut)

	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func parkingID(w http.ResponseWriter, r *http.Request) {
	generateID := parkIn()
	fmt.Fprintf(w, generateID)
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

func parkingOut(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var parkExit Parked
	err = json.Unmarshal(reqBody, &parkExit)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	entryTime, exitTime, parkingFee := parkOut(parkExit.ParkID, parkExit.Vehicle, parkExit.Plate)
	result := ParkBill{parkExit.Vehicle, parkExit.Plate, entryTime, exitTime, parkingFee}

	js, _ := json.Marshal(result)
	w.Write([]byte(js))
}

func parkOut(vType, vPlate, parkID string) (string, string, int) {
	var parkTime int64
	var currentTime int64
	var duration int64
	var totalFee int = 0
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
