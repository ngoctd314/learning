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

func main() {
	add23(1, 2)
}
func add23(x, y int) {
	fmt.Println(x, y)
}
