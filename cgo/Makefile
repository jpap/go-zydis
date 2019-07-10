# Copyright 2019 John Papandriopoulos.  All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE.md file.

# This Makefile was designed to run on macOS, using a cross-compiler to
# build the Linux and Windows versions of the zydis library.
#   1. Xcode
#   2. brew install mingw-w64
#   3. brew install musl-cross

CC_DARWIN_AMD64=clang
CC_LINUX_AMD64=x86_64-linux-musl-gcc
CC_WINDOWS_386=i686-w64-mingw32-gcc
CC_WINDOWS_AMD64=x86_64-w64-mingw32-gcc

CXX_DARWIN_AMD64=clang++
CXX_LINUX_AMD64=x86_64-linux-musl-g++
CXX_WINDOWS_386=i686-w64-mingw32-g++
CXX_WINDOWS_AMD64=x86_64-w64-mingw32-g++

LD_DARWIN_AMD64=ld
LD_LINUX_AMD64=x86_64-linux-musl-ld
LD_WINDOWS_386=i686-w64-mingw32-ld
LD_WINDOWS_AMD64=x86_64-w64-mingw32-ld

LDFLAGS_GCC=--whole-archive 
LDFLAGS_CLANG=-all_load -arch x86_64

.PHONY: all prepare clean install_headers darwin_amd64 linux_amd64

all: install

# Prepare the build
prepare: clean
	git clone https://github.com/zyantific/zydis.git .build
	cd .build && \
	git submodule update --init

darwin_amd64: prepare
	cd .build && \
	env \
		AR=$(AR_DARWIN_AMD64) \
		CC=$(CC_DARWIN_AMD64) \
		CXX=$(CXX_DARWIN_AMD64) \
	cmake . && \
	make && \
	$(LD_DARWIN_AMD64) -r -o ../../libzydis_darwin_amd64.syso $(LDFLAGS_CLANG) libZydis.a

linux_amd64: prepare
	cd .build && \
	env \
		CC=$(CC_LINUX_AMD64) \
		CXX=$(CXX_LINUX_AMD64) \
	cmake \
		-DCMAKE_SYSTEM_NAME=linux \
		-DZYDIS_BUILD_EXAMPLES=OFF \
		-DCMAKE_POSITION_INDEPENDENT_CODE=ON \
		. && \
	make && \
	$(LD_LINUX_AMD64) -r -o ../../libzydis_linux_amd64.syso $(LDFLAGS_GCC) libZydis.a

windows_amd64: prepare
	cd .build && \
	env \
		CC=$(CC_WINDOWS_AMD64) \
		CXX=$(CXX_WINDOWS_AMD64) \
	cmake \
		-DCMAKE_SYSTEM_NAME=windows \
		-DZYDIS_BUILD_EXAMPLES=OFF \
		-DCMAKE_POSITION_INDEPENDENT_CODE=ON \
		. && \
	make && \
	$(LD_WINDOWS_AMD64) -r -o ../../libzydis_windows_amd64.syso $(LDFLAGS_GCC) libZydis.a

windows_386: prepare
	cd .build && \
	env \
		CC=$(CC_WINDOWS_386) \
		CXX=$(CXX_WINDOWS_386) \
	cmake \
		-DCMAKE_SYSTEM_NAME=windows \
		-DZYDIS_BUILD_EXAMPLES=OFF \
		-DCMAKE_POSITION_INDEPENDENT_CODE=ON \
		. && \
	make && \
	$(LD_WINDOWS_386) -r -o ../../libzydis_windows_386.syso $(LDFLAGS_GCC) libZydis.a

# Clean the build
clean:
	rm -rf .build

install_headers: build_darwin_amd64
	# Copy include folder
	cp -Rvp .build/include/* include/
