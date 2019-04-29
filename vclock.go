package awset

import (
	"math"
)


type vclock struct {
	dots map[string]uint64
}



func (V *vclock) contains(actor string, tick uint64) bool {

	val, e := V.dots[actor];

	if e && val >= tick {
		return true 
	}

	return false 
} 


func (V *vclock) increment(actor string) uint64 {

	
	current, e := V.dots[actor]

	if e {
		next :=  (current + 1)
		V.dots[actor] = next
		return next 
	}



	V.dots[actor] = 1
	return 1	
}



func init_vclock() vclock {
	return vclock {make(map[string]uint64)}
}




func (L *vclock) merge (R *vclock) {

	for actor, tick := range R.dots {
		t, exists := L.dots[actor] 

		if exists {
			L.dots[actor] = uint64( math.Max( float64(t), float64(tick) ) ) 
			
			continue
		}

		L.dots[actor] = tick
		
	}	


}


func (V *vclock) set(id string, value uint64) {
	V.dots[id] = value 
}




func (L *vclock) subset_of(R *vclock) bool {


	var rv bool

	rv = true 

	for k, v := range L.dots {

		if !R.contains(k, v) {
			rv = false
			break
		}		

	}


	return rv 
}
