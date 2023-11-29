module github.com/DataDog/datadog-agent/pkg/util/winutil

go 1.20

replace (
	github.com/DataDog/datadog-agent/pkg/util/log => ../log/
	github.com/DataDog/datadog-agent/pkg/util/scrubber => ../scrubber/
)

require (
	github.com/DataDog/datadog-agent/pkg/util/log v0.0.0-00010101000000-000000000000
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/fsnotify/fsnotify v1.7.0
	github.com/stretchr/testify v1.8.4
	go.uber.org/atomic v1.11.0
	golang.org/x/sys v0.14.0
)

require (
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.50.0-rc.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)