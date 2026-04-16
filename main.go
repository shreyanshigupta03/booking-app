package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const confTickets = 50

var ConfName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Printf("These are all our bookings:%v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out.Come again next Year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The Emal you entered doesnot contains @ sign")
			}
			if !isValidTickets {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v ticktes\n", remainingTickets, userTickets)
			}

		}

	}
	wg.Wait()

}

func greetUser() {
	fmt.Println("Welcome to ", ConfName, "booking application!")
	fmt.Println("We have total", confTickets, "ticket and", remainingTickets, "tickets are still remaining.")
	fmt.Println("Get your tickets here to attend the", ConfName, ".....!")

}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter your number of tickets you want to book:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//bookings[0] = firstName + " " + lastName

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will get confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, ConfName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("##################################")
	fmt.Printf("Sending tickets:\n %v to email address %v\n", ticket, email)
	fmt.Println("##################################")
	wg.Done()
}
