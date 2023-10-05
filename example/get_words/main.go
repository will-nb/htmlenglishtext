package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/will-nb/htmlenglishtext/pkg/filter"
)

func saveJSONFile(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	jsonFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// 读取命令行参数中的文件名
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: extract-english-text <filename>")
		os.Exit(1)
	}
	filename := args[0]

	// 读取 HTML 文件并提取其中的英文单词
	englishText := filter.EnglishText{}
	err := englishText.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = englishText.FilterHTML()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = englishText.ExtractSentences()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = englishText.ExtractWords()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = englishText.UniqueWords()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 将提取的 uniqueWords 保存成一个 json 文件
	err = saveJSONFile("unique_words.json", englishText.GetUniqueWords())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 将 sentence 保存成一个 json 文件
	err = saveJSONFile("sentence.json", englishText.GetSentences())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
