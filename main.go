/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"pwm/cmd"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

var verifiedMasterPassword string

func main() {
	if _, err := os.Stat("master_password.csv"); os.IsNotExist(err) {
		createMasterPassword()
	} else {
	
		if passwordVerify() {
			verifiedMasterPassword = "YourVerifiedPassword"
			cmd.Execute(verifiedMasterPassword)
		}
	
	}
}

func createMasterPassword() {
	fmt.Printf("Enter master password you will use: ")
	var masterPassword string
	fmt.Scanln(&masterPassword)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(masterPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating password hash:", err)
		return
	}


	file, err := os.Create("master_password.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()


	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{string(hashedPassword)})
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func passwordVerify()  bool{
	fmt.Print("Enter master password: ")
	masterPasswordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}
	masterPassword := string(masterPasswordBytes)

	file, err := os.Open("master_password.csv")
	if err != nil {
		log.Fatalf("Error while opening the file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error while reading CSV records: %v", err)
	}

	if len(records) == 0 || len(records[0]) == 0 {
		log.Fatal("CSV file is empty or does not contain the expected data")
	}

	hashedPassword := records[0][0]

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(masterPassword))
	if err != nil {
		fmt.Println("Incorrect master password.")
		return false
	} else {
	
		verifiedMasterPassword = masterPassword
		return true
	}
}
