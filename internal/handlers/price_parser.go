package handlers

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func PriceParser(address string) (int, error) {
	resp, err := http.Get(address)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("can't read body", err)
		return 0, err
	}

	r, _ := regexp.Compile(`content="(\d{6,9})\.\d+"`)
	preRes := r.FindString(string(body))

	r2, _ := regexp.Compile(`(\d{6,9})`)
	result := r2.FindString(preRes)

	intResult, err := strconv.Atoi(result)
	if err != nil {
		log.Println("can't ascii to integer", err)
		return 0, err
	}
	//log.Println(intResult)
	return intResult, nil
}
