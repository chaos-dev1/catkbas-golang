package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	read := flag.String("read", "", "read data in file")
	erase := flag.String("erase", "", "erase file data")
	write := flag.String("write", "", "write sth in file")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage -flag argument \n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *read != "" {
		file_data, err := os.ReadFile(*read)
		if err != nil {
			fmt.Print("fatal error! maybe file not found or im without rights! \n error: ", err)
			return
		}
		fmt.Println(string(file_data))
	}
	if *erase != "" {
		err := os.Remove(*erase)
		if err != nil {
			fmt.Print("fatal error! maybe file not found or im without rights! \n error: ", err)
		}
	}
	if *write != "" {

		fmt.Print("write your text>")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			text := scanner.Text()
			err := os.WriteFile(*write, []byte(text), 0644)
			if err != nil {
				fmt.Print("fatal error! maybe file not found or im without rights! \n error: ", err)
			}
		}

	}
}
