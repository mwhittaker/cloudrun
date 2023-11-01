# Cloud Run

This repository shows how to run a Service Weaver application on Cloud Run.
First, create a Docker container and upload it to Artifact Registry according to
[these instructions](https://cloud.google.com/run/docs/building/containers).

```shell
$ docker build -t REGION-docker.pkg.dev/PROJECT_ID/REPO_NAME/PATH:TAG .
$ docker push REGION-docker.pkg.dev/PROJECT_ID/REPO_NAME/PATH:TAG
```

Then, deploy the container to Cloud Run using `gcloud run deploy`:

```shell
$ gcloud run deploy NAME --image=REGION-docker.pkg.dev/PROJECT_ID/REPO_NAME/PATH:TAG --region=REGION --allow-unauthenticated
```

This command should print out a URL that you can use to access the service.
Alternatively, you can curl the service from the command line:

```shell
$ curl -H "Authorization: Bearer $(gcloud auth print-identity-token)" URL
```
