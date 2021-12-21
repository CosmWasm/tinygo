# Building the Docker files

We want to support both amd64 (most machines) and arm64 (mac m1, rasperry pi?)

## Prep Work

From https://www.docker.com/blog/multi-arch-images/

```shell
docker buildx create --name mybuilder
docker buildx use mybuilder
docker buildx inspect --bootstrap
```

## Actual Build

```shell
VERSION=0.19.2
docker buildx build --platform linux/amd64,linux/arm64 -t cosmwasm/tinygo:${VERSION} -f Dockerfile.wasm --push .

docker inspect cosmwasm/tinygo:${VERSION} | jq '.[] | {Arch: .Architecture, Os: .Os}'

docker buildx imagetools inspect cosmwasm/tinygo:${VERSION}
```
