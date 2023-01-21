proto_dir = ./proto

protogen:
	rm -rf $(proto_dir)/*.go
	protoc -I ./proto --go_out=$(proto_dir) --go_grpc_out=$(proto_dir) $(proto_dir)/*.proto

buildimg:
	# docker rmi $(shell docker images -q alex/bookserv:1.0.0)
	docker build -t alex/bookserv:1.0.0 .
	# docker exec -it $(kind get clusters | head -1)-control-plane crictl images --
	# docker exec -it $(kind get clusters | head -1)-control-plane crictl rmi <image id>
	# kind load docker-image alex/bookserv:1.0.0

runclient:
	# kubectl run busybox -i --tty --image=busybox --restart=Never --rm -- sh
	
	# wget https://github.com/moparisthebest/static-curl/releases/download/v7.80.0/curl-amd64
	# chmod +x ./curl-amd64

	# ./curl-amd64 -sSL -k  "https://github.com/fullstorydev/grpcurl/releases/download/v1.8.7/grpcurl_1.8.7_linux_x86_64.tar.gz" | tar -xz

	# grpcurl -plaintext bookserv-svc:9000 list
	# grpcurl -d '{"id": 1234, "tags": ["foo","bar"]}' \
    # grpc.server.com:443 my.custom.server.Service/Method
 

.PHONY: protogen buildimg runclient
