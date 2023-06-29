# local-docker-lambda

The template for emulating AWS Lambda functions locally using Docker.

## useage

Build image.

```sh
make build
```

Run Container.

```sh
make start
```

Invoke function.

```sh
curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{ "name": "RIE" }'
```
