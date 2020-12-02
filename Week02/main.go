package main

import (
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

var d dao

func main() {
	var (
		user *User
		err  error
	)
	user, err = d.Get(1)
	if err == nil {
		fmt.Println(user)
		return
	}
	fmt.Println(err)
}

var (
	ErrNoRows = errors.New(" sql.ErrNoRows ")
)

type User struct {
	Id   int
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("id:%v,name:%v", u.Id, u.Name)
}

type dao struct {
}

func (d *dao) Get(id int) (*User, error) {
	var err error
	if id == 1 {
		err = ErrNoRows
	}

	if err == nil {
		return &User{Id: id, Name: "song"}, err
	}
	return nil, xerrors.Wrapf(err, "id:%v not found", id)
}
