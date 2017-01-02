package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"


)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("infos")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		panic(err)
	}
	//
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone:", result.Phone)
	//var arr = []Person{}
	////err = c.FindId("58692de07241b7122a1d6719").All(&arr)
	//err = c.Find(bson.M{"name": "Ale"}).All(&arr)
	//if err != nil {
	//	panic(err)
	//}
	////c.FindId()
	//fmt.Println(arr[0])
	////err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
	err = c.Update(bson.M{"name": "Ale"}, bson.M{"$set": bson.M{"name": "abc"}})
	_,err = c.UpdateAll(bson.M{"name": "Ale"}, bson.M{"$set": bson.M{"name":"abc"}})

	if err != nil {
		panic(err)
	}

	err = c.Insert(&Person{
		Name:  "Jimmy Kuu",
		Phone: "3"},
	)

	if err != nil {
		panic(err)
	}

}
