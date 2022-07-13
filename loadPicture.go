package main

import "io/ioutil"

type pictures []string

func LoadPictures(n string) (pictures, error) {
	files, err := ioutil.ReadDir(n)
	if err != nil {
		return nil, err
	}

	var pictures []string
	for _, file := range files {
		pictures = append(pictures, n+"/"+file.Name())
	}

	return pictures, nil
}
