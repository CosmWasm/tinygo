# CosmWasm

This is a minor patch on TinyGo to build CosmWasm-compatible contracts.
It adds a new target `cosmwasm`, which can be used in place of `wasm` in order to build WASM blob that is compatible
with CosmWasm.

In particular, it doesn't require a number of imports like `putc` and fixes false-positive float usage warnings from the
CosmWasm VM.

## Building a Docker image

We use this custom docker image in [cosmwasm-go](https://github.com/CosmWasm/cosmwasm-go) to build WASM files.
[CosmWasm.Dockerfile](./CosmWasm.Dockerfile) is a stripped down version of the original [Dockerfile](./Dockerfile)
TinyGo image (refer to file's header for more details).
When updating to a newer TinyGo, we need to update our build container:

```shell
git checkout cw-0.23.x
docker build -f CosmWasm.Dockerfile -t cosmwasm/tinygo:0.23.0 .
```

### Multi-arch Docker image build

We want to support both amd64 (most machines) and arm64 (Mac M1, Rasperry PI?).

#### Prep Work

Do this once to set up a multi-platform builder ([from](https://www.docker.com/blog/multi-arch-images/)).

```shell
docker buildx create --name mybuilder
docker buildx use mybuilder
docker buildx inspect --bootstrap
```

### Actual Build

#### Using buildx

```shell
VERSION=0.23.0

docker buildx build --platform linux/amd64,linux/arm64 -t cosmwasm/tinygo:${VERSION} -f CosmWasm.Dockerfile --push .

docker inspect cosmwasm/tinygo:${VERSION} | jq '.[] | {Arch: .Architecture, Os: .Os}'

docker buildx imagetools inspect cosmwasm/tinygo:${VERSION}
```

#### No cross-compile

We need to compile on each machine (each architecture):

```shell
VERSION_ARM=0.23.0-arm

docker build -t cosmwasm/tinygo:${VERSION_ARM} -f CosmWasm.Dockerfile . && docker push cosmwasm/tinygo:${VERSION_ARM}

docker manifest create cosmwasm/tinygo:${VERSION} --amend cosmwasm/tinygo:${VERSION_ARM}

docker manifest inspect cosmwasm/tinygo:${VERSION}

docker manifest push cosmwasm/tinygo:${VERSION}
```

## Contract compile

An example of building WASM blob using CosmWasm version of TinyGo:

```shell
tinygo build -tags "cosmwasm tinyjson_nounsafe" -no-debug -target wasi -o "${OUTPUT_WASM_FILE}" "${CONTRACT_PROJECT_DIR}/main.go"
```
