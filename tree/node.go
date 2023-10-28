// Package tree provides an implementation of an interval tree data structure.

package tree

import (
	"time"
)

// Node represents a node in the interval tree.
type Node struct {
	interval  *IntervalPeriod
	maxEnd    time.Time
	leftNode  *Node
	rightNode *Node
}

// NewNode creates a new node with the given interval and sets its maximum end time to the end time of the interval.
func NewNode(interval IntervalPeriod) *Node {
	return &Node{&interval, interval.End, nil, nil}
}

// UpdateMaxEnd updates the maximum end time of the node based on the maximum end times of its children.
func (n *Node) UpdateMaxEnd() {
	n.maxEnd = max(n.GetInterval().End, max(getMaxEnd(n.leftNode), getMaxEnd(n.rightNode)))
}

// GetMaxEnd returns the maximum end time of the node.
func (n *Node) GetMaxEnd() time.Time {
	return n.maxEnd
}

// GetLeftNode returns the left child of the node.
func (n *Node) GetLeftNode() *Node {
	return n.leftNode
}

// SetLeftNode sets the left child of the node to the given node.
func (n *Node) SetLeftNode(node *Node) {
	n.leftNode = node
}

// GetRightNode returns the right child of the node.
func (n *Node) GetRightNode() *Node {
	return n.rightNode
}

// SetRightNode sets the right child of the node to the given node.
func (n *Node) SetRightNode(node *Node) {
	n.rightNode = node
}

// GetInterval returns the interval represented by the node.
func (n *Node) GetInterval() IntervalPeriod {
	return *n.interval
}

// SetInterval sets the interval represented by the node to the given interval.
func (n *Node) SetInterval(interval IntervalPeriod) {
	n.interval = &interval
}

// Set max end time
func (n *Node) SetMaxEnd(time time.Time) {
	n.maxEnd = time
}

// getMaxEnd returns the maximum end time of the given node, or the zero time if the node is nil.
func getMaxEnd(node *Node) time.Time {
	if node == nil {
		return time.Time{}
	}
	return node.maxEnd
}

// max returns the maximum of two given times.
func max(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}
