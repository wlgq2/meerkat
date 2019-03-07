package meerkat

type RadixTreeNode struct {
	Key string
	father *RadixTreeNode
	Child []*RadixTreeNode
	Value interface{}
}

type RadixTree struct {
	root *RadixTreeNode
}


func GetStringHeadCommonCnt(str1 string,str2 string) int{
	cnt :=0
	for ;cnt< len(str1) && cnt< len(str2);cnt++{
		if str1[cnt]!= str2[cnt] {
			break
		}
	}
	return cnt
}

func (tree *RadixTree) RootNode() *RadixTreeNode{
    return tree.root
}

func (tree *RadixTree) Insert(key string,value interface{}){
	//如果根节点为空，则插入根节点。
	if tree.root ==nil{
		tree.root = &(RadixTreeNode{key,nil,nil,value})
	}else{
		//如果根节点都为找到，则添加新的根节点
		if !(tree.insertInNode(tree.root,key,value)){
			root := &(RadixTreeNode{"", nil, nil, nil})
			root.Child = append(root.Child , &(RadixTreeNode{key, root, nil, value}))
			root.Child = append(root.Child , tree.root)
			tree.root = root
		}
	}

}

func (tree *RadixTree) insertInNode(node *RadixTreeNode,key string,value interface{}) bool{
	//比较与node.Key的长度
	commonHeadCnt := GetStringHeadCommonCnt(node.Key,key)

	if commonHeadCnt == len(node.Key){
        //如果key node.key相等则直接赋值
        if commonHeadCnt == len(key) {
            node.Value = value
            return true
        }
		childKey := key[commonHeadCnt:]
		//如果共同长度等于节点长度，则递归子节点。
		for i:=0;i< len(node.Child);i++{
			rst := tree.insertInNode(node.Child[i],childKey,value)
			//如果子节点找到则返回
			if rst {
				return rst
			}
		}
		//如果子节点找不到则添加节点并返回
		node.Child = append(node.Child, &(RadixTreeNode{childKey,node,nil,value}))
		return true
	}else if commonHeadCnt ==0 {
		//未找到共同字符串则返回
		return false
	}else {
		//共同长度大于0，且小于node.Key长度，则分裂该节点。
        commonKey := node.Key[:commonHeadCnt]
        newKey := node.Key[commonHeadCnt:]
        node.Key = commonKey
        newNode :=  &(RadixTreeNode{newKey,node,node.Child,node.Value})
        node.Child = nil
        node.Child = append(node.Child, newNode)
        node.Value = nil
        
        //如果key大于共同长度，则插入新节点
        if commonHeadCnt < len(key) {
            childKey := key[commonHeadCnt:]
            childNode :=  &(RadixTreeNode{childKey,node,nil,value})
            node.Child = append(node.Child, childNode)
        } else {
            //key等于共同长度则直接赋值
            node.Value = value
        }
        return true
	}
}

func (tree *RadixTree) GetNode(key string) *RadixTreeNode{
    if nil == tree.root{
        return nil
    }
    return tree.getNode(key,tree.root)
}

func (tree *RadixTree) getNode(key string, node *RadixTreeNode) *RadixTreeNode{
    
    commonHeadCnt := GetStringHeadCommonCnt(node.Key,key)
    //该key节点不存在
    if commonHeadCnt < len(node.Key){
        return nil
    }
    //找到该节点
    if commonHeadCnt == len(key){
        return node
    }
    //遍历子节点
    newKey := key[commonHeadCnt:]
    for i:=0;i<len(node.Child);i++ {
        rst := tree.getNode(newKey,node.Child[i])
        if nil !=  rst{
            return rst
        }
    }
    
    //找不到该节点
    return nil
}


func (tree *RadixTree) GetValue(key string) interface{}{
	node := tree.GetNode(key)
	if nil == node{
		return nil
	}
	return node.Value
}