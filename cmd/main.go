package main

import (
	"context"

	"github.com/k0kubun/pp"

	"github.com/bregydoc/gaea/backends/mongo"

	"github.com/bregydoc/gaea"
)

func main() {

	uri := "mongodb+srv://bregymr:" + pass + "@bombo0-l7ebf.gcp.mongodb.net/test?retryWrites=true"

	repo, err := mongo.NewRepo(uri, "gaea", "people", "accounts", "sessions")
	if err != nil {
		panic(err)
	}

	err = repo.Connect(context.TODO())
	if err != nil {
		panic(err)
	}

	service, err := gaea.NewService(repo)

	bregy, err := service.SignUp(context.TODO(), &gaea.MinimalPersonInformation{
		Name: "Bregy Malpartida Ramos",
		Account: &gaea.Account{
			Type: gaea.Email,
			ID:   "bregymr2@gmail.com",
		},
		Password: "malpartida1",
	})

	if err != nil {
		panic(err)
	}

	pp.Println(bregy)
}

const pass = "6Ntgf3YQBXi0WWj8"
