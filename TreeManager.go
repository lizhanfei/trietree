package trietree

type TreeManager struct {
	rootNode *TreeNode
}

func (tree *TreeManager) Create() {
	tree.rootNode = &TreeNode{Child: make(map[rune]*TreeNode)}
}

func (tree *TreeManager) Append(word string) bool {
	if 0 == len(word) {
		return false
	}
	if nil == tree.rootNode {
		tree.Create()
	}
	var nowNode *TreeNode
	nowNode = tree.rootNode
	wordArr := []rune(word)
	var run rune
	for wordStep := 0; wordStep < len(wordArr); wordStep++ {
		//当前字符 wordArr[wordStep]
		//判断当前字符  是否在nowNode 节点的子节点中存在
		run = wordArr[wordStep]
		if _, exist := (*nowNode).Child[run]; exist {
			//存在
			if wordStep == (len(wordArr) - 1) {
				//是最后一个
				(*nowNode).Child[run].IsEnd = true
			}
			nowNode = (*nowNode).Child[run]
		} else {
			//不存在
			//如果没有，则插入一个节点
			var newTreeNode = &TreeNode{Child: make(map[rune]*TreeNode)}

			if wordStep == (len(wordArr) - 1) {
				newTreeNode.IsEnd = true
			} else {
				newTreeNode.IsEnd = false
			}
			newTreeNode.Content = wordArr[wordStep]
			(*nowNode).Child[wordArr[wordStep]] = newTreeNode
			nowNode = newTreeNode
		}
	}

	return true
}

func (tree *TreeManager) GetTreeNode() TreeNode {
	return *tree.rootNode
}

/**
匹配一个句子中包含的词
*/
func (tree *TreeManager) Search(sentence string) (words []string) {
	if 0 == len(sentence) {
		return
	}
	if nil == tree.rootNode {
		tree.Create()
		return
	}
	wordArr := []rune(sentence)
	var sentenceTmp []rune
	for wordbeginStep := 0; wordbeginStep < len(wordArr); wordbeginStep++ {
		for wordEndStep := wordbeginStep + 1; wordEndStep <= len(wordArr); wordEndStep++ {
			sentenceTmp = wordArr[wordbeginStep:wordEndStep]
			//fmt.Printf("start %d : %d end  %s \n", wordbeginStep, wordEndStep, string(sentenceTmp))
			if word, exist := tree.hasWord(sentenceTmp); exist {
				words = append(words, word)
			}
		}
	}

	return
}

func (tree *TreeManager) hasWord(wordArr []rune) (word string, exit bool) {
	if 0 == len(wordArr) {
		exit = false
		return
	}
	if nil == tree.rootNode {
		tree.Create()
		exit = false
		return
	}
	var nowNode *TreeNode
	nowNode = tree.rootNode

	for wordStep := 0; wordStep < len(wordArr); wordStep++ {
		if _, exist := nowNode.Child[wordArr[wordStep]]; exist {
			word = word + string(nowNode.Child[wordArr[wordStep]].Content)
			if wordStep == len(wordArr)-1 && nowNode.Child[wordArr[wordStep]].IsEnd {
				exit = true
				return
			}
			nowNode = nowNode.Child[wordArr[wordStep]]
		} else {
			break
		}
	}
	exit = false
	return
}

/**
从树中移除一个词
将词节点的isEnd 改为false
*/
func (tree *TreeManager) Delete(word string) bool {
	if 0 == len(word) {
		return true
	}
	if nil == tree.rootNode {
		tree.Create()
		return true
	}
	var nowNode *TreeNode
	nowNode = tree.rootNode
	var run rune
	wordArr := []rune(word)
	for wordStep := 0; wordStep < len(wordArr); wordStep++ {
		//当前字符 wordArr[wordStep]
		//判断当前字符  是否在nowNode 节点的子节点中存在
		run = wordArr[wordStep]
		if _, exist := (*nowNode).Child[run]; exist {
			if wordStep == len(wordArr)-1 {
				(*nowNode).Child[run].IsEnd = false
			}
			nowNode = (*nowNode).Child[run]
		}
	}
	return true
}

/**
根据输入获取树上包含当前词语的词
*/
func (tree *TreeManager) GetTreeWord(word string) (resultWords []string) {
	if 0 == len(word) {
		return
	}
	if nil == tree.rootNode {
		tree.Create()
		return
	}
	var nowNode *TreeNode
	nowNode = tree.rootNode
	var run rune
	wordArr := []rune(word)
	var resultWordPrefix string
	for wordStep := 0; wordStep < len(wordArr); wordStep++ {
		run = wordArr[wordStep]
		if _, exist := nowNode.Child[run]; exist {
			//存在
			resultWordPrefix = resultWordPrefix + string(nowNode.Child[wordArr[wordStep]].Content)
			if wordStep == len(wordArr)-1 {
				//当前是最后一次词
				//遍历 nowNode.Child[run] 节点下的所有词
				if nowNode.Child[run].IsEnd {
					resultWords = append(resultWords, resultWordPrefix)
				}
				tree.deepGetTreeWWord(resultWordPrefix, nowNode.Child[run], &resultWords)
				return
			}
			nowNode = nowNode.Child[run]
		} else {
			return
		}
	}

	return
}

/**
以wordPrefix 为前缀
遍历 当前节点的子节点，组装成结果
*/
func (tree *TreeManager) deepGetTreeWWord(wordPrefix string, nowNode *TreeNode, resultword *[]string) {
	for content, node := range nowNode.Child {
		if node.IsEnd {
			//如果是最后一个词
			*resultword = append(*resultword, wordPrefix+string(content))
		}
		//不是最后一个
		tree.deepGetTreeWWord(wordPrefix+string(content), node, resultword)
	}
}
