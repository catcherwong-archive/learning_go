package models

type Person struct {
	Id int 
	Name string
	Gender int
}

func GetAllPersons() (lists []Person){
	lists = make([]Person, 0)

	lists = append(lists,Person{Id:1,Name:"catcher",Gender:1})
	lists = append(lists,Person{Id:2,Name:"tom",Gender:1})
	lists = append(lists,Person{Id:3,Name:"james",Gender:1})
	lists = append(lists,Person{Id:4,Name:"lisa",Gender:2})
	lists = append(lists,Person{Id:5,Name:"lily",Gender:2})

	return lists
}

func GetPersonById(id int) (obj Person){
	return Person{Id:id,Name:"name",Gender:1}
}