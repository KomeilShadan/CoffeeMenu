package myPackage

import "fmt"

var PackageVar = "packageVar"

func PrintMe(text string) {
	fmt.Println(text, PackageVar)
}
