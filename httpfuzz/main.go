package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var (
	ErrNameRequired = errors.New("Name is required")

	ErrEmailRequired = errors.New("Email is required")

	ErrEmailInvalid = errors.New("Email is invalid")

	ErrPwdRequired = errors.New("Password is required")

	ErrPwdMinChars = errors.New("Password minimum is 6 chars")

	ErrAgeMin = errors.New("Minimum age is 18")

	ErrInvalidJSON = errors.New("Invalid JSON")

	ErrInvalidPayload = errors.New("Invalid payload")
)

type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint8  `json:"age"`
}

func (p *Person) IsValid() error {
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Email == "" {
		return ErrEmailRequired
	} else {
		rgx := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:.[a-zA-Z0-9-]+)*$")
		if !rgx.MatchString(p.Email) {
			return ErrEmailInvalid
		}
	}
	if p.Password == "" {
		return ErrPwdRequired
	}
	if len(p.Password) < 6 {
		return ErrPwdMinChars
	}
	if p.Age < 18 {
		return ErrAgeMin
	}
	return nil
}

func validate(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInsufficientStorage)
		return
	}

	// if !json.Valid(data) {
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprintf(rw, ErrInvalidJSON.Error())
	// 	return
	// }

	var person Person

	err = json.Unmarshal(data, &person)
	if err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(rw, ErrInvalidPayload.Error())
		return
	}

	err = person.IsValid()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, err.Error())
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "")
}

func main() {
	http.HandleFunc("/person/validate", validate)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
