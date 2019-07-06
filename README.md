[![Board Status](https://mchirico.visualstudio.com/c7e39c31-ad05-41c9-8b5b-22ac9a522c14/bde6ae0a-fc94-4b19-a70c-2389949f30e8/_apis/work/boardbadge/54cb834f-ff69-4ab9-b365-6625aa2a9081?columnOptions=1)](https://mchirico.visualstudio.com/c7e39c31-ad05-41c9-8b5b-22ac9a522c14/_boards/board/t/bde6ae0a-fc94-4b19-a70c-2389949f30e8/Microsoft.RequirementCategory/)

[![Build Status](https://mchirico.visualstudio.com/ipblock/_apis/build/status/mchirico.ipblock?branchName=master)](https://mchirico.visualstudio.com/ipblock/_build/latest?definitionId=9&branchName=master)



[![codecov](https://codecov.io/gh/mchirico/ipblock/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/ipblock)
# ipblock

## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


