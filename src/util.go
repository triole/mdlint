package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func isFolder(target string) bool {
	info, err := os.Stat(target)
	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
	}
	if info.IsDir() {
		return true
	}
	return false
}

func find(basedir string, rxFilter string) []string {
	_, err := os.Stat(basedir)
	if err != nil {
		fmt.Printf("Fail access md folder %q\n", err)
		os.Exit(1)
	}
	filelist := []string{}
	rxf, _ := regexp.Compile(rxFilter)

	err = filepath.Walk(basedir, func(path string, f os.FileInfo, err error) error {
		if rxf.MatchString(path) {
			inf, err := os.Stat(path)
			if err == nil {
				if inf.IsDir() == false {
					filelist = append(filelist, path)
				}
			} else {
				print("Fail stat file %q", err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Fail find files %q\n", err)
		os.Exit(1)
	}
	return filelist
}

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
