
# webserve

[![Go Report Card](https://goreportcard.com/badge/github.com/Pandademic/webserve)](https://goreportcard.com/report/github.com/Pandademic/webserve)
----

my tiny little webserver

### why?
____
faster and more performant than the other webserver's I could find.At least for my tasks.
### Usage
` webserve DIR PORT`
where `DIR` is the dir containg the files and `PORT` is the port to run it on.
## Known issues
### 1.0.0
___
- shutting down a running server is a quite wonky proccess on stable , as in `ctrl-c` may or may not kill the server.The only way for certain is to close the currenly running terminal.
### 1.1
____
😄 none! 😄
### Install
___
##### Stable
download the binary for your platform and put it in a folder avilable on your path
**IF** no binary is provided , just request one via [issues](https://github.com/Pandademic/webserve/issues)
##### Dev/canary
>⚠️ **CAUTION ⚠️**

clone the repo and then run `task buildall` to obtain a dev (unreleased , possibly unstable) builds for all 
platforms to use
### Getting Help
____
File an [issue.](https://github.com/Pandademic/webserve/issues)
### Licence
The Unlicense
```
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
```
For more information, please refer to <http://unlicense.org/>
