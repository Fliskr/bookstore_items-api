module github.com/gervi/bookstore_items-api/src

go 1.15

require (
	github.com/gervi/bookstore_oauth-go v0.0.0-00010101000000-000000000000
	github.com/gervi/bookstore_utils-go v0.0.0-20200406081703-a9b52b6e34e6
	github.com/gorilla/mux v1.8.0
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/olivere/elastic/v7 v7.0.18
	go.mongodb.org/mongo-driver v1.4.3
	go.uber.org/zap v1.16.0
)

replace github.com/gervi/bookstore_oauth-go => ../../bookstore_oauth-go

replace github.com/gervi/bookstore_utils-go => ../../bookstore_utils-go
