package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var help string = "Usage: gora [word] [language_codes]"

var helpFlag = flag.Bool("help", false, help)

func RemoveBasicHTML(text string) string {

	re := regexp.MustCompile("<.*?>")
	return re.ReplaceAllString(text, "")

}

func GetRegexResult(content string, language_selection string) ([]string, error) {

	// tr <-> en
	pattern := fmt.Sprintf("(?s)<div class=\".*?\" id=\"dc_%s\" >(.*?)</div>", language_selection)

	re, err := regexp.Compile(pattern)

	if err != nil {
		return nil, err
	}

	results := re.FindStringSubmatch(content)

	return results, nil

}

func GetTranslation(word string, language_selection string) (string, error) {
	url := fmt.Sprintf("http://m.seslisozluk.com/?word=%s&lang=%s", word, language_selection)

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	// close it at the end of the execution of function.
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	content := string(body)

	result, err := GetRegexResult(content, language_selection)

	if len(result) == 0 {
		return "", fmt.Errorf("Translation not found.")
	}

	return RemoveBasicHTML(result[1]), nil

}

func main() {

	flag.Parse()

	if flag.NArg() == 0 && flag.NFlag() == 0 {
		fmt.Println(help)
		os.Exit(0)
	}

	if *helpFlag {
		fmt.Println(help)
		os.Exit(0)
	}

	word := flag.Arg(0)
	language_selection := "en_tr"

	if flag.NArg() == 2 {
		language_selection = flag.Arg(1)
	}

	translation, err := GetTranslation(word, language_selection)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(translation)
}
