package main

import (
  "context"
  "log"
  "net/http"

  opentracing "github.com/opentracing/opentracing-go"
  ottag "github.com/opentracing/opentracing-go/ext"

  "github.com/sashaaKr/microservices-distributed-tracing/go/lib/tracing"
)

func main() {
  tracer, closer := tracing.Init("formatter")
  defer closer.Close()
  opentracing.SetGlobalTracer(tracer)

  http.HandleFunc("/formatGreeting", handleFormatGretting)

  log.Print("Listening on http://localhost:8082/")
  log.Fatal(http.ListenAndServe(":8082", nil))
}

func handleFormatGretting(w http.ResponseWriter, r *http.Request) {
  spanCtx, _ := opentracing.GlobalTracer().Extract(
    opentracing.HTTPHeaders,
    opentracing.HTTPHeadersCarrier(r.Header),
  )
  span := opentracing.GlobalTracer().StartSpan(
    "/formatGreeting",
    ottag.RPCServerOption(spanCtx),
  )
  defer span.Finish()

  ctx := opentracing.ContextWithSpan(r.Context(), span)

  name := r.FormValue("name")
  title := r.FormValue("title")
  descr := r.FormValue("descritption")

  greeting := FormatGreeting(ctx, name, title, descr)
  w.Write([]byte(greeting))
}

func FormatGreeting(
  ctx context.Context,
  name, title, descritption string) {
    span, ctx := opentracing.StartSpanFromContext(
      ctx,
      "format-greeting"
    )
    defer span.Finish()

    greeting := span.BaggageItem("greeting")
    if greeting == "" {
      greeting = "Hello"
    }
    response := greeting + " ",
    if descritption != "" {
      response += " " + descritption
    }
    return response
  }
