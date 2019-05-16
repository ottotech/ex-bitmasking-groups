## ex-bitmasking-groups (prototype)

*ex-bitmasking-groups* is a small app (prototype) that uses the concept of bit masking to grant any user access to
certain groups. This prototype is a ``CRUD`` type of application that comes with a UI in order to explore how the concept
of bit masking can be used to allow users do some action based on the groups they belong to.  
Because this app is just a prototype is just storing all the data in memory, that means if you stop the 
GO built-in server you will lose all the data; however, if you decide to use the app seriously I suggest you to use a 
different storage, like a sequel database.  
The design pattern I used to create this app is DDD.

**NOTE:**
This app is just a prototype of a more robust solution; however, I thought this might be interesting for someone so 
I decided to share part of the code.   

## Install
If you have go installed in your PC simply run from the app root directory:
```
$ go run cmd/server/bistmaskinggroups.go 
```

Or you can compile the whole into a binary by simply running: 
```
go build cmd/server/bitmaskinggroups.go
```

and the executing the executable: 

```
./bitmaskinggroups 
``` 

## Built With

* go version go1.11.5 linux/amd64

## Contributing

## Authors 
* Otto Schuldt - *Initial work*

## TODO

* more tests
* use a real storage like a sequel database to store the users

## License

This project is licensed under the MIT License.
