package main

import "fmt"

// type user struct {
// 	name string
// 	age int
// 	isLoggedIn bool
// 	greet func() 
// }


type user struct {
	name string
	age int
	isLoggedIn bool
}


func main() {
	
	// user1 := user{
	// 	name: "John",
	// 	age: 20,
	// 	isLoggedIn: false,
	// 	}

	// 	user1.greet = func() {
	// 		println("Hello", user1.name)
	// 	}

	// 	user1.greet()


	user1 := user{
		name: "John",
		age: 20,
		isLoggedIn: false,
	}

	user2 := user{
		name: "Jane",
		age: 30,
		isLoggedIn: false,
	}

	user1.greet()
	pointerUser2 := &user2
	pointerUser2.login()

	fmt.Printf("User1: %+v\n", user1)
	fmt.Printf("User2: %+v\n", user2)

}


func (u *user) login() {
	(*u).isLoggedIn = true
}

func (u user) greet() {
    println("Hello", u.name)
}