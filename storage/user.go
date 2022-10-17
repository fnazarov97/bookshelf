package storage

import (
	"bookshelf/models"
	"errors"
	"fmt"
)

const (
	SQL = `
            INSERT INTO users
            (name, key, secret)
            VALUES ($1,$2,$3) RETURNING *
        `
	UPD = `
            UPDATE books
            SET isbn = $1, title = $2, gener = $3, description = $4, author = $5, publishYear = $6, created_at = $7, updated_at = $8
            WHERE id = $9 RETURNING *
        `
	DEL = `DELETE FROM books WHERE id = $1 RETURNING *`
)

func (im Inmemory) AddUser(id string, entity models.CreateUserModel) error {
	db, err := Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = im.GetUser(entity.Key)
	if err != nil {
		_, err = db.Query(SQL, entity.Name, entity.Key, entity.Secret)
		return err
	}
	return errors.New("user already exist in base")

}

func (im Inmemory) GetUser(key string) (models.GetUserModel, error) {
	var result models.GetUserModel
	db, err := Connect()
	if err != nil {
		panic(err)
	}
	query := `select * from users where key = $1;`
	row := db.QueryRow(query, key)
	err = row.Scan(&result.ID, &result.Name, &result.Key, &result.Secret, &result.CreatedAt)
	fmt.Println(result, "result", result.ID)
	if err != nil {
		return result, errors.New("scan error")
	}
	defer db.Close()
	return result, nil
}
