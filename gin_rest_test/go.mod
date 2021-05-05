module github.com/getveryirichet/gin_rest_test

go 1.16

replace github.com/getveryrichet/gin_rest_test/controllers => ./controllers

replace github.com/getveryrichet/gin_rest_test/models => ./models

require (
	github.com/getveryrichet/gin_rest_test/controllers v0.0.0-00010101000000-000000000000
	github.com/getveryrichet/gin_rest_test/models v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.1
	github.com/jinzhu/gorm v1.9.16 // indirect
)
