# cfDDNS

A Cloudflare DDNS Tool

## Build Environment

1. install docker

2. make cross-platform build environment

    ```bash
    docker buildx create --use --name default-cross default
    ```
   
    or directly run:

    ```bash
   make cross-build-env
    ```