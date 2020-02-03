# easy-rstudio-k8s: Quick-and-dirty RStudio Setup in your Kubernetes Cluster

## TL;DR

### Minikube setup example

1. Edit `/etc/hosts`, add one line such as:

    ```
    ${INGRESS_IP_ADDRESS}   easy-rstudio-default    easy-rstudio-johndoe
    ```

   The value of `${INGRESS_IP_ADDRESS}` would be available by running:

    > minikube ip

2. Applying resources by running:

    > kubectl apply -f kubernetes/statefulset.yml

3. Access RStudio at http://easy-rstudio-default and http://easy-rstudio-johndoe

## NOTICE

This code base is in very early stage, I make it just because I want to make some tasks easier, there is no guarantee no warranty for it.