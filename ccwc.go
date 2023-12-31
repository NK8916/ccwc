package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	valid_map := make(map[string]bool)
	valid_map["-l"] = true
	valid_map["-c"] = true
	valid_map["-w"] = true
	filepath := args[len(args)-1]
	args = args[:len(args)-1]
	is_valid, invalid_arg := check_for_valid_args(valid_map, args)
	if !is_valid {
		log.Fatal("Invalid argument(s) ", invalid_arg)
	} else {
		result := evaluate_file(args, filepath)
		fmt.Println(filepath, " ", result)
	}
}

func check_for_valid_args(valid_map map[string]bool, args []string) (bool, string) {
	for _, arg := range args {
		if !valid_map[arg] {
			return false, arg
		}
	}
	return true, ""
}

func evaluate_file(ops []string, file_path string) string {
	default_ops := []string{"-c", "-l", "-w"}
	result := []string{}
	if len(ops) == 0 {
		ops = default_ops
	}
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal("Error while opening file ", err)
	} else {
		{
			for _, op := range ops {
				if op == "-c" {
					fi, err := file.Stat()
					if err != nil {
						log.Fatal("Error while fetching file stats ", err)
					}
					result = append(result, strconv.Itoa(int(fi.Size())))
				} else {
					data, err := os.ReadFile(file_path)
					if err != nil {
						log.Fatal("Error while reading file ", err)
					}
					string_data := string(data)
					data_array := strings.Split(string_data, "\n")
					if op == "-l" {
						result = append(result, strconv.Itoa(len(data_array)))
					} else {
						word_count := 0
						for _, line := range data_array {
							word_array := strings.Split(line, " ")
							for _, word := range word_array {
								if string(word) != "" && string(word) != " " {
									word_count += 1
								}
							}

						}
						result = append(result, strconv.Itoa(word_count))
					}
				}
			}
		}
	}
	return strings.Join(result, " ")

}
