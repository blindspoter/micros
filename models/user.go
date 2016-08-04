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

func (this *User) GetUserNameById(int64 userId) string {
	// for rpc test
	name := "rpc server"
	return name
}

func (this *User) Add() {
	//TODO:
}

func (this *User) Get() {
	//TODO:
}

func (this *User) Update() {
	//TODO:
}
