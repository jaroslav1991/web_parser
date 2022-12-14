package main

import (
	"log"
	"net/http"
	"time"
	"yla_parser/internal/config"
	"yla_parser/internal/handlers"
	"yla_parser/pkg/repository"
)

func main() {
	dbConf, err := config.GetDbConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(dbConf)
	if err != nil {
		log.Fatal(err)
	}

	storage := handlers.NewStorage(db)

	//go func() {
	//	var address = "https://edc.sale/ru/mozhajsk/real-estate/sale/rooms/prodaju-2-komnaty-v-3-kh-komn-kv-v-g-ruza-moskovskoj-oblasti-812065.html"
	//	handlers.WriteNewAddress(storage, address)
	//}()

	go func() {
		var email = "lox@gmail.com"

		for {
			http.HandleFunc("/events", handlers.ServiceComparePrice(storage, email))

			//handlers.Testing(storage, email)
			time.Sleep(time.Second * 10)
		}
	}()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
