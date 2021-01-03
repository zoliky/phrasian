// Copyright 2020 Zoltán Király. All rights reserved.

package models

import (
	"database/sql"
	"time"
)

type Phrase struct {
	ID            string
	Phrase        string
	Translation   sql.NullString
	Pronunciation sql.NullString
	Created       time.Time
	UserID        string
}
