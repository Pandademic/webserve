# CONTRIBUTING TO WEBSERVE
------
First up , thank you for the sentiment to help.
Here's how you could help:
- file an issue , of a bug you found. If possible open a PR to fix it.
- open a PR to fix an existing issue.
if you do decide to help .... 
the coding style guide is below.
Other than that....
please make sure you are following are following the standards for working with us:
## Our Standards
- Use welcoming and inclusive language
- Respect each other
- Provide and gracefully accept constructive criticism
- Show empathy towards other community members

Examples of unacceptable behavior by participants include:

- Trolling, insulting/derogatory comments, and personal or political attacks
- Public or private harassment
- Publishing others' private information, such as a physical or electronic address, without explicit permission
- The use of sexualized language or imagery
- Unwelcome sexual attention or advances
- Other conduct which could reasonably be considered inappropriate in a professional setting

## Enviorment setup

Requirments:

- [go 1.18](https://go.dev) 

- [task](taskfile.dev)

- a editor you can work with

- [git , of course](https://git-scm.com/)

Enjoy!
## Style Guide

### Code

- use `:=` whenever possible , unless you are declaring a value.
EG:
```go

func main(){
    // CORRECT:
    x:= 9.5
    // INCORRECT:
    var x float 
}
// BUT 
var(
  y int
)
y = 7
// is okay
```
- don't use fmt.... use log

EX:
``` GO
fmt.Println("this is bad")
log.Println("good")
```
- don't use `os.Exit()` instead use `log.Fatal`
- don't use `panic` either.... It look's disgusting , and can be confusing
- when prefixing a message(ex: in a log) use no caps , eg: `info:` **NOT** `INFO:` , it looks better
### Git 

use this commit msg format:
`(type) add/remove/improve filename/function`
types:
- docs
- refactor (edit the code , but no  new features)
- feat (new feature)
- ci (build)
- cleanup
- rm (removing a feature)
- bugfix
remeber you **CAN** use multiple , just seperate them with comma's
then , in the extended description describe all the great things you did!
use `*` as bullet points

## Best practice's

- run `task cleanup` before commiting , to make sure everything looks good
- check your grammar if it's a docs change
- makes sure `task buildall` works
