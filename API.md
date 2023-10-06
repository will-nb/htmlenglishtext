## htmlenglishtext

`htmlenglishtext` 是一个用于处理英文文本的 Go 库，可以从 HTML 中提取文本、分解文本为句子和单词、过滤掉长度小于 3 的单词和罗马数字，以及去重单词等功能。

### EnglishText

`EnglishText` 是一个结构体，包含要过滤的文本。

#### ReadFile

`ReadFile` 读取文件内容并设置 `EnglishText` 结构体的 `html` 字段。

```go
func (f *EnglishText) ReadFile(filePath string) error
```

参数：

- `filePath`：要读取的文件路径。

返回值：

- `error`：如果读取文件时出现错误，则返回相应的错误信息。

#### SetHtml

`SetHtml` 设置 `EnglishText` 结构体的 `html` 字段。

```go
func (f *EnglishText) SetHtml(html string)
```

参数：

- `html`：要设置的 HTML 内容。

#### GetHtml

`GetHtml` 返回 `EnglishText` 结构体的 `html` 字段。

```go
func (f *EnglishText) GetHtml() string
```

返回值：

- `string`：HTML 内容。

#### FilterHTML

`FilterHTML` 从字符串中删除所有 HTML 标记并设置 `EnglishText` 结构体的 `text` 字段。

```go
func (f *EnglishText) FilterHTML(args ...string) error
```

参数：

- `args`：可选参数，用于指定要过滤的 HTML 内容。

返回值：

- `error`：如果过滤 HTML 时出现错误，则返回相应的错误信息。

#### GetText

`GetText` 返回 `EnglishText` 结构体的 `text` 字段。

```go
func (f *EnglishText) GetText() string
```

返回值：

- `string`：过滤后的文本。

#### ExtractSentences

`ExtractSentences` 从过滤后的文本中提取句子并设置 `EnglishText` 结构体的 `sentences` 字段。

```go
func (f *EnglishText) ExtractSentences() error
```

返回值：

- `error`：如果提取句子时出现错误，则返回相应的错误信息。

#### GetSentences

`GetSentences` 返回 `EnglishText` 结构体的 `sentences` 字段。

```go
func (f *EnglishText) GetSentences() []string
```

返回值：

- `[]string`：从过滤后的文本中提取的句子。

#### ExtractWords

`ExtractWords` 从过滤后的文本中提取单词并设置 `EnglishText` 结构体的 `words` 字段。

```go
func (f *EnglishText) ExtractWords() error
```

返回值：

- `error`：如果提取单词时出现错误，则返回相应的错误信息。

#### GetWords

`GetWords` 返回 `EnglishText` 结构体的 `words` 字段。

```go
func (f *EnglishText) GetWords() []string
```

返回值：

- `[]string`：从过滤后的文本中提取的单词。

#### UniqueWords

`UniqueWords` 去掉 `EnglishText` 结构体的 `words` 字段中的重复单词并设置 `uniqueWords` 字段。

```go
func (f *EnglishText) UniqueWords() error
```

返回值：

- `error`：如果去重时出现错误，则返回相应的错误信息。

#### GetUniqueWords

`GetUniqueWords` 返回 `EnglishText` 结构体的 `uniqueWords` 字段。

```go
func (f *EnglishText) GetUniqueWords() []string
```

返回值：

- `[]string`：去重后的单词。