package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func FuzzValidatePaylod(f *testing.F) {
	testcases := []Person{
		{"Tiago Temporin", "tiago@aprendagolang.com.br", "1234", 32},
		{"Maria Castro", "maria.castro@algumdominio.com", "1Av6s#", 22},
		{"Daniela Fernandez", "dani@teste.com.br", "1234", 16},
	}

	for _, tc := range testcases {
		data, _ := json.Marshal(tc)

		f.Add(data)
	}

	knowErrs := map[string]bool{
		ErrNameRequired.Error():   true,
		ErrEmailRequired.Error():  true,
		ErrPwdRequired.Error():    true,
		ErrPwdMinChars.Error():    true,
		ErrAgeMin.Error():         true,
		ErrInvalidPayload.Error(): true,
	}

	srv := httptest.NewServer(http.HandlerFunc(validate))
	defer srv.Close()

	f.Fuzz(func(t *testing.T, data []byte) {
		resp, err := http.DefaultClient.Post(srv.URL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			if _, ok := knowErrs[string(body)]; ok {
				t.Skip(fmt.Sprintf("skiping knowing error: %s", body))
			}

			t.Errorf("Expected status code %d, got %d with error :%s", http.StatusOK, resp.StatusCode, body)
		}
	})
}

func FuzzValidateData(f *testing.F) {
	testcases := []Person{
		{"Tiago Temporin", "tiago@aprendagolang.com.br", "1234", 32},
		{"Maria Castro", "maria.castro@algumdominio.com", "1Av6s#", 22},
		{"Daniela Fernandez", "dani@teste.com.br", "1234", 16},
	}

	for _, tc := range testcases {
		f.Add(tc.Name, tc.Email, tc.Password, tc.Age)
	}

	knowErrs := map[string]bool{
		ErrNameRequired.Error():   true,
		ErrEmailRequired.Error():  true,
		ErrEmailInvalid.Error():   true,
		ErrPwdRequired.Error():    true,
		ErrPwdMinChars.Error():    true,
		ErrAgeMin.Error():         true,
		ErrInvalidPayload.Error(): true,
	}

	srv := httptest.NewServer(http.HandlerFunc(validate))
	defer srv.Close()

	f.Fuzz(func(t *testing.T, name, email, password string, age uint8) {
		p := Person{
			Name:     name,
			Email:    email,
			Password: password,
			Age:      age,
		}

		data, _ := json.Marshal(p)

		resp, err := http.DefaultClient.Post(srv.URL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			if _, ok := knowErrs[string(body)]; ok {
				t.Skip(fmt.Sprintf("skiping knowing error: %s", body))
			}

			t.Errorf("Expected status code %d, got %d with error :%s", http.StatusOK, resp.StatusCode, body)
		}
	})
}
