package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	helmfluxv1 "github.com/fluxcd/helm-operator/pkg/apis/helm.fluxcd.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"

	"github.com/fluxcd/helm-operator/model/pkg/schemagen"
	//"github.com/fabric8io/kubernetes-client/kubernetes-model/pkg/schemagen"
)

type Schema struct {
	TypeMeta   metav1.TypeMeta
	ObjectMeta metav1.ObjectMeta

	Spec   helmfluxv1.HelmReleaseSpec
	Status helmfluxv1.HelmReleaseStatus
	//HelmRelease                              helmfluxv1.HelmRelease
	//sHelmReleaseList                          helmfluxv1.HelmReleaseList
	//JSONSchemaPropsorStringArray             apiextensions.JSONSchemaPropsOrStringArray

}

func main() {
	customTypeNames := map[string]string{
		"JSONSchemaPropsorStringArray": "JSONSchemaPropsOrStringArray",
	}
	packages := []schemagen.PackageDescriptor{
		{"k8s.io/api/core/v1", "", "io.fabric8.kubernetes.api.model", "kubernetes_core_"},
		{"k8s.io/apimachinery/pkg/apis/meta/v1", "", "io.fabric8.kubernetes.api.model.apis", "kubernetes_apimachinery_pkg_apis_"},
		{"k8s.io/apimachinery/pkg/apis/meta/v1", "", "io.fabric8.kubernetes.api.model", "kubernetes_apimachinery_"},
	}

	typeMap := map[reflect.Type]reflect.Type{
		reflect.TypeOf(time.Time{}): reflect.TypeOf(""),
		reflect.TypeOf(struct{}{}):  reflect.TypeOf(""),
	}
	schema, err := schemagen.GenerateSchema(reflect.TypeOf(Schema{}), packages, typeMap, customTypeNames)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %v", err)
		return
	}

	args := os.Args[1:]
	if len(args) < 1 || args[0] != "validation" {
		schema.Resources = nil
	}

	b, err := json.Marshal(&schema)
	if err != nil {
		log.Fatal(err)
	}
	result := string(b)
	result = strings.Replace(result, "\"additionalProperty\":", "\"additionalProperties\":", -1)

	/**
	 * Hack to fix https://github.com/fabric8io/kubernetes-client/issues/1565
	 *
	 * Right now enums are having body as array of jsons rather than being array of strings.
	 * (See https://user-images.githubusercontent.com/13834498/59852204-00d25680-938c-11e9-91b6-74f6bc3ae65b.png)
	 *
	 * I could not find any other way of fixing this since I'm not sure where it's coming from.
	 * So doing this search and replace of whole enum json object block hence converting it to an array of plain
	 * strings rather than of json objects.
	 */
	result = strings.Replace(result, "\"enum\":{\"type\":\"array\",\"description\":\"\",\"javaOmitEmpty\":true,\"items\":{\"$ref\":\"#/definitions/kubernetes_apiextensions_JSON\",\"javaType\":\"io.fabric8.kubernetes.api.model.apiextensions.JSON\"}},",
		"\"enum\":{\"type\":\"array\",\"description\":\"\",\"javaOmitEmpty\":true,\"items\":{\"type\": \"string\"}},", -1)

	var out bytes.Buffer
	err = json.Indent(&out, []byte(result), "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}
