package mdb

import (
	"database/sql"
	"log"
	"time"

	"github.com/mattn/go-sqlite3"
)

type EmailEntry struct {
	Id          int64
	Email       string
	ConfirmedAt *time.Time
	OptOut      bool
}

func TryCreate(db *sql.DB) {
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS emails (
				id INTEGER PRIMARY KEY,
				email TEXT UNIQUE,
				confirmed_at DATETIME,
				opt_out BOOLEAN
		)
	`)

	if err != nil {
		if sqlError, ok := err.(sqlite3.Error); ok {
			if sqlError.Code != 1 { // table already exists
				log.Fatal(sqlError)
			} else {
				log.Fatal(err)
			}
		}
	}
}

func emailEntryFromRow(row *sql.Rows) (*EmailEntry, error) {
	var id int64
	var email string
	var confirmedAt int64
	var optOut bool

	err := row.Scan(&id, &email, &confirmedAt, &optOut)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	t := time.Unix(confirmedAt, 0)
	return &EmailEntry{id, email, &t, optOut}, nil
}

func CreateEmail(db *sql.DB, email string) error {
	_, err := db.Exec(`
		INSERT INTO emails
			(email, confirmed_at, opt_out) VALUES (?, 0, false)
	`, email)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetEmail(db *sql.DB, email string) (*EmailEntry, error) {
	rows, err := db.Query(`
		SELECT id, email, confirmed_at, opt_out
		FROM emails
		WHERE email = ?`, email)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		entry, err := emailEntryFromRow(rows)

		if err != nil {
			return nil, err
		}

		return entry, nil
	}
	return nil, nil
}

func UpdateEmail(db *sql.DB, entry EmailEntry) error {
	t := entry.ConfirmedAt.Unix()
	_, err := db.Exec(`
		UPDATE emails
		SET confirmed_at = ?, opt_out = ?
		WHERE id = ?
	`, t, entry.OptOut, entry.Id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteEmail(db *sql.DB, email string) error {
	_, err := db.Exec(`UPDATE emails SET opt_out = true WHERE email = ?`, email)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type GetEmailBatchQueryParams struct {
	Page  int
	Count int
}

func GetEmailBatch(db *sql.DB, params GetEmailBatchQueryParams) ([]EmailEntry, error) {
	rows, err := db.Query(`
		SELECT id, email, confirmed_at, opt_out
		FROM emails
		ORDER BY id ASC
		LIMIT ? OFFSET ?
	`, params.Count, params.Count*(params.Page-1))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	entries := make([]EmailEntry, 0, params.Count)

	for rows.Next() {
		entry, err := emailEntryFromRow(rows)

		if err != nil {
			return nil, err
		}

		entries = append(entries, *entry)
	}

	return entries, nil
}
