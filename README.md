# TCS
## TCP Client Shenanigans (simple tcp testing)
This repository contains a simple Go TCP testing client
Useful for testing the Morphux Package Server ([Morphux Package Server](//github.com/morphux/mps/))

## Installation
    
    go get github.com/mrgosti/tcs

## Usage

* Simple
```
    echo "test" | tcs 127.0.0.1:4242 # Send "test" to 127.0.0.1:4242 
```


* Whitespace splitted
```
    echo "test tata" | tcs 127.0.0.1:4242 # Send "testtata" to 127.0.0.1:4242 
```

* Bytes
```
    echo "42" | tcs 127.0.0.1:4242 # Send "0x42" to 127.0.0.1:4242 
```
