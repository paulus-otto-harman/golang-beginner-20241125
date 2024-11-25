package main

import (
	"flag"
	"fmt"
	l "log"
	"net/http"
	"os"
	"os/exec"
	"project/config"
	"project/router"
)

func main() {
	appConfig, err := config.Load()
	if err != nil {
		panic(err)
	}

	dumpDb := flag.Bool("dump", false, "use this flag to dump database")
	flag.Parse()

	if *dumpDb {
		dbdump(appConfig.Database, "ecommerce-db.sql")
	}

	r, log, db := router.Init(appConfig)
	defer db.Close()

	l.Printf("Server started on port %s", appConfig.Port)
	log.Info(fmt.Sprintf("Server started on port %s", appConfig.Port))
	http.ListenAndServe(fmt.Sprintf(":%s", appConfig.Port), r)
}

func dbdump(dbConfig config.DatabaseConfig, fileName string) {
	dbHost := fmt.Sprintf("--host=%s", dbConfig.Host)
	dbPort := fmt.Sprintf("--port=%d", 5432)
	dbName := fmt.Sprintf("--dbname=%s", dbConfig.Name)
	dbUsername := fmt.Sprintf("--username=%s", dbConfig.Username)
	dumpFormat := fmt.Sprintf("--format=%s", "plain")
	file := fmt.Sprintf("--file=../%s", fileName)
	cmd := exec.Command("pg_dump", dbHost, dbPort, dbName, dbUsername, dumpFormat, file)
	_, err := cmd.Output()

	if err != nil {
		panic(err)
	}
	l.Println("Database dumped successfully")
	os.Exit(0)
}
