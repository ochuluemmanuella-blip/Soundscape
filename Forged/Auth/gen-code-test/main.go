package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"database/sql"
)

func generateCode() (string, error) {
	max := big.NewInt(1000000)
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%06d", num.Int64()), nil
}
func storeCode(db *sql.DB, email string) (string, error) {
	storedcode, err := generateCode();
		if err != nil {
			return "", err
		
	}

}
func main() {
	for i := 0; i < 11; i++ {
		code, err := generateCode()
		if err != nil {
			continue
		}
		fmt.Println(code)
	}
}
