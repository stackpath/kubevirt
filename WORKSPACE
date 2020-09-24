load(
    "@bazel_tools//tools/build_defs/repo:http.bzl",
    "http_archive",
    "http_file",
)

# Additional bazel rules
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "301c8b39b0808c49f98895faa6aa8c92cbd605ab5ad4b6a3a652da33a1a2ba2e",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.0/rules_go-0.18.0.tar.gz"],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "7949fc6cc17b5b191103e97481cf8889217263acf52e00b560683413af204fcb",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.16.0/bazel-gazelle-0.16.0.tar.gz"],
)

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "e0b5b400cfef17d65886365dc7289cb4ef8dfe07066165607413a271a32aa2a4",
    strip_prefix = "buildtools-db073457c5a56d810e46efc18bb93a4fd7aa7b5e",
    # version 0.20.0
    url = "https://github.com/bazelbuild/buildtools/archive/db073457c5a56d810e46efc18bb93a4fd7aa7b5e.zip",
)

load(
    "@bazel_tools//tools/build_defs/repo:git.bzl",
    "git_repository",
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "aed1c249d4ec8f703edddf35cbe9dfaca0b5f5ea6e4cd9e83e99f3b0d1136c3d",
    strip_prefix = "rules_docker-0.7.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.7.0.tar.gz"],
)

http_archive(
    name = "com_github_atlassian_bazel_tools",
    sha256 = "e4737fd3636d23f12cd3f9880b1cfa75c1bbdd4a967852785e227f3b0ab11844",
    strip_prefix = "bazel-tools-7d296003f478325b4a933c2b1372426d3a0926f0",
    urls = ["https://github.com/atlassian/bazel-tools/archive/7d296003f478325b4a933c2b1372426d3a0926f0.zip"],
)

# Libvirt dependencies
http_file(
    name = "libvirt_libs",
    sha256 = "681a5fdb707027a9ae52fb403117856613505cf72ba5e72fbcd629eb6231a28e",
    urls = [
        "https://copr-be.cloud.fedoraproject.org/results/%40virtmaint-sig/for-kubevirt/fedora-29-x86_64/01113098-libvirt/libvirt-libs-5.0.0-2.fc29.x86_64.rpm",
    ],
)

http_file(
    name = "libvirt_devel",
    sha256 = "fabcd2c4aaf36aaf530c37591dbba7104aa7410069c0b31dc6248d73eb5e23bb",
    urls = [
        "https://copr-be.cloud.fedoraproject.org/results/%40virtmaint-sig/for-kubevirt/fedora-29-x86_64/01113098-libvirt/libvirt-devel-5.0.0-2.fc29.x86_64.rpm",
    ],
)

# Disk images
http_file(
    name = "alpine_image",
    sha256 = "5a4b2588afd32e7024dd61d9558b77b03a4f3189cb4c9fc05e9e944fb780acdd",
    urls = [
        "http://dl-cdn.alpinelinux.org/alpine/v3.7/releases/x86_64/alpine-virt-3.7.0-x86_64.iso",
    ],
)

http_file(
    name = "cirros_image",
    sha256 = "a8dd75ecffd4cdd96072d60c2237b448e0c8b2bc94d57f10fdbc8c481d9005b8",
    urls = [
        "https://download.cirros-cloud.net/0.4.0/cirros-0.4.0-x86_64-disk.img",
    ],
)

http_file(
    name = "fedora_image",
    sha256 = "72b6ae7b4ed09a4dccd6e966e1b3ac69bd97da419de9760b410e837ba00b4e26",
    urls = [
        "http://mirror.sfo12.us.leaseweb.net/fedora/linux/releases/30/Cloud/x86_64/images/Fedora-Cloud-Base-30-1.2.x86_64.qcow2",
    ],
)

http_file(
    name = "virtio_win_image",
    sha256 = "97e9f9c7a47cb3d2ba744b8e1c78008a20a2804c89657fcae47113a96885c566",
    urls = [
        "https://fedorapeople.org/groups/virt/virtio-win/direct-downloads/latest-virtio/virtio-win.iso",
    ],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.11.5",
    nogo = "@//:nogo_vet",
)

load("@com_github_atlassian_bazel_tools//goimports:deps.bzl", "goimports_dependencies")

goimports_dependencies()

load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies",
    "go_repository",
)

gazelle_dependencies()

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

# Winrmcli dependencies
go_repository(
    name = "com_github_masterzen_winrmcli",
    commit = "c85a68ee8b6e3ac95af2a5fd62d2f41c9e9c5f32",
    importpath = "github.com/masterzen/winrm-cli",
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_image",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

# Pull base image fedora28
container_pull(
    name = "fedora",
    digest = "sha256:d6a6d60fda1b22b6d5fe3c3b2abe2554b60432b7b215adc11a2b5fae16f50188",
    registry = "index.docker.io",
    repository = "library/fedora",
    #tag = "28",
)

# Pull base image libvirt
container_pull(
    name = "libvirt",
    digest = "sha256:71d10b2ae0af286e4ff0674a7ecfb86ac2455e529fafc97810101f5700b4a416",
    registry = "index.docker.io",
    repository = "kubevirt/libvirt",
    #tag = "5.0.0",
)

# Pull kubevirt-testing image
container_pull(
    name = "kubevirt-testing",
    digest = "sha256:eb86f7388217bb18611c8c4e6169af3463c2a18f420314eb4d742b3d3669b16f",
    registry = "index.docker.io",
    repository = "kubevirtci/kubevirt-testing",
    #tag = "28",
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

http_archive(
    name = "io_bazel_rules_container_rpm",
    sha256 = "151261f1b81649de6e36f027c945722bff31176f1340682679cade2839e4b1e1",
    strip_prefix = "rules_container_rpm-0.0.5",
    urls = ["https://github.com/rmohr/rules_container_rpm/archive/v0.0.5.tar.gz"],
)

# Get container-disk-v1alpha RPM's
http_file(
    name = "qemu-img",
    sha256 = "e02c4181a61cc48de4a0248fe56248d3c2082853e8eb16453bf89d65d204007f",
    urls = [
        "https://dl.fedoraproject.org/pub/fedora/linux/updates/30/Everything/x86_64/Packages/q/qemu-img-3.1.1-2.fc30.x86_64.rpm",
    ],
)

http_file(
    name = "bzip2",
    sha256 = "be206924278c745fc04b6b299ac46b4d0c9fe3d979bf34232e0a4071c434741a",
    urls = [
        "https://dl.fedoraproject.org/pub/fedora/linux/releases/30/Everything/x86_64/os/Packages/b/bzip2-1.0.6-29.fc30.x86_64.rpm",
    ],
)

http_file(
    name = "capstone",
    sha256 = "847ebb3a852f82cfe932230d1700cb8b90f602acbe9f9dcf5de7129a1d222c6b",
    urls = [
        "https://archives.fedoraproject.org/pub/archive/fedora/linux/updates/28/Everything/x86_64/Packages/c/capstone-3.0.5-1.fc28.x86_64.rpm",
    ],
)

http_file(
    name = "libaio",
    sha256 = "48ada36d31c735b5a4abddd76914d7509fd8784e121cc1b8fe8705dcefd9803c",
    urls = [
        "https://dl.fedoraproject.org/pub/fedora/linux/releases/30/Everything/x86_64/os/Packages/l/libaio-0.3.111-4.fc30.x86_64.rpm",
    ],
)

http_file(
    name = "libstdc",
    sha256 = "f7271e8986d90ca2d573d984e4d5bb180172b55cd8ad02be48525aeeebe66f96",
    urls = [
        "https://dl.fedoraproject.org/pub/fedora/linux/updates/30/Everything/x86_64/Packages/l/libstdc++-9.3.1-2.fc30.x86_64.rpm",
    ],
)

http_file(
    name = "qemu-guest-agent",
    sha256 = "84385346dbaadb92d7e6e3063878de8d72489c848a7aa0b8b1c7457cbc9fcc1e",
    urls = [
        "https://dl.fedoraproject.org/pub/fedora/linux/updates/30/Everything/x86_64/Packages/q/qemu-guest-agent-3.1.1-2.fc30.x86_64.rpm",
    ],
)

http_file(
    name = "stress",
    sha256 = "a3f7e2c6e4e77d980e493b5816da813116a1c3d8209a85e604b9a2ad52436452",
    urls = [
        "https://dl.fedoraproject.org/pub/fedora/linux/releases/30/Everything/x86_64/os/Packages/s/stress-1.0.4-22.fc30.x86_64.rpm",
    ],
)

http_file(
    name = "e2fsprogs",
    sha256 = "b0ef4faa1abcc38230687bfb95bd4e0f326669b8262bc1cd612ad58385299bf4",
    urls = [
        "https://dl.fedoraproject.org/pub/fedora/linux/updates/30/Everything/x86_64/Packages/e/e2fsprogs-1.44.6-2.fc30.x86_64.rpm",
    ],
)
