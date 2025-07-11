load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "ok_lib",
    srcs = ["ok.go"],
    data = [
    ],
    importpath = "ok.build/cli/cmd/ok",
    visibility = ["//visibility:private"],
    deps = [
        "//cli/arg",
        "//cli/bazelisk",
        "//cli/claude",
        "//cli/command",
        "//cli/command/register",
        "//cli/help",
        "//cli/log",
        "//cli/picker",
    ],
)

go_binary(
    name = "ok",
    embed = [":ok_lib"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "ok-darwin-amd64",
    out = "ok-darwin-amd64",
    embed = [":ok_lib"],
    gc_linkopts = [
        "-s",
        "-w",
    ],
    goarch = "amd64",
    goos = "darwin",
    pure = "on",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "ok-darwin-arm64",
    out = "ok-darwin-arm64",
    embed = [":ok_lib"],
    gc_linkopts = [
        "-s",
        "-w",
    ],
    goarch = "arm64",
    goos = "darwin",
    pure = "on",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "ok-linux-amd64",
    out = "ok-linux-amd64",
    embed = [":ok_lib"],
    gc_linkopts = [
        "-s",
        "-w",
    ],
    goarch = "amd64",
    goos = "linux",
    pure = "on",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "ok-linux-arm64",
    out = "ok-linux-arm64",
    embed = [":ok_lib"],
    gc_linkopts = [
        "-s",
        "-w",
    ],
    goarch = "arm64",
    goos = "linux",
    pure = "on",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "ok-windows-amd64",
    out = "ok-windows-amd64.exe",
    embed = [":ok_lib"],
    goarch = "amd64",
    goos = "windows",
    pure = "on",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)

package(default_visibility = ["//cli:__subpackages__"])
