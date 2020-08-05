package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var park string
var parkTimes = make([]int64, 0)
var parkList = make([]string, 0)

func main() {
	var input string

	for input != "99" {
		fmt.Println("1. Parking Entry")
		fmt.Println("2. Parking Out")
		fmt.Println("3. Parking ID List")
		fmt.Println("99. Done")
		fmt.Println("What to choose?")
		fmt.Scanln(&input)

		switch input {
		case "1":
			fmt.Println("Welcome Customer! Parking ID:", parkIn(), "\n")

		case "2":
			var carType string
			var carPlate string
			var parkID string
			fmt.Println("Vehicle Type? (Car/Bike)")
			fmt.Scanln(&carType)
			fmt.Println("Vehicle plate?")
			fmt.Scanln(&carPlate)
			fmt.Println("Parking ID : ")
			fmt.Scanln(&parkID)
			fmt.Println(parkOut(carType, carPlate, parkID))

		case "3":
			if len(parkList) > 0 {
				fmt.Println(parkList)
			} else {
				fmt.Println("No parking vehicle right now \n")
			}

		case "99":
			fmt.Println("Thank You")

		default:
			fmt.Println("Opsi yang anda masukan salah")

		}
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

func parkOut(carType, carPlate, parkID string) string {
	var parkTime int64
	var currentTime int64
	var totalFee int = 0
	var bool = false

	for input, val := range parkList {
		if parkID == val {
			bool = true
			parkTime = parkTimes[input]
			// fmt.Println("parking", parkTimes)
			RemoveParkTimes(&parkTimes, input)
			RemoveParkID(&parkList, input)

			currentTime = time.Now().Unix() - parkTime
			switch carType {
			case "car":

				for i := 1; i <= int(currentTime); i++ {
					if i == 1 {
						totalFee += 5000
						continue
					}
					totalFee += 3000
				}
			case "bike":

				for i := 1; i <= int(currentTime); i++ {
					if i == 1 {
						totalFee += 3000
						continue
					}
					totalFee += 2000
				}
			default:

				return "Plese enter valid type!"
			}

		}
	}

	if !bool {
		return "Parking ID " + parkID + " is not found"
	}

	return "Park Time: " + strconv.Itoa(int(currentTime)) + " second(s)\n" + "Total Fee: " + strconv.Itoa(totalFee) + "\n"
}

func RemoveParkTimes(arr *[]int64, index int) {
	newArr := *arr
	*arr = append(newArr[:index], newArr[index+1:]...)
}
func RemoveParkID(arr *[]string, index int) {
	newArr := *arr
	*arr = append(newArr[:index], newArr[index+1:]...)
}
