package main

type Emulator struct {
	paused  bool
	running bool
	ticks   uint64
}
