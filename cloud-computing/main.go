package main

import "fmt"

type virtualizedResources struct{}
type physicalResources struct{}

type virtualization map[virtualizedResources]physicalResources

func newVirtualization(processsType string) virtualization {
	switch processsType {
	case "using hardware functionality":
		// partitioning, partition controller
		var virtualized virtualization
		return virtualized

	case "software functionality":
		// hypervisor
		var virtualized virtualization
		return virtualized
	}

	return nil

}

type Person struct {
	Name string
}

func main() {
	someparam := "a string"
	fmt.Println("hello motion world")
	fmt.Println(someparam)
}
