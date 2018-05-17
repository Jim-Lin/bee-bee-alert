# architecture
![architecture](architecture.png?raw=true)

## service
- elasticsearch

  search similar product

- redis

  counter for hit product

- golang backend
  - check product price and notify when it's expensive
  - **notify hit product for marketing**
- nginx reverse proxy

## chrome extension
### notification
![notification](notification.png?raw=true)

# local env.
requirements
- Docker

## Docker Compose
### build & run
```bash
$ docker-compose up -d
```

### cleanup
```bash
$ docker-compose rm -s
```

## Elasticsearch
### init fake data
```bash
$ curl -X PUT 'localhost:9200/bee' -H 'Content-Type: application/json' -d @./elasticsearch/fake/schema.json

$ curl -X POST 'localhost:9200/_bulk?pretty' -H 'Content-Type: application/x-ndjson' --data-binary @./elasticsearch/fake/water.json

# check index information
$ curl -X GET "localhost:9200/_cat/indices?v"
```

### cleanup
```bash
$ curl -X DELETE "localhost:9200/bee"
```

# cloud env. with GCP
![gcp](gcp.png?raw=true)

requirements
- Cloud SDK

### create project
```bash
$ gcloud projects create bee-bee-alert-jimlin --name=bee-bee-alert --set-as-default

$ gcloud config set project bee-bee-alert-jimlin

$ gcloud config set compute/zone asia-east1-b
```

Take the following steps to enable the Kubernetes Engine API:
1. Visit the https://console.cloud.google.com/projectselector/kubernetes in the Google Cloud Platform Console.
1. Create or select a project.
1. Wait for the API and related services to be enabled. This can take several minutes.

## GKE
### create cluster
```bash
$ gcloud container clusters create bee-bee-alert --machine-type n1-standard-1 --num-nodes 3 --enable-autoscaling --min-nodes 3 --max-nodes 7 --zone asia-east1-b

$ gcloud config set container/cluster bee-bee-alert

$ gcloud container clusters get-credentials bee-bee-alert
```

### create configmap
```bash
$ kubectl create configmap nginx-config --from-file=./nginx/default.conf

$ kubectl create configmap backend-config --from-file=./backend/backend.properties
```

## Continuous Deployment with Google Container Registry (GCR)
![cd](cd.png?raw=true)

### Cloud Container Builder
enable the Container Builder API  
https://console.cloud.google.com/flows/enableapi?apiid=cloudbuild.googleapis.com

give Container Builder Service Account `container.developer` role access to your Kubernetes Engine clusters
```bash
$ PROJECT="$(gcloud projects describe \
    $(gcloud config get-value core/project -q) --format='get(projectNumber)')"

$ gcloud projects add-iam-policy-binding $PROJECT \
    --member=serviceAccount:$PROJECT@cloudbuild.gserviceaccount.com \
    --role=roles/container.developer
```

### add trigger
![trigger](trigger.png?raw=true)

when push tag to GitHub, it will auto run by builder with **cloudbuild.yaml**  
![builder](builder.png?raw=true)

### init fake data
```bash
$ kubectl get pods
NAME                             READY     STATUS    RESTARTS   AGE
backend-6c78d66f56-8gkvc         1/1       Running   0          1m
backend-6c78d66f56-k7fkv         1/1       Running   0          1m
elasticsearch-5b7dc6659b-b98tf   1/1       Running   0          9m
nginx-8db8dbcdb-vpfkr            1/1       Running   0          40m
redis-54868fc78b-5jq9c           1/1       Running   0          2h

# run shell in the elasticsearch container
$ kubectl exec -it elasticsearch-5b7dc6659b-b98tf sh

$ curl -X PUT 'localhost:9200/bee' -H 'Content-Type: application/json' -d @./fake/schema.json

$ curl -X POST 'localhost:9200/_bulk?pretty' -H 'Content-Type: application/x-ndjson' --data-binary @./fake/water.json

$ exit
```

## cleanup
```bash
$ kubectl delete configmap nginx-config
$ kubectl delete configmap backend-config

$ kubectl delete --all svc
$ kubectl delete --all deployment

# wait for the load balancer delete
$ gcloud compute forwarding-rules list

$ gcloud container clusters delete bee-bee-alert
```
