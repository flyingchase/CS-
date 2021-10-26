package DataStructure

//
// import "fmt"
//
// const (
// 	// 顶点数量
// 	MAXVEX    int = 9
// 	MAXWEIGHT int = 100
// )
//
// // v0 to vx 最短路径
// var shortestPath = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
//
// func main() {
// 	graph := NewGraph()
// 	shortestV := make(map[int]int)
// 	// shortpath 未遍历的最小节点值
// 	var pathMin int
// 	//  shortestPath 未遍历最小节点的下标
// 	var Vx int
// 	// 记录结点是否找到 v0到 vx 的最小路径
// 	var isgetpath [MAXVEX]bool
//
// 	for v := 0; v < len(graph); v++ {
// 		shortestPath[v] = graph[0][v]
//
// 	fmt.Println("shortestTablePath is : ", shortestPath)
// }
// func NewGraph() [MAXVEX][MAXVEX]int {
// 	var graph [MAXVEX][MAXVEX]int
// 	var v0 = [MAXVEX]int{0, 1, 5, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
// 	var v1 = [MAXVEX]int{1, 0, 3, 7, 5, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
// 	var v2 = [MAXVEX]int{5, 3, 0, MAXWEIGHT, 1, 7, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
// 	var v3 = [MAXVEX]int{MAXWEIGHT, 7, MAXWEIGHT, 0, 2, MAXWEIGHT, 3, MAXWEIGHT, MAXWEIGHT}
// 	var v4 = [MAXVEX]int{MAXWEIGHT, 5, 1, 2, 0, 3, 6, 9, MAXWEIGHT}
// 	var v5 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, 7, MAXWEIGHT, 3, 0, MAXWEIGHT, 5, MAXWEIGHT}
// 	var v6 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 3, 6, MAXWEIGHT, 0, 2, 7}
// 	var v7 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 9, 5, 2, 0, 4}
// 	var v8 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 7, 4, 0}
//
// 	graph[0] = v0
// 	graph[1] = v1
// 	graph[2] = v2
// 	graph[3] = v3
// 	graph[4] = v4
// 	graph[5] = v5
// 	graph[6] = v6
// 	graph[7] = v7
// 	graph[8] = v8
// 	return graph
//
// }
