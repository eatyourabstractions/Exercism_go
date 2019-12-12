package forth

import (
	"errors"
	"strconv"
	"strings"
)

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value int
		prev *node
	}	
)
// Create a new stack
func NewStack() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
func (this *Stack) Peek() (int, error) {
	if this.length == 0 {
        return 0, errors.New("empty stack")
	}
	return this.top.value, nil
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() (int, error) {
	if this.length == 0 {
        return 0, errors.New("empty stack")
	}
	
	n := this.top
	this.top = n.prev
	this.length--
	return n.value, nil
}
// Push a value onto the top of the stack
func (this *Stack) Push(value int) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}


func (this *Stack) PrintStack()(result []int){
    size := this.Len()
    if size > 0 {
    for i := 0; i < size; i++{
        val, _ := this.Pop()
        result = append([]int{val}, result...) 
            } 
        } else {
            result = []int{}
    }
    return
}

type Lang struct {
    stack *Stack
    ops map[string][]string
}

func NewF() *Lang{
    f := Lang{
        stack: NewStack(),
        ops: make(map[string][]string),
    }
    f.ops["+"] = []string{"+"}
    f.ops["-"] = []string{"-"}
    f.ops["*"] = []string{"*"}
    f.ops["/"] = []string{"/"}
    f.ops["dup"] = []string{"dup"}
    f.ops["drop"] = []string{"drop"}
    f.ops["swap"] = []string{"swap"}
    f.ops["over"] = []string{"over"}
    
    return &f
}

func (f *Lang) sum()(result int, err error){
    size := f.stack.Len()
    if size >= 2{
        one, _ := f.stack.Pop()
        two, _ := f.stack.Pop()
        result = one + two
        f.stack.Push(result)
    } else{
        err = errors.New("il manque des parametres")
    }
    return
    
}

func (f *Lang) sub()(result int, err error){
     size := f.stack.Len()
    if size >= 2{
        one, _ := f.stack.Pop()
        two, _ := f.stack.Pop()
        result = two - one
        f.stack.Push(result)
    } else{
        err = errors.New("il manque des parametres")
    }
    return
    
}

func (f *Lang) mult()(result int, err error){
     size := f.stack.Len()
    if size >= 2{
        one, _ := f.stack.Pop()
        two, _ := f.stack.Pop()
        result = one * two
        f.stack.Push(result)
    } else{
        err = errors.New("il manque des parametres")
    }
    return
    
}

func (f *Lang) div()(result int, err error){
	size := f.stack.Len()
   if size >= 2{
	   one, _ := f.stack.Pop()
	   two, _ := f.stack.Pop()
	   if one != 0 {
	   result = two / one
	   f.stack.Push(result)
	   } else{ err = errors.New("div by zero!!")}
   } else{
	   err = errors.New("il manque des parametres")
   }
   return
   
}

func (f *Lang) swap()(err error){
     size := f.stack.Len()
    if size >= 2{
        one, _ := f.stack.Pop()
        two, _ := f.stack.Pop()
        f.stack.Push(one)
        f.stack.Push(two)
        err = nil
    } else{
        err = errors.New("il manque des parametres")
    }
    return
    
}

func (f *Lang) over()(err error){
     size := f.stack.Len()
    if size >= 2{
        one, _ := f.stack.Pop()
        two, _ := f.stack.Peek()
        f.stack.Push(one)
        f.stack.Push(two)
        err = nil
    } else{
        err = errors.New("il manque des parametres")
    }
    return
    
}

func (f *Lang) dup()(err error){
    if val, e := f.stack.Peek(); e == nil {
        f.stack.Push(val)
        err = nil
    } else{
        err = e
    }
    return
    
}

func (f *Lang) drop()(err error){
    if _, e := f.stack.Pop(); e == nil {
        err = nil
    } else{
        err = e
    }
    return
    
}


func (f *Lang) execute(str string)(result int, err error){
     if v, e := strconv.Atoi(str); e == nil {
            f.stack.Push(v)
         return
        }
    
    switch str {
        case "+":
          result, err = f.sum()
        case "-":
          result, err = f.sub()
        case "*":
          result, err = f.mult()
        case "/":
          result, err = f.div()
        case "dup":
          err = f.dup()
          result = 0
        case "drop":
          err = f.drop()
          result = 0
        case "swap":
          err = f.swap()
          result = 0
        case "over":
          err = f.over()
          result = 0
        default:
            result, err = f.eval(str)
    }
    return
}


func (f *Lang) eval(input string)(result int, erreur error){
    in := strings.Fields(input)
    if in[0] == ":" {
        var evaluation_list []string
        for _, val := range in[2:]{
             if _, e := strconv.Atoi(val); e == nil {
                 evaluation_list = append(evaluation_list, val)
             } else {
                 evaluation_list = append(evaluation_list, f.ops[val]...)
             }
                                                              
            }
        f.ops[in[1]] = evaluation_list
        return
        }                                                   
        
    
    
    
        for _, symbol := range in{
            if v, e := strconv.Atoi(symbol); e == nil {
                   f.stack.Push(v)
            } else if _, ok := f.ops[symbol]; ok {
               for _, word := range f.ops[symbol]{
                   if _, erreur = f.execute(string(word)); erreur != nil {
                       f.stack = NewStack()
                   break } 
           }
              
         } else { erreur = errors.New("cmd not in f.ops") }
       }
    return
}

func (engine *Lang) action(input string)(erreur error){
    str := strings.ToLower(input)
    in := strings.Fields(str)
    size := len(in)
    if size > 1{
        if _, e := strconv.Atoi(in[1]); e == nil && in[0] == ":" {
            erreur = errors.New("can't redefine numbers pal!")
        } else{
            if _, err := engine.eval(str); err == nil{
                erreur = nil
            } else{ erreur = err }
        } 
    } else { 
        _, erreur = engine.eval(str)
    }
    
    return
                           
    
}

func Forth(input []string)(result []int, erreur error){
    engine := NewF()
    for _, item := range input{
        if err := engine.action(item); err != nil{
			erreur = err
            engine.stack = NewStack()
            break
        }
    }
    result = engine.stack.PrintStack()
    return
}

