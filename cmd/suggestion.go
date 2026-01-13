package cmd

import (
	"database/sql"
	"fmt"

	"github.com/arfadmuzali/restui/internal/config"
	"github.com/spf13/cobra"
)

var clear bool

// suggestionCmd represents the suggestion command
var suggestionCmd = &cobra.Command{
	Use:   "suggestion",
	Short: "Manage suggestion in RESTUI app",
	Run: func(cmd *cobra.Command, args []string) {

		err := config.ConfigInitialization()
		if err != nil {
			panic(err)
		}

		db, err := config.DatabaseInitialize()
		if err != nil {
			panic(err)
		}
		defer db.Close()

		if clear {
			err := clearSuggestion(db)
			if err != nil {
				panic(err)
			}
			fmt.Println("Successfully deleted all suggestions.")
			return
		}

		suggestions, err := getSuggestion(db)

		if len(suggestions) == 0 {
			fmt.Println("No result.")
		} else {
			for i, suggestion := range suggestions {
				s := fmt.Sprintf("%v. %v\n", i+1, suggestion)
				fmt.Print(s)
			}
		}

	},
}

func init() {
	suggestionCmd.Flags().BoolVarP(&clear, "clear", "c", false, "Clear all suggestion")
	rootCmd.AddCommand(suggestionCmd)
}

func getSuggestion(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`
		SELECT text FROM suggestions ORDER BY updated_at DESC;
		`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []string

	for rows.Next() {
		var suggestion string
		if err := rows.Scan(&suggestion); err != nil {
			return nil, err
		}
		result = append(result, suggestion)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func clearSuggestion(db *sql.DB) error {
	rows, err := db.Query(`
		DELETE FROM suggestions
		`)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
