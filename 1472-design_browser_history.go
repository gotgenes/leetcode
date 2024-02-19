type SiteNode struct {
    Address string
    Prev *SiteNode
    Next *SiteNode
}

type BrowserHistory struct {
    current *SiteNode
}


func Constructor(homepage string) BrowserHistory {
    node := &SiteNode{Address: homepage}
    return BrowserHistory{ current: node }
}


func (this *BrowserHistory) Visit(url string)  {
    node := &SiteNode{
        Address: url,
        Prev: this.current,
    }
    this.current.Next = node
    this.current = node
}


func (this *BrowserHistory) Back(steps int) string {
    node := this.current
    for i := 0; i < steps && node.Prev != nil; i++ {
        node = node.Prev
    }
    this.current = node
    return node.Address
}


func (this *BrowserHistory) Forward(steps int) string {
    node := this.current
    for i := 0; i < steps && node.Next != nil; i++ {
        node = node.Next
    }
    this.current = node
    return node.Address
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */