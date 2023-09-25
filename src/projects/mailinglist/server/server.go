package main

import (
	"database/sql"
	"log"
	"mailinglist/jsonapi"
	"mailinglist/mdb"
	"sync"

	"github.com/alexflint/go-arg"
)

var args struct {
	DbPath   string `arg:"env:MAILINGLIST_DB" help:"Path to the database file"`
	BindJson string `arg:"env:MAILINGLIST_BIND_JSON" help:"Bind address for the JSON API"`
}

func main() {
	arg.MustParse(&args)

	if args.DbPath == "" {
		args.DbPath = "list.db"
	}

	if args.BindJson == "" {
		args.BindJson = ":8080"
	}

	log.Printf("Using database: %s\n", args.DbPath)
	db, err := sql.Open("sqlite3", args.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mdb.TryCreate(db)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		log.Println("starting API sserver")
		jsonapi.Serve(db, args.BindJson)
		wg.Done()
	}()

	wg.Wait()
}
