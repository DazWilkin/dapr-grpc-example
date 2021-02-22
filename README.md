# Dapr: gRPC Golang example

A simplified dupe of the [Dapr SDK for Go](https://github.com/dapr/go-sdk) example

## Notes

+ Need to create ImagePullSecret if using private repo


```bash
NAMESPACE="example"

kubectl create namespace ${NAMESPACE}

kubectl create secret generic ghcr \
--from-file=.dockerconfigjson=${HOME}/snap/docker/471/.docker/config.json \
--type=kubernetes.io/dockerconfigjson \
--namespace=${NAMESPACE}

kubectl apply \
--filename=./kubernetes/server.yaml \
--namespace=${NAMESPACE}

kubectl get all --namespace=${NAMESPACE}
NAME                          READY   STATUS
pod/server-7bfdb46b95-f6b48   2/2     Running

NAME                  TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)
service/server-dapr   ClusterIP   None         <none>        80/TCP,50001/TCP,50002/TCP,9090/TCP

NAME                     READY   UP-TO-DATE   AVAILABLE
deployment.apps/server   0/1     1            0

NAME                                DESIRED   CURRENT   READY
replicaset.apps/server-7bfdb46b95   1         1         0
```

From e.g. `ashapp`, you can invoke a REST (!) endpoint against the gRPC server:

```bash
curl \
--request POST \
--header "Content-Type: application/json" \
--data '{"message":"Hello Freddie!"}' \
http://localhost:3500/v1.0/invoke/server.${NAMESPACE}/method/echo
{"message":"Hello Freddie!"}
```

And:

```bash
kubectl logs deployment/client \
--container=app \
--namespace=${NAMESPACE}

2021/02/19 21:22:44 [main] Entered (appID: server.example
dapr client initializing for: 127.0.0.1:50001
[main] Response: Hello Freddie
[main] Response: Hello Freddie
[main] Response: Hello Freddie

kubectl logs deployment/server \
--container=app \
--namespace=${NAMESPACE}

2021/02/19 21:07:32 [main] Entered (port: 50051)
2021/02/19 21:07:32 [main] Start gRPC service: :50051
2021/02/19 21:09:44 [main:echo] Entered
2021/02/19 21:09:44 [main:echo] ContentType:application/json, Verb:POST, QueryString:map[], {"message":"Hello Freddie!"}
2021/02/19 21:09:45 [main:echo] Entered
2021/02/19 21:09:45 [main:echo] ContentType:application/json, Verb:POST, QueryString:map[], {"message":"Hello Freddie!"}
2021/02/19 21:19:51 [main:echo] Entered
2021/02/19 21:19:51 [main:echo] ContentType:application/json, Verb:POST, QueryString:map[], {"message":"Hello Freddie!"}
...
```
