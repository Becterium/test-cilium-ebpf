package main

import (
//    "fmt"
    "github.com/cilium/ebpf"
)

func main(){
        spec, err := ebpf.LoadCollectionSpec("bpf_prog.o")
        if err != nil{
                panic(err)
        }

//        m := spec.Maps["my_map"]
//        p := spec.Programs["my_prog"]
//
//        fmt.Println(m.Type, p.Type)
//
//        fmt.Println(m.Key, m.Value)
//
//        fmt.Println(m)
//
//        fmt.Println(p)
//
//        fmt.Println(p.Instructions)


        coll, err := ebpf.NewCollection(spec)
        if err != nil {
             panic(err) 
        }
        defer coll.Close()

        m := coll.Maps["my_map"]

        if err := m.Put(uint32(1), uint64(2)); err != nil{
                panic(err)
        }

}
