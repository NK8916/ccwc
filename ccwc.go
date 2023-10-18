package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	valid_map := make(map[string]bool)
	valid_map["-l"] = true
	valid_map["-c"] = true
	valid_map["-w"] = true
	is_valid, invalid_arg := check_for_valid_args(valid_map, args)
	if !is_valid {
		log.Fatal("Invalid arguments ", invalid_arg)
	}
	fmt.Println(args, valid_map)
}

func check_for_valid_args(valid_map map[string]bool, args []string) (bool, string) {

	for _, arg := range args {
		if !valid_map[arg] {
			return false, arg
		}
	}
	return true, ""
}
