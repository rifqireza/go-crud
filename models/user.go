package models

import (
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeypc/lat1/config"
	"github.com/jeypc/lat1/entity"
)

type UserModel struct {
	conn *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &UserModel{
		conn: conn,
	}
}

func (u *UserModel) GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *UserModel) AuthenticateUser(username string, password string) (bool, error) {
	var user entity.User
	err := u.conn.QueryRow("SELECT id, username, password FROM users WHERE username=$1 AND password=$2", username, password).Scan(&user.Id, &user.UserName, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
