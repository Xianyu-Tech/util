package sensitive

import (
	"strings"
	"sync"
	"unicode"
)

type Node struct {
	_node map[rune]*Node

	_end bool
}

func newNode() *Node {
	node := &Node{
		_node: make(map[rune]*Node),
		_end:  false,
	}

	return node
}

type Trie struct {
	_root  *Node
	_mutex sync.RWMutex
}

// 通过敏感词表初始化Trie
func NewTrieByWords(sensitivewords []string) *Trie {
	sensTree := newTrie()

	for _, sensitiveWord := range sensitivewords {
		sensTree.Add(sensitiveWord)
	}
	return sensTree
}

func newTrie() *Trie {
	trie := &Trie{
		_root: newNode(),
	}

	return trie
}

func (trie *Trie) Add(data string) {
	keys := []rune(strings.ToLower(strings.TrimSpace(data)))

	if len(keys) == 0 {
		return
	}

	trie._mutex.Lock()
	defer trie._mutex.Unlock()

	node := trie._root

	for _, key := range keys {
		if _, ret := node._node[key]; ret == false {
			node._node[key] = newNode()
		}

		node = node._node[key]
	}

	node._end = true
}

func (trie *Trie) cycleDel(node *Node, keys []rune, index int) bool {
	key := keys[index]
	lens := len(keys)

	if tmpNode, ret := node._node[key]; ret == true {
		if index+1 < lens {
			// 未到达末节点, 递归到下一个节点
			flag := trie.cycleDel(tmpNode, keys, index+1)

			// 下一个节点可删除,且当前节点无更长敏感词
			if flag == true && len(tmpNode._node) == 0 {
				// 当前节点无更短敏感词, 删除节点
				if tmpNode._end == true {
					return false
				} else {
					delete(node._node, key)
				}
			}
		} else {
			// 敏感词末节点
			// 无更长的敏感词, 删除节点
			// 有更长的敏感词, 重置_end值
			if tmpNode._end == true {
				if len(tmpNode._node) == 0 {
					delete(node._node, key)

					return true
				} else {
					tmpNode._end = false
				}
			}
		}
	}

	return false
}

func (trie *Trie) Del(data string) {
	keys := []rune(strings.ToLower(strings.TrimSpace(data)))

	if len(keys) == 0 {
		return
	}

	trie._mutex.Lock()
	defer trie._mutex.Unlock()

	node := trie._root
	trie.cycleDel(node, keys, 0)
}

func (trie *Trie) replace(chars []rune, start, end int) string {
	word := chars[start : end+1]

	for index := start; index <= end; index++ {
		chars[index] = 42 // *的rune为42
	}

	return string(word)
}

func (trie *Trie) HasSensitive(data string) bool {
	chars := []rune(data)
	lens := len(chars)

	if lens == 0 {
		return false
	}

	var ret bool

	for i := 0; i < lens; i++ {
		node := trie._root

		node, ret = node._node[unicode.ToLower(chars[i])]

		if ret == false {
			continue
		}

		end := 0

		for j := i + 1; j < lens; j++ {
			node, ret = node._node[unicode.ToLower(chars[j])]

			if ret == false {
				if end > 0 {
					// 匹配结束, 判断是否有屏蔽字
					return true

					//i = end
				}

				break
			}

			// 记住上次匹配位置，寻找最大匹配串
			if node._end == true {
				end = j

				// 匹配串最后节点或者最后字符
				if len(node._node) == 0 || j+1 == lens {
					return true
				}
			}
		}
	}

	return false
}

func (trie *Trie) Replace(data string) string {
	chars := []rune(data)
	lens := len(chars)

	if lens == 0 {
		return string(chars)
	}

	var ret bool

	for i := 0; i < lens; i++ {
		node := trie._root

		node, ret = node._node[unicode.ToLower(chars[i])]

		if ret == false {
			continue
		}

		end := 0

		for j := i + 1; j < lens; j++ {
			node, ret = node._node[unicode.ToLower(chars[j])]

			if ret == false {
				if end > 0 {
					// 匹配结束, 判断是否有屏蔽字
					trie.replace(chars, i, end)

					i = end
				}

				break
			}

			// 记住上次匹配位置，寻找最大匹配串
			if node._end == true {
				end = j

				// 匹配串最后节点或者最后字符
				if len(node._node) == 0 || j+1 == lens {
					trie.replace(chars, i, j)

					i = j
					break
				}
			}
		}
	}

	return string(chars)
}
