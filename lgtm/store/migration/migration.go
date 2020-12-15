package migration

//go:generate go-bindata -pkg migration -o bindata.go sqlite3/ mysql/ postgres/
//go:generate go fmt bindata.go
//go:generate sed -i.bak "s/Sql/SQL/" bindata.go
//go:generate rm bindata.go.bak
