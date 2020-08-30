# Minikube を起動
minikube start --kubernetes-version v1.16.13
eval $(minikube docker-env)

# context を設定
kubectl config set-context hatena-intern-2020 --cluster=minikube --user=minikube --namespace=hatena-intern-2020
kubectl config use-context hatena-intern-2020

# 起動
make up

minikube -n hatena-intern-2020 service blog
