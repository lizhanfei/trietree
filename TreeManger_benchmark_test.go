package trietree_test

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"
	"github.com/lizhanfei/trietree"
)

func BenchmarkTreeManager_Search(b *testing.B) {
	var treemanager trietree.TreeManager
	initAppendWord(&treemanager)
	words := getTestSearchWord()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		treemanager.Search(words[rand.Intn(3)])
	}
}

/**
获取测试用户的词
 */
func getTestSearchWord() (words []string) {
	filerPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败 ")
		return
	}
	file, err := os.Open(filerPath + "/testSearch.txt")
	if err != nil {
		fmt.Printf("打开文件失败 ")
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		data, _, err := reader.ReadLine()
		if err != nil {
			break;
		}
		words = append(words, string(data))
	}
	return
}

func initAppendWord(treeManager *trietree.TreeManager) {
	filerPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败 ")
		return
	}
	file, err := os.Open(filerPath + "/word_bench.txt")
	if err != nil {
		fmt.Printf("打开文件失败 ")
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		data, _, err := reader.ReadLine()
		if err != nil {
			break;
		}
		treeManager.Append(string(data))
	}
}
