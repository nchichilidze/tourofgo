package main

import ( 
	"golang.org/x/tour/tree"
)

func WalkWrapper(t *tree.Tree, ch chan int) { 
	Walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if (t == nil) {
		return
	}
	if (t.Left != nil) { 
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if (t.Right != nil) { 
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int) 

	go WalkWrapper(t1, ch1)
	go WalkWrapper(t2, ch2)
	
	count := 0
	
	for {
		if (count > 10) { 
			return true
		}
		count += 1
		v1, ok1 := <- ch1
		v2, ok2 := <- ch2
		if (ok1 && !ok2 || !ok1 && ok2 || ((ok1 && ok2) && (v1 !=v2)) ) { 
			return false
		} 
	}
	return true
}

func main() {
	Same(tree.New(1), tree.New(1))
	Same(tree.New(1), tree.New(2))
}
