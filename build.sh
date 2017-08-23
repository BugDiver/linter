cleanBuild () {
    echo "removing bin dir...";
    rm -rf ./bin;
    echo "building linter in bin dir";
    go build -ldflags="-s -w" -o ./bin/linter
    wd=`pwd`
    echo "Binaries available at $wd/bin"
}

cleanBuild