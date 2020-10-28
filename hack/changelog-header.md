### Linux

```shell     

curl -L https://github.com/jenkins-x-plugins/kuberhealthy-terraform-drift-check/releases/download/v{{.Version}}/kuberhealthy-terraform-drift-check-linux-amd64.tar.gz | tar xzv 
sudo mv kh-terraform-drift /usr/local/bin
```

### macOS

```shell
curl -L  https://github.com/jenkins-x-plugins/kuberhealthy-terraform-drift-check/releases/download/v{{.Version}}/kuberhealthy-terraform-drift-check-darwin-amd64.tar.gz | tar xzv
sudo mv kh-terraform-drift /usr/local/bin
```

