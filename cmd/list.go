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
    _ "github.com/mattn/go-sqlite3"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
        defer db.Close()
		listService(db)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listService(db *sql.DB) {
    query := `SELECT service, username FROM passwords`
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var service, username string
        err := rows.Scan(&service, &username)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf( colors.InfoColor("- %s  Username: %s\n"), service, username)
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }
}
