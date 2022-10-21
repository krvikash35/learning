## VM vs Container

![vm host arch](./vm-container-host-arch.png)

## Concepts
* Container runtime/engine: Docker, CRI-O, containerd, podman, lima etc. High level runtime that manage container lifycycle using low level OCI comlaint runtime.
* OCI(open container initiative): linux foundatin work to establish standard for container. OCI complaint runtime runc, crun, runv, etc
* runc: is a CLI tool for running containers on Linux according to the OCI specification
* containerd: 

## docker vs podman
* docker has daemon that take command from docker cli through socket, while podman is daemonless.
* docker container has been mostly run as root while same can run rootless under podman.
* many dev are used to docker-compose which are not well supported by podman. we can use colima for the same if we can't use docker due to license or other reason.

![docker layer arch](./docker-layer-arch.png)

## Podman

```
brew install podman
podman machine init
podman machine start
podman info
```


