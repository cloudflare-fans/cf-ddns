# cfDDNS

A Cloudflare DDNS Tool

## Build This Project

This project supports cross-platform building with container environment.
You can build your packages with any local dev or CI/CD environment.

Before you start your building, you should finish the following steps to set
up your build environment.

1. install docker
   
   [How to install docker?](https://docs.docker.com/get-started/get-docker/)

2. make cross-platform build environment with docker
   
   following cross-platform installation step is included in `make` command now.

    ```bash
    docker buildx create --use --name default-cross default
    ```
   
    or directly run:

    ```bash
   make cross-build-env
    ```
   
3. use UNIX make tool to make all packages
   
   ```bash
   make
   ```