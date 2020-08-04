package main

import "fmt"

func test() {
	var name string
	var age int = 5

	var (
		firstName string
		lastName  string
		fullname  string
	)

	var (
		address, sex string
	)

	name = "Administrator"
	firstName = "Admin"
	lastName = "Istrator"
	address = "Jakarta"
	sex = "Male"
	fullname = firstName + " " + lastName

	fmt.Println(firstName, lastName)
	fmt.Println(address, sex)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("Hello " + name)
	fmt.Println(fullname)

	var firstNo int = 300
	fmt.Println("The first no is", firstNo)

	var secondNo float64 = 12.22
	fmt.Println("The second no is", secondNo)

	var thirdNo, fourthNo float64 = 11.11, 22.22
	fmt.Println("third number", thirdNo, "fourth number", fourthNo)

	fifthNo := "28" // kalau dirubah jadi int fifthNo di bawah error
	fmt.Println("fifth number", fifthNo)

	fifthNo = "asd"
	fmt.Println("fifth number", fifthNo)

	const a int = 345
	fmt.Println("a", a)

	if a == 10 {
		fmt.Println("a sama dengan", a)
	} else if a == 20 {
		fmt.Println("a sama dengan", a)
	} else {
		fmt.Println("a bukan 10 maupun 20")
	}

	b := 10 + 10
	switch b {
	case 10:
		fmt.Println("b sama dengan", b)
	case 20:
		fmt.Println("b sama dengan", b)
	default:
		fmt.Println("b bukan 10 maupun 20")
	}

	var arrayTest [3]int
	arrayTest[0] = 10
	arrayTest[1] = 20
	arrayTest[2] = 345

	fmt.Println("index 0:", arrayTest[0])
	fmt.Println("index 1:", arrayTest[1])
	fmt.Println("index 2:", arrayTest[2])

	var multiArr [2][4]int
	multiArr[1][3] = 9
	fmt.Println("multi array:", multiArr)
	fmt.Println("multi array:", multiArr[0])
	fmt.Println("multi array [1]:", multiArr[1])
	fmt.Println("multi array [1][3]:", multiArr[1][3])

	// normal for di golang
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// while di golang
	j := 20
	for true {
		fmt.Println("j:", j)
		if j == 20 {
			break
		}
		j++
	}

	var numbers = make([]int, 3, 5)
	fmt.Println("numbers", numbers)
	fmt.Println("length", len(numbers))
	fmt.Println("capacity", cap(numbers))
	numbers = append(numbers, 1, 2, 3)
	fmt.Println("numbers", numbers)
	fmt.Println("length", len(numbers))
	fmt.Println("capacity", cap(numbers))
	fmt.Println(numbers[2:6])

	// a := [...]int{12, 78, 50, 65, 34}
	// fmt.Println("a", a)
}
