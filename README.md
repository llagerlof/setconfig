# setconfig
Easily change **'variable = value'** pair in configuration files.

## Usage example

**> example.conf**
```
a = 1
b = 2
c = 3
```

Change **b = 2** to **b = 7**
```
$ go run setconfig.go example.conf b 7
```

To add a new variable to the end of file
```
$ go run setconfig.go example.conf myvar 10
```

## Magic

### It can identify if config file has spaces before and after the = sign when editing existing variable:

**> file1.conf**
```
a = 1
b=2
```

```
$ go run setconfig.go file1.conf b 7
```

***Result:***
```
a = 1
b=7
```

```
$ go run setconfig.go file1.conf a 3
```

***Result:***
```
a = 3
b=7
```

### It can identify if config file has spaces before and after the = sign when adding a new variable:

**> file2.conf**
```
port = 1234
host = my.localhost
```

```
$ go run setconfig.go file2.conf firewall true
```

***Result:***
```
port = 1234
host = my.localhost
firewall = true
```

**> file3.conf**
```
port=1234
host=my.localhost
```

```
$ go run setconfig.go file3.conf firewall true
```

***Result:***
```
port=1234
host=my.localhost
firewall=true
```

### It can identify CR or CRLF line endings:

**> file4.conf**
```
variable1 = value1<LF>
variable2 = value2<LF>
```

```
$ go run setconfig.go file4.conf variable3 value3
```

***Result:***
```
variable1 = value1<LF>
variable2 = value2<LF>
variable3 = value3<LF>
```

**> file5.conf**
```
variable1 = value1<CRLF>
variable2 = value2<CRLF>
```

```
$ go run setconfig.go file5.conf variable3 value3
```

***Result:***
```
variable1 = value1<CRLF>
variable2 = value2<CRLF>
variable3 = value3<CRLF>
```

*Plus, if the file hasn't line endings, identify by OS.*

### It can keep, if exist, blank spaces before the variable:

**> file6.conf**
```
····home = /home/user
→   sync = yes
```

```
$ go run setconfig.go file6.conf home /home/marcvin
```

***Result:***
```
····home = /home/marcvin
→   sync = yes
```

```
$ go run setconfig.go file6.conf sync no
```

***Result:***
```
····home = /home/marcvin
→   sync = no
```
