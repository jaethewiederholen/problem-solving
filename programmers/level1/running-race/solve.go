package runningRace

func Solution(players []string, callings []string) []string {
    order := make(map[string]int)
    for i, p := range players {
        order[p] = i
    }
    for _, call := range callings {
        idx := order[call]
        var loser string
        
        players, loser = swap(players, idx)
        order[call] -=1
        order[loser] +=1
    } 
    return players
}

func getIndex(arr []string, name string) int {
    for i, _ := range arr {
        if arr[i] == name {return i}
    }
    return 0
}

func swap(arr []string, index int) ([]string, string) {
    loser := arr[index-1]
    if index > 0 && index < len(arr) {
        // Perform the swap
        arr[index], arr[index-1] = arr[index-1], arr[index]
    }
    return arr, loser
}