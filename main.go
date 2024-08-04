package main;import("fmt";i"os";h"os/exec");func main(){r,_:=h.Command(i.Args[1],i.Args[2]).Output();fmt.Printf("%s", r)}
