package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type OneTimeFunction struct{
	Key []string 
	Values []string
}

type Person struct{
	Id string
	Name string
}

func main() {
	person := Person{Id: "0", Name: "Ali",}
	data, err := yaml.Marshal(&person)	
	if err != nil{
		return
	}
	fmt.Println("Marshal Data: ",string(data))
	
	err = yaml.Unmarshal(data, &person)
	if err != nil{
		return
	}
	fmt.Println("Unmarshal data:" , person)
	
	storage := OneTimeFunction{}
	storage.Key = make([]string, 1)
	storage.Values = make([]string, 1)
	
	storage.Key = []string{person.Id}
	storage.Values = []string{person.Name} 
	local , err := yaml.Marshal(&storage)
	if err != nil{
		return
	}
	fmt.Println("Marshal Keys and their values", storage)
	err = yaml.Unmarshal(local, &storage)
	if err != nil{
		return
	}
	fmt.Println("Conversion:", storage)
	
}


