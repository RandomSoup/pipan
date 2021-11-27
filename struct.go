package main

type PipanProfile struct {
	// Core //

	Name   string
	Modded bool

	// Preferences //

	RendDist string
	Username string
	Features map[string]bool
}
