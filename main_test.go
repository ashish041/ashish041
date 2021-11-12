package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"testing"
	"time"
	"util/encode"
)

var httpClient = &http.Client{Timeout: 5 * time.Second}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func fetchData(url string) (string, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		log.Printf("Error client.Get(%s): %v", url, err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error ioutil.ReadAll(%v): %v", resp.Body, err)
		return "", err
	}
	return string(body), nil
}

func TestSimplePath(t *testing.T) {
	for i := 5; i < 100; i = i + 5 {
		expected := RandomString(i)
		url := fmt.Sprintf("http://localhost:3333/%s", expected)
		got, err := fetchData(url)
		if err != nil {
			t.Error(
				"For", url,
				"expected", expected,
				"got", err,
			)
		}
		if !reflect.DeepEqual(got, expected) {
			t.Error(
				"For", url,
				"expected", expected,
				"got", got,
			)
		}
	}
}

func TestEncoding_1(t *testing.T) {
	_for := "discovergy"
	expected := "joyiubkxme"
	got := encode.EncodeString(_for)
	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"For", _for,
			"expected", expected,
			"got", got,
		)
	}
}

func TestDecoding_1(t *testing.T) {
	_for := "joyiubkxme"
	expected := "discovergy"
	got := encode.DecodeString(_for)
	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"For", _for,
			"expected", expected,
			"got", got,
		)
	}
}

func TestEncodingDecoding(t *testing.T) {
	for i := 5; i < 100; i = i + 5 {
		_for := RandomString(i)
		expected := _for
		enc := encode.EncodeString(_for)
		got := encode.DecodeString(enc)
		if !reflect.DeepEqual(got, expected) {
			t.Error(
				"For", _for,
				"expected", expected,
				"got", got,
			)
		}
	}
}
