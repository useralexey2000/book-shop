proto_dir = ./proto
srv_img_name = bookserv
gw_img_name = bookservgw
remote_test_host = 192.168.56.3
remote_test_dir = /home/al/k8s/test
remote_user = al

protogen:
	rm -rf $(proto_dir)/*.go
	protoc -I ./proto --go_out=$(proto_dir) --go_grpc_out=$(proto_dir) $(proto_dir)/*.proto

build_books_grpc:
	# @docker rmi -f $(shell docker images -q alex/bookserv:1.0.0)
 	@docker build -f Dockerfile -t alex/$(srv_img_name):1.0.0 .
	
build_books_gw:
	docker build -f gw.Dockerfile -t alex/$(gw_img_name):1.0.0 .

clean_cluster_reg:
	# docker exec -it $(kind get clusters | head -1)-control-plane crictl images --
	# docker exec -it $(kind get clusters | head -1)-control-plane crictl rmi <image id>

load_images_to_reg:
	# kind load docker-image alex/$(srv_img_name):1.0.0
	# kind load docker-image alex/$(gw_img_name):1.0.0

runclient:
	kubectl run -n book-shop -i --tty book-client --rm --image=busybox --restart=Never -- /bin/sh -c 'wget --header="Content-Type:application/json" --post-data="{\"offset\":0, \"limit\":100}" -q -O- http://bookserv-svc:8080/books/list'

autoscale_test:
	kubectl run -n book-shop -i --tty load-generator --rm --image=busybox --restart=Never -- /bin/sh -c 'while sleep 0.01; do wget --header="Content-Type:application/json" --post-data="{\"offset\":0, \"limit\":100}" -q -O- http://bookserv-svc:8080/books/list; done'
	# check scaling
	# kubectl get pods -n book-shop
	# kubectl top pod -n book-shop book-deployment-<hash> --containers

test:
	echo ${remote_user}@${remote_test_host}:${remote_test_dir}
	rsync -r test/* ${remote_user}@${remote_test_host}:${remote_test_dir}
	#TODO create docker-compose
	docker run -it --rm -v ${remote_test_dir}:/behave:ro  williamyeh/behave

.PHONY: protogen build_books_grpc build_books_gw clean_cluster_reg load_images_to_reg runclient test