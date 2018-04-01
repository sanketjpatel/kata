package main

import "fmt"

// Point represents a point on cartesian plane
type Point struct {
	x, y int
}

func (point Point) String() string {
	return fmt.Sprintf("(%v, %v)", point.x, point.y)
}

// PointWithDistance is a node that has two properties:
//   distance - square of euclidean distance from some reference point
//   point - the point which is at the given distance from the reference point
type PointWithDistance struct {
	distance int
	point    *Point
}

// MaxHeap represents a max heap of points.
//   Each node in the heap is a point with distance from some reference point
type MaxHeap struct {
	capacity int
	size     int
	nodes    []*PointWithDistance
}

// Create a MaxHeap of capacity k points closest to the target from an array of Points
func createMaxHeap(points []*Point, target *Point, k int) *MaxHeap {
	newHeap := MaxHeap{k, 0, make([]*PointWithDistance, k)}

	for _, point := range points {
		pointWithDistance := createNode(point, target)
		newHeap.addNodeToHeap(pointWithDistance)
	}

	return &newHeap
}

func createNode(point, target *Point) *PointWithDistance {
	delX := point.x - target.x
	delY := point.y - target.y
	distance := (delX * delX) + (delY * delY)

	return &PointWithDistance{distance, point}
}

// Add node to heap
func (heap *MaxHeap) addNodeToHeap(newNode *PointWithDistance) {
	if heap.size != heap.capacity {
		heap.nodes[heap.size] = newNode
		heap.size++
		heap.heapifyFromBottom()
	} else if newNode.distance < heap.nodes[0].distance {
		heap.nodes[0] = newNode
		heap.heapifyFromTop()
	}
}

func getLeftChild(i int) int {
	return (2 * i) + 1
}

func getRightChild(i int) int {
	return (2 * i) + 2
}

func getParent(i int) int {
	return (i - 1) / 2
}

func (heap *MaxHeap) swap(i, j int) {
	temp := heap.nodes[i]
	heap.nodes[i] = heap.nodes[j]
	heap.nodes[j] = temp
}

func (heap *MaxHeap) isLeaf(nodeIndex int) bool {
	return getLeftChild(nodeIndex) > heap.size-1
}

// gets index of the child node with max distance
func (heap *MaxHeap) getMaxChildIndex(nodeIndex int) int {
	if heap.isLeaf(nodeIndex) {
		panic("Cannot get max child for a leaf node")
	}

	if getRightChild(nodeIndex) >= heap.size || heap.nodes[getLeftChild(nodeIndex)].distance > heap.nodes[getRightChild(nodeIndex)].distance {
		return getLeftChild(nodeIndex)
	}

	return getRightChild(nodeIndex)
}

// Last node may not be at the correct place in the heap. Trickle it up
func (heap *MaxHeap) heapifyFromBottom() {
	trickleNodeIndex := heap.size - 1
	for (trickleNodeIndex != 0) && (heap.nodes[trickleNodeIndex].distance > heap.nodes[getParent(trickleNodeIndex)].distance) {
		heap.swap(trickleNodeIndex, getParent(trickleNodeIndex))
		trickleNodeIndex = getParent(trickleNodeIndex)
	}
}

// Heap root may not be at the correct plact. Trickle it down
func (heap *MaxHeap) heapifyFromTop() {
	trickleNodeIndex := 0

	for trickleNodeIndex < heap.size {
		if heap.isLeaf(trickleNodeIndex) {
			return
		}

		maxChildIndex := heap.getMaxChildIndex(trickleNodeIndex)
		if heap.nodes[trickleNodeIndex].distance > heap.nodes[maxChildIndex].distance {
			return
		}
		heap.swap(trickleNodeIndex, maxChildIndex)
		trickleNodeIndex = maxChildIndex
	}
}

// Finds k closest points from target point
func findClosest(k int, points []*Point, target *Point) []*Point {
	if points == nil || k >= len(points) {
		return points
	}

	maxHeap := createMaxHeap(points, target, k)

	closestPoints := make([]*Point, k)
	for i, node := range maxHeap.nodes {
		closestPoints[i] = node.point
	}

	return closestPoints
}

// A utility function that returns all the points with distance from target point
func getPointsWithDistance(points []*Point, target *Point) []*PointWithDistance {
	pwds := make([]*PointWithDistance, len(points))
	for i, point := range points {
		pwds[i] = createNode(point, target)
	}

	return pwds
}

func main() {
	points := []*Point{
		{0, -2},
		{-2, 4},
		{1, 3},
		{-1, -3},
		{5, 2},
	}

	origin := &Point{1, 2}
	// fmt.Println(getPointsWithDistance(points, origin))
	fmt.Println(findClosest(1, points, origin))
}
