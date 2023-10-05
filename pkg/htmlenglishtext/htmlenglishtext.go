package htmlenglishtext

import (
	"os"
	"regexp"
	"strings"
	"unicode"
)

// EnglishText是一个结构体，包含要过滤的文本。
type EnglishText struct {
	html        string
	text        string
	sentences   []string
	words       []string
	uniqueWords []string
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
func (f *EnglishText) GetText() string {
	return f.text
}

// 把文本分解为句子
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
func (f *EnglishText) GetSentences() []string {
	return f.sentences
}

// 从f.text中删除所有标点符号,提取所有单词并设置EnglishText结构体的words字段。
// 本函数会把首字母大写的单词转换为首字母小写。
// 本函数会过滤所有长度小于3的单词。
// 本函数会过滤所有罗马数字。
// 把文本分解为单词
func (f *EnglishText) ExtractWords() error {
	re := regexp.MustCompile("[^a-zA-Z'-]+|\\b[IVXLCDM]+\\b")
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

// contains 检查字符串数组中是否包含指定的字符串。
func contains(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

// 去掉所有重复的单词并返回结果。
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
func (f *EnglishText) GetUniqueWords() []string {
	return f.uniqueWords
}
