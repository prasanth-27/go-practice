package myPackage

import "fmt"

const packageValue = 100

func PackageFunc() {
	fmt.Printf("Type and value of the value :%T, \t %d\n", packageValue, packageValue)

}
