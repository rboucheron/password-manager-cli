package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"pwm/colors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [service] [username] [password]",
	Short: "",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		username := args[1]
		password := args[2]

		db := initDB()
		defer db.Close()
		addPassword(db, service, username, password)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./passwords.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS passwords (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        service TEXT NOT NULL,
        username TEXT NOT NULL,
        password TEXT NOT NULL
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func addPassword(db *sql.DB, service, username, password string) {

	insertStmt := `INSERT INTO passwords (service, username, password) VALUES (?, ?, ?)`
	_, err := db.Exec(insertStmt, service, username, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(colors.SuccessColor("Password added successfully!"))
}
