## build
```bash
docker build -t guyarb/example .
```

## run locally
```bash
cd example
go run ./cmd/example/service.go
```

## run using docker
```bash
docker run -it -p 8080:8080 guyarb/example
```

## deploy
```bash
kubectl apply -f ./k8s/example.yaml -n <ns>
```
