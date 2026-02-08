ln -sf /usr/include/asm-generic/ /usr/include/asm

//go:build ignore

go get -tool github.com/cilium/ebpf/cmd/bpf2go

:p
