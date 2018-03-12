// Package handler
// 9 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package handler

import (
	"os"
	"testing"

	"github.com/KaiserGald/logger"
)

var lg *logger.Logger

func TestMain(m *testing.M) {
	lg = logger.New()
	r := m.Run()
	os.Exit(r)
}

func TestStart(t *testing.T) {
	err := Start(lg)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddRoute(t *testing.T) {

}

func TestHandle(t *testing.T) {

}
