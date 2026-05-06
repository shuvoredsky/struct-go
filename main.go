package main


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

	user1.greet()


}


func (u user) greet() {
    println("Hello", u.name)
}