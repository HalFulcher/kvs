package store

import (
	"testing"
)

func TestGetEntry(t *testing.T) {

	AddEntry("James", "078383838", "username1")
	AddEntry("Sarah", "074893983", "username2")
	result := GetEntry("Sarah")

	if result != "074893983" {
		t.Errorf("got %q, wanted %q", result, "074893983")
	}
}

func TestGetAllEntries(t *testing.T) {
	AddEntry("James", "078383838", "username1")
	AddEntry("Sarah", "074893983", "username2")
	result := GetAllEntries()

	if len(result) != 2 {
		t.Errorf("got %v, wanted %v", result, "2")
	}
}

func TestAddEntry(t *testing.T) {

	AddEntry("James", "078383838", "username1")
	result := GetEntry("James")

	if result != "078383838" {
		t.Errorf("got %v, wanted %q", result, "078383838")
	}
}

func TestDeleteEntry(t *testing.T) {
	AddEntry("James", "078383838", "username1")
	AddEntry("Sarah", "074893983", "username2")
	DeleteEntry("James", "username1")

	result := false
	for k, v := range DB {
		if k == "James" && v.Value == "075383838" {
			result = true
		}
	}
	if result != false {
		t.Errorf("got %v, wanted %v", result, false)
	}
}
