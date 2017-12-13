package utils

import (
	"crypto/rand"
	"fmt"
	"strconv"

	"github.com/nikhosagala/soapi/db"
	"golang.org/x/crypto/bcrypt"
)

// CreatePassword ...
func CreatePassword(password string) string {
	bytePlain := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePlain, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// ComparePassword ...
func ComparePassword(hashPassword string, plainPassword string) bool {
	byteHash := []byte(hashPassword)
	bytePlain := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}

// Token is a function to generate token to save into database
func Token() string {
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Code ...
func Code(code string) string {
	c := []byte(code)[0:3]
	d := len(db.ReadServices(-1, 0)) + 1
	return string(c) + appentInt(d)
}

func appentInt(code int) string {
	var result string
	c := strconv.Itoa(code)
	switch len(c) {
	case 1:
		result = "00" + c
	case 2:
		result = "0" + c
	case 3:
		result = c
	default:
		result = c
	}
	return result
}
