## build server image
```bash
docker build -f example.Dockerfile -t guyarb/example .
```

## build client image
```bash
docker build -f exampleclient.Dockerfile -t guyarb/example .
```

## run server locally
```bash
cd example
go run ./cmd/example/service.go
```

## run client locally
```bash
cd example
go run ./cmd/exampleclient/service.go localhost:8080
```

## run using docker
```bash
docker run -it -p 8080:8080 guyarb/example
```

## deploy
```bash
kubectl apply -f ./k8s/example.yaml -n <ns>
kubectl apply -f ./k8s/exampleclient.yaml -n <ns>
```
