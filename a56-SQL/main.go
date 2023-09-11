package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id string
	name string
	age int
	grade int
}