//go:build licenses && generate

//go:generate /bin/sh -c "go-licenses save ./... --force --ignore=github.com/eriner/burr --ignore=github.com/eriner/burr/vendor --ignore=github.com/eriner/burr/licenses --save_path=./licenses_new || true"
//go:generate /bin/cp -R ./licenses_new/ licenses/
//go:generate /bin/rm -rf ./licenses_new
//go:generate /bin/sh -c "find ./licenses -type f -iname '*.go' -exec rm {} \\;"
//go:generate /bin/sh -c "go-licenses csv ./... --ignore=github.com/eriner/burr --ignore=github.com/eriner/burr/vendor --ignore=github.com/eriner/burr/licenses 2>/dev/null > ./licenses/licenses.csv"
package main
