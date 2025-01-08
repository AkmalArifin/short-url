package models

import (
	"errors"
	"time"

	"github.com/AkmalArifin/short-url/internal/db"
	"github.com/AkmalArifin/short-url/pkg"
	"github.com/guregu/null/v5"
)

type URL struct {
	ID          int64        `json:"id"`
	URL         null.String  `json:"url"`
	ShortCode   null.String  `json:"short_code"`
	AccessCount null.Int64   `json:"access_count"`
	CreatedAt   pkg.NullTime `json:"created_at"`
	UpdatedAt   pkg.NullTime `json:"updated_at"`
}

func GetURLS() ([]URL, error) {
	query := `SELECT id, url, short_code, access_count, created_at, updated_at FROM shorten_url`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []URL
	for rows.Next() {
		var url URL
		err = rows.Scan(&url.ID, &url.URL, &url.ShortCode, &url.AccessCount, &url.CreatedAt, &url.UpdatedAt)
		if err != nil {
			return nil, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}

func (u *URL) Save() error {
	query := `
		INSERT INTO shorten_url(url, short_code, created_at, updated_at)
		VALUES (?, ?, ?, ?)
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	u.CreatedAt.SetValue(time.Now())
	u.UpdatedAt.SetValue(time.Now())

	res, err := stmt.Exec(u.URL, u.ShortCode, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}

	u.ID, err = res.LastInsertId()

	return err
}

func (u *URL) Delete() error {
	query := `
		DELETE FROM shorten_url
		WHERE short_code = ?
		`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ShortCode)
	return err
}

func (u *URL) Update() error {
	query := `
		UPDATE shorten_url
		SET url = ?, short_code = ?, access_count = ?, updated_at = ?
		WHERE id = ?
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	u.UpdatedAt.SetValue(time.Now())
	_, err = stmt.Exec(u.URL, u.ShortCode, u.AccessCount, u.UpdatedAt, u.ID)
	return err
}

func (u *URL) GetShortCode() error {
	query := `SELECT id, url, short_code, access_count, created_at, updated_at FROM shorten_url WHERE url = ?`
	row := db.DB.QueryRow(query, u.URL.ValueOrZero())

	var url URL
	err := row.Scan(&url.ID, &url.URL, &url.ShortCode, &url.AccessCount, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		return err
	}

	if u.URL.ValueOrZero() != url.URL.ValueOrZero() || !u.URL.Valid || !url.URL.Valid {
		return errors.New("is not valid")
	}

	u.ID = url.ID
	u.ShortCode.SetValid(url.ShortCode.ValueOrZero())
	u.AccessCount.SetValid(url.AccessCount.ValueOrZero())
	u.CreatedAt.SetValue(url.CreatedAt.Time)
	u.UpdatedAt.SetValue(url.UpdatedAt.Time)

	return nil
}

func GetOriginURL(shortCode string) (string, error) {
	query := `SELECT url FROM shorten_url WHERE short_code = ?`
	row := db.DB.QueryRow(query, shortCode)

	var url string
	err := row.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func GetStat(shortCode string) (URL, error) {
	query := `SELECT id, url, short_code, access_count, created_at, updated_at FROM shorten_url WHERE short_code = ?`
	row := db.DB.QueryRow(query, shortCode)

	var url URL
	err := row.Scan(&url.ID, &url.URL, &url.ShortCode, &url.AccessCount, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		return URL{}, err
	}

	return url, nil
}
