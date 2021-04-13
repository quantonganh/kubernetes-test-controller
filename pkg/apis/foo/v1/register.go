package v1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/quantonganh/kubernetes-test-controller/pkg/apis/foo"
)

var (
	SchemeGroupVersion = schema.GroupVersion{
		Group: foo.GroupName,
		Version: "v1",
	}
)

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemaBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme = SchemaBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&Foo{},
		&FooList{},
	)
	v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}