/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [service]",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()
		service := args[0]
		delete(db, service)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func delete(db *sql.DB, service string) {

	query := `DELETE FROM passwords WHERE service = ?`
	result, err := db.Exec(query, service)
	if err != nil {
		fmt.Println("Error: password not deleted")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error: could not retrieve the number of affected rows")
		return
	}

	if rowsAffected == 0 {
		fmt.Println("No password found for the given service")
	} else {
		fmt.Println("Password deleted successfully")
	}
}

