## snap-hello-world

### Prepare environment and build
```
# install snapcraft
sudo snap install snapcraft --classic 

# initialize snapcraft
snapcraft init 

# build the code and edit snap/snapcraft.yaml accordingly

# build snap using lxd. Multipass not working on EC2s
sg lxd -c 'snapcraft --use-lxd' 

#install unsigned snap locally
sudo snap install nockberge_0.0.2_amd64.snap --dangerous 
```

### Publish snap
```
snapcraft login
snapcraft register nockberge

# Push to Snap Store withou publishing
snapcraft push nockberge_0.0.2_amd64.snap

# Push and release immediately
snapcraft upload --release=edge nockberge_0.0.3_amd64.snap

```


### Check and install snap
```
snap info nockberge
sudo snap install nockberge
```