package hamming

import "errors"

func Distance(a, b string) (int,error) {
    d := 0
    fst := []rune(a)
    snd := []rune(b)
    if len(a) == len(b) {
    for idx, _ := range a{
        if fst[idx] != snd[idx] {
            d++
            }
        }
         return d, nil
    } else {
        return 0, errors.New("strands of diff size!.")
    }
}