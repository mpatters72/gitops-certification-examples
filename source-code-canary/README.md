# Mike's Customizations

- Added `/health` and `/version` json response.

- Updated aline version

## Creating images

- comment out `version` and `color`, and set docker tag to corresponding v1, v2, v3

- create 3 images

```bash
export myrepo="replaceme.com"
# comment out except v2 version and blue 
docker build . -t ${myrepo}/gitops-canary-app:v1.0
# comment out except v2 version and yellow
docker build . -t ${myrepo}/gitops-canary-app:v2.0
# comment out except v3 version and magenta
docker build . -t ${myrepo}/gitops-canary-app:v3.0
```

- test

```bash
docker run --rm -p 8080:8080 ${myrepo}/gitops-canary-app:v1.0
curl http://localhost:8080/health
curl http://localhost:8080/version
curl http://localhost:8080/callme
curl http://localhost:8080/
```

- then push

