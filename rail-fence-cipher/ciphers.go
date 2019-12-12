package railfence


func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func reverse(s []int)[]int{
    res := make([]int,len(s))
    for idx, _ := range s{
        res[idx] = s[(len(s) - 1) - idx]
    }
    
    return res
}

func sine_traverse(slc []int, n int)(updown []int){
	
	
	slc_reversed := reverse(slc)
	for i := 1; i <= n ; i++{
		if i % 2 == 0 {
		updown = append(updown,slc_reversed[:len(slc) - 1]...)
		} else {
		updown = append(updown,slc[:len(slc) - 1]...)
			}
	}
	updown = updown[:n]
	return  
}

func build_grid(r, c int)[][]string{
	grid := make([][]string, r)

	for i := 0; i < r; i++ {
	  grid[i] = make([]string, c)
	}
	
      return grid
}


func Encode(msg string, r int)(encoded string) {
	
	c := len(msg)
	
	m := build_grid(r, c)
	
	sine := sine_traverse(makeRange(0,r - 1),c)
	
	for i, num := range sine {
	  
	   m[num][i] = string(msg[i])
	}
	
	for _, r := range m{
	  for _, item := range r {
	    encoded += item
	  }
	}
	return	

}

func Decode(msg string, r int)(original string){
	
	c := len(msg)
	
	m := build_grid(r, c)
	g := build_grid(r, c)
	
	
	sine := sine_traverse(makeRange(0,r - 1),c)
	
	for i, num := range sine {
	  
	   m[num][i] = "X"
	}
	
	n := 0
	for ridx, r := range m {
	  for cidx, _ := range r {
	  	
	    if m[ridx][cidx] == "X" {
	        g[ridx][cidx] += string(msg[n])
		n++
	    }
	  }
	
	}
	
	for i, num := range sine {
	  
	   original += g[num][i]
	}
	return
}
