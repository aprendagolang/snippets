load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_push")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/aprendagolang/gorilla
gazelle(name = "gazelle")

go_library(
    name = "gorilla_lib",
    srcs = ["main.go"],
    importpath = "github.com/aprendagolang/gorilla",
    visibility = ["//visibility:private"],
    deps = [
        "//handlers",
        "@com_github_gorilla_mux//:go_default_library",
    ],
)

go_binary(
    name = "gorilla",
    embed = [":gorilla_lib"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "api_image",
    binary = ":gorilla",
    importpath = "github.com/aprendagolang/gorilla",
)

container_push(
   name = "push_api",
   image = ":api_image",
   format = "Docker",
   registry = "index.docker.io",
   repository = "aprendagolang/gorilla-mux",
   tag = "bazel",
)
