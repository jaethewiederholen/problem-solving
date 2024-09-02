package main

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Stack struct {
    v *list.List
  }
  
  func NewStack() *Stack {
    return &Stack{list.New()}
  }
  
  func (q *Stack) Push(v interface{}) {
    q.v.PushBack(v)
  }
  
  func (q *Stack) Pop() interface{} {
    back := q.v.Back()
    if back == nil {
      return nil
    }
  
    return q.v.Remove(back)
  }

  type Queue struct {
    v *list.List
  }
  
  func NewQueue() *Queue {
    return &Queue{list.New()}
  }
  
  func (q *Queue) Push(v interface{}) {
    q.v.PushBack(v)
  }
  
  func (q *Queue) Pop() interface{} {
    front := q.v.Front()
    if front == nil {
      return nil
    }
  
    return q.v.Remove(front)
  }

func solution(plansInput [][]string) []string {
    type plan struct {
        name string
        start int
        time int
    }

    type delayed struct {
        name string
        left int
    }

    plans := make([]*plan,len(plansInput))
    for i, pi := range plansInput {
        arr := strings.Split(pi[1], ":")
        hour,_ := strconv.Atoi(arr[0]) 
        min,_  := strconv.Atoi(arr[1])
        time,_ := strconv.Atoi(pi[2])
        p := plan{
            name: pi[0],
            start: hour*60 + min,
            time: time,
        }
       plans[i] = &p
    }
    
    //sort by starttime
    sort.Slice(plans, func(i,j int )bool {
        return plans[i].start < plans[j].start
    })
    workQueue := NewQueue()
    for _, p := range plans {
        workQueue.Push(p)
    }

    delayedStack := NewStack()
    var ans []string
    currWork, _ := workQueue.Pop().(*plan)
    currTime := currWork.start

    for {
        if len(ans) == len(plans) { return ans}
        var nextStart int
        newWork, ok := workQueue.Pop().(*plan)
        if !ok {
            // 새롭게 할 work 가 없음.
            nextStart = 24*60
        } else {
            nextStart = newWork.start
        }
        HandleDelay: 
        delayedPlan, ok := delayedStack.Pop().(delayed)
        if ok {
            // delay 된게 있음
            fmt.Println("handleDelay", delayedPlan.name, "curr", currTime)
            var start int
            if currWork == nil {
                start = 24*60
            } else {
                start = currWork.start
            }
            if currTime + delayedPlan.left <= start {
                // delay 처리 가능
                currTime += delayedPlan.left
                ans = append(ans, delayedPlan.name)
                goto HandleDelay
            } else {
                // 또 딜레이
                delayedPlan.left = (currTime + delayedPlan.left) -  start
                currTime =start
                delayedStack.Push(delayedPlan)
            }
        }
        // curr Work 처리
        if currWork != nil {
            currTime = currWork.start
            fmt.Println("handleCurrWork", currWork.name, "curr", currTime)
            if currTime + currWork.time <= nextStart {
                ans = append(ans, currWork.name)
                currTime += currWork.time
            } else {
                d := delayed { name : currWork.name, left : currTime + currWork.time - nextStart}
                delayedStack.Push(d)
                currTime += currWork.time
            }
        }
        currWork = newWork
    }

}


func main() {
    //var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
    fmt.Println(solution([][]string{{"korean", "11:40", "20"}, {"english", "12:10", "30"}, {"math", "12:30", "40"}}))
//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
//     [[0, 4], [0, 1], [2, 3]] => 2
	
}
