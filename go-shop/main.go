package main // Определяет, что это главный пакет программы

import ( // Начало блока импорта библиотек
	"go-shop/controllers" // Импортируем пакет с контроллерами
	"go-shop/models"      // Импортируем пакет с моделями

	"github.com/gin-gonic/gin" // Импортируем библиотеку Gin для работы с веб-сервером
)

func main() { // Главная функция программы
	r := gin.Default() // Создаем новый роутер с настройками по умолчанию
	models.InitDB()    // Инициализируем базу данных

	// Добавьте middleware для CORS
	r.Use(corsMiddleware()) // Добавляем middleware для обработки CORS

	r.POST("/register", controllers.Register)    // Определяем маршрут для регистрации пользователей
	r.POST("/login", controllers.Login)          // Определяем маршрут для логина пользователей
	r.GET("/users", controllers.GetAllUsers)     // Определяем маршрут для получения всех пользователей
	r.GET("/users/:id", controllers.GetUserByID) // Определяем маршрут для получения пользователя по ID

	r.Run(":8080") // Запускаем веб-сервер на порту 8080
}

func corsMiddleware() gin.HandlerFunc { // Функция, которая возвращает middleware для обработки CORS
	return func(c *gin.Context) { // Анонимная функция, обрабатывающая запросы
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                                                                                                            // Устанавливаем заголовок для разрешения всех источников
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")                                                                                                                    // Разрешаем использование учетных данных
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With") // Устанавливаем разрешенные заголовки
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")                                                                                                     // Устанавливаем разрешенные методы

		if c.Request.Method == "OPTIONS" { // Если метод запроса OPTIONS, значит это preflight запрос
			c.AbortWithStatus(204) // Прерываем обработку и возвращаем статус 204 (No Content)
			return                 // Выходим из функции
		}

		c.Next() // Продолжаем обработку запроса
	}
}
