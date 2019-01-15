package main

import (
	"context"
	"github.com/bregydoc/gaea"
	"github.com/k0kubun/pp"
	"time"
)

func main() {

	uri := "mongodb+srv://bregymr:" + pass + "@bombo0-l7ebf.gcp.mongodb.net/test?retryWrites=true"

	repo, err := gaea.NewMongoRepo(&gaea.Config{
		UriConnection:  uri,
		UsersDBName:    "people",
		SessionsDBName: "sessions",
	})
	if err != nil {
		panic(err)
	}

	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = repo.Connect(c)
	if err != nil {
		panic(err)
	}

	p, err := repo.SignUp(c, &gaea.MinimalPersonInformation{
		Name: "Bregy Malpartida",
		Account: &gaea.Account{
			ID:   "bregy.malpartida@utec.edu.pe",
			Type: gaea.Email,
		},
		Password: "malpartida1",
	})

	if err != nil {
		panic(err)
	}

	_, _ = pp.Println(p)
}

const pass = "6Ntgf3YQBXi0WWj8"
