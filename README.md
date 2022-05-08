## snap-hello-world

```
sudo snap install multipass

sudo snap install snapcraft --classic

cd snap-hello-world
snapcraft init

# Run from root (outside from snap/). Destructive mode allows nested virtualization on EC2s, recommended for CI systems as well.
snapcraft --destructive-mode
```