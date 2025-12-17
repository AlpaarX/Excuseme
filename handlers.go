package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Excuse struct {
	Phrase string
	Author string
}
type Excuses []Excuse

func RandomExcuse(data *Excuses) Excuse {
	return (*data)[rand.Intn(len(*data))]
}

var (
	staticDir    = getAbsDirPath() + "/static/"
	templatesDir = staticDir + "/templates"
	templates    = template.Must(template.ParseFiles(
		templatesDir + "/index.html",
	))
)

func fetchDB() (Excuses, error) {
	excuses := Excuses{{Phrase: "I'm sorry, I have to take my cat to the vet.", Author: "John"},
		{Phrase: "I'm sorry, I have to take my dog to the vet.", Author: "Jane"},
		{Phrase: "I'm sorry, I have to take my fish to the vet.", Author: "Jack"},
		{Phrase: "I'm sorry, my internet is down.", Author: "Alice"},
		{Phrase: "I'm sorry, I have a dentist appointment.", Author: "Bob"},
		{Phrase: "I'm sorry, my car broke down.", Author: "Charlie"},
		{Phrase: "I'm sorry, I have a family emergency.", Author: "Dave"},
		{Phrase: "I'm sorry, I have a headache.", Author: "Eve"},
		{Phrase: "I'm sorry, I have to pick up my kids.", Author: "Frank"},
		{Phrase: "I'm sorry, I have a prior engagement.", Author: "Grace"},
		{Phrase: "I'm sorry, I have to help a friend.", Author: "Hank"},
	}
	return excuses, nil
}

func getAbsDirPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	return pwd
}

func index(w http.ResponseWriter, r *http.Request) {
	// fake data fetch
	data, err := fetchDB()
	if err != nil {
		log.Fatal(err)
	}

	templates.Execute(w, RandomExcuse(&data))
}
