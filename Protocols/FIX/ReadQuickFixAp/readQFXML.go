package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/bhbosman/Application/Protocols/FIX/ReadQuickFix"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	readers, err := GetInput(args)
	if err != nil {
		os.Exit(101)
	}
	if len(readers) > 0 {
		xmlDecoder := xml.NewDecoder(readers[0])
		fix := &ReadQuickFix.FixDecl{}
		err := xmlDecoder.Decode(fix)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(len(fix.Header.Fields))
		fmt.Println(len(fix.Trailer.Fields))
		fmt.Println(len(fix.Components.Fields))
		fmt.Println(len(fix.Fields.Fields))
	}
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetInput(args []string) ([]io.ReadCloser, error) {
	if len(args) > 0 {
		list := make([]io.ReadCloser, 0, len(args))
		for _, item := range args {
			if Exists(item) {
				file, err := os.Open(item)
				if err != nil {
					return nil, err
				}
				list = append(list, file)
				continue
			} else {
				readerCloser := ioutil.NopCloser(strings.NewReader(item))
				list = append(list, readerCloser)
			}
		}
		return list, nil
	}
	return nil, fmt.Errorf("no input files or strings")
}
