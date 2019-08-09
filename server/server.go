package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/handler"
	"github.com/julz/kngraphql"
	"github.com/julz/kngraphql/resolver"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"knative.dev/serving/pkg/client/clientset/versioned"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	clientset := versioned.NewForConfigOrDie(MustGetConfig())
	resolver := resolver.NewResolver(clientset)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(kngraphql.NewExecutableSchema(kngraphql.Config{Resolvers: resolver})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func MustGetConfig() *rest.Config {
	if config, err := rest.InClusterConfig(); err == nil {
		return config
	}

	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		log.Fatal(err)
	}

	return config
}
