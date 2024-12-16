package models

type api struct {
	Status  int      `json:"status"`  // Код ответа
	Cash    bool     `json:"cash"`    // Кэш
	Message string   `json:"message"` // Сообщение
	Image   image    `json:"image"`   // Изображение
	Kebord  []kebord `json:"kebord"`  // Клавиатура
}

type kebord struct {
	Intline    bool   `json:"intline"`    // Инлайн
	ButtonName string `json:"buttonName"` // Название кнопки
	ButtonData string `json:"buttonData"` // Дата для кнопок INLINE !!!
}

type image struct {
	Url []string `json:"url"` // URL изображения
}
