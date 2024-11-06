package main

import (
	"fmt"
)

func main() {
	//router := gin.Default()
	//router.GET("test", getAlbums)
	//router.POST("test", postAlbums)

	//err := router.Run("localhost:8080")
	var numberA = 4
	var numberB = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	var test = make(map[string]Car)
	test["brand"] = Car{
		Name:  "test",
		Color: "Red",
		Price: 1000,
	}

	fmt.Println(test)

	cars := []Car{
		{
			Name:  "BMW",
			Price: 200,
			Color: "Blue",
		},
		{
			Name:  "BMW",
			Price: 200,
			Color: "Blue",
		},
		{
			Name:  "BMW",
			Price: 2_00,
			Color: "Blue",
		},
	}

	for i := 0; i < len(cars); i++ {
		if i == len(cars)-1 {
			break
		}
		fmt.Print(addCar(cars[i].Price, cars[i+1].Price))
	}

	user := User{
		Name: "Ridho",
		Age:  12,
	}

	fmt.Println(user.Name)
	user.printUser()

	users := []User{
		{
			"ridho1",
			12,
		},
		{
			"ridho2",
			13,
		},
		{
			"ridho3",
			14,
		},
	}

	fmt.Println(users)
	//slice := users[1:2]

	//fmt.Println(users)
	//fmt.Println(slice)
	//fmt.Println(slice[0].Name)
	//fmt.Println(cap(slice))
	//
	//if users[0].Name == "ridho1" {
	//	fmt.Println(&users)
	//}
	//
	//for i := 0; i < len(users); i++ {
	//	fmt.Println(time.Now())
	//}

	iCar := ICar(Car{
		Name:  "TEST",
		Color: "WEW",
		Price: 0,
	})
	fmt.Println(iCar)
}

func (user User) printUser() {
	fmt.Println(user.Name, " ", user.Age)
}

type User struct {
	Name string
	Age  int
}
