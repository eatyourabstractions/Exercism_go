package luhn

import (
    "strconv"
	"strings" 
)

func Valid(s string) bool {
    doubleMe := func(n int) int {
        if (n * 2) > 9{
            return (n * 2) - 9
        } else { 
            return n * 2}
    }
    
    str := strings.Join(strings.Fields(s),"")
    if len(str) <= 1 {
        return false
    }
    alternate := false
    sum := 0
    for i := len(str) - 1; i >= 0; i-- {
		item, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return false
		}
        if alternate {
             sum = sum + doubleMe(item)
            alternate = !alternate
        } else {
            sum = sum + (item)
            alternate = !alternate
            }
        
        
    }
    return sum % 10 == 0
}