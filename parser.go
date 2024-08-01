// take the html document
// Parse it into struct of strings
package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// Parser take a reader and converts it into simplified version of link
func Parser(r io.Reader) ([]Link, error) {
	n, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(n)
	for _, node := range nodes {
		fmt.Println(node)
	}
	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
