package main

import "fmt"

type dog struct {
	name string
}

func (d *dog) prun() {
	fmt.Println("Woof Woooof, I am prunning")
}
func (d dog) run() {
	fmt.Println("Woof Woooof, I am running")
}

type younging interface {
	walk()
	run()
}

type pyounging interface {
	prun()
}

func youngRun(y younging) {
	y.run()
}

func youngPRun(y pyounging) {
	y.prun()
}

type person struct {
	first string
}

func main() {
	fmt.Println("Practicing pointers")

	v := 10
	var i *int = &v

	fmt.Printf("address of i: %p, value held by i: %v, dereferencing pointer:%v\n",
		&i, i, *i)

	var (
		a, b, c *string
		d       *int
	)

	p := "THis is a string pointer"
	q := "values is dereferenced"
	r := "I don't know what I am doing"
	s := 10
	a = &p
	b = &q
	c = &r
	d = &s

	fmt.Printf("Type of a:%T, value of a:%#v, deref: %#v\n", a, a, *a)
	fmt.Printf("Type of a:%T, value of a:%#v, deref: %#v\n", b, b, *b)
	fmt.Printf("Type of a:%T, value of a:%#v, deref: %#v\n", c, c, *c)
	fmt.Printf("Type of a:%T, value of a:%#v, deref: %#v\n", d, d, *d)

	d1 := dog{
		name: "Dolly",
	}

	d2 := &dog{
		name: "pDolly",
	}

	d1.run()
	d2.run()
	fmt.Printf("%T\n", d2)
	fmt.Println("===")
	(*d2).run()
	fmt.Println("====")

	youngRun(&d1)
	youngRun(d2)

	youngPRun(&d1)
	youngPRun(d2)

	p1 := person{
		first: "person",
	}

	fmt.Println(p1)

	fmt.Println(namer(p1))
	fmt.Println(p1)
	pnamer(&p1)
	fmt.Println(p1)

}

func (d dog) walk() {
	fmt.Println("Dog walking")
}

func pnamer(p *person) {
	p.first = "pnammer"
}

func namer(p person) person {
	p.first = "namer"
	return p
}
