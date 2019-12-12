package series

import "strings"

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func scalar_sum(vector []int, scalar int)[]int{
    new_vector := make([]int, len(vector))
	for i, _ := range vector {
		new_vector[i] = vector[i] + scalar
	}
	return new_vector
}


func windowing(in string, idx_list []int) string{

    window := make([]string, len(idx_list))    
    for i, v := range idx_list {
        window[i] = string(in[v])
    }
    return strings.Join(window,"")
    
}

func All(n int, s string)(nlist []string){
    str_len := len(s) 
    idxs := makeRange(0, n - 1)
    
    for i := 0; i <= (str_len - n); i++{
        indices := scalar_sum(idxs, i)
        nlist = append(nlist, windowing(s, indices))

    }
    return

}

func UnsafeFirst(n int, s string) string{
    return windowing(s, makeRange(0, n - 1))
}