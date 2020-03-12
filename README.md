# druid-prometheus exporter

##### MOTIVATION
- Scale druid middlemanagers on the basis of custom metrics such as pending, waiting tasks. 
- These are the metrics which currently this exporter emits count for

    1. Running Tasks
    2. Pending Tasks
    3. Waiting Tasks
    4. Completed Tasks
    5. Succeded Tasks
    5. Count of All DataSources

##### GETTING STARTED
 - To deploy a druid cluster on kuberentes, refer to this awesome operator ```https://github.com/druid-io/druid-operator```.
 - This exporter can be ran as a simple deployment in kubernetes. 
 - The deployment files expects a env for the druid endpoint ```DRUID_EP```, this can be an overlord or a router endpoint. 
 - Build your own docker image.

 ```
 - docker build -t your_repo/image_name:tag .
 - docker push image_name
 ```

 ```

 ## Deploy exporter
 - kubectl create -f deploy/exporter.yaml 
 ## Deploy prometheus
 - kubectl create -f cm.yaml 
 - kubectl create -f prometheus.yaml
 ```
- Run tasks on druid and check the metrics. Expose prometheus UI ```kubectl port-forward prometheus_pod 9090:9090```
```
- druid_running_tasks
- druid_completed_tasks
- druid_pending_tasks
- druid_waiting_tasks
- druid_datasources_count_all
```

# Custom Scaling Of Middle Managers

  - Deploy prometheus custom metrics adapter. Refer to this  ```https://github.com/stefanprodan/k8s-prom-hpa``` to get started.
  - Deploy the HPA yaml to scale on the basis of pending tasks.
 ```
 ##  Deploy HPA 
 - kubectl create -f deploy/hpa.yaml
 - kubectl get hpa
 ```
 
 ######  NOTE: This is tested for HPA apiVersion ```autoscaling/v2beta1```. Exporter is still under development.



 

