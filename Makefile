
all:
	go fmt ./...
	go test ./...
	go mod tidy

minikube: all
	eval $$(minikube docker-env) && docker build -t 2020-01-kubernetes-client-go-job .
	kubectl config use-context minikube
	kubectl delete job/pi && kubectl wait --for=delete job/pi || echo ignore error
	kubectl delete -f minikube.yaml && kubectl wait --for=delete pod/example || echo ignore error
	kubectl apply -f minikube.yaml
	sleep 5
	kubectl wait --for=condition=complete job/pi
	kubectl logs pod/example
	kubectl logs job/pi
