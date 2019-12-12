package pythagorean


type Triplet [3]int

func newTriplet(a, b, c int) Triplet{
	return [3]int{a,b,c}
}

 
func isPythagorean(a, b, c int) bool {
    return (a * a) + (b * b) == c * c
}


func Range(start, end int)(triplets []Triplet) {
    for a:= start; a <= end; a++ {
        for b := a + 1; b <= end; b++ {
            for c := b + 1; c <= end; c++ {
                if isPythagorean(a, b, c){
                    triplets = append(triplets, newTriplet(a, b, c))
                }
            }
        } 
	}
    return 
}

func is_sum(triplets Triplet, num int) bool{
	sum := 0
	for _, v := range triplets {
		sum += v
	}
	return sum == num
}

func Sum(num int)(ans []Triplet) {
	for _, triplet := range Range(1, num) {
		if is_sum(triplet, num){
			ans = append(ans, triplet)
		}
	}
	return
}
