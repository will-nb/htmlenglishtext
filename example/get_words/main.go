package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/will-nb/htmlenglishtext/pkg/filter"
)

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
	jsonData, err := json.MarshalIndent(englishText.GetUniqueWords(), "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jsonFile, err := os.Create("unique_words.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
