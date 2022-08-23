# my-first-k8-operator
My First Kubernetes Operator

## Install and Setup
To install and setup make sure that your DNS can recognize the URL in your current kube-context. Run 'make install' and ensure that all install processes run without error, then run 'make run' to start the service. Confirm that everything is running smoothly running 'kubectl api-resources' and look for your api. If you cannot find it confirm install ran successfully but looking for the CRD in 'kubectl get crd'.

## Upload your first Custom Resource (CR)
In the folder config/samples there should be a file "{GROUP}_v1_test.yaml" which you can use to define a deployment for your new CR. By default, this new resource will only have the string parameter 'Foo' so you can only add this after spec. You can review the parameter on the etcd version of the yaml.