pkg_origin=fenrirunbound
pkg_name=mplater
pkg_version=0.1.0
pkg_maintainer="Darren Matsumoto <aeneascorrupt@gmail.com>"
#pkg_source="https://github.com/FenrirUnbound/mplater"
pkg_upstream_url=https://github.com/FenrirUnbound/mplater
pkg_deps=()
pkg_build_deps=()
pkg_exports=()
pkg_bin_dirs=()
pkg_exposes=()

do_download() {
  wget -O $HAB_CACHE_SRC_PATH/mplater.tar.gz https://github.com/FenrirUnbound/mplater/releases/download/v0.0.7/mplater_0.0.7_Tux_64-bit.tar.gz
}

do_unpack() {
  tar -zxvf $HAB_CACHE_SRC_PATH/mplater.tar.gz -C $HAB_CACHE_SRC_PATH
}

do_build() {
  return 0
}

do_install() {
  mkdir -p $pkg_prefix/bin
  mv $HAB_CACHE_SRC_PATH/mplater $pkg_prefix/bin/mplater
  chmod +x $pkg_prefix/bin/mplater
}
