package main

import (
    "fmt"
//    "github.com/cilium/ebpf"
)

func panicf(format string, args ...any) {
    panic(fmt.Sprintf(format, args...))
}

func main(){
    // Load the object file from disk using a bpf2go-generated scaffolding.
    spec, err := loadVariables()
    if err != nil {
        panicf("loading CollectionSpec: %s", err)
    }
    
//    // Set the 'const_u32' variable to 42 in the CollectionSpec.
//    want := uint32(42) 
//    if err := spec.Variables["const_u32"].Set(want); err != nil {
//        panicf("setting variable: %s", err)
//    }
//    
//    // Load the CollectionSpec.
//    //
//    // Note: modifying spec.Variables after this point is ineffectual!
//    // Modifying *Spec resources does not affect loaded/running BPF programs.
//    var obj variablesPrograms
//    if err := spec.LoadAndAssign(&obj, nil); err != nil {
//        panicf("loading BPF program: %s", err)
//    }
//    
//    fmt.Println("Running program with const_u32 set to", want)
//    
//    // Dry-run the BPF program with an empty context.
//    ret, _, err := obj.ConstExample.Test(make([]byte, 15)) 
//    if err != nil {
//        panicf("running BPF program: %s", err)
//    }
//    
//    if ret != want {
//        panicf("unexpected return value %d", ret)
//    }
//    
//    fmt.Println("BPF program returned", ret)
//    
//    // Output:
//    // Running program with const_u32 set to 42
//    // BPF program returned 42

    
    set := uint16(9000)
    if err := spec.Variables["global_u16"].Set(set); err != nil {
        panicf("setting variable: %s", err)
    }
    
    for range 3 {
        ret, _, err := spec.Programs["global_example"].Test(make([]byte, 15))
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

}
