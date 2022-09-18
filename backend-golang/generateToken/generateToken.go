package generateToken

import (
	"ChangeContactWeb/db"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateSecureToken() {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		fmt.Println(err)
	}
	db := db.DB()
	a := hex.EncodeToString(b)
	db.Exec("INSERT INTO tokens (token, is_use, count) VALUES ($1, false, 0)", a)

}
