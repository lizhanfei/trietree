# 字典树实现


### 增加词
```go
var tree tiretree.TreeManager
ret := tree.Append("abcd")
//结果为true
```

### 语句中词语匹配
```go
var tree tiretree.TreeManager
ret := tree.Append("abcd")
ret := tree.Append("ef")
serachRet := tree.Search("abcdefg")
//结果为： ["abcd", "ef"]
```

### 删除一个词
```go
var tree tiretree.TreeManager
ret := tree.Append("abcd")
ret = tree.Append("abc")
serachRet := tree.Search("abcdefg")
//结果为： ["abcd", "abc"]
ret = tree.Delete("abc")
serachRet = tree.Search("abcdefg")
//结果为： ["abcd"]
```

### 词语联想
```go
var tree tiretree.TreeManager
ret := tree.Append("abcd")
ret = tree.Append("abcdefg")
searchRest := tree.GetTreeWord("ab")
//结果为： ["abcd", "abcdefg"]
```

### 基准测试数据
    
    词典树使用地名和随机人名9000+。测试短语，search操作平均耗时为 0.139ms
    参考./src/tiretree/TreeManger_benchmark_test.go

更多使用方式可以参考测试代码 ./src/tiretree/TreeManager_test.go