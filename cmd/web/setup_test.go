package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	// Set up the path to templates var for all handler tests
	pathToTemplates = "./../../templates/"
	// create a session for our tests to use
	app.Session = getSession()

	os.Exit(m.Run())
}
