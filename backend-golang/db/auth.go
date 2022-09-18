package db

import (
	"ChangeContactWeb/model"
	"fmt"
)

func Compare(token *model.Token) error {
	var id int
	db := DB()
	res := db.QueryRow(`SELECT id FROM "tokens" WHERE token=$1;`, token.Token)
	_ = res.Scan(&id)
	if id == 0 {
		return fmt.Errorf("Wrong token")
	}
	return nil
}
