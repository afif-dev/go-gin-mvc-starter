# Go GIN MVC Starter
![go gin mvc starter](https://raw.githubusercontent.com/afif-dev/go-gin-mvc-starter/c7566d77c903dfe8cfe5bc98009e5baa58bb0f7e/static/img/go-gin-mvc-starter.png)

Gin starter with MVC, ORM, template library, Bootstrap 5 and cronjob for Web and API. Build with package or module: GIN, GORM, GoDotEnv, Bootstrap 5, Goview, Cron.

For document check this url: https://gin-mvc-starter.herokuapp.com/

##

**Folders and files structure** 

https://gin-mvc-starter.herokuapp.com/structure

![go gin mvc starter - folders & files](https://raw.githubusercontent.com/afif-dev/go-gin-mvc-starter/main/static/img/go-gin-mvc-structure.jpg)

##

### Note

.env usage:
- rename .env-example to .env
- uncomment code in router/router.go & console/main_console.go
```
// # Uncomment this code to load .env file
// err := godotenv.Load()
// if err != nil {
// 	log.Fatal("Error loading .env file")
// }
```
