# setconfig
Easily change 'variable = value' pair in configuration files.

## Usage example

**file.conf**
```
a = 1
b = 2
c = 3
```

To change b = **2** to b = **7**
```
$ go run setconfig.go file.conf b 7
```

To add a new variable to the end of file
```
$ go run setconfig.go file.conf myvar 10
```
