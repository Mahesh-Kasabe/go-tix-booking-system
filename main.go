package main

import ("fmt" )

type UserData struct {
	firstName       string
	lastName        string
	email           string
	tickets uint
}

var conferenceName string = "Kubecon"
var remainingTickets uint = 50
var bookings = make([]UserData , 0)


func main(){

	greetusers(conferenceName)

	firstName, lastName, email, tickets := getusers()

	 isValidname , isValidemail , isValidTix := ValidateUserInput(firstName,lastName,email,tickets)

	 if  isValidname && isValidemail && isValidTix {

		if remainingTickets == 0{
			fmt.Println("Our Confference is booked out , come back next year ... ")
		

		booktix(firstName, lastName, email, tickets)

}
	 }else {
		if !isValidname {
			fmt.Println("Your name is incorrect")
		
		if !isValidemail {
			fmt.Println("Your email is incorrect")
		}}
		if !isValidTix {
			fmt.Println("Your ticket is incorrect")
		}
	 }
}

func greetusers(conferenceName string) {
	fmt.Printf("Get your %v tix here. \n", conferenceName)
}

func getusers() (string, string, string, uint) {
	var fname string 
	var lname string
	var email string
	var tix uint 

	fmt.Println("Enter First Name : ")
	fmt.Scan(&fname)

	fmt.Println("Enter Last Name : ")
	fmt.Scan(&lname)

	fmt.Println("Enter Email : ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets : ")
	fmt.Scan(&tix)

	return fname , lname , email , tix

}

func booktix(firstName string, lastName string, email string, tickets uint) {
	remainingTickets = remainingTickets - tickets

	var userdata = UserData  {
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		tickets: tickets,
	}

	bookings = append(bookings, userdata)
	fmt.Printf("Thanks %v %v for booking %v tickets for kubecon, You will recieve a confirmation email on %v \n", firstName, lastName, tickets, email)
	fmt.Printf("The list of bookings : %v \n", bookings)
	fmt.Printf("from 50 Tickets  %v remaining \n", remainingTickets)

}