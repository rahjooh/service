# service

based on the makefile , you can run running cluster : 
```shell
make dev-up
```

it will use /zarf/k8s/dev/king-config.yaml , to run the cluster and its wait 120s 

```
Creating cluster "hadi-cluster" ...
✓ Ensuring node image (kindest/node:v1.27.1) 🖼
✓ Preparing nodes 📦  
✓ Writing configuration 📜
✓ Starting control-plane 🕹️
✓ Installing CNI 🔌
✓ Installing StorageClass 💾
Set kubectl context to "kind-hadi-cluster"
You can now use your cluster with:

kubectl cluster-info --context kind-hadi-cluster

Thanks for using kind! 😊
kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner`
```

after your cluster up you can check services and pods status with thanks of the makefile : 
```shell
 make dev-status
 ```

```
kubectl get nodes -o wide
NAME                         STATUS   ROLES           AGE     VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE                         KERNEL-VERSION    CONTAINER-RUNTIME
hadi-cluster-control-plane   Ready    control-plane   3m33s   v1.27.1   172.18.0.2    <none>        Debian GNU/Linux 11 (bullseye)   6.10.0-linuxkit   containerd://1.6.21
kubectl get svc -o wide
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE     SELECTOR
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   3m31s   <none>
kubectl get pods -o wide --watch --all-namespaces
NAMESPACE            NAME                                                 READY   STATUS    RESTARTS   AGE     IP           NODE                         NOMINATED NODE   READINESS GATES
kube-system          coredns-5d78c9869d-94wxd                             1/1     Running   0          3m16s   10.244.0.2   hadi-cluster-control-plane   <none>           <none>
kube-system          coredns-5d78c9869d-h5d46                             1/1     Running   0          3m16s   10.244.0.3   hadi-cluster-control-plane   <none>           <none>
kube-system          etcd-hadi-cluster-control-plane                      1/1     Running   0          3m31s   172.18.0.2   hadi-cluster-control-plane   <none>           <none>
kube-system          kindnet-j9xpz                                        1/1     Running   0          3m16s   172.18.0.2   hadi-cluster-control-plane   <none>           <none>
kube-system          kube-apiserver-hadi-cluster-control-plane            1/1     Running   0          3m30s   172.18.0.2   hadi-cluster-control-plane   <none>           <none>
kube-system          kube-controller-manager-hadi-cluster-control-plane   1/1     Running   0          3m31s   172.18.0.2   hadi-cluster-control-plane   <none>           <none>
kube-system          kube-proxy-r5zpl                                     1/1     Running   0          3m16s   172.18.0.2   hadi-cluster-control-plane   <none>           <none>
kube-system          kube-scheduler-hadi-cluster-control-plane            1/1     Running   0          3m30s   172.18.0.2   hadi-cluster-control-plane   <none>           <none>
local-path-storage   local-path-provisioner-6bc4bddd6b-nbn6f              1/1     Running   0          3m16s   10.244.0.4   hadi-cluster-control-plane   <none>           <none>
```

and finally you can down and remove the cluster :

``` shell
make dev-down    
```
```
kind delete cluster --name hadi-cluster
Deleting cluster "hadi-cluster" ...
Deleted nodes: ["hadi-cluster-control-plane"]
```