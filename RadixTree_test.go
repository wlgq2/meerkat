package meerkat

import  "testing"
	
func TestRedisTree(t *testing.T) {
	tree := RadixTree{}
	strs :=[]string{
		"abcd",
		"1234",
		"abcdef",
		"abcde",
		"abctt",
		"1r45334",
		"123456",
		"12",
		"test123",
		"testabc",
		"test",
	}
	for i:=0;i<len(strs);i++ {
		tree.Insert(strs[i],strs[i])
	}
	for i:=0;i<len(strs);i++ {
		if tree.GetValue(strs[i]).(string) != strs[i]{
			t.Error("Redix tree error :",strs[i]) 
		}
	}
	t.Log("success~")
}
