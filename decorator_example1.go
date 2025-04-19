package main

import (
	"fmt"
)

type DB interface {
	Store(string) error
}

type Store struct{}

// This is coming from a third party lib
type ExecuteFn func(string)

func (s *Store) Store(value string) error {
	fmt.Println("Store into DB", value)
	return nil

}

func Excute(fn ExecuteFn) {
	fn("FOO BAR BAZ")
}

func myExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("my ex func", s)
		db.Store(s)
	}
}

func RunDecoratorExample1() {
	s := &Store{}
	Excute(myExecuteFunc(s))
}
