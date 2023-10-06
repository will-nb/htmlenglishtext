package htmlenglishtext

import (
	"os"
	"regexp"
	"strings"
	"unicode"
)

// EnglishText是一个结构体，包含要过滤的文本。
type EnglishText struct {
	html        string   // HTML 内容。
	text        string   // 过滤后的文本。
	sentences   []string // 从过滤后的文本中提取的句子。
	words       []string // 从过滤后的文本中提取的单词。
	uniqueWords []string // 去重后的单词。
}

// ReadFile 读取文件内容并设置 EnglishText 结构体的 html 字段。
func (f *EnglishText) ReadFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	f.html = string(content)
	return nil
}

// SetHtml 设置 EnglishText 结构体的 html 字段。
func (f *EnglishText) SetHtml(html string) {
	f.html = html
}

// GetHtml 返回 EnglishText 结构体的 html 字段。
func (f *EnglishText) GetHtml() string {
	return f.html
}

// FilterHTML 从字符串中删除所有 HTML 标记并设置 EnglishText 结构体的 text 字段。
// 如果提供了 html 参数，则将其用作要过滤的文本，否则将使用 EnglishText 结构体的 html 字段中的文本。
//
// 参数：
//
//	args: 可选参数，用于指定要过滤的 HTML 内容。
func (f *EnglishText) FilterHTML(args ...string) error {
	var html string
	if len(args) > 0 {
		f.html = args[0]
	}
	html = f.html
	re := regexp.MustCompile("<[^>]*>")
	f.text = re.ReplaceAllString(html, "")
	return nil
}

// GetText 返回 EnglishText 结构体的 text 字段。
//
// 返回值：
//
//	string: 过滤后的文本。
func (f *EnglishText) GetText() string {
	return f.text
}

// ExtractSentences 从过滤后的文本中提取句子并设置 EnglishText 结构体的 sentences 字段。
//
// 返回值：
//
//	error: 如果提取句子时出现错误，则返回相应的错误信息。
func (f *EnglishText) ExtractSentences() error {
	re := regexp.MustCompile(`\b[A-Z][^.\n?!]*[.?!]`)
	f.sentences = re.FindAllString(f.text, -1)
	for i := 0; i < len(f.sentences); i++ {
		if !strings.Contains(f.sentences[i], " ") {
			f.sentences = append(f.sentences[:i], f.sentences[i+1:]...)
			i--
		}
	}
	return nil
}

// GetSentences 返回 EnglishText 结构体的 sentences 字段。
//
// 返回值：
//
//	[]string: 从过滤后的文本中提取的句子。
func (f *EnglishText) GetSentences() []string {
	return f.sentences
}

// ExtractWords 从过滤后的文本中提取单词并设置 EnglishText 结构体的 words 字段。
// 本函数会把首字母大写的单词转换为首字母小写。
// 本函数会过滤所有长度小于3的单词。
// 本函数会过滤所有罗马数字。
//
// 返回值：
//
//	error: 如果提取单词时出现错误，则返回相应的错误信息。
func (f *EnglishText) ExtractWords() error {
	re := regexp.MustCompile("[^a-zA-Z-]+|\\b[IVXLCDM]+\\b")
	words := re.Split(f.text, -1)
	var filteredWords []string
	for _, word := range words {
		if len(word) >= 3 && !unicode.IsDigit([]rune(word)[0]) {
			if strings.ToUpper(word) == word {
				filteredWords = append(filteredWords, word)
			} else {
				filteredWords = append(filteredWords, strings.ToLower(word[:1])+word[1:])
			}
		}
	}
	f.words = filteredWords
	return nil
}

// GetWords 返回 EnglishText 结构体的 words 字段。
//
// 返回值：
//
//	[]string: 从过滤后的文本中提取的单词。
func (f *EnglishText) GetWords() []string {
	return f.words
}

// contains 检查字符串数组中是否包含指定的字符串。
func contains(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

// UniqueWords 去掉 EnglishText 结构体的 words 字段中的重复单词并设置 uniqueWords 字段。
//
// 返回值：
//
//	error: 如果去重时出现错误，则返回相应的错误信息。
func (f *EnglishText) UniqueWords() error {
	var uniqueWords []string
	for _, word := range f.words {
		if !contains(uniqueWords, word) {
			uniqueWords = append(uniqueWords, word)
		}
	}
	f.uniqueWords = uniqueWords
	return nil
}

// GetUniqueWords 返回 EnglishText 结构体的 uniqueWords 字段。
//
// 返回值：
//
//	[]string: 去重后的单词。
func (f *EnglishText) GetUniqueWords() []string {
	return f.uniqueWords
}
