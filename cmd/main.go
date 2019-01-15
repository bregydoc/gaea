package main

import (
	"context"
	"github.com/bregydoc/gaea"
	"github.com/k0kubun/pp"
	"net/url"
	"time"
)

func main() {

	uri := "mongodb+srv://bregymr:" + pass + "@bombo0-l7ebf.gcp.mongodb.net/test?retryWrites=true"
	url, err := url.ParseRequestURI(uri)
	if err != nil {
		panic(err)
	}
	//_, _ = pp.Println(url.String())
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//client, err := mongo.Connect(ctx, url.String())
	//if err != nil {
	//	panic(err)
	//}
	//
	////_, _ = pp.Println(client)
	//fmt.Print("ping...")
	//err = client.Ping(ctx, readpref.Primary())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print("pong...")
	repo, err := gaea.NewMongoRepo(&gaea.Config{
		UriConnection:  url.String(),
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
