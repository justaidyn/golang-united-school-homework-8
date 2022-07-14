package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   uint   `json:"age"`
}

type Arguments map[string]string

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	flagOperation := flag.String("operation", "", "")
	flagItem := flag.String("item", "", "")
	flagFileName := flag.String("fileName", "", "")

	flag.Parse()

	return Arguments{
		"operation": *flagOperation,
		"item":      *flagItem,
		"fileName":  *flagFileName,
	}
}
func Perform(args Arguments, writer io.Writer) error {
	id := args["id"]
	// item := args["item"]
	fileName := args["fileName"]
	operation := args["operation"]

	//check operation
	if operation == "" {
		return errors.New("-operation flag has to be specified")
	}

	switch operation {
	case "list":
		if args["fileName"] == "" {
			return errors.New("-fileName flag has to be specified")
		}
		return list(fileName, writer)
	case "add":
		if args["fileName"] == "" {
			return errors.New("-fileName flag has to be specified")
		}
		if args["item"] == "" {
			return errors.New("-item flag has to be specified")
		}
		return nil
	case "remove":
		if args["fileName"] == "" {
			return errors.New("-fileName flag has to be specified")
		}
		return nil
	case "findById":
		if args["fileName"] == "" {
			return errors.New("-fileName flag has to be specified")
		}
		if id == "" {
			return errors.New("-id flag has to be specified")
		}
		return findById(id, fileName, writer)
	default:
		return errors.New("Operation " + operation + " not allowed!")
	}
}

func list(fileName string, writer io.Writer) error {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	writer.Write(bytes)

	return nil
}

func findById(id, fileName string, writer io.Writer) error {
	var users []User
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed read file: %v", err)
	}

	if err := json.Unmarshal(bytes, &users); err == nil {
		for _, user := range users {
			if user.Id == id {
				bytes, err := json.Marshal(user)
				if err != nil {
					return fmt.Errorf("failed to marshal user as JSON: %v", err)
				}

				writer.Write(bytes)
			}
		}
	}

	return nil
}
