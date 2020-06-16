### helm commands

https://habr.com/ru/company/flant/blog/420437/
https://cloud.croc.ru/blog/byt-v-teme/kubernetes-helm-chart/

## install
helm install my my-example-chart
helm upgrade nsqadmin nsqadmin_my

## sh
kubectl exec -it nsqlookup-nsqlookupd-0 sh

## templ
helm template . --debug  > tmp.yaml