package models

// type api struct {
// 	Status  int      `json:"status"`  // Код ответа
// 	Cash    bool     `json:"cash"`    // Кэш
// 	Message string   `json:"message"` // Сообщение
// 	Image   image    `json:"image"`   // Изображение
// 	Kebord  []kebord `json:"kebord"`  // Клавиатура
// }

// type kebord struct {
// 	Intline    bool   `json:"intline"`    // Инлайн
// 	ButtonName string `json:"buttonName"` // Название кнопки
// 	ButtonData string `json:"buttonData"` // Дата для кнопок INLINE !!!
// }

// type image struct {
// 	Url []string `json:"url"` // URL изображения
// }

type ControllerResponce struct {
	Answer    string   `json:"answer"`
	Delay     int      `json:"delay"`
	Keyboard  Keyboard `json:"keyboard"`
	IsKb      bool     `json:"isKb"`
	IsNextMsg bool     `json:"isNextMsg"`
	Id        int      `json:"id"`
}

type Keyboard struct {
	Button []Button `json:"buttons"`
	Type   string   `json:"type"`
}

type Button struct {
	Caption string `json:"caption"`
	Data    string `json:"data"`
	Order   int    `json:"order"`
	Row     int    `json:"row"`
}

// {
// 	"answer": "Но для начала давай ты расскажешь немного о себе!",
// 	"delay": 2,
// 	"id": 5,
// 	"isKb": true,
// 	"isNextMsg": false,
// 	"keyboard": {
// 		"buttons": [
// 			{
// 				"caption": "Зарегистрироваться",
// 				"data": null,
// 				"order": 0,
// 				"row": 0
// 			}
// 		],
// 		"type": "reply"
// 	},
// 	"nextMsg": null,
// 	"nextState": null,
// 	"set_value": null,
// 	"set_variable": null,
// 	"state": "start"
// }
