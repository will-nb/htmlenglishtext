# HTML 英文文本处理库

这是一个 Golang 库，提供了对 HTML 中的英文文本进行处理所需的所有功能。

## 功能列表

- 从 HTML 中提取英文单词
- 去除 HTML 标记
- 将文本分解为句子
- 提取所有单词并去除标点符号和罗马数字
- 去除所有重复的单词


### 安装

使用以下命令安装本库：

```bash
go get github.com/will-nb/htmlenglishtext
```
## 示例代码
以下是一个示例代码，演示了如何使用本库提取 HTML 中的英文单词并保存成一个 JSON 文件：
```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"

    "github.com/will-nb/htmlenglishtext"
)

func main() {
    // 读取 HTML 文件并提取其中的英文单词
    englishText := &htmlenglishtext.EnglishText{}
    err := englishText.ReadFile("example.html")
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
    err = englishText.GetUniqueWords()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // 将提取的 uniqueWords 保存成一个 json 文件
    jsonData, err := json.MarshalIndent(englishText.UniqueWords, "", "  ")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    err = ioutil.WriteFile("unique_words.json", jsonData, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

## 命令行工具
本库还提供了一个命令行工具，用于从 HTML 文件中提取英文单词并保存成一个 JSON 文件，它的源码在example/get_words下，可执行文件在dist文件夹下。

使用以下命令运行命令行工具：
```shell
get_words test.html
```