package main

import (
	"fmt"
	"hex-arch-template/internal/adapters/app/api"
	"hex-arch-template/internal/adapters/core/arithmetic"
	"hex-arch-template/internal/adapters/framework/right/db"
	"hex-arch-template/internal/ports"
	"log"
)

func main() {
	// ports
	var core ports.ArithmeticPort = arithmetic.NewAdapter()
	db, err := db.NewAdapter("mysql", "")

	if err != nil {
		log.Fatal(err)
	}

	var app ports.APIPort = api.NewAdapter(db, core)

	fmt.Println(app.GetAddition(1, 4))
}
