package sublist

func Sublist(a, b []int) string {
	// slice comparison
	 equal := func(x, y []int) bool {
 
	 for i, v := range x {
		 if v != y[i] {
			 return false
		 }
	 }
	 return true
 }
	 // sliding window
	 loop := func(small, big []int) bool {
		start := 0
		fin := len(small) 
		stop := (len(big) - len(small)) 
		ans := false
		for i:=0; i <= stop; i++ {
			if equal(small,big[start:fin]){
				ans = true
				break
			} else {
				start += 1
				fin += 1
				}
			}
			return ans
		}
	 if len(a) < len(b){
		 if loop(a,b){
			 return "sublist"
		 } else {
			 return "unequal"
		 }
	 } else if len(b) < len(a) {
		 if loop(b,a){
			 return "superlist"
		 } else {
			 return "unequal"
		 }
	 } else {
		 if equal(a,b){
			 return "equal"
		 } else {
			 return "unequal"
		 }
	 }
 }