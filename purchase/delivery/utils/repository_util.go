package utils

import (
	"crypto/rand"
	//"github.com/jinzhu/gorm"
	//"time"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

/*
func Generatetoken(db *gorm.DB, email string)(*Response, error){
	//user, err := FindByEmail(db, email)
	token, err := generateRandomString(20)
	expirationTime := time.Now().Add(5 * time.Minute)
	expirationTime.Unix()

	if err != nil{
		return nil, err
	}
	newRecovery := &PassRecovery{
		//UserID			:	user.ID,
		Token		 	:	string(token),
		ExpiredDate		:	expirationTime,
	}
	id, err := CreateToken(db, newRecovery)
	if err != nil{
		return nil, err
	}
	return &Response{Id: id, Token: token}, nil
}
*/