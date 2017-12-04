run-api:
	export KUBERNETES_SERVICE_HOST='localhost' && \
    export KUBERNETES_SERVICE_PORT='9443' && \
    bin/apiserver --etcd-servers=http://localhost:2379 --secure-port=9443 --delegated-auth=false 
