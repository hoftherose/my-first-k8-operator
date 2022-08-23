# my-first-k8-operator
My First Kubernetes Operator

## Install and Setup
To install and setup make sure that your DNS can recognize the URL in your current kube-context. Run 'make install' and ensure that all install processes run without error, then run 'make run' to start the service. Confirm that everything is running smoothly running 'kubectl api-resources' and look for your api. If you cannot find it confirm install ran successfully but looking for the CRD in 'kubectl get crd'.
