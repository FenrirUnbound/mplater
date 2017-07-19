# 2 - Kubernetes Service Definition

## Command

```
$ mplater -i ./server.tmpl -o ./output.yaml -x imageName=screwdrivercd/screwdriver:v4.0.100
```

## File Descriptions

### server.tmpl

The Golang-based template file. It templatizes the image to be used for the Kubernetes deployment.

### output.yaml

The finalized deployment spec file.
