// Package main
// 9 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package main

import (
	"os"
	"testing"

	"github.com/KaiserGald/logger"
	"github.com/KaiserGald/unlicht-server/daemon"
)

func TestMain(m *testing.M) {
	log = logger.New()
	r := m.Run()
	os.Exit(r)
}

func TestProcessCLI(t *testing.T) {
	os.Setenv("RUN_ENV", "DEV")
	os.Args = []string{"cmd", "-c", "-v", "-q", "-p=2112"}
	exp := &daemon.Config{
		DevMode:    true,
		ListenSpec: ":2112",
	}
	res := processCLI()
	if !color {
		t.Errorf("Color flag not parsed.")
	}

	if !verbose {
		t.Errorf("Verbose flag not parsed.")
	}

	if !quiet {
		t.Errorf("Quiet flag not parsed.")
	}
	if port != 2112 {
		t.Errorf("Port flag not parsed.")
	}

	if res.DevMode != exp.DevMode {
		t.Errorf("Daemon config DevMode not correct. Expected '%v' got '%v'.", exp.DevMode, res.DevMode)
	}

	if res.ListenSpec != exp.ListenSpec {
		t.Errorf("Daemon config ListenSpec not correct. Expected '%v' got '%v'.", exp.ListenSpec, res.ListenSpec)
	}
}

func TestConfigureDaemonInDevMode(t *testing.T) {
	cfg := &daemon.Config{}
	os.Setenv("RUN_ENV", "DEV")
	configureDaemon(cfg)
	if !cfg.DevMode {
		t.Errorf("Daemon not configured correctly. Expected 'true' got %v.\n", cfg.DevMode)
	}
}

func TestConfigureDaemonInProdMode(t *testing.T) {
	cfg := &daemon.Config{}
	os.Setenv("RUN_ENV", "PROD")
	configureDaemon(cfg)
	if cfg.DevMode {
		t.Errorf("Daemon not configured correctly. Expected 'false' got %v.\n", cfg.DevMode)
	}
}
