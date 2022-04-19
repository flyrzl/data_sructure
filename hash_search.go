package main

import "fmt"

// Record表示一个键值对
type Record struct {
	name    string
	country string
	next    *Record
}

// HashTableSize指哈希表底层数组的大小
const HashTableSize = 6

// 创建哈希表
// 即一个包含HashTableSize个元素的数组，每个元素存一个指向Record对象的指针
var hashTable [HashTableSize]*Record

// 哈希函数将name映射到哈希表底层数组的某个元素
func hashFunction(name string) int {
	sum := 0
	for i := 0; i < len(name); i++ {
		sum += int(name[i])
	}
	return sum % HashTableSize
}

// 查找
func lookup(name string) string {
	hashIndex := hashFunction(name)
	entry := hashTable[hashIndex]
	cur := entry
	for cur != nil {
		if cur.name == name {
			return cur.country
		}
		cur = cur.next
	}
	return "not found"
}

// 测试数据
// 为了方便，使用内置的map存储，也可以用列表或其他数据结构
var studentsMap = map[string]string{
	"Berners-Lee, Tim": "UK", "Wu, Wenjun": "China",
	"Godel, Kurt": "Austria", "Turing, Alan": "UK",
	"Knuth, Donald": "USA", "Leibniz, Gottfried": "Germany",
	"Von Neumann, John": "Hungary", "Amdahl, Gene": "USA",
	"Yao, Andrew": "China", "Moore, Gordon": "USA",
	"Yang, Xiong": "China", "Karp, Richard": "USA",
	"Boole, George": "UK", "Makimoto, Tsugio": "Japan",
	"Torvalds, Linus": "Finland",
}

func main() {
	// 输入测试数据
	for key, value := range studentsMap {
		// hashIndex是该键值对应该插入到数组的位置索引
		hashIndex := hashFunction(key)
		// 创建节点
		newRecord := &Record{
			name:    key,
			country: value,
			next:    hashTable[hashIndex], // 将该键值对插入到相应链表的头部
		}
		hashTable[hashIndex] = newRecord
	}

	// 查找s,t
	s, t := "Knuth, Donald", "Babayan, Boris"
	fmt.Println(s, "is from", lookup(s))
	fmt.Println(t, "is", lookup(t))
}
