# ThreadSafeList
Thread safe wrapper for go's List type
The purpose of this project is to hopefully save someone else the time required to implement a lock around the standard go List methods.

## Installing

Either copy the contents of the file ThreadSafeList.go into your project or use go get

```bash
go get github.com/BenKnigge/ThreadSafeList
```

## Importing into your project


```golang
import "github.com/BenKnigge/ThreadSafeList"
```

## Methods not in go's List

In addition to the standard methods  in go's list which can be found here

https://golang.org/pkg/container/list/

The methods Lock and UnLock have been added

```golang
func (l *ThreadSafeList) Lock() {
	
}

func (l *ThreadSafeList) Unlock() {
	
}
```


## Contributing

If you would like to improve this project please make feel free to make a pull request.


## Questions an comments

If you have a comment or question about this project please use the issue tracker. 


## Created by

Benjamin Knigge 

If you would like to hire me I can be contacted via the contact form on my blog at https://odinsql.com/contact/


