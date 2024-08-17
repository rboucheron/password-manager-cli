package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// addCmd représente la commande "add"
var addCmd = &cobra.Command{
	Use:   "add [service] [username] [password]",
	Short: "Ajoute un nouveau mot de passe",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		username := args[1]
		password := args[2]

		// Récupérer le mot de passe maître depuis les flags
		masterPassword, _ := cmd.Flags().GetString("master-password")

		// Chiffrer le mot de passe utilisateur
		encryptedPassword, err := encryptPassword(masterPassword, password)
		if err != nil {
			log.Fatalf("Failed to encrypt password: %v", err)
		}

		db := initDB()
		defer db.Close()
		addPassword(db, service, username, encryptedPassword)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func encryptPassword(masterPassword, password string) (string, error) {
	key := []byte(masterPassword + "effcd21a0000")
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(password), nil)
	return hex.EncodeToString(ciphertext), nil
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

	fmt.Println("Password added successfully!")
}
