package db

import (
	"ChangeContactWeb/model"
)

func UpdateCount(token *model.Token) error {
	db := DB()
	var count int
	var isUse bool
	_, err := db.Exec(`UPDATE "tokens" SET count = count + 1 WHERE token=$1 `, token.Token)
	if err != nil {
		return err
	}

	row := db.QueryRow(`SELECT count, is_use FROM "tokens" WHERE token=$1`, token.Token)
	err = row.Scan(&count, &isUse)
	if err != nil {
		return err
	}
	if !isUse {
		_, err := db.Exec(`UPDATE "tokens" SET is_use = true WHERE token=$1 `, token.Token)
		if err != nil {
			return err
		}
	}
	if count > 50 {
		_, err := db.Exec(`DELETE FROM "tokens" WHERE token=$1`, token.Token)
		if err != nil {
			return err
		}
	}
	return nil
}
