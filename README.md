## Zydis Bindings for Go

[Zydis](http://zydis.re) is a fast and lightweight x86/x86-64 disassembler library.  This repo provides bindings to Go via cgo and is considered a "complete" wrapper of the Zydis API, ready for production use.

It was created because the [pure-Go disassembler](https://godoc.org/golang.org/x/arch/x86/x86asm) is signficantly lacking in AMD64 support.  Decoding x86 is complex business, and it was more straightforward to make this port instead of digging deep into the pure-Go package.

**This repository uses [Git LFS](https://git-lfs.github.com/) to store a precompiled version of the Zydis library (see below), so please make sure you have it installed before getting this package.**

### Import

```
import "go.jpap.org/zydis"
```

### Sample code

See the file `cmd/demo.go`.

### Upgrading Zydis library

The Zydis library is packaged as a static syso object file so that this package is go gettable.  Precompiled macOS (amd64, arm64), Linux (amd64, arm64), and Windows (amd64, 386) binaries are provided.

Use the Makefile in the `lib/` folder to upgrade to a newer version, rebuild, or add support for another platform.  The default Makefile target clones the Zydis repo and its submodule, performs the build, and creates the syso files for Go linkage under macOS with suitable cross-compilers installed.

### License

MIT, see the `LICENSE.md` file.
