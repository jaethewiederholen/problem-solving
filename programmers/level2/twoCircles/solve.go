package twoCircles

import (
	"math"
)

func Solution(r1 int, r2 int) int {
	var min float64 = 0.0000001
	
	r1Integers := integers(float64(r1) - min)
	r2Integers := integers(float64(r2))
							
	return r2Integers - r1Integers
}

func integers (fRadius float64) int {
	var ans float64 
	for i:=float64(0); i <= fRadius ; i++ {
		fi := float64(i)
		num := math.Floor(math.Sqrt(math.Pow(fRadius,2) - math.Pow(fi,2))) + 1
		//fmt.Println(num)
		ans += num
	}
	//fmt.Println("check", int(ans))
	return (int(ans) - (int(fRadius)*2 +1 )) * 4 + 1 + 4*int(fRadius)
}