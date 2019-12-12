package circular

import "errors"


// These are bunch of helper functions:
func All(vs []byte, f func(byte) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

func isIt(n byte) func(byte) bool{
    
    intent := func(other byte) bool {
        return (n == other)
    }
    return intent
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func rearrange(arr []int, n int)(rearranged []int){
    rearranged = append(rearranged, append(arr[n:], arr[:n]...)...)
    return
}

func (b *Buffer) findMeASlot()(int, error){
    for _, v := range rearrange(b.indices, b.oldest) {
        if b.arr[v] == 0 {
            return v, nil
        } 
    }
    return 0, errors.New("everything is full, sorry!!")
    
}
// end of list of helper functions

type Buffer struct {
    length int
    indices []int
    next int
    oldest int
    arr []byte
}
func NewBuffer(size int) *Buffer{
    ls := make([]byte, size)
    B := Buffer{
        length: size,
        indices: makeRange(0, size - 1),
        next: 0,
        oldest: 0,
        arr: ls,
    }
    return &B
}

func (buf *Buffer) ReadByte()(byte, error){
    var b byte
    if All(buf.arr, isIt(0)){
        return b, errors.New("empty array, nothing to look here!!")
    } else{
            res := buf.arr[buf.next]
            buf.arr[buf.next] = b
            if (buf.next + 1) == buf.length{
                buf.next = 0
            } else{
                buf.next += 1
            }
            return res, nil
    }
}


func (buf *Buffer) WriteByte(c byte) error{
    slot, err := buf.findMeASlot()
    if err != nil{
        return err
    }
    
    if(slot == buf.length - 1){
        buf.arr[slot] = c
        buf.oldest = 0
        return nil
    } else {
        buf.arr[slot] = c
        buf.oldest = slot + 1
        return nil
    }
}


func (buf *Buffer) Overwrite(c byte){
    if (buf.oldest == buf.length - 1){
        buf.arr[buf.oldest] = c
        buf.oldest = 0
        buf.next = 0
    } else {
        buf.arr[buf.oldest] = c
        buf.oldest += 1
        buf.next += 1
    }
    
}

func (buf *Buffer) Reset(){
    buf.next = 0
    buf.oldest = 0
    buf.arr =  make([]byte, buf.length)
}


