load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "register",
    srcs = ["register.go"],
    importpath = "ok.build/cli/command/register",
    visibility = ["//visibility:public"],
    deps = [
        "//cli/command",
        "//cli/please",
        "//cli/version",
    ],
)

package(default_visibility = ["//cli:__subpackages__"])
