##To see minikube IP or service URLs, use the following commands
```sh
minikube ip
minikube service --url mongo-nodeport-svc
```

#to create secrets keys 
echo "mega_secret_key" | base64

Command to connect mongo:
mongo --host <ip> --port <port of nodeport svc> -u adminuser -p password123

Exec into the client.
kubectl exec deployment/mongo-client -it -- /bin/bash

Login into the MongoDB shell
mongo --host mongo-nodeport-svc --port 27017 -u adminuser -p password123

#to clean up the deployment objects, execute the following
kubectl delete -f .