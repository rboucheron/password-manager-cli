package cmd

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/spf13/cobra"
    _ "github.com/mattn/go-sqlite3"
)

// getCmd représente la commande "get"
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
    var username, hashedPassword string
    query := `SELECT username, password FROM passwords WHERE service = ?`
    err := db.QueryRow(query, service).Scan(&username, &hashedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("Service non trouvé.")
        } else {
            log.Fatal(err)
        }
        return
    }

    fmt.Printf("Username: %s\nPassword: %s\n", username, hashedPassword)
}
