module github.com/tilt-dev/tilt-trigger-on-ready

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/tilt-dev/tilt v0.20.2
	github.com/tilt-dev/wmclient v0.0.0-20201109174454-1839d0355fbc
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
)

replace (
	github.com/pkg/browser v0.0.0-00010101000000-000000000000 => github.com/pkg/browser v0.0.0-20210115035449-ce105d075bb4

	k8s.io/apimachinery => github.com/tilt-dev/apimachinery v0.20.2-tilt-20210505
)
