# wait-for
wait-for is designed to synchronize services like docker containers. It was inspired by [eficode/wait-for](https://github.com/eficode/wait-for), but completely rewritten in golang. It is extremely lightweight (based on the docker [scratch](https://hub.docker.com/_/scratch) image).

```sh
% ./wait-for
Usage: ./wait-for [OPTIONS] host:port
  -interval int
    	minimum interval (in second) between connection attempts (default 1)
  -timeout int
    	timeout (in second) before the program exits with an error code
```