package palindrome

import (
	"strconv"
	"errors"
)

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}


func reverse(s string)string{
    runes := []rune(s)
    res := make([]rune,len(runes))
    for idx, _ := range runes{
        res[idx] = runes[(len(runes) - 1) - idx]
    }
    
    return string(res)
}

func comb_helper(p int, n []int, c []int, cc [][]int) [][]int {
    if len(n) == 0 || p <= 0 {
        return cc
    }
    p--
    for i := range n {
        r := make([]int, len(c)+1)
        copy(r, c)
        r[len(r)-1] = n[i]
        if p == 0 {
            cc = append(cc, r)
        }
        cc = comb_helper(p, n[i+1:], r, cc)
    }
    return cc
}

func comb(p int, n []int) [][]int {
    return comb_helper(p, n, nil, nil)
}

func moiMeme(n int) []int{
    return []int{n,n}
}

func Map(vs []int, f func(int) []int) [][]int {
    vsm := make([][]int, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func is_palindrome(num int)bool{
    return strconv.Itoa(num) == reverse(strconv.Itoa(num))
}

type Product struct {
	Product int 
	Factorizations [][2]int
 }
func Products(start, end int)(pmin, pmax Product, err error){
	if start > end {
	  err = errors.New("fmin > fmax")
	  return
	}
	
    myFactors := make(map[int][][2]int)
    myRange := makeRange(start, end)
    myCombinations := append(comb(2, myRange), Map(myRange, moiMeme)... )
    min := myCombinations[len(myCombinations) - 1][0] * myCombinations[len(myCombinations) - 1][1]
    var max int
    var arr [2]int
    
    for i := 0; i < len(myCombinations); i++{
        aPair := myCombinations[i]
        aProduct := aPair[0] * aPair[1]
        if is_palindrome(aProduct){
    	   copy(arr[:], aPair)
           myFactors[aProduct] = append(myFactors[aProduct], arr)
           if aProduct > max {
			  max = aProduct
		}
	   	    if aProduct < min {
			  min = aProduct
		}

        }
    }
  	if (len(myFactors[min]) + len(myFactors[max])) > 0 {
		pmin = Product{min, myFactors[min]}
		pmax = Product{max, myFactors[max]}
		err = nil
	} else {
		err = errors.New("no palindromes")
	}
    return 
    
}