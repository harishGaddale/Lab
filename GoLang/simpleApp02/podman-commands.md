# build the image
podman build -t simple-go-app-02:v1.0.0
# run the above built image as a container
podman run -d -p 8852:8852 --name simple-go-app-02 simple-go-app-02:v1.0.0
