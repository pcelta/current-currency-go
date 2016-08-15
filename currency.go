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
	}
}

func main() {

	base := flag.String("base", "USD", "Currency Symbol")	
	flag.Parse()

	response, error := http.Get("http://api.fixer.io/latest?symbols=BRL&base="+*base)
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
			output(*base, message)
		}
	}
}

func output(base string, message *Message) {
	fmt.Printf("1 %s is equivalent to R$%f\n", base, message.Rates.Brl)
}
