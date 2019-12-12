package summultiples

func Any(vs []int, f func(int) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

func devides(n int) func(int) bool{
    
    intent := func(divisor int) bool {
		if divisor == 0 {
			return false
		}
        return (n % divisor == 0)
    }
    return intent
}




func SumMultiples(limit int, divisors... int)(factors_sum int){
    for i := 1; i < limit; i++{
        if Any(divisors, devides(i)) {
            factors_sum += i
        }
    }
    return
}
