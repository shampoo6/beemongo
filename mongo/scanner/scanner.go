package scanner

import (
	"beemongo/mongo/connection/pool"
	"bufio"
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var modelsDir string

// 扫描文件夹，通常是 domains 文件夹，自动创建文件夹下对应文件的表
func ScanDir(dir string) {
	logs.Info("开始扫描 domains")
	logs.Info(`获取 domains 路径：%s`, dir)
	modelsDir = dir

	dirObj, err := ioutil.ReadDir(modelsDir)
	if err != nil {
		panic(err)
	}

	// PthSep := string(os.PathSeparator) // 系统目录分割符

	for _, fi := range dirObj {
		if !fi.IsDir() { // 处理文件
			readFile(fi)
		}
	}
}

// 一次性读完文件
//bytes, err := ioutil.ReadFile(filePath)
//if err != nil {
//	logs.Error(err)
//	return
//}
//logs.Debug(string(bytes))

func readFile(fileInfo os.FileInfo) {
	filePath := modelsDir + string(os.PathSeparator) + fileInfo.Name()
	logs.Debug(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	startRead := false
	over := false
	br := bufio.NewReader(file)
	documentReader := readDocument()
	for {
		bytes, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		var line string
		line = string(bytes)
		// logs.Debug(line)
		if !startRead && !over {
			startRead = strings.Contains(line, "@Document")
		} else if !over {
			over = documentReader(line)
		}
	}
	_ = file.Close()
}

func readDocument() func(line string) bool {
	lineNum := 1
	readIndexOver := false
	indexReader := readIndex()
	var collectionName string
	readLogic := func(line string) bool {
		logs.Debug("[%d] %s", lineNum, line)
		// 第一行 找到并创建对应表
		if lineNum == 1 {
			collectionName = createDocument(line)
		} else {
			readIndexOver = indexReader(line, collectionName)
			if readIndexOver {
				indexReader = readIndex()
				readIndexOver = false
			}
		}
		lineNum++
		isLastLine := strings.Contains(line, "}")
		//if isLastLine {
		//	// 读取完成后 给表加入 createTime 和 updateTime 的索引
		//	createIndex(collectionName, "CreateTime", false)
		//	createIndex(collectionName, "UpdateTime", false)
		//}
		return isLastLine
	}
	return readLogic
}

// 返回的方法： 读取文档中的索引，返回是否读取完一个索引
// 判断是否是 @Index 开头
// 不是就跳过
// 是就读取接下来的字段
func readIndex() func(line string, collectionName string) bool {
	// 记录读取的行数
	count := 0
	unique := false
	fn := func(line string, collectionName string) bool {
		if count == 0 {
			// 读取第一行 判断是否有 @Index
			containIndex := strings.Contains(line, "@Index")
			if containIndex {
				logs.Debug("找到 @Index ：%s", line)
				unique = strings.Contains(line, "unique")
				if unique {
					logs.Debug("找到 @Index unique")
				}
				count++
			}
			return !containIndex
		} else if count == 1 {
			// 读取第二行 获取字段
			logs.Debug("读取字段：%s", line)
			lineArr := strings.Fields(strings.TrimSpace(line))
			if len(lineArr) < 1 {
				panic("Document 格式错误，字段名")
			}
			// 创建索引
			createIndex(collectionName, lineArr[0], unique)
			return true
		} else {
			return true
		}
	}
	return fn
}

// 建表
func createDocument(line string) string {
	logs.Debug("第一行：%s", line)
	lineArr := strings.Fields(strings.TrimSpace(line))
	if len(lineArr) < 2 {
		panic("Document 格式错误，无法获取表名")
	}
	collectionName := lineArr[1]
	fn := func(db *mgo.Database) interface{} {
		db.C(collectionName)
		return nil
	}
	pool.GetConnectionPool().ExecDbFn(fn)
	return collectionName
}

func createIndex(collectionName string, fieldName string, unique bool) {
	//// 字段首字母小写
	//fieldStr := ""
	//for i, c := range fieldName {
	//	//fieldStr += i == 0? unicode.ToLower(c):c
	//	if i == 0 {
	//		fieldStr += string(unicode.ToLower(c))
	//	} else {
	//		fieldStr += string(c)
	//	}
	//}
	fn := func(db *mgo.Database) interface{} {
		collection := db.C(collectionName)
		index := mgo.Index{
			Key:        []string{fieldName}, // 索引字段， 默认升序,若需降序在字段前加-
			Unique:     unique,              // 唯一索引 同mysql唯一索引
			DropDups:   true,                // 索引重复替换旧文档,Unique为true时失效
			Background: true,                // 后台创建索引
		}
		if err := collection.EnsureIndex(index); err != nil {
			panic(err)
		}
		return nil
	}
	pool.GetConnectionPool().ExecDbFn(fn)
}
