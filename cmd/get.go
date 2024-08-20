package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"log"
	"pwm/colors"
)

var getCmd = &cobra.Command{
	Use:   "get [service]",
	Short: "Récupère un mot de passe",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]

		db := initDB()
		defer db.Close()
		getPassword(db, service)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getPassword(db *sql.DB, service string) {
	var username, password string
	query := `SELECT username, password FROM passwords WHERE service = ?`
	err := db.QueryRow(query, service).Scan(&username, &password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(colors.ErrorColor("Service not found"))
		} else {
			log.Fatal(err)
		}
		return
	}

	fmt.Printf(colors.InfoColor("Username: %s\nPassword: %s\n"), username, password)
}
