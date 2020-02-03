package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func scanDir(dirName string) ([]string, []string, error) {
	// check if dir exists
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		log.Errorf("Directory %s does not exist", dirName)
		return []string{}, []string{}, err
	}
	// check if dirName ends with slash
	if !strings.HasSuffix(dirName, "/") {
		dirName += "/"
	}

	var files, dirs []string

	dirFiles, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Errorf("Unable to scan dir %s: %v", dirName, err)
		return []string{}, []string{}, err
	}

	for _, f := range dirFiles {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
			continue
		}
		files = append(files, f.Name())
	}

	/*
		err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				dirs = append(dirs, strings.Replace(path, dirName, "", -1))
				return nil
			}

			files = append(files, strings.Replace(path, dirName, "", -1))
			return nil
		})

		if err != nil {
			log.Errorf("Unable to scan dir %s: %v", dirName, err)
			return []string{}, []string{}, err
		}
	*/

	log.Infof("Successfully scanned %s", dirName)
	return files, dirs, nil
}

func getVideos(dirName string) []string {
	videoExtensions := []string{".mp4", ".mkv", ".mov", ".avi", "flv", "wmv"}
	videoFiles := getFilesByType(dirName, videoExtensions)
	return videoFiles
}

func getFilesByType(dirName string, extensions []string) []string {
	// check if dir exists
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		log.Errorf("Directory %s does not exist", dirName)
		return []string{}
	}
	// check if dirName ends with slash
	if !strings.HasSuffix(dirName, "/") {
		dirName += "/"
	}

	var files []string

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if stringInSlice(filepath.Ext(path), extensions) {
			files = append(files, strings.Replace(path, dirName, "", -1))
		}
		return nil
	})

	if err != nil {
		log.Errorf("Unable to scan dir %s: %v", dirName, err)
		return []string{}
	}

	log.Infof("Successfully scanned %s", dirName)
	return files
}

func stringInSlice(str string, list []string) bool {
	for _, x := range list {
		if x == str {
			return true
		}
	}
	return false
}
