package main

import "fmt"

type T int

func main() {
	fmt.Println("Understanding method sets")

	var t T = 100
	t.Pprint() // this is lega as it can be rewritten as (&t).Pprint()
	pt := &t
	pt.Pprint() // this is legal as receiver is *pt.Pprintln

	t.print()  // this is understandable as receiver is t.print
	pt.print() // this is legal as rewritten as (*pt).print()

	interfacePrint(t)  //this I can do as the receiver for print is type t.print()
	interfacePrint(pt) // this also works (*pt).print()
	interfacePrint(*pt)
	interfacePrint(&t)

	//pinterfacePrint(t) // this doesn't work as t here is not addressable in concrete
	pinterfacePrint(&t) //this works as we are passing a pointer and receiver is *type
	pinterfacePrint(pt) // this works as receiver is  *type
	//pinterfacePrint(*pt) doesn't work
}

//Method sets:
//still not clear need to check more
/*
   t T:
		print()

	t *T:
		print()
		Pprint()

	printer: print
		print()
		Pprint()

	pprinter pprint
		Pprint()

*/

func (t *T) Pprint() {
	fmt.Println(*t)
}

func (t T) print() {
	fmt.Println(t)
}

type printer interface {
	print()
}

type pPrinter interface {
	Pprint()
}

func interfacePrint(i printer) {
	i.print()
}

func pinterfacePrint(i pPrinter) {
	i.Pprint()
}
