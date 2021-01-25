load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "com_github_go_ping_ping",
        importpath = "github.com/go-ping/ping",
        sum = "h1:jI2GiiRh+pPbey52EVmbU6kuLiXqwy4CXZ4gwUBj8Y0=",
        version = "v0.0.0-20201115131931-3300c582a663",
    )
