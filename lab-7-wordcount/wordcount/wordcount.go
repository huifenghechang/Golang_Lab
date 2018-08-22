package main


import (
	"unicode"
	"strings"
	"unicode/utf8"
	"fmt"
	"sort"
)

type Pair struct {
	Key string
	Value int
}

//PairList实现了Sort接口
type PairList []Pair

func (p PairList)Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}

func (p PairList)Len()int  {
	return len(p)
}

func (p PairList)Less(i,j int) bool{
	return p[j].Value < p[i].Value
}

//提取单词

/*
Go相关知识：
	Go语言中byte和rune实质上就是uint8和int32类型。
    	- byte用来强调数据是raw data，而不是数字；
		- 而rune用来表示Unicode的code point
*/

func SplitOnNonLetters(s string) []string  {
	notALetter := func(char rune) bool {
		//  IsLetter 判断 r 是否为一个字母字符 (类别 L)
		return !unicode.IsLetter(char)
	}
	//  strings.Fields("  foo bar  baz   ")) //["foo" "bar" "baz"] 返回一个列表
	return strings.FieldsFunc(s,notALetter)
}

/*
基于Map实现了类型wordCount,并对其实现了Merge（）、Report、SortReport、UpdateFreq()等方法
*/

type WordCount map[string]int

//用于合并统计两个wordCount
func (source WordCount)Merge(wordCount WordCount)WordCount  {
	for k,v := range wordCount{
		source[k] += v
	}

	return source
}

//打印词频统计情况
func(wordCount WordCount)Report(){
	words := make([]string,0,len(wordCount))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range wordCount{
		words = append(words, word)
		// utf8.RuneCountInString 用来统计字符串中字符的个数
		if width := utf8.RuneCountInString(word);width> wordWidth{
			wordWidth = width
		}
		// frequency表示词频。fmt.Sprint，将结果以字符串的形式返回。
		if width := len(fmt.Sprint(frequency));width > frequencyWidth{
			frequencyWidth = width
		}
	}
	// sort.Strings ,对字符进行升序排序！
	sort.Strings(words)

	gap := wordWidth + frequencyWidth - len("Word")-len("Frequency")

	fmt.Printf("Word %*s%s\n",gap," ","Frequency")
	for _,word := range words{
		fmt.Printf("%-*s%*d\n",wordWidth,word,frequencyWidth,wordCount[word])
	}
}

func main(){
	workCount := make(WordCount)
	workCount["suqiancheng"] = 11
	workCount["lianghui"] = 8
	workCount.Report()

}
