package goweb

type Node struct {
	// path不为空表示存在路径
	pattern string
	// 当前路径段
	part string
	// 树的子节点
	children map[string]*Node
	// 是否模糊匹配，比如*.html匹配2.html\3.html
	isWild bool
}

func newNode() *Node {
	return &Node{
		children: make(map[string]*Node),
	}
}

//
func (n *Node) search(parts []string, depth int) *Node {
	// root不存part
	if len(parts) == depth {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[depth]
	if _, ok := n.children[part]; !ok {
		return nil
	}

	// todo 正则的时候怎么处理

	return n.children[part].search(parts, depth+1)
}

func (n *Node) insert(parts []string, pattern string, depth int) {
	// todo 正则的情况
	if len(parts) == depth {
		n.pattern = pattern
		return
	}

	part := parts[depth]
	if _, ok := n.children[part]; !ok {
		n.children[part] = newNode()
		n.children[part].part = part
	}
	child := n.children[part]
	child.insert(parts, pattern, depth+1)
}
