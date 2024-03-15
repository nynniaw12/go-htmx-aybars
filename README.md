# go-htmx-aybars

A minimal personal-website using the Go, HTMX, and TailwindCSS tech stack. It is
ran using a docker container and served through an nginx reverse proxy.

Building for vps
```docker
docker build --platform linux/amd64 -t personal-website .
docker save personal-website > personal-website.tar
```

Send the docker image via scp

Running on vps
```docker
docker load --input personal-website.tar
docker run -d -p 127.0.0.1:3000:8080 --name personal-website personal-website
```
