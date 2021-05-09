# Backend - GO

## Instalar Dependencias
```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/bson/primitive
go get golang.org/x/crypto/bcrypt
go get github.com/gorilla/mux
go get github.com/rs/cors
go get github.com/dgrijalva/jwt-go
```

## Initialize a git repository in a new or existing directory
```bash
heroku git:remote -a bakendgolang
# Deploy your application
git add .
git commit -am "make it better"
git push heroku master
# Pagina web heroku 
https://bakendgolang.herokuapp.com/
# Revision Logs
heroku logs --tail
```