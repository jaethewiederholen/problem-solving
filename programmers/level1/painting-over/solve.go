package PaintingOver


func Solution(n int, m int, section []int) int {
    isPainted := make(map[int]bool)
    for i:=0; i<n; i++ {
        isPainted[i+1] = true
    }
    for _, s := range section {
        isPainted[s] = false
    }
    fromFront := true
    cnt := 0
    for {
        if len(section) == 0 { break}
        if fromFront {
            for i:= 0; i < len(section); i++ {
                if !isPainted[section[i]] {
                    //paint
                    from := section[i]
                    to := from + m -1
                    paintedIdx := 0
                    for i:=0; i < len(section); i++ {
                        if to >= section[i] {
							isPainted[section[i]]= true
                            paintedIdx = i
						
                        }
                    }
                    section = section[paintedIdx+1:]
                    cnt  += 1
                    fromFront = false
                    break
                }
            }
        } else {
            for i := len(section)-1; i>=0; i-- {
                if !isPainted[section[i]] {
                    // paint
                    to := section[i]
                    from := to -m +1
                    paintedIdx := len(section)-1
                    for i:= len(section)-1; i >= 0; i-- {
                        if from <= section[i] {
							isPainted[section[i]]= true
                            paintedIdx = i
				
                        }
                    }
                    section = section[:paintedIdx]
                    cnt+=1
                    fromFront = true
                    break
                }
            }
        }
    }
    return cnt
}