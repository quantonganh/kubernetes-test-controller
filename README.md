How to build a custom controller using [code-generator](https://github.com/kubernetes/code-generator)?

Create required files:

```shell
pkg/apis/foo/v1/
├── doc.go
├── register.go
├── types.go
```

Generate code:

```shell
./generate-groups.sh all github.com/quantonganh/kubernetes-test-controller/pkg/client github.com/quantonganh/kubernetes-test-controller/pkg/apis foo:v1 --go-header-file hack/boilerplate.go.txt

├── pkg
│   ├── apis
│   │   └── foo
│   │       ├── register.go
│   │       └── v1
│   │           ├── doc.go
│   │           ├── register.go
│   │           ├── types.go
│   │           └── zz_generated.deepcopy.go
│   └── client
│       ├── clientset
│       │   └── versioned
│       │       ├── clientset.go
│       │       ├── doc.go
│       │       ├── fake
│       │       ├── scheme
│       │       └── typed
│       ├── informers
│       │   └── externalversions
│       │       ├── factory.go
│       │       ├── foo
│       │       ├── generic.go
│       │       └── internalinterfaces
│       └── listers
│           └── foo
│               └── v1
```

Create a CustomResourceDefinition:

```shell
kubectl create -f crd.yaml
```

Create a custom resource of type Foo:

```shell
kubectl create -f example-foo.yaml
```

Build and run:

```shell
$ go build -o test-controller -v cmd/controller/main.go
$ ./test-controller 
I0414 08:26:20.889250   37404 main.go:94] Waiting cache to be synced
I0414 08:26:20.931570   37404 main.go:51] Added: &{{Foo foo.com/v1} {example-foo  default  04f95f07-0e18-470b-9d27-e3615a861b01 135626 1 2021-04-14 08:18:14 +0700 +07 <nil> <nil> map[] map[] [] []  [{kubectl-create Update foo.com/v1 2021-04-14 08:18:14 +0700 +07 FieldsV1 {"f:spec":{".":{},"f:deploymentName":{},"f:replicas":{}}}}]} {example-foo 0xc00004bd8c} {0}}
I0414 08:26:20.989578   37404 main.go:104] Starting custom controller
I0414 08:27:20.935377   37404 main.go:54] Updated: &{{Foo foo.com/v1} {example-foo  default  04f95f07-0e18-470b-9d27-e3615a861b01 135626 1 2021-04-14 08:18:14 +0700 +07 <nil> <nil> map[] map[] [] []  [{kubectl-create Update foo.com/v1 2021-04-14 08:18:14 +0700 +07 FieldsV1 {"f:spec":{".":{},"f:deploymentName":{},"f:replicas":{}}}}]} {example-foo 0xc00004bd8c} {0}}

```