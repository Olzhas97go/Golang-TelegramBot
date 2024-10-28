package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var db *sql.DB

// QuizQuestion структура для хранения данных из базы
type QuizQuestion struct {
	Id       int
	Question string
	Answer   string
	Note     string
	ImageUrl string // Add this line
}

func main() {
	// Загружаем переменные окружения
	botToken := os.Getenv("TELEGRAM_TOKEN")

	connectDatabase()

	// Создаем инициализацию бота
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Получаем обновление
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Создаем обработчик бота
	bh, _ := th.NewBotHandler(bot, updates)

	// Остановка бота
	defer bh.Stop()

	// Остановка получение обновлении
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Отправка сообщения

		message := tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Привет %s! Нажми на кнопку Рандом", update.Message.From.FirstName),
		)

		keyboard := tu.Keyboard(tu.KeyboardRow(tu.KeyboardButton("Рандом")))

		_, _ = bot.SendMessage(message.WithReplyMarkup(keyboard))
	}, th.CommandEqual("start"))

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		// Отправка текста с инлайн клавиатурой
		if message.Text == "Рандом" {

			question, _ := getRandomQuestion()
			if err != nil {
				log.Printf("Error retrieving random question: %v", err)
				return
			}

			if question.ImageUrl != "" {
				// Send as photo with caption
				photo := tu.Photo(
					tu.ID(message.Chat.ID),
					tu.FileFromURL(question.ImageUrl),
				).WithCaption(question.Question)
	
				keyboard := tu.InlineKeyboard(
					tu.InlineKeyboardRow(
						tu.InlineKeyboardButton("Показать ответ").WithCallbackData("answer:" + strconv.Itoa(question.Id)),
					),
				)
				_, err = bot.SendPhoto(photo.WithReplyMarkup(keyboard))
				if err != nil {
					log.Printf("Error sending photo and inline keyboard: %v", err)
				}
			} else {
				// Send as a simple text message
				msg := tu.Message(
					tu.ID(message.Chat.ID),
					question.Question,
				)
	
				keyboard := tu.InlineKeyboard(
					tu.InlineKeyboardRow(
						tu.InlineKeyboardButton("Показать ответ").WithCallbackData("answer:" + strconv.Itoa(question.Id)),
					),
				)
				_, err = bot.SendMessage(msg.WithReplyMarkup(keyboard))
				if err != nil {
					log.Printf("Error sending message and inline keyboard: %v", err)
				}
			}
		}
		if message.Text == "Сәлем" {
			message := tu.Message(
				tu.ID(message.Chat.ID),
				"Сәлем. Қалың қалай?",
			)

			_, _ = bot.SendMessage(message)
		}

		if message.Text == "Жақсы" {
			message := tu.Message(
				tu.ID(message.Chat.ID),
				"Менде жақсымын",
			)

			_, _ = bot.SendMessage(message)
		}

		if message.Text == "Golang" || message.Text == "GO" {
			imageUrl := "https://habrastorage.org/r/w1560/getpro/habr/upload_files/dd2/c20/cd3/dd2c20cd39d84c6b374588e72c9eae27.png"

			photo := tu.Photo(
				tu.ID(message.Chat.ID),
				tu.FileFromURL(imageUrl),
			)

			_, _ = bot.SendPhoto(photo)
		}
	})

	// Обработка колбэк запросов
	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		if query.Data != "" {
			quizQuestionId := query.Data[7:]

			question, _ := getQuestionById(quizQuestionId)

			message := tu.Message(
				tu.ID(query.Message.GetChat().ID),
				question.Answer,
			)

			_, _ = bot.SendMessage(message)
		}
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataContains("answer:"))

	// Запуск обработчика
	bh.Start()
}

func connectDatabase() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
}

func getRandomQuestion() (QuizQuestion, error) {
	row := db.QueryRow("SELECT id, question, answer, note, image_url FROM quiz ORDER BY RAND() LIMIT 1")
	
	var question QuizQuestion
	err := row.Scan(&question.Id, &question.Question, &question.Answer, &question.Note, &question.ImageUrl)
	return question, err
}

func getQuestionById(quiz_id string) (QuizQuestion, error) {
	row := db.QueryRow("SELECT id, question, answer, note, image_url FROM quiz WHERE id = ?", quiz_id)
	
	var question QuizQuestion
	err := row.Scan(&question.Id, &question.Question, &question.Answer, &question.Note, &question.ImageUrl)
	return question, err
}
