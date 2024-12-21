package v1

import "github.com/gin-gonic/gin"



// {
// 	"username": "string",
// 	"password": "string",
// 	"email": "string"
//   }

//   Ответ:
//   • Успех: ```201 Created```
//   ```json
//   {
// 	"message": "User registered successfully"
//   }

//   • Ошибка: ```400 Bad Request```
//   ```json
//   {
// 	"error": "Username or email already exists"
//   }

//   ▎Описание

//   Регистрация нового пользователя.
//   Проверяется уникальность имени пользователя и адреса электронной почты.
//   Пароль должен быть зашифрован перед сохранением в базе данных.

func (r *Router) register(ctx *gin.Context) {

}

// 	"username": "string",
// 	"password": "string"
// 	}

// 	Ответ:

// 	• Успех: ```200 OK```
// 	```json
// 	{
// 	  "token": "JWT_TOKEN"
// 	}

// 	• Ошибка: ```401 Unauthorized```
// 	```json
// 	{
// 	  "error": "Invalid username or password"
// 	}

// 	▎Описание

// 	Авторизация пользователя.
// 	При успешной авторизации возвращается JWT-токен, который будет использоваться для аутентификации последующих запросов.

func (r *Router) login(ctx *gin.Context) {

}
