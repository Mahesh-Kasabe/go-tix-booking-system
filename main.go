package main

import "fmt"

func main(){

	greetusers()

	 getusers()
}

func greetusers() {
	fmt.Println("Get your kubecon tix here")
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