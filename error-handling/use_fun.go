package erratum


func Use(o ResourceOpener, input string) error {
	ro_err_handler := func(ro ResourceOpener)(Resource, error) {
		for {
			resource, err := ro()
			if err != nil {
				if _, ok := err.(TransientError); ok {
					continue
				} else {
					return nil, err
				}
			} else {
				return resource, nil
			} 
		}
	}
	res, err := ro_err_handler(o)
	if err != nil {
		return err
	} 

	var finalRes error

	defer res.Close()

	defer func(){
		if r := recover(); r != nil{
			if e, ok := r.(FrobError); ok{
				res.Defrob(e.defrobTag)
				finalRes = e
				return
			} else {
				finalRes = r.(error)
				return
			}
		}
	}()
	
	res.Frob(input)
	return finalRes
}