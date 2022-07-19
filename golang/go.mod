module github.com/nurrizkyimani/longlife_learning/golang

go 1.14

replace github.com/nurrizkyimani/longlife_learning/golang/utils => ./utils

replace github.com/nurrizkyimani/longlife_learning/golang/greetings => ./greetings

require (
	github.com/nurrizkyimani/longlife_learning/golang/greetings v0.0.0-00010101000000-000000000000
	github.com/nurrizkyimani/longlife_learning/golang/utils v0.0.0-00010101000000-000000000000
)

replace github.com/nurrizkyimani/longlife_learning/golang/model => ./model
