#### use make file format

example : [cluster-capacity Makefile](https://github.com/kubernetes-incubator/cluster-capacity/blob/master/Makefile)

```
build:
	go build -o hypercc github.com/kubernetes-incubator/cluster-capacity/cmd/hypercc
	ln -sf hypercc cluster-capacity
	ln -sf hypercc genpod
run:
	@./cluster-capacity --kubeconfig ~/.kube/config --podspec=examples/pod.yaml --verbose

test:
	./test.sh

integration-tests:
	./integration-tests.sh

image:
	docker build -t cluster-capacity .

clean:
	rm -f cluster-capacity genpod hypercc
```
