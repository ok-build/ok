load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "please",
    srcs = ["please.go"],
    importpath = "ok.build/cli/please",
    deps = [
        "//cli/claude",
    ],
)

package(default_visibility = ["//cli:__subpackages__"])
