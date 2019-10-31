# How to use this guid generator
1. Import the package in go script
    ```
package main
import (
	"github.com/alessandromr/GoGuid/guid"
)
    ```
2. Simple call the package in your function
```
//...Function Code
   token := guidutil.CreateToken()
//...Function Code
```
Or use custom token format `(length int, timeFormat string)`
```
//...Function Code
   token := guidutil.CreateCustomToken(256, "20060102150405.999999")
//...Function Code
```


## How Token is formed:
Tokens are formed with a URL safe encoding, they can be used in cookie or in URL without any problems.
The token are encoded with base64 (URL Friendly).


###### Example Token
```
MjAxODEwMjUxMDM3NTkuNjM0__lRuOZxsIgkVHXC5aoebCrtYyYj-XRXiIWqjhJhZ6rMgQcin-3KDbqe8IoRnx0CVlfwrQG7TbGOPPlVHUu77s1Q==
```


#### Timestamp
Timestamp make the token sequential, before encoding timestamp is a simple datetime string:
```
20060102150405.999
```
This datetime is formed in this way: **YearMonthDayHourMinuteSecond.Milliseconds**
Timestamp is the part of the token before the separator
In example token the timestamp is:
```
MjAxODEwMjUxMDM3NTkuNjM0
```
#### Separator
Timestamp and random string are separeted by two underscore:
```
__
```

#### Random string
Random string is the part of the token after the separator
In example token the random string is:
```
lRuOZxsIgkVHXC5aoebCrtYyYj-XRXiIWqjhJhZ6rMgQcin-3KDbqe8IoRnx0CVlfwrQG7TbGOPPlVHUu77s1Q==
```




