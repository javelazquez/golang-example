package server

import "time"

type ConfigServer struct {
	Port            string
	IdleTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
	Env             string `env:"ENV"`
}
