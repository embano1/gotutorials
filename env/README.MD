Use simple CLI flags and environment variables to show how Kubernetes pods can be templated (e.g. Java maxHeap)  
Becomes interesting when referencing reservations/ limits in ENV:

```yaml
env:
      - name: HEAP  
        valueFrom:  
          resourceFieldRef:  
            resource: limits.memory
```

Run the pod (requires Kubernetes, e.g. minikube) with `kubectl create -f pod.yaml` and then check output with `kubectl logs env`