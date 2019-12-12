package react


import (
    "crypto/rand"
    "encoding/base64"
)

func GenerateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }

    return b, nil
}

func GenerateRandomString(s int) (string, error) {
    b, err := GenerateRandomBytes(s)
    return base64.URLEncoding.EncodeToString(b), err
}

type myReactor struct{
    engine *myReactor 
    observers map[string][]*myComputeCell
    observables map[string]int
    ancient_values map[string]int
    callback_list map[string]map[string]func(int)
    callbacks_to_execute *[]string
    
}
func New() Reactor{
    Re := myReactor{
              observers: make(map[string][]*myComputeCell),
              observables: make(map[string]int),
              ancient_values: make(map[string]int),
            callback_list: make(map[string]map[string]func(int)),
            callbacks_to_execute: &[]string{},
        }
    Re.engine = &Re
    return Re
}

func (r myReactor) CreateInput(val int) InputCell{
    name, _ := GenerateRandomString(10)
    newic := myInputCell{
        id: name,
		engine: r.engine,
    }
     r.engine.observables[name] = val
     r.engine.ancient_values[name] = val
    return newic

}

// ----- inputCell ----
type myInputCell struct {
    id string
	engine *myReactor
}

func (ic myInputCell) SetValue(val int){
    ic.engine.observables[ic.id] = val
    for _, cc := range ic.engine.observers[ic.id]{
        cc.update()
    }

    ancients := ic.engine.ancient_values
    listeners := ic.engine.observables

    for _, cell := range *ic.engine.callbacks_to_execute{
        if ancients[cell] != listeners[cell]{
            for _, f := range ic.engine.callback_list[cell]{
                f(ic.engine.observables[cell])
                ic.engine.ancient_values[cell] = listeners[cell]
            }
        }
    }
    *ic.engine.callbacks_to_execute = nil
}

func (ic myInputCell) Value() int{
	 return ic.engine.observables[ic.id]
}

// ----- myComputeCell -----
type myComputeCell struct{
    engine *myReactor
    id string
    update func() 

}

func (cc myComputeCell) Value() int{
    return cc.engine.observables[cc.id]
}


// ===== myCanceler impl =====
type myCanceler struct {
    computeCell_id string
    engine *myReactor
    func_to_be_erased string
}

func (canceler myCanceler) Cancel(){
    delete(canceler.engine.callback_list[canceler.computeCell_id], canceler.func_to_be_erased)
    }


func (c myComputeCell) AddCallback(f func(int)) Canceler {

    fun_id, _ := GenerateRandomString(7)
    if _, ok := c.engine.callback_list[c.id]; ok{
        c.engine.callback_list[c.id][fun_id] = f 
    } else {
        c.engine.callback_list[c.id] = make(map[string]func(int))
        c.engine.callback_list[c.id][fun_id] = f 
    }
  

    
    return myCanceler{
        computeCell_id: c.id,
        engine: c.engine,
        func_to_be_erased: fun_id,
    }
}


// ===== myReactor impl =====


func whoami(c Cell)(name string) {
	inputC, fst := c.(myInputCell)
    computeC, _ := c.(myComputeCell)
	if fst {
	    name = inputC.id
	} else { 
        name = computeC.id
    }
    return
	
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}


func (reactor myReactor) CreateCompute1(input Cell, fn func(int) int) ComputeCell{
    
    
    watchMe := whoami(input)
    compute_cell_id, _ := GenerateRandomString(10)
    reactor.observables[compute_cell_id] = fn(input.Value())
    reactor.ancient_values[compute_cell_id] = fn(input.Value())
    
    update := func(){
        reactor.observables[compute_cell_id] = fn(input.Value())
        for _, subsequent_observers := range reactor.observers[compute_cell_id]{
            subsequent_observers.update()
        }
        if !contains(*reactor.callbacks_to_execute, compute_cell_id){
            *reactor.callbacks_to_execute = append(*reactor.callbacks_to_execute, compute_cell_id)

        }

    }
        
    
    newCCell := myComputeCell{
        id: compute_cell_id,
        update: update,
        engine: reactor.engine,
	}

    reactor.observers[watchMe] = append(reactor.observers[watchMe], &newCCell)
    
    return newCCell
   
}

func (reactor myReactor) CreateCompute2(fst_in Cell, snd_in Cell, fn func(int, int) int) ComputeCell{
    
    
    watch_fst := whoami(fst_in)
    watch_snd := whoami(snd_in)
    compute_cell_id, _ := GenerateRandomString(10)
    first_val := fn(fst_in.Value(), snd_in.Value())
    reactor.observables[compute_cell_id] = first_val
    reactor.ancient_values[compute_cell_id] = first_val

    update := func(){
        reactor.observables[compute_cell_id] = fn(fst_in.Value(), snd_in.Value())
        for _, subsequent_observers := range reactor.observers[compute_cell_id]{
            subsequent_observers.update()
        }
        if !contains(*reactor.callbacks_to_execute, compute_cell_id){
            *reactor.callbacks_to_execute = append(*reactor.callbacks_to_execute, compute_cell_id)
   

        }

    } 
    
    newCCell := myComputeCell{
        id: compute_cell_id,
        update: update,
        engine: reactor.engine,
	}

    reactor.observers[watch_fst] = append(reactor.observers[watch_fst], &newCCell)
    reactor.observers[watch_snd] = append(reactor.observers[watch_snd], &newCCell)
    
    return newCCell

}
