package models

import (
	"fmt"
	"time"
)

type User struct {
	Id      int64     `sql:"id" json:"id"`
	Name    string    `sql:"name" json:"name"`
	Created time.Time `sql:"created" json:"created"`
	Uptated time.Time `sql:"updated" json:"updated"`
}

func (this *User) String() string {
	return fmt.Sprintf("<User name:%s>", this.Name)
}

func GetUserNameById(userId int64) (string, error) {
	// for test
	return "hello, you get me!", nil
}
