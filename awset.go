package awset
//import "sync"




type AWSet struct {
	actor string
	elements map[interface{}]vclock
	dots vclock
}



func New(actor string) AWSet {
	elements := make(map[interface{}]vclock)
	return AWSet {actor, elements, init_vclock() }
}



func FromSlice(actor string, slice []interface{}) AWSet {

	os := New(actor)

	for x, _ := range slice {
		os.Insert(x)
	}

	return os
} 



func (OS *AWSet) Insert(e interface{}) {
	ctr := OS.dots.increment(OS.actor) 

	s, exists := OS.elements[e]


	if exists == true {
		s.set(OS.actor, ctr)
		return 
	}


	dots := init_vclock()

	dots.set(OS.actor, ctr) 
	OS.elements[e] = dots
}




func (OS *AWSet) Remove(e interface{}) {
	
	delete(OS.elements, e)
	OS.dots.increment(OS.actor)
}





func (OS *AWSet) View() []interface{} {


	var ret []interface{}

	
	for key, _ := range OS.elements {
		ret = append(ret, key) 
	}

	return ret
	
}




func (L *AWSet) Merge(R *AWSet) {


	for e, dots := range R.elements {
		ldots, exists := L.elements[e]

		has_dots := dots.subset_of(&L.dots)
		
		if !exists && !has_dots {
			L.elements[e] = dots 
			continue 
		}


		if exists {
			ldots.merge(&dots)
			L.elements[e] = ldots
		}

		
	}

	
	for e, vclock := range L.elements {
		_, exists := R.elements[e]
		has_dots := vclock.subset_of(&R.dots)
		
		if !exists && has_dots {
			delete(L.elements, e)
		}
	}

	

	L.dots.merge(&R.dots) 

	
}

	
	


	

