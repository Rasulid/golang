package main

import (
	"fmt"
	"regexp"
)

func main() {
	//reStr := "The ape was at the apex"
	//match, _ := regexp.MatchString("(ape[^ ]?", reStr)
	//fmt.Println(match)

	reStr := "Cat rat mat fat pat"
	re, _ := regexp.Compile("([crmfp]at)")
	fmt.Println("Metchstring", re.MatchString(reStr))
	fmt.Println("FindString", re.FindString(reStr))
	fmt.Println("FindStringIndex", re.FindStringIndex(reStr))
	fmt.Println("FingAllString", re.FindAllString(reStr, -1))
	fmt.Println("1st to secomd", re.FindAllString(reStr, 2))
	fmt.Println("All submuich index", re.FindAllStringIndex(reStr, -1))
	fmt.Println("Replace all string", re.ReplaceAllString(reStr, "Dog"))

}
