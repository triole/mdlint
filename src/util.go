package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

func isFolder(target string) bool {
	info, err := os.Stat(target)
	if os.IsNotExist(err) {
		fmt.Printf("file does not exist: %s", err)
		os.Exit(1)
	}
	if info.IsDir() {
		return true
	}
	return false
}

func find(basedir string, rxFilter string) []string {
	_, err := os.Stat(basedir)
	if err != nil {
		fmt.Printf("can not access folder %q\n", err)
		os.Exit(1)
	}
	filelist := []string{}
	rxf, _ := regexp.Compile(rxFilter)

	err = filepath.Walk(basedir, func(path string, f os.FileInfo, err error) error {
		if rxf.MatchString(path) {
			inf, err := os.Stat(path)
			if err == nil {
				if !inf.IsDir() {
					filelist = append(filelist, path)
				}
			} else {
				print("stat file failed %q", err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("unable to detect files %q\n", err)
		os.Exit(1)
	}
	return filelist
}

func makeAlphaIteratorItf(m map[string]interface{}) (arr []string) {
	for k := range m {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	return
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
