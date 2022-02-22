package main

import "fmt"

func newMysql() mysqlImplement {
	return mysqlImplement{}
}

type mysqlImplement struct{}

func (mysqlImplement) insert() {
	fmt.Println("insert mysql")
}
