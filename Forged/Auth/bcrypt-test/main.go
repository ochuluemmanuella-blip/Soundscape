package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func main() {
	password := "ok29211"
	hash, err := HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash:", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:", match)

	wrongMatch := CheckPasswordHash("someWrongPassword", hash)
	fmt.Println("Wrong password match:", wrongMatch)
}
