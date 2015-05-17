# `webserver-debug`

A debug webserver for Docker. Will return a response that is the path that is
requested + the port + an environmental variables

```shell
docker run --rm -e PORT=8000 -e TEXT="hey there" -p 8080:8080 saulshanabrook/webserver-debug
```
