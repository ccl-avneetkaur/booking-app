package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//var conferenceName string = "Go Conference"

	greetUser()

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// for {
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("These are all the bookings: %v\n", firstNames)

		// var noTicketsRemaining bool = remainingTickets == 0
		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("Your first name or Last name is too short!")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesntcontain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
		// fmt.Printf("We have only %v tickets remaining, So you cannot book %v tickets\n", remainingTickets, userTickets)
		// fmt.Println("Your input data is invalid, try again!")
	}
	// }
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ is replaced by index as it is unused
		// var names = strings.Fields(booking)
		//var firstName = names[0]
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a map for a user

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //type conversion
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v are remaining tickets for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// fmt.Printf("%v tickets for %v %v", userTickets, firstName, lastName)
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket\n %v \nto email address: %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}

// var bookings = [50]string{"Avneet", "Shehbaz", "Anahat"}
// var bookings [50]string
// var bookings []string

// firstName = "Tom"

// bookings[0] = firstName + " " + lastName

// fmt.Printf("The whole array: %v\n", bookings)
// fmt.Printf("The first value: %v\n", bookings[0])
// fmt.Printf("Array type: %T\n", bookings)
// fmt.Printf("Array length: %v\n", len(bookings))

// fmt.Printf("The whole slice: %v\n", bookings)
// fmt.Printf("The first value: %v\n", bookings[0])
// fmt.Printf("Slice type: %T\n", bookings)
// fmt.Printf("Slice length: %v\n", len(bookings))

// fmt.Println("Welcome to our", conferenceName, "booking application")
// fmt.Println("We have total of", totalTickets, "tickets and", remainingTickets, "tickets are still available")
// fmt.Println("Get your tickets here to attend")

// fmt.Println(conferenceName)
