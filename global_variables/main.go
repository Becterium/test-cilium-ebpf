package main

import (
    "fmt"
    "github.com/cilium/ebpf"
)

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func main() {
	// Load the object file from disk using a bpf2go-generated scaffolding.
	spec, err := loadVariables()
	if err != nil {
		panicf("loading CollectionSpec: %s", err)
	}
    
    set := uint16(9000)
    if err := spec.Variables["global_u16"].Set(set); err != nil {
        panicf("setting variable: %s", err)
    }

    coll, err := ebpf.NewCollection(spec)
    if err != nil {
        panic(err)
    }
    
    // Close the Collection before the enclosing function returns.
    defer coll.Close()

    for range 3 {
        ret, _, err := coll.Programs["global_example"].Test(make([]byte, 15))
        if err != nil {
            panicf("running BPF program: %s", err)
        }
        fmt.Println("BPF program returned", ret)
    }
    
    // Output:
    // Running program with global_u16 set to 9000
    // BPF program returned 9000
    // BPF program returned 9001
    // BPF program returned 9002

    var global_u16 uint16
    if err := coll.Variables["global_u16"].Get(&global_u16); err != nil {
            panicf("getting variable: %s", err)
    }
    fmt.Println("Variable global_u16 is now", global_u16)

    // Output:
    // Variable global_u16 is now 9003
}
