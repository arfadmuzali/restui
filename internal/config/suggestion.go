package config

func AddSuggestion(text string) error {

	_, err := DB.Exec(`
		INSERT INTO suggestions (text)
		VALUES (?)
		ON CONFLICT(text) DO UPDATE SET
    		updated_at = CURRENT_TIMESTAMP;
		`, text)
	if err != nil {
		return err
	}

	return nil
}

func GetSuggestions() ([]string, error) {

	rows, err := DB.Query(`
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
