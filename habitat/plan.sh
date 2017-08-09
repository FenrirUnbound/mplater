pkg_origin=fenrirunbound
pkg_name=mplater
pkg_version=0.1.0
pkg_maintainer="Darren Matsumoto <aeneascorrupt@gmail.com>"
#pkg_source="https://github.com/FenrirUnbound/mplater"
pkg_upstream_url=https://github.com/FenrirUnbound/mplater
pkg_deps=()
pkg_build_deps=(core/git core/go17 core/which)
pkg_exports=()
pkg_bin_dirs=()
pkg_exposes=()

do_prepare() {
  export GOROOT="/hab/pkgs/core/go17/1.7.5/20170514000935"
  export GOPATH="$HOME/go"
  export PATH="$PATH:$GOROOT/bin"
}

do_build() {
  go install github.com/Masterminds/glide/...
  build_line "Here"
  ls /hab/pkgs/core/go17/1.7.5/20170514000935
}
