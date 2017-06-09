package main

import "flag"

type hosts []string

var host hosts

func (i *hosts) String() string {
	return "" //fmt.Sprintf("%s", *i)
}

func (h *hosts) Set(value string) error {
	*h = append(*h, value)
	return nil
}
func main() {
	flag.Var(&host, "hosts", "List of hosts")
	flag.Parse()

}
