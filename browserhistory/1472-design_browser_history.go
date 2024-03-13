package browserhistory

type SiteNode struct {
	Prev    *SiteNode
	Next    *SiteNode
	Address string
}

type BrowserHistory struct {
	current *SiteNode
}

func Constructor(homepage string) BrowserHistory {
	node := &SiteNode{Address: homepage}
	return BrowserHistory{current: node}
}

func (b *BrowserHistory) Visit(url string) {
	node := &SiteNode{
		Address: url,
		Prev:    b.current,
	}
	b.current.Next = node
	b.current = node
}

func (b *BrowserHistory) Back(steps int) string {
	node := b.current
	for i := 0; i < steps && node.Prev != nil; i++ {
		node = node.Prev
	}
	b.current = node
	return node.Address
}

func (b *BrowserHistory) Forward(steps int) string {
	node := b.current
	for i := 0; i < steps && node.Next != nil; i++ {
		node = node.Next
	}
	b.current = node
	return node.Address
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
