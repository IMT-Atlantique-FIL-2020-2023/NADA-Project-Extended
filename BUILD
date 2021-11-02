load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended
gazelle(name = "gazelle")

go_library(
    name = "NADA-extended_lib",
    srcs = ["main.go"],
    importpath = "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "NADA-extended",
    embed = [":NADA-extended_lib"],
    visibility = ["//visibility:public"],
)
