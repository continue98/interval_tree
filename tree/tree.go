// Package tree provides an implementation of an interval tree data structure.

package tree

// IntervalTree represents an interval tree data structure.
type IntervalTree struct {
	Root *Node // Root node of the tree.
}

// NewIntervalTree creates a new empty interval tree.
func NewIntervalTree() *IntervalTree {
	return &IntervalTree{nil}
}

// Insert inserts the given interval into the tree.
func (it *IntervalTree) Insert(interval IntervalPeriod) {
	it.Root = it.insertNode(it.Root, interval)
}

// insertNode inserts the given interval into the subtree rooted at the given node and returns the root of the updated subtree.
func (it *IntervalTree) insertNode(node *Node, interval IntervalPeriod) *Node {
	if node == nil {
		return NewNode(interval)
	}

	insIntervalStart := interval.GetStart()

	if insIntervalStart.Before(node.interval.Start) {
		node.SetLeftNode(it.insertNode(node.GetLeftNode(), interval))
	} else {
		node.SetRightNode(it.insertNode(node.GetRightNode(), interval))
	}

	node.UpdateMaxEnd()
	return node
}

// Delete deletes the given interval from the tree.
func (it *IntervalTree) Delete(interval IntervalPeriod) {
	it.Root = it.deleteNode(it.Root, interval)
}

// deleteNode deletes the given interval from the subtree rooted at the given node and returns the root of the updated subtree.
func (it *IntervalTree) deleteNode(node *Node, interval IntervalPeriod) *Node {
	if node == nil {
		return nil
	}

	deleteIntervalStart := interval.GetStart()
	deleteIntervalEnd := interval.GetEnd()

	leftNode := node.GetLeftNode()
	rightNode := node.GetRightNode()

	intervalNodeStart := node.GetInterval().Start
	intervalNodeEnd := node.GetInterval().End

	if deleteIntervalStart.Before(intervalNodeStart) {
		node.SetLeftNode(it.deleteNode(leftNode, interval))
	} else if deleteIntervalStart.After(intervalNodeEnd) {
		node.SetRightNode(it.deleteNode(rightNode, interval))
	} else if deleteIntervalEnd.Before(intervalNodeEnd) {
		node.SetLeftNode(it.deleteNode(leftNode, interval))
	} else if deleteIntervalEnd.After(intervalNodeEnd) {
		node.SetRightNode(it.deleteNode(rightNode, interval))
	} else {
		if leftNode == nil {
			return rightNode
		} else if rightNode == nil {
			return leftNode
		}

		minRight := it.findMin(rightNode)
		node.SetInterval(*minRight.interval)
		node.SetRightNode(it.deleteNode(rightNode, *minRight.interval))
	}

	node.UpdateMaxEnd()
	return node
}

// Search returns all intervals in the tree that intersect with the given interval.
func (it *IntervalTree) Search(interval IntervalPeriod) []IntervalPeriod {
	var result []IntervalPeriod
	it.searchNodes(it.Root, interval, &result)
	return result
}

// searchNodes recursively searches for intervals in the subtree rooted at the given node that intersect with the given interval and appends them to the result slice.
func (it *IntervalTree) searchNodes(node *Node, interval IntervalPeriod, result *[]IntervalPeriod) {
	if node == nil {
		return
	}

	intervalNode := node.GetInterval()

	rightNode := node.GetRightNode()
	leftNode := node.GetLeftNode()

	intervalStart := intervalNode.Start
	intervalEnd := intervalNode.End

	searchIntervalStart := interval.GetStart()
	searchIntervalEnd := interval.GetEnd()

	if searchIntervalStart.Before(intervalEnd) && searchIntervalEnd.After(intervalStart) {
		*result = append(*result, intervalNode)
	}

	if leftNode != nil && searchIntervalStart.Before(leftNode.maxEnd) {
		it.searchNodes(leftNode, interval, result)
	}

	if rightNode != nil && searchIntervalEnd.After(intervalStart) {
		it.searchNodes(rightNode, interval, result)
	}
}

// findMin returns the node with the minimum interval in the subtree rooted at the given node.
func (it *IntervalTree) findMin(node *Node) *Node {
	for node.GetLeftNode() != nil {
		node = node.GetLeftNode()
	}
	return node
}
