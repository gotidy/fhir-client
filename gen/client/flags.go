package main

import "flag"

func ParseFlags() Config {
	config := Config{}
	flag.StringVar(&(config.Output), "o", "", "Output file, else output to STDOUT")
	flag.Parse()
	return config
}
