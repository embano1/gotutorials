# About
Use simple CLI flags and environment variables to show how Kubernetes pods can be templated (e.g. Java maxHeap)  
Becomes interesting when referencing reservations/ limits in ENV:

```yaml
env:
      - name: HEAP  
        valueFrom:  
          resourceFieldRef:  
            resource: limits.memory
```

## In case you want to play with the code (otherwise skip to "Run")
`git clone https://github.com/embano1/gotutorials.git`

## Run (requires Kubernetes, e.g. minikube) 
`kubectl create -f https://raw.githubusercontent.com/embano1/gotutorials/master/env/pod.yaml` and then check output with  
`kubectl logs env`


# If you don´t have Minikube running (or don´t want to install it)
## Access the environment (from your browser)
Go to [http://labs.play-with-k8s.com] and get a virtual Kubernetes setup in your browser.

```
1. Initializes cluster master node:

 kubeadm init --apiserver-advertise-address $(hostname -i)


 2. Initialize cluster networking:

 kubectl apply -n kube-system -f \
    "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"


 3. (Optional) Initialize kube-dashboard:

 curl -L -s https://git.io/kube-dashboard  | sed 's/targetPort: 9090/targetPort: 9090\n  type: LoadBalancer/' | \
    kubectl apply -f -
```

## Install some packages (in case you want to play with the code)
`yum install git go`  
`git clone https://github.com/embano1/gotutorials.git`

## Run
`kubectl create -f https://raw.githubusercontent.com/embano1/gotutorials/master/env/pod.yaml` and then check output with  
`kubectl logs env`