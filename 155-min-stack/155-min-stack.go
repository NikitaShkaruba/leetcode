type MinStack struct {
    Stack []int
    Min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{Stack: []int{}, Min: []int{}}
}


func (this *MinStack) Push(x int)  {
    this.Stack = append(this.Stack, x)
    if len(this.Min) == 0 {
        this.Min = append(this.Min, 0)       
    } else {
        if this.Stack[this.Min[len(this.Min)-1]] < x {
            this.Min = append(this.Min, this.Min[len(this.Min)-1])
        } else {
            this.Min = append(this.Min, len(this.Stack)-1)
        }
    }
}


func (this *MinStack) Pop()  {
    this.Stack = this.Stack[:len(this.Stack)-1]
    this.Min = this.Min[:len(this.Min)-1]
}


func (this *MinStack) Top() int {    
    return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.Stack[this.Min[len(this.Min)-1]]
}