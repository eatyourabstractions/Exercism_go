package perfect

import "errors"

func Filter(vs []int64, f func(int64) bool) []int64 {
    vsf := make([]int64, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func makeRange(min, max int64) []int64 {
    a := make([]int64, max-min+1)
    for i := range a {
        a[i] = min + int64(i)
    }
    return a
}


func intention(n int64) func(int64) bool{
    
    intent := func(divisor int64) bool {
        return (n % divisor == 0)
    }
    return intent
}

var ErrOnlyPositive = errors.New("Can't classify zero or negative.")


type Classification string
const ClassificationDeficient Classification = "deficient"
const ClassificationPerfect Classification = "perfect"
const ClassificationAbundant Classification = "abundant"

func Classify(num int64)(Classification, error){
    
    if num < 1{
        return "", ErrOnlyPositive
    }

	pfs := Filter(makeRange(1, int64(num/2)), intention(num))
	var sum int64 = 0
	for _, factor := range pfs {
		sum = sum + int64(factor)
	}

	if sum < num {
		return ClassificationDeficient, nil
	} else if sum == num {
	    return ClassificationPerfect ,nil
	} else {
		return ClassificationAbundant, nil
	}

}