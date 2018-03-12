// Package 404
// 9 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package fnftest

import (
	"testing"

	baloo "gopkg.in/h2non/baloo.v3"
)

var test = baloo.New("http://localhost:8080")

func TestBaloo404Handler(t *testing.T) {
	test.Get("/bad").
		Expect(t).
		Status(404).
		Done()
}
