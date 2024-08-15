/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"password-manager/cmd"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

var sessionActive bool

func main() {

	cmd.Execute()
}

func creatMasterPassword() {
	fmt.Printf("Enter master password you will use ")
	var masterPassword string
	fmt.Scanln(&masterPassword)
	hashedPassword := hashPassword(masterPassword)
	file, err := os.Create("master_password.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write([]string{hashedPassword})
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hex.EncodeToString(hash[:])
}


func passwordVerify() {

	fmt.Print("Enter master password: ")
	masterPasswordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}
	masterPassword := string(masterPasswordBytes)

	file, err := os.Open("Students.csv")
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


	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(masterPassword)); err != nil {
		fmt.Println("Incorrect master password.")
	} else {
		fmt.Println("Master password verified successfully.")

	}
}
