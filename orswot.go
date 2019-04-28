package orswot
//import "sync"




type Orswot struct {
	actor string
	elements map[interface{}]vclock
	dots vclock
}





func (OS *Orswot) Insert(e interface{}) {
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




func (OS *Orswot) Remove(e interface{}) {
	
	delete(OS.elements, e)
	OS.dots.increment(OS.actor)
}





func (OS *Orswot) View() []interface{} {


	var ret []interface{}

	
	for key, _ := range OS.elements {
		ret = append(ret, key) 
	}

	return ret
	
}




func (L *Orswot) Merge(R *Orswot) {


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

	
	


	

