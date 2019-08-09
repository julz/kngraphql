kngraphql - some simple graphql bindings for knative

## Usage

~~~~
# locally
go run github.com/julz/kngraphql/server

# in-cluster (assuming you have ko installed from https://github.com/google/ko)
# .. yes, it installs as a knative service :-)
ko apply -f hack/deploy.yaml
~~~~
