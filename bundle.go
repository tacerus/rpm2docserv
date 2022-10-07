package bundle

//go:generate sh -c "go run goembed.go -package bundled -var assets assets/chameleon/header.tmpl assets/chameleon/footer.tmpl assets/chameleon/style.css assets/chameleon/chameleon.css assets/chameleon/manpage.tmpl assets/chameleon/manpageerror.tmpl assets/chameleon/manpagefooterextra.tmpl assets/chameleon/contents.tmpl assets/chameleon/pkgindex.tmpl assets/chameleon/srcpkgindex.tmpl assets/chameleon/index.tmpl assets/chameleon/about.tmpl assets/chameleon/notfound.tmpl assets/chameleon/favicon.ico assets/chameleon/fallback-icon.svg > pkg/bundled/GENERATED_bundled.go"
