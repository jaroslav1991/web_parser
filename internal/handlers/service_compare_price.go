package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func ServiceComparePrice(storage *Storage, email string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		events, err := GetEvents(storage)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(events); i++ {

			actualPrice, err := PriceParser(events[i].Address)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("actualPrice = %d", actualPrice)
			dbPrice, err := GetEventByAddress(storage, events[i].Address)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("dbPrice = %d", dbPrice.Price)
			res := ComparePrices(actualPrice, dbPrice.Price)

			if res {
				SendingPriceToEmail(email)
				_, err := UpdateEvents(storage)
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("not sending")
				}
			}
			s := fmt.Sprintf("actualPrice = %d dbPrice = %d, ", actualPrice, dbPrice.Price)
			//io.WriteString(w, s)
			_, err = w.Write([]byte(s))
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}

func WriteNewAddress(storage *Storage, address string) {
	giveAddress, err := PriceParser(address)
	if err != nil {
		log.Fatal(err)
	}
	_, err = CreateEvent(storage, address, giveAddress)
	if err != nil {
		log.Fatal(err)
	}
}

func Testing(storage *Storage, email string) {
	events, err := GetEvents(storage)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(events); i++ {

		actualPrice, err := PriceParser(events[i].Address)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("actualPrice = %d  ", actualPrice)

		dbPrice, err := GetEventByAddress(storage, events[i].Address)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("dbPrice = %d  ", dbPrice.Price)

		res := ComparePrices(actualPrice, dbPrice.Price)

		if res {
			SendingPriceToEmail(email)
			_, err := UpdateEvents(storage)
			if err != nil {
				log.Fatal(err)
			} else {
				log.Println("not sending")
			}
		}
		fmt.Sprintf("actualPrice = %d, dbPrice = %d", actualPrice, dbPrice.Price)
	}
}
