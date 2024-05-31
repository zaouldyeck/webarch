package main

import (
	"fmt"

	architecture "github.com/zaouldyeck/webarch"
	"github.com/zaouldyeck/webarch/storage/mongo"
	"github.com/zaouldyeck/webarch/storage/postgres"
)

func main() {
	dbm := mongo.Db{}
	dbpg := postgres.Db{}

	p1 := architecture.Person{
		First: "John",
	}
	p2 := architecture.Person{
		First: "Paul",
	}

	ps := architecture.NewPersonService(dbm)

	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)

	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbm, 2))

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	architecture.Put(dbpg, 1, p1)
	architecture.Put(dbpg, 2, p2)

	fmt.Println(architecture.Get(dbpg, 1))
	fmt.Println(architecture.Get(dbpg, 2))
}
