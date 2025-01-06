package main

import "fmt"

func main() {
	fmt.Println("Practicing Structs!!")

	type person struct {
		name string
		age  int
		job  string
	}

	p1 := person{
		"John",
		32,
		"ITA",
	}

	p2 := person{
		"Mia",
		28,
		"Intern",
	}

	xp := []person{p1, p2}

	for _, p := range xp {
		fmt.Println(p)
		fmt.Printf("%#v", p)
	}

	fmt.Println(p1.name)

	fmt.Println("=======================")

	mp := make(map[string]person)
	for _, p := range xp {
		mp[p.name] = p
	}

	fmt.Println(len(mp))
	for k, v := range mp {
		fmt.Printf("Details of \"%v\" are %#v\n", k, v)
	}

	if p, ok := mp["Jane"]; ok {
		fmt.Printf("Found %v\n", p)
	}

	fmt.Println("======================")

	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine
		make  string
		model string
		doors int
		color string
	}

	nexon := vehicle{
		engine{
			false,
		},
		"SUV",
		"xms",
		4,
		"Red",
	}

	fmt.Println(nexon)
	fmt.Printf("%#v\n", nexon)

	fmt.Println(nexon.engine.electric)

	mustang := vehicle{
		engine: engine{false},
		make:   "Muscle",
		model:  "GT",
		doors:  2,
		color:  "custom",
	}
	fmt.Println(mustang)

	fmt.Println("==============")

	panon := struct {
		first   string
		friends map[string]int
		drinks  []string
	}{
		first: "james",
		friends: map[string]int{
			"moneypenny": 32,
			"q":          54,
		},
		drinks: []string{"water", "ice cream"},
	}

	fmt.Println(panon)
	fmt.Printf("%#v", panon)
}
