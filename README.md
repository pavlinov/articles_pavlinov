go mod init your_project_name
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/gorm/dialects/sqlite
go get -u github.com/dgrijalva/jwt-go

.
├── Makefile
├── README.md
├── articles_pavlinov
├── controllers
│   ├── article.go
│   ├── auth.go
│   └── preference.go
├── database
│   ├── setup.go
│   └── test.db
├── go.mod
├── go.sum
├── main.go
├── main_test.go
├── middleware
│   ├── auth.go
│   └── auth_middleware_test.go
├── models
│   ├── article.go
│   ├── preference.go
│   └── user.go
├── routes
│   └── routes.go
├── test.db
├── tests_curl.sh
└── utils
    ├── token.go
    └── token_test.go

curl -X POST http://localhost:8084/auth/login -H "Content-Type: application/json" -d '{"username": "user2", "password": "password"}'
curl -X POST http://localhost:8084/articles/ -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MTU3Nzk0MjZ9.swtdOehf4LoqTnOv1XwisNR7aBPDvRFq-DQ4ha11wE0" -H "Content-Type: application/json" -d '{"title": "Article 22", "content": "Content of Article 22"}'\n
curl -X POST http://localhost:8084/articles/ -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MTU3Nzk0MjZ9.swtdOehf4LoqTnOv1XwisNR7aBPDvRFq-DQ4ha11wE0" -H "Content-Type: application/json" -d '{"title": "Article 22", "content": "Content of Article 22"}'\n
curl -X POST http://localhost:8084/articles/ -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MTU3Nzk0MjZ9.swtdOehf4LoqTnOv1XwisNR7aBPDvRFq-DQ4ha11wE0" -H "Content-Type: application/json" -d '{"title": "Article 333", "content": "Content of Article 333"}'\n
curl -X POST http://localhost:8084/articles/ -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHAiOjE3MTU3Nzk0MjZ9.swtdOehf4LoqTnOv1XwisNR7aBPDvRFq-DQ4ha11wE0" -H "Content-Type: application/json" -d '{"title": "Article 333", "content": "Content of Article 333"}'\n
