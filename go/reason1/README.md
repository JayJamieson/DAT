# Go Docker

Example here shows building a minimal Go docker image.

## Build

- `docker build --build-arg ENV=production -t main .` - builds docker image with argument for environment
- `docker build --build-arg ENV=scratch -t main_scratch -f Dockerfile.scratch .` even smaller container

## Inpspect the docker image

- `docker image ls | grep main` shows size of docker image
- `docker image ls | grep main_scratch` shows size of docker image
- `docker image history main` shows build history and size of each layer
- `docker image history main_scratch` shows build history and size of each layer
