kngraphql - some simple graphql bindings for knative

## Usage

1. Run locally (assumes your `.kube/config` is set):

   ~~~~
   go run github.com/julz/kngraphql/server
   ~~~~~

1. Or: deploy in-cluster (assuming you have `ko` installed & configured, from https://github.com/google/ko)

   (.. yes, of course it deploys as a knative service :-))

   ~~~~
   ko apply -f hack/deploy.yaml 
   ~~~~

   After deploying, access at `http://kngraphql.default.{your app domain}`.
