package main

import "fmt"
import "net/http"
import "flag"
import "encoding/json"

type Message struct {
	Base string
	Date string
	Rates struct {
		Brl float32 `json:"BRL"`
		Cad float32 `json:"CAD"`
	}
}

func main() {

	base := flag.String("base", "USD", "Currency Symbol")	
	flag.Parse()

	fmt.Println("Recovering current situation...")
	response, error := http.Get("http://api.fixer.io/latest?base="+*base)
	defer response.Body.Close()
	if error != nil {
		fmt.Println(error)
	} else {
		body := response.Body
		message := new(Message)
		err := json.NewDecoder(body).Decode(message)
		if err != nil {
			fmt.Println(err)
		} else {
			output(message)
		}
	}
}

func output(message *Message) {
	fmt.Printf("Dolar Canadense: %f\n", message.Rates.Cad);
	fmt.Printf("Real: %f", message.Rates.Brl);
}
