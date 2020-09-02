你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
  
  
