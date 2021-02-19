# dapr: gRPC Golang example

## Notes

+ Need to create ImagePullSecret if using private repo
+ Client address Server as `server.${NAMESPACE}`

```bash
NAMESPACE="example"

kubectl create namespace ${NAMESPACE}

kubectl create secret generic ghcr \
--from-file=.dockerconfigjson=/home/dazwilkin/snap/docker/471/.docker/config.json \
--type=kubernetes.io/dockerconfigjson \
--namespace=${NAMESPACE}

kubectl apply --filename=./kubernetes/server.yaml --namespace=${NAMESPACE}

kubectl get all --namespace=${NAMESPACE}
NAME                          READY   STATUS    RESTARTS   AGE
pod/server-7bfdb46b95-f6b48   2/2     Running   0          8s

NAME                  TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                               AGE
service/server-dapr   ClusterIP   None         <none>        80/TCP,50001/TCP,50002/TCP,9090/TCP   8s

NAME                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/server   0/1     1            0           8s

NAME                                DESIRED   CURRENT   READY   AGE
replicaset.apps/server-7bfdb46b95   1         1         0       8s
```


From e.g. `ashapp`, you can invoke a REST (!) endpoint against the gRPC server:

```bash
curl \
--request POST \
--header "Content-Type: application/json" \
--data '{"message":"Hello Freddie!"}' \
http://localhost:3500/v1.0/invoke/server.example/method/echo
{"message":"Hello Freddie!"}

curl --request GET http://localhost:3500/v1.0/invoke/server.example/method/healthz
ok
```

And:

```bash
kubectl logs deployment/client --namespace=${NAMESPACE} client
2021/02/19 21:22:44 [main] Entered (appID: server.example
dapr client initializing for: 127.0.0.1:50001
[main] Response: Hello Freddie
[main] Response: Hello Freddie
[main] Response: Hello Freddie
[main] Response: Hello Freddie

2021/02/19 21:07:32 [main] Entered (port: 50051
2021/02/19 21:07:32 [main] Start gRPC service: :50051
2021/02/19 21:09:33 [main:healthz] Entered
2021/02/19 21:09:33 [main:healthz] ContentType:application/json, Verb:GET, QueryString:map[], 
2021/02/19 21:09:44 [main:echo] Entered
2021/02/19 21:09:44 [main:echo] ContentType:application/json, Verb:POST, QueryString:map[], {"message":"Hello Freddie!"}
2021/02/19 21:09:45 [main:echo] Entered
2021/02/19 21:09:45 [main:echo] ContentType:application/json, Verb:POST, QueryString:map[], {"message":"Hello Freddie!"}
2021/02/19 21:19:51 [main:echo] Entered
2021/02/19 21:19:51 [main:echo] ContentType:application/json, Verb:POST, QueryString:map[], {"message":"Hello Freddie!"}
2021/02/19 21:20:01 [main:healthz] Entered
2021/02/19 21:20:01 [main:healthz] ContentType:application/json, Verb:GET, QueryString:map[], 
2021/02/19 21:22:44 [main:echo] Entered
2021/02/19 21:22:44 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:22:49 [main:echo] Entered
2021/02/19 21:22:49 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:22:54 [main:echo] Entered
2021/02/19 21:22:54 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:22:59 [main:echo] Entered
2021/02/19 21:22:59 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:23:04 [main:echo] Entered
2021/02/19 21:23:04 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:23:09 [main:echo] Entered
2021/02/19 21:23:09 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:23:14 [main:echo] Entered
2021/02/19 21:23:14 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:23:19 [main:echo] Entered
2021/02/19 21:23:19 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:23:24 [main:echo] Entered
2021/02/19 21:23:24 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
2021/02/19 21:23:29 [main:echo] Entered
2021/02/19 21:23:29 [main:echo] ContentType:text/plain, Verb:POST, QueryString:map[], Hello Freddie
```