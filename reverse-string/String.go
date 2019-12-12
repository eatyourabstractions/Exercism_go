package reverse

func String(s string)string{
    runes := []rune(s)
    res := make([]rune,len(runes))
    for idx, _ := range runes{
        res[idx] = runes[(len(runes) - 1) - idx]
    }
    
    return string(res)
}