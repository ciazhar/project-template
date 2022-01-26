#mock
go get github.com/vektra/mockery/v2/.../

# linter
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.32.2
curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin vX.Y.Z
go get -u github.com/mgechev/revive
GO111MODULE=on go get -v -u github.com/go-critic/go-critic/cmd/gocritic
go get github.com/quasilyte/go-consistent

# migration
go get -v github.com/rubenv/sql-migrate/...

# pm2
npm install pm2 -g
