package main

import (
	"strconv"
	"regexp"
	"fmt"
	"github.com/fatih/color"
)

func convertStringToInt(str string) (int, error) {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	num, err := strconv.Atoi(nonAlphanumericRegex.ReplaceAllString(str, ""))
	
	if err != nil{
		return -1, err
	}
	
	return num, nil
}

func convertIntToCoords(pos int) string{
	return fmt.Sprintf("row %d, col %d", (int)(pos / 8), pos % 8)
}

func colorPrint(text string, c1 color.Attribute, c2 color.Attribute){
	c := color.New(c1).Add(c2)
	c.Print(text)
}