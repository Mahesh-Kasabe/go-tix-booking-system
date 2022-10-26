package main

import "strings"

func validateUserInput(fname string, lname string, email string, tix int) (bool, bool, bool) {
	isValidname := len(fname) >= 2 && len(lname) >= 2
	isValidemail := strings.Contains(email , "@")
	isValidTix := tix > 0 
	return isValidemail , isValidname , isValidTix
}
