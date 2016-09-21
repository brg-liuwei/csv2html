package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var csvFile string
var reg *regexp.Regexp = regexp.MustCompile("(\\s+)")

func init() {
	flag.StringVar(&csvFile, "csv", "", "csv file")
	flag.Parse()
}

func split(line string) []string {
	regLine := reg.ReplaceAllString(line, " ")
	finalLine := strings.TrimSpace(regLine)
	return strings.Split(finalLine, " ")
}

func main() {
	if len(csvFile) == 0 {
		flag.Usage()
		return
	}

	f, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	t := NewTableGen()

	// read first line
	line, prefix, err := reader.ReadLine()
	if err != nil {
		if err == io.EOF {
			return
		} else {
			panic(err)
		}
	}
	if prefix {
		panic("too long a line")
	}
	t.AddHeader(split(string(line))...)

	for {
		line, prefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		if prefix {
			panic("too long a line")
		}
		t.AddBody(split(string(line))...)
	}

	fmt.Println(string(t.Gen()))
}
