package trietree_test

import (
	"testing"
	"github.com/lizhanfei/trietree"
)

/**
添加，简单检索
添加一个词
构造一个成功检索
一个失败检索
*/
func TestAppendSimple(t *testing.T) {
	var tree trietree.TreeManager
	ret := tree.Append("abcd")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("abcd")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("")
	if ret != false {
		t.Error("加入失败", ret)
	}

	serachRet := tree.Search("abcdefg")
	if len(serachRet) != 1 || serachRet[0] != "abcd" {
		t.Error("search  异常， 结果： ", serachRet)
	}

	serachRet = tree.Search("abc")
	if len(serachRet) != 0 {
		t.Error("结果错误， 结果： ", serachRet)
	}
}

func TestSearch(t *testing.T) {
	var tree trietree.TreeManager
	ret := tree.Append("abcd")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("abc")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("bcdef")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("aaaaa")
	if ret != true {
		t.Error("加入失败 ", ret)
	}

	serachRet := tree.Search("abcdefg")
	if len(serachRet) != 3 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("abcd", serachRet) || !in_array("abc", serachRet) || !in_array("bcdef", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}

	serachRet = tree.Search("")
	if len(serachRet) != 0 {
		t.Error("结果错误， 结果： ", serachRet)
	}

	serachRet = tree.Search("bcdefgh")
	if len(serachRet) != 1 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("bcdef", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}

	serachRet = tree.Search("ef")
	if len(serachRet) != 0 {
		t.Error("结果错误， 结果： ", serachRet)
	}

	//删除一个词
	ret = tree.Delete("")
	if ret != true {
		t.Error("删除失败 ", ret)
	}

	ret = tree.Delete("abc")
	if ret != true {
		t.Error("删除失败 ", ret)
	}

	ret = tree.Delete("rrrr")
	if ret != true {
		t.Error("删除失败 ", ret)
	}

	serachRet = tree.Search("abcdefg")
	if len(serachRet) != 2 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("abcd", serachRet) || !in_array("bcdef", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}

	serachRet = tree.Search("abc")
	if len(serachRet) != 0 {
		t.Error("结果错误， 结果： ", serachRet)
	}
}

func TestChineseSearch(t *testing.T) {
	var tree trietree.TreeManager
	ret := tree.Append("吕布")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("赵云")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("典韦")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("关羽")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("孙策")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("吕布居")
	if ret != true {
		t.Error("加入失败 ", ret)
	}

	serachRet := tree.Search("赵云马超关羽刘备都是三国名人")
	if len(serachRet) != 2 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("赵云", serachRet) || !in_array("关羽", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}
	serachRet = tree.Search("")
	if len(serachRet) != 0 {
		t.Error("结果错误， 结果： ", serachRet)
	}

	serachRet = tree.Search("吕布居住在什么地方？")
	if len(serachRet) != 2 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("吕布", serachRet) || !in_array("吕布居", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}

	serachRet = tree.Search("关羽和典韦可以打几个回合？")
	if len(serachRet) != 2 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("典韦", serachRet) || !in_array("关羽", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}

	ret = tree.Delete("关羽")
	if ret != true {
		t.Error("删除失败 ", ret)
	}

	serachRet = tree.Search("关羽和典韦可以打几个回合？")
	if len(serachRet) != 1 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("典韦", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}
}

func TestChineseMixSearch(t *testing.T) {
	var tree trietree.TreeManager
	serachRet := tree.Search("adcs赵abc云degf起")
	if len(serachRet) != 0 {
		t.Error("结果错误， 结果： ", serachRet)
	}

	ret := tree.Append("abcd吕布")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("关羽abcd吕布")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("赵abc云")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("abc云动")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("关羽abcd")
	if ret != true {
		t.Error("加入失败 ", ret)
	}

	serachRet = tree.Search("adcs赵abc云degf起")
	if len(serachRet) != 1 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("赵abc云", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}
	serachRet = tree.Search("")
	if len(serachRet) != 0 {
		t.Error("结果错误， 结果： ", serachRet)
	}

	serachRet = tree.Search("adcs赵abc云动degf起？")
	if len(serachRet) != 2 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("赵abc云", serachRet) || !in_array("abc云动", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}

	serachRet = tree.Search("adujs非常时期关羽abcdefg？")
	if len(serachRet) != 1 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("关羽abcd", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}

	serachRet = tree.Search("11关羽abcd吕布22？")
	if len(serachRet) != 3 {
		t.Error("结果错误， 结果： ", serachRet)
	}
	if !in_array("abcd吕布", serachRet) || !in_array("关羽abcd吕布", serachRet) || !in_array("关羽abcd", serachRet) {
		t.Error("结果缺失， 结果： ", serachRet)
	}

	serachRet = tree.Search("关羽和吕布13123都是三国名将abchdjsk")
	if len(serachRet) != 0 {
		t.Error("结果错误， 结果： ", serachRet)
	}
}

func TestTreeNode(t *testing.T) {
	var tree trietree.TreeManager
	ret := tree.Append("吕布")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	treeNode := tree.GetTreeNode()
	if 0 != treeNode.Content {
		t.Error("节点错误, 当前节点内容 ", treeNode.Content)
	}
	wordArr := []rune("吕布")
	if _, exist := treeNode.Child[wordArr[0]]; !exist {
		t.Error("节点信息缺失, 当前节点内容 ", treeNode.Child)
	}
}

func TestTreeword(t *testing.T) {
	var tree trietree.TreeManager

	ret := tree.Delete("any code")
	if ret != true {
		t.Error("删除操作失败 ", ret)
	}

	ret = tree.Append("abcd")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("abc")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("fgr")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	searchRest := tree.GetTreeWord("ab")
	if len(searchRest) != 2 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	if !in_array("abc", searchRest) || !in_array("abcd", searchRest) {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("fg")
	if len(searchRest) != 1 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	if !in_array("fgr", searchRest) {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("")
	if len(searchRest) != 0 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("ad")
	if len(searchRest) != 0 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
}

func TestChineseTreeword(t *testing.T) {
	var tree trietree.TreeManager
	searchRest := tree.GetTreeWord("")
	if len(searchRest) != 0 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("anycode")
	if len(searchRest) != 0 {
		t.Error("查找失败 ，实际结果", searchRest)
	}


	ret := tree.Append("吕布")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("吕方")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("诸葛大力")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("诸葛亮")
	if ret != true {
		t.Error("加入失败 ", ret)
	}

	searchRest = tree.GetTreeWord("")
	if len(searchRest) != 0 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("吕")
	if len(searchRest) != 2 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	if !in_array("吕布", searchRest) || !in_array("吕方", searchRest) {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("诸葛")
	if len(searchRest) != 2 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	if !in_array("诸葛大力", searchRest) || !in_array("诸葛亮", searchRest) {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("张")
	if len(searchRest) != 0 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("葛大")
	if len(searchRest) != 0 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
}

func TestChineseMixTreeword(t *testing.T) {
	var tree trietree.TreeManager
	ret := tree.Append("吕布1")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("吕方22")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("aa诸葛大力")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("aa诸葛大亮")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("aa诸葛b'b大亮")
	if ret != true {
		t.Error("加入失败 ", ret)
	}
	ret = tree.Append("aa诸葛b'b大亮sosos")
	if ret != true {
		t.Error("加入失败 ", ret)
	}

	searchRest := tree.GetTreeWord("")
	if len(searchRest) != 0 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("吕")
	if len(searchRest) != 2 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	if !in_array("吕方22", searchRest) || !in_array("吕布1", searchRest) {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("aa")
	if len(searchRest) != 4 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	if !in_array("aa诸葛b'b大亮", searchRest) || !in_array("aa诸葛大亮", searchRest) || !in_array("aa诸葛大力", searchRest) || !in_array("aa诸葛b'b大亮sosos", searchRest) {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	searchRest = tree.GetTreeWord("aa诸葛b")
	if len(searchRest) != 2 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	if !in_array("aa诸葛b'b大亮", searchRest) || !in_array("aa诸葛b'b大亮sosos", searchRest) {
		t.Error("查找失败 ，实际结果", searchRest)
	}

	searchRest = tree.GetTreeWord("aa诸葛b'b大亮")
	if len(searchRest) != 2 {
		t.Error("查找失败 ，实际结果", searchRest)
	}
	if !in_array("aa诸葛b'b大亮", searchRest) || !in_array("aa诸葛b'b大亮sosos", searchRest) {
		t.Error("查找失败 ，实际结果", searchRest)
	}
}

/**
判断 word 是否在 words 数组中
*/
func in_array(word string, words []string) (inArray bool) {
	inArray = false
	for _, oneWord := range words {
		if oneWord == word {
			inArray = true
			return
		}
	}
	return
}
