/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
    "database/sql"
    "fmt"
    "log"
    "pwm/colors"

    "github.com/spf13/cobra"
)

var Username, Password bool

// updateCmd represents the update command
var updateCmd = &cobra.Command{
    Use:   "update [service] [newvalue]",
    Short: "Update a username or password for a specific service",
    Long:  `This command updates either the username or the password for a given service, based on the flag provided.`,
    Run: func(cmd *cobra.Command, args []string) {

        if len(args) < 2 {
            log.Fatal("Error: You must provide service and newvalue arguments")
        }

        service := args[0]
        newvalue := args[1]

        if !Username && !Password {
            log.Fatal("Error: You must specify whether to update the username (-u) or password (-p)")
        }

        if Username && Password {
            log.Fatal("Error: You cannot update both the username and password at the same time")
        }

        db := initDB()
        defer db.Close()

        if Username {
            updateUsername(db, service, newvalue)
        } else if Password {
            updatePassword(db, service, newvalue)
        }

        fmt.Println("Update command executed successfully")
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)

    updateCmd.Flags().BoolVarP(&Username, "username", "u", false, "Update the username for the service")
    updateCmd.Flags().BoolVarP(&Password, "password", "p", false, "Update the password for the service")
}

func updatePassword(db *sql.DB, service, password string) {

    insertStmt := `UPDATE passwords SET password = ? WHERE service = ?`
    _, err := db.Exec(insertStmt, password, service)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(colors.SuccessColor("Password added successfully!"))

}

func updateUsername(db *sql.DB, service, username string) {

    insertStmt := `UPDATE passwords SET username = ? WHERE service = ?`
    _, err := db.Exec(insertStmt, username, service)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(colors.SuccessColor("Username added successfully!"))

}

