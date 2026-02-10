//go:build ignore

#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>


//volatile const __u32 const_u32;

//SEC("socket") int const_example() {
//    return const_u32;
//}

volatile __u16 global_u16;

SEC("socket") int global_example() {
    global_u16++;
    return global_u16;
}
