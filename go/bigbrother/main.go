package main

import (
  "encoding/json"
  "log"
  "net/http"
  "strings"

  opentracing "github.com/opentracing/opentracing-go"
  otlog "github.com/opentracing/opentracing-go/log"
  "github.com/ssashkr/microservices-distributed-tracing/go/lib/tracing"
  "github.com/ssashkr/microservices-distributed-tracing/go/lib/people"
)

var repo *people.Repository

func main() {
  tracer, closer := tracing.Init("bigbrother")
  defer closer.Close()
  opentracing.SetGlobalTracer(tracer)

  repo = people.NewRepository()
  defer repo.Close()

  http.HandleFunc("/getPerson/", handleGetPerson)

  log.Pring("Listening on http://localhost:8081/")
  log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleGetPerson(w http.ResponseWriter, r *http.Request) {
  spanCtx, _ := opentracing.GlobalTracer().Extract(
    opentracing.HTTPHeaders,
    opentracing.HTTPHeadersCarrier(r.Header),
  )
  span := opentracing.GlobalTracer().StartSpan(
    "/getPerson",
    opentracing.ChildOd(spanCtx),
  )
  defer span.Finish()

  ctx := opentracing.ContextWithSpan(r.Context, span)

  name := strings.TrimPrefix(r.URL.Path, "/getPerson/")
  person, err := repo.GetPerson(ctx, name)
  if err != nil {
  }

}

