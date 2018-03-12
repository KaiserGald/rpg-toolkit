// Package index
// 9 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package index

import (
	"testing"

	"gopkg.in/h2non/baloo.v3"
)

var test = baloo.New("http://localhost:8080")

func TestBalooIndexHandler(t *testing.T) {
	test.Get("/").
		Expect(t).
		Status(200).
		Done()
}
