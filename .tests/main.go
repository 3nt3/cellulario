package main

import (
	"net"
	"fmt"
)

func main() {
	fmt.Println(net.LookupAddr("84.153.102.66"))
}

