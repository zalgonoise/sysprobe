load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "dbf5a9ef855684f84cac2e7ae7886c5a001d4f66ae23f6904da0faaaef0d61fc",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.11/rules_go-v0.24.11.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.11/rules_go-v0.24.11.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "222e49f034ca7a1d1231422cdb67066b885819885c356673cb1f72f748a3c9d4",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_bazelbuild_bazel_gazelle",
    importpath = "github.com/bazelbuild/bazel-gazelle",
    sum = "h1:cfvF5dcjFKTehIQkLWioFWpGWi41Q38/WBM2d/2QCTs=",
    version = "v0.22.3",
)

go_repository(
    name = "com_github_bazelbuild_buildtools",
    importpath = "github.com/bazelbuild/buildtools",
    sum = "h1:Et1IIXrXwhpDvR5wH9REPEZ0sUtzUoJSq19nfmBqzBY=",
    version = "v0.0.0-20200718160251-b1667ff58f71",
)

go_repository(
    name = "com_github_bazelbuild_rules_go",
    importpath = "github.com/bazelbuild/rules_go",
    sum = "h1:wzbawlkLtl2ze9w/312NHZ84c7kpUCtlkD8HgFY27sw=",
    version = "v0.0.0-20190719190356-6dae44dc5cab",
)

load("//:repositories.bzl", "go_repositories")

# gazelle:repository_macro repositories.bzl%go_repositories
go_repositories()

gazelle_dependencies()

# external dependencies

go_repository(
    name = "com_github_go-ping_ping",
    importpath = "github.com/go-ping/ping",
    sum = "h1:jI2GiiRh+pPbey52EVmbU6kuLiXqwy4CXZ4gwUBj8Y0=",
    version = "v0.0.0-20201115131931-3300c582a663",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:003p0dJM77cxMSyCPFphvZf/Y5/NXf5fzg6ufd1/Oew=",
    version = "v0.0.0-20210119194325-5f4716e94777",
)
