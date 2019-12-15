package say


import (
	"strconv"
)

func take(arr []int, num int)(h []int, rest []int){
    if len(arr) > num {
    h = arr[:num]
    rest = arr[num:]
        return
    } else {
        h = []int{}
        rest = arr
        return
    }
}

func grouped(arr []int, num int, f func([]int, int)([]int,[]int))(result [][]int){
    if len(arr) <= num{
        result = append(result, arr)
        return
    } else {
        h, t := f(arr, num)
        result = append(result, grouped(h, num, f)...)
        result = append(result, grouped(t, num, f)...)
    }
    return
}

func takeRight(arr []int, num int)(h []int, taken []int){
    size := len(arr)
    if size > num{
        taken = arr[size - num : size]
        h = arr[:size - num]
        return
    } else{
        h = []int{}
        taken = arr
        return
    }
}

func preprocessing(num int, f func([]int, int)([]int,[]int) )(resultat [][]int){
    tmp := []int{}
    for _, val := range strconv.Itoa(num){
        n, _ := strconv.Atoi(string(val))
        tmp = append(tmp, n)
    }
    resultat = grouped(tmp, 3, f)
    return
}

 func InBetween(i, min, max int) bool {
         if (i >= min) && (i <= max) {
                 return true
         } else {
                 return false
         }
 }

var myDict = map[int]string{
	0 : "zero",
	1 : "one",
	2 : "two",
	3 : "three",
	4 : "four",
	5 : "five",
	6 : "six",
	7 : "seven",
	8 : "eight",
	9 : "nine",
	10 : "ten",
	11 : "eleven",
	12 : "twelve",
	13 : "thirteen",
	14 : "fourteen",
	15 : "fifteen",
	16 : "sixteen",
	17 : "seventeen",
	18 : " eighteen",
	19 : "nineteen",
	20 : "twenty",
	30 : "thirty",
	40 : "forty",
	50 : "fifty",
	60 : "sixty",
	70 : "seventy",
	80 : "eighty",
	90 : "ninety",
    100: "one hundred",
    1000 : "one thousand",
    10000: "ten thousand",
    100000: "one hundred thousand",
    1000000 : "one million",
    10000000 : "ten million",
    100000000 : "one hundred million",
    1000000000 : "one billion",
}


func name_num(myArr []int, myNums map[int]string)(name string){
    
   
    defaults := func(key int)(res string, isHere bool) {
            if value, ok := myNums[key]; ok{
                res = value
                isHere = true
                return
            } else{
                res = "nothing to look here pal!!"
                isHere = false
                return
            }
        return
    }
    arr2int := func(ar []int)(num int) {
        str := ""
        for i := 0; i < len(ar); i++{
            str += strconv.Itoa(ar[i])
        }
        key, _ := strconv.Atoi(str)
        num = key
        return
    }
    
    separate := func(num int)(m, n int){
        n = num % 10
        m = num - n
        return
    }
    op := func(arr []int){
        num := arr2int(arr)
        if ans, yesno := defaults(num); yesno{
            name = ans
        } else {
            big, small := separate(num)
            name = myNums[big] + "-" + myNums[small]
        }
        return
    }
    
    handle_trailing_zeros := func(arr []int)(newar []int){
        
         if arr[0] == 0 {
            newar = arr[1:]
            return
         } else if (len(arr) > 1) && (arr[0] + arr[1]) == 0 {
            newar = []int{arr[2]}
            return
        } else { 
            newar = arr
            return
        }
    }
    
    
    myArr = handle_trailing_zeros(myArr)

    if len(myArr) == 3{
        op(myArr[1:])
        name = myNums[myArr[0]] + " hundred " + name
    } else if len(myArr) == 1 {
       name = myNums[myArr[0]] 
    } else {
        op(myArr) }
  

    return
}

var echelle = []string{" billion ", " million ", " thousand ","" }


func Say(num int)(str string, err bool){
    if val, ok := myDict[num]; ok{
        str = val
        err = true
        return
    } else if !InBetween(num,0,999999999999) {
        err = false
        return
    }
    
    in := preprocessing(num, takeRight)
    size := len(in)
    _, scale_indices := takeRight([]int{0, 1, 2, 3}, size)

    for i := size - 1; i >= 0; i--{
        str = name_num(in[i], myDict) + echelle[scale_indices[i]] + str
		}
		err = true
    return
}