package main

import (
	"errors"
	"fmt"
	"github.com/tendryll/goassert/validate"
	"strings"
)

type Latitude struct {
	Degrees   float64 `json:"degrees"assert:"required=true,min=0.0,max=90.0"`
	Direction string  `json:"direction"assert:"required=true,pattern:^(N|S)$"`
}

type Location struct {
	Latitude  Latitude `json:"latitude"assert:"required=true"`
	Longitude float64  `json:"longitude"assert:"required=true,min=-180.0,max=180.0"`
}

type Address struct {
	Address1 string   `json:"address1"assert:"required=true"`
	Address2 string   `json:"address2"`
	Address3 string   `json:"address3"`
	City     string   `json:"city"`
	State    string   `json:"state"assert:"required=true"`
	Country  string   `json:"country"assert:"required=true,maxlength:3"`
	ZipCode  string   `json:"zipcode"assert:"required=true"`
	Location Location `json:"location"`
}

type Person struct {
	FirstName  string     `json:"firstName"assert:"required=true"`
	MiddleInit string     `json:"middleInitial"assert:"required=true,minlength=1,maxlength=1"`
	LastName   string     `json:"lastName"assert:"required=true"`
	Address    []*Address `json:"address"assert:"required=true"`
}

func main() {
	p := Person{
		FirstName:  "James",
		MiddleInit: "T",
		LastName:   "Kirk",
		Address: []*Address{
			{
				Address1: "755 Crossover Lane",
				City:     "Memphis",
				State:    "TN",
				Country:  "USA",
				ZipCode:  "38107",
				Location: Location{
					Latitude: Latitude{
						// Degrees:   35.1098212,
						Direction: "N",
					},
					Longitude: -89.9077976,
				},
			},
		},
	}

	fmt.Printf("%+v\n", validate.Validate(p))

	//t := `"required=true,min=10"`
	//
	//tagPairs := strings.Split(t, ",")
	//m := make(map[string]string)
	//
	//for _, pair := range tagPairs {
	//	key, value, err := asKeyValue(pair, "=")
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	m[key] = value
	//}
	//
	//fmt.Println(m)
}

func asKeyValue(pair, sep string) (string, string, error) {
	if pair == "" {
		return "", "", errors.New("pair is an empty string")
	}

	if sep == "" {
		return "", "", errors.New("sep is an empty string")
	}

	if !strings.Contains(pair, sep) {
		return "", "", errors.New(fmt.Sprintf("%s doesn't contain the '%s' separator", pair, sep))
	}

	keyValue := strings.Split(pair, "=")

	if len(keyValue) != 2 {
		return "", "", errors.New("pair doesn't contain both a key and value")
	}

	return keyValue[0], keyValue[1], nil
}
