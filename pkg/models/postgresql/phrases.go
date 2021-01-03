// Copyright 2020 Zoltán Király. All rights reserved.

package postgresql

import (
	"database/sql"

	"github.com/zoliky/phrasian/pkg/models"
)

type PhraseModel struct {
	DB *sql.DB
}

func (p *PhraseModel) GetAll(status string, userID string) (*[]models.Phrase, error) {
	stmt := `
		SELECT
			id,
			phrase,
			translation,
			pronunciation
		FROM
			phrase
		WHERE
			status=$1
		AND
			user_id=$2
		ORDER BY
			id DESC;`

	rows, err := p.DB.Query(stmt, status, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	phrases := []models.Phrase{}
	for rows.Next() {
		p := models.Phrase{}
		err := rows.Scan(
			&p.ID,
			&p.Phrase,
			&p.Translation,
			&p.Pronunciation)
		if err != nil {
			return nil, err
		}
		phrases = append(phrases, p)
	}

	return &phrases, nil
}

func (p *PhraseModel) GetPhrase(id string) (*models.Phrase, error) {
	stmt := `
		SELECT
			id,
			phrase,
			translation,
			pronunciation
		FROM
			phrase
		WHERE
			id=$1;`
	row := p.DB.QueryRow(stmt, id)
	phr := models.Phrase{}
	err := row.Scan(&phr.ID, &phr.Phrase, &phr.Translation, &phr.Pronunciation)
	if err != nil {
		return nil, err
	}
	return &phr, nil
}
