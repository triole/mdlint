package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
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

func makeAlphaIterator(m map[string]string) (arr []string) {
	for k := range m {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	return
}

func rxFind(rx string, str string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(str)
	return
}

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
