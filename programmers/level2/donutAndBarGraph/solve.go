package donutandbargraph

func Solution(edges [][]int) []int {
    //key - node num, val - [in, out]
    
    nodeDic := make(map[int][]int)
    for _, edge := range edges {
        inNode := edge[1]
        outNode := edge[0] 
        if _, ok := nodeDic[inNode]; !ok {
            nodeDic[inNode] = []int{0,0}
        } 
        nodeDic[inNode][0] += 1
        if _, ok := nodeDic[outNode]; !ok {
            nodeDic[outNode] = []int{0,0}
        }
        nodeDic[outNode][1] += 1
    }
    generateNode, donut, bar, eight := 0, 0,0, 0
    
    for node, inout := range nodeDic {
        if inout[0] == 0 && inout[1] >= 2{
            generateNode = node
        }
        if inout[0] >= 2 && inout[1] == 2 {
            eight +=1
        }
        if inout[0] >= 1 && inout[1] == 0 {
            bar +=1
        } 
    }
    donut = nodeDic[generateNode][1] - bar - eight

    return []int{generateNode, donut, bar, eight}
}