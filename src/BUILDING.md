# Building the Docker files

We want to support both amd64 (most machines) and arm64 (mac m1, rasperry pi?)

## Prep Work

Do this once to set up a multi-platform builder. From https://www.docker.com/blog/multi-arch-images/

```shell
docker buildx create --name mybuilder
docker buildx use mybuilder
docker buildx inspect --bootstrap
```

## Actual Build

This seems to hit issues... 

```shell
VERSION=0.19.2

docker buildx build --platform linux/amd64,linux/arm64 -t cosmwasm/tinygo:${VERSION} -f Dockerfile.wasm --push .

docker inspect cosmwasm/tinygo:${VERSION} | jq '.[] | {Arch: .Architecture, Os: .Os}'

docker buildx imagetools inspect cosmwasm/tinygo:${VERSION}
```

No cross-compile, we need to compile on each machine:

```
VERSION=0.19.2-arm
docker build -t cosmwasm/tinygo:${VERSION} -f Dockerfile.wasm . && docker push cosmwasm/tinygo:${VERSION}
```

```
docker manifest create cosmwasm/tinygo:0.19.3 --amend cosmwasm/tinygo:0.19.3-amd64 --amend cosmwasm/tinygo:0.19.3-arm64 

docker manifest inspect cosmwasm/tinygo:0.19.3

docker manifest push cosmwasm/tinygo:0.19.3

```