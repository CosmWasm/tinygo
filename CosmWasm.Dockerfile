# tinygo-base stage install Go 1.18, LLVM 13 and the TinyGo compiler itself.
## Comparing to the original Dockerfile:
##   - using prebuild LLVM version;
##   - using only WASI platform dependencies (wasi-libc, binaryen /lib submodules);
FROM golang:1.18-bullseye AS tinygo-base

## Add the LLVM 13 repo for Debian 11 Bullseye
RUN wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add - && \
    echo "deb http://apt.llvm.org/bullseye/ llvm-toolchain-bullseye-13 main" >> /etc/apt/sources.list.d/llvm.list

## Install LLVM toolchain packages
RUN apt-get update && \
    apt-get install -y clang-13 llvm-13-dev lld-13 libclang-13-dev cmake ninja-build

## Repo copy
COPY . /tinygo

## Update submodule
RUN cd /tinygo/ && \
    rm -rf ./lib/*/ && \
    git submodule sync && \
    git submodule update --init --recursive --force lib/wasi-libc && \
    git submodule update --init --recursive --force lib/binaryen

## Build WASI libs, Bynaryen (wasm-opts dependency) and the TinyGo compiler
RUN cd /tinygo/ && \
    make wasi-libc binaryen && \
    go install .

CMD ["tinygo"]
