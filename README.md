# `webserver-debug`

A debug webserver for Docker. Will return a response that is the path that is
requested + the port + an environmental variables

```shell
docker run --rm -e SERVE_PORT=8000 -e TEXT="hey there" -p 8000:8000 saulshanabrook/webserver-debug
```
