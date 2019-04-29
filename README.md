# AWSet

A simple, efficient, implementation of add win sets or as it is known in Riak orsets without tombstones. 



## Usage 
```



```

  


## Why did I bother making this / Why should you use it. 

I was working on another project in go and all I needed was the one crdt, and when I was for the relevant packages all of them either:

- Used strings to represent elements and I didn't like the expense of serialization
 
- Used tomb stones meaning that it kept two sets and the delete operation put it in the remove set which means that the data structure takes up 2N space. 

- Were tied in with a large codebase, so it incurred a big footprint. 

- Used reflection under the hood which incurred extra overhead. 



So if you're looking for a stand alone implementation and the stuff I mentioned makes you uneasy use this.






## Notes

If you intend to use it, do note that it isn't thread safe, so be sure to wrap an RWMutex. 



Add win semantics have the following properties: 
```

  
  //concurrent adds (A did not observe B's add before removing value)
  
  set_a, set_b
  
  set_a.insert("x"), set_b.insert("x")
  set_a.remove("x")
  set_a.merge(set_b) 
  
  set_b.remove(set_a)
  
  
  but it will work after the fact 
  
  set_a.remove("x").merge(set_b) 
  
  //false 
  set_a.contains("x")
  
  
  //removes are only temporary 
  
  set_a, set_b = aw_set {x..z}, aw_set{x..z}
  
  set_a.merge(set_b) 
  
  set_a.remove("x") 
  
  set_b.merge(a)
  
  set_b.add("x")
  
  //true 
  set_b.contains("x")
  
  
  set_a.merge(set_b) 
  
  //true
  set_a.contains("x")
  
```
  
  
