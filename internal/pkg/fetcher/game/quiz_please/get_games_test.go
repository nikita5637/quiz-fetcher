package quiz_please

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	mocks "github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
)

const (
	json1 = `{
"gameId": 50069,
"nameGame": "Квиз, плиз!.jpeg SPB",
"max_players": 9,
"titleGame": "#1",
"numberGame": "#1",
"game_type_id": 1,
"blockData": "15 января",
"blockNumberIs": 169,
"blockOf": 230,
"place": "Grand Bianco",
"special": "",
"address": "Набережная канала Грибоедова, д. 36",
"place_description": "<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">Метро: Сенная площадь, Невский проспект</span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Сайт заведения: </span></span><span style=\"font-size: 11pt;\"><a style=\"text-decoration: none;\" href=\"https://grandbianco.ru/\"><span style=\"color: #1155cc;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #1155cc; background-color: #ffffff; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: underline; text-decoration-skip: none; vertical-align: baseline; white-space: pre-wrap;\">https://grandbianco.ru</span></span></a></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-weight: 400; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Есть в Питере одно местечко, которое является для нас самым родным. Это бар, в котором 2 июня 2017 года состоялась самая первая игра Квиз, плиз! в Санкт-Петербурге. Помимо этого Grand Bianco представляет из себя стильное мультиформатное пространство, расположенное в самом излюбленном (и, следовательно, самом красивом) месте для туристов. После игры, если идете к Невскому проспекту, можно устроить себе настоящую культурную прогулку вдоль канала Грибоедова, мимо банковского моста и величавого Казанского собора, прямиком в сердце города. Только перед этим обязательно насладитесь малиновым сидром бара Grand Bianco!</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">&nbsp;</span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">&nbsp;<span>Спэшл:&nbsp;</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- понедельник скидка 20% на десерты</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- вторник скидка 20% для женских компаний от 2-х человек</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- среда три кружки пенного по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- четверг винный безлимит за 900 рублей</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- пятница 3 шота по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- суббота 3 коктейля лонг-дринк по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- воскресенье скидка 20% на бутылку вина</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><strong id=\"docs-internal-guid-e4bc13cb-7fff-fbaa-1a23-528219620953\"><span style=\"font-size: 11pt; font-weight: normal;\">&nbsp;</span></strong></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-weight: 400; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Здесь можно посмотреть меню: <a style=\"color: #954f72;\" href=\"https://grandbianco.ru/#!/tab/422553998-2\">https://grandbianco.ru/#!/tab/422553998-2</a></span></span></p>",
"option": "Если вы желаете определенный стол, укажите это при регистрации в комментариях или позвоните нам.",
"menu": null,
"photos": [
	"files/2022%2F09%2F6324b6588628a.jpeg",
	"files/2022%2F09%2F6324b65896e20.jpeg",
	"files/2022%2F09%2F6324b658b54fc.jpeg",
	"files/2022%2F09%2F6324b658d478f.jpeg"
],
"time": "19:30",
"price": "400₽",
"text": "с человека, наличные или карта",
"payment_icon": 2,
"at": "в",
"description": "Вам просто необходима эта игра, если вы обожали, когда в школьных заданиях на логику попадались загадки, в журналах вы не пролистывали игры “найди 10 отличий”, а сегодня каламбурные мемы стали частью вашей жизни.",
"cityName": "Санкт-Петербург",
"status": 1,
"count": 61,
"is_teens": false,
"is_past": false,
"text_block": "<p><span style=\"font-size: 12pt; font-family: Calibri, sans-serif;\">В этой игре вам пригодится ассоциативное мышление, логика, насмотренность. Короче, включайте креатив на максимум!</span></p>\r\n<p><span style=\"font-size: 12pt; font-family: Calibri, sans-serif;\"><strong>Квиз, плиз!</strong> &mdash; это не только интеллектуальная игра, но и отличный повод для встречи с друзьями. Вы приходите командой, сидите за столом в баре, общаетесь, вкусно пьёте и едите, при этом отвечаете на интересные вопросы, серьёзные и не очень.</span></p>",
"custom_fields": false,
"imageData": null,
"free_status": 2,
"success": true,
"game_type": 0,
"covid_free": 0,
"no_covid": 1,
"map_type": "yandex",
"latitude": "59.931450000000000",
"longitude": "30.321463000000000",
"special_mobile_banner": null,
"datetime": "15.01.23 19:30",
"is_little_place": 0,
"link_to_bar": null,
"no_covid_general": ""
}`
	json2 = `{
"gameId": 50071,
"nameGame": "Квиз, плиз! SPB",
"max_players": 9,
"titleGame": "#608",
"numberGame": "#608",
"game_type_id": 1,
"blockData": "15 января",
"blockNumberIs": 145,
"blockOf": 230,
"place": "Дворец Олимпия",
"special": "",
"address": "Литейный пр., д. 14",
"place_description": "<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Метро: Чернышевская</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Сайт заведения:&nbsp;</span><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\"><a href=\"https://olympia-palace.ru/\"><span style=\"color: #1155cc;\">https://olympia-palace.ru</span></a></span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Дворец Олимпия может порадовать своей атмосферой девятнадцатого века: парадная лестница, лепнина, настенные росписи. Поэтому если вы любите путешествовать во времени так же, как любите бывать на играх Квиз, плиз!, это место точно вам понравится! А еще Дворец Олимпия может похвастаться блюдами смешанной кухни. Особые ценители могут откушать миндальное пирожное по рецепту личного повара самой княгини Долгорукой.&nbsp;<br />После игры вы можете неспешно прогуляться по изысканному району города, так как ресторан находится близко к станции метро Чернышевская, или уехать на своем автомобиле, рядом с рестораном расположена городская парковка.&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Спэшл:</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- десерт для именинников</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- пиво на гостевом кране 3 по цене 2-х</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- самое вкусное блюдо в меню по мнению шефа - миндальное пирожное</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Здесь можно посмотреть меню:&nbsp;</span><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\"><a href=\"http://dostavka.olympia-palace.ru/\"><span style=\"color: #1155cc;\">http://dostavka.olympia-palace.ru</span></a></span></p>",
"option": "Если вы желаете определенный стол, укажите это при регистрации в комментариях или позвоните нам.",
"menu": null,
"photos": [
"files/2022%2F09%2F632dc9346d4dc.jpg",
"files/2022%2F09%2F632dc93490acc.jpg",
"files/2022%2F09%2F632dc934ac332.jpg",
"files/2022%2F09%2F632dc934c63a6.jpg",
"files/2022%2F09%2F632dc934dfabf.jpg",
"files/2022%2F09%2F632dc9350f1bb.jpg"
],
"time": "19:30",
"price": "400₽",
"text": "с человека, наличные или карта",
"payment_icon": 2,
"at": "в",
"description": "Наша классическая игра – 7 раундов, более 39 вопросов, 2+ часа и море позитива. В команде может быть от 2 до 9 человек, игра открыта для всех: новых и опытных участников.",
"cityName": "Санкт-Петербург",
"status": 1,
"count": 85,
"is_teens": false,
"is_past": false,
"text_block": "<p class=\"p1\">Квиз, плиз! &mdash; это командная игра, отличный способ провести вечер с друзьями в уютном баре и получить новые позитивные впечатления.</p>\r\n<p class=\"p1\">К игре невозможно подготовиться, а вопросы могут быть на абсолютно разные темы: от покемонов до истории Древнего Рима. Более того, пригодиться может любая информация, которую вы встречали в своей жизни. Более 50 тысяч команд играет в Квиз, плиз! по всему миру, попробуйте и вы.</p>",
"custom_fields": false,
"imageData": null,
"free_status": 2,
"success": true,
"game_type": 0,
"covid_free": 0,
"no_covid": 1,
"map_type": "yandex",
"latitude": "59.945008000000000",
"longitude": "30.349022000000000",
"special_mobile_banner": null,
"datetime": "15.01.23 19:30",
"is_little_place": 0,
"link_to_bar": null,
"no_covid_general": ""
}`
	json3 = `{
"gameId": 50068,
"nameGame": "[music party] кринж эдишн SPB",
"max_players": 9,
"titleGame": "#1",
"numberGame": "#1",
"game_type_id": 9,
"blockData": "15 января",
"blockNumberIs": 173,
"blockOf": 190,
"place": "Puberty",
"special": "",
"address": "наб. Выборгская, д.47",
"place_description": "<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Метро: Выборгская</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span id=\"docs-internal-guid-7c8f09c6-7fff-1a78-6273-20c5061ad5b9\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; font-variant-ligatures: normal; font-variant-east-asian: normal; font-variant-position: normal; vertical-align: baseline; white-space: pre-wrap;\">Сайт заведения: </span><a style=\"text-decoration: none;\" href=\"http://puberty-spb.ru/\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #1155cc; font-variant-ligatures: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: underline; text-decoration-skip: none; vertical-align: baseline; white-space: pre-wrap;\">http://puberty-spb.ru/</span></a></span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm 0cm 12pt; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Отличный выбор для любителей классики не только на играх Квиз, плиз! Аппетитные блюда смешанной кухни и напитки на любой вкус, а самое главное - шикарное пенное собственного производства. Ресторан-пивоварня Puberty варят пиво по чешским технологиям, используя секреты самого Мартина Матушки. Рядом расположена бесплатная парковка, а еще из окна бара открывается чудесный вид на Выборгскую набережную.&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Спэшл:</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- скидка для именинников 20%</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- скидка на блюда навынос 15%</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- в&nbsp;заведении действует бонусная карта:</span><span style=\"font-size: 11pt; font-family: Roboto;\">&nbsp;</span><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">с кешбеком от 10% до 20%, оплачивать баллами можно до 50% от счета</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- самое вкусное блюдо в меню по мнению шефа &ndash; обязательно попробуйте все</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Здесь можно посмотреть меню:&nbsp;</span><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\"><a href=\"https://puberty-spb.ru/wp-content/uploads/2022/07/menyu-puberty.pdf\"><span style=\"color: #1155cc;\">https://puberty-spb.ru/wp-content/uploads/2022/07/menyu-puberty.pdf</span></a></span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm 0cm 12pt; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">&nbsp;</span></p>",
"option": "Если вы желаете определенный стол, укажите это при регистрации в комментариях или позвоните нам.",
"menu": null,
"photos": [
"files/2022%2F09%2F632dcba3f2984.jpg",
"files/2022%2F09%2F632dcba425afc.jpg",
"files/2022%2F09%2F632dcba461b5d.jpg",
"files/2022%2F09%2F632dcba4c1a25.jpg",
"files/2022%2F09%2F632dcba5019f5.jpg",
"files/2022%2F09%2F632dcba5399db.jpg"
],
"time": "16:00",
"price": "400₽",
"text": "с человека, наличные или карта",
"payment_icon": 2,
"at": "в",
"description": "Экспериментальная игра, где главный эксперимент проводится над участниками. Вам придется столкнуться с лавиной русскоязычной современной и хорошо забытой музыки прошлого, которую сложно слушать на полном серьезе. А иногда и просто сложно слушать дольше формата Тик-Тока. ",
"cityName": "Санкт-Петербург",
"status": 2,
"count": 17,
"is_teens": false,
"is_past": false,
"text_block": "<p><span style=\"font-size: 12pt; font-family: Calibri, sans-serif;\">Про такие треки говорят: &ldquo;Не знаю, откуда я это знаю&rdquo;. В игре будет все от Инстасамки и Олега Майами до лучших песен Богдана Титомира и Кая Метова&hellip;</span></p>\r\n<p><span style=\"font-size: 12pt; font-family: Calibri, sans-serif;\"><br /><strong>Квиз, плиз!</strong> &mdash; это не только интеллектуальная игра, но и отличный повод для встречи с друзьями. Вы приходите командой, сидите за столом в баре, общаетесь, вкусно пьёте и едите, при этом отвечаете на интересные вопросы, серьёзные и не очень.<br /></span></p>",
"custom_fields": false,
"imageData": "/files/2022%2F12%2F639affd62f8e3.png",
"free_status": 2,
"success": true,
"game_type": 0,
"covid_free": 0,
"no_covid": 1,
"map_type": "yandex",
"latitude": "59.973572000000000",
"longitude": "30.333764000000000",
"special_mobile_banner": "/storage/source/1/x_3hedbpdB2lqSqpMBBC3En86XrcOpcM.png",
"datetime": "15.01.23 16:00",
"is_little_place": 1,
"link_to_bar": null,
"no_covid_general": ""
}`
	json4 = `{
"gameId": 50502,
"nameGame": "Кино и музыка [стрим] SPB",
"max_players": 9,
"titleGame": "#70",
"numberGame": "#70",
"game_type_id": 5,
"blockData": "15 января",
"blockNumberIs": 25,
"blockOf": 500,
"place": "Везде, где есть интернет",
"special": "",
"address": "Невский проспект",
"place_description": "",
"option": "Если вы желаете определенный стол, укажите это при регистрации в комментариях или позвоните нам.",
"menu": null,
"photos": [],
"time": "20:00",
"price": "1000₽",
"text": "с команды, оплата онлайн",
"payment_icon": 3,
"at": "в",
"description": "Регистрируйтесь и подключайтесь к стриму помериться, кто посмотрел больше сериалов, услышал больше новых синглов и урвал максимальное количество бесплатных подписок. Подходит и новичкам, и матёрым игрокам, а в команде может быть сколько угодно человек: созванивайтесь, создавайте отдельный чат и побеждайте!",
"cityName": "Санкт-Петербург",
"status": 1,
"count": 470,
"is_teens": false,
"is_past": false,
"text_block": "<p>Кино и музыка [стрим] &mdash; это старые добрые игры [кино и музыка], но теперь онлайн. Вы сидите на кухне или в любимом кресле, подпеваете знакомым песням и смотрите отрывки из фильмов и сериалов со своей командой в Зуме или Скайпе, вкусно пьёте и едите (что успели урвать в магазине) и отвечаете на интересные вопросы, серьёзные и не очень.</p>\r\n<p>&nbsp;</p>\r\n<p>Стоимость участия &mdash; 1000 рублей с команды, ограничений по количеству игроков в команде нет.</p>\r\n<p>&nbsp;</p>\r\n<p><span style=\"font-family: Calibri, sans-serif; font-size: 16px;\">Играете впервые? Посмотрите наши ответы в&nbsp;</span><a style=\"font-family: Calibri, sans-serif; font-size: 16px; color: #0563c1;\" href=\"https://info.quizplease.ru/what_the_stream\">FAQ</a><span style=\"font-family: Calibri, sans-serif; font-size: 16px;\">, чтобы сразу врубиться, как все это работает.</span></p>",
"custom_fields": false,
"imageData": "/files/media_old/qp_files_30.png",
"free_status": 1,
"success": true,
"game_type": 1,
"covid_free": 0,
"no_covid": 0,
"map_type": "yandex",
"latitude": null,
"longitude": null,
"special_mobile_banner": "/storage/source/1/jbESqmonCGXaeg36hDnOWnCmh0jZqxN1.png",
"datetime": "15.01.23 20:00",
"is_little_place": 0,
"link_to_bar": null,
"sdk_shop_id": "755554",
"sdk_key": "live_NzU1NTU07gLSO5hNG4ORAWfm2xDZmiS3lLBoXLpF3JQ"
}`
	json5 = `{
"gameId": 50484,
"nameGame": "[18+] SPB",
"max_players": 9,
"titleGame": "#6",
"numberGame": "#6",
"game_type_id": 2,
"blockData": "18 января",
"blockNumberIs": 247,
"blockOf": 230,
"place": "Grand Bianco",
"special": "",
"address": "Набережная канала Грибоедова, д. 36",
"place_description": "<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">Метро: Сенная площадь, Невский проспект</span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Сайт заведения: </span></span><span style=\"font-size: 11pt;\"><a style=\"text-decoration: none;\" href=\"https://grandbianco.ru/\"><span style=\"color: #1155cc;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #1155cc; background-color: #ffffff; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: underline; text-decoration-skip: none; vertical-align: baseline; white-space: pre-wrap;\">https://grandbianco.ru</span></span></a></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-weight: 400; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Есть в Питере одно местечко, которое является для нас самым родным. Это бар, в котором 2 июня 2017 года состоялась самая первая игра Квиз, плиз! в Санкт-Петербурге. Помимо этого Grand Bianco представляет из себя стильное мультиформатное пространство, расположенное в самом излюбленном (и, следовательно, самом красивом) месте для туристов. После игры, если идете к Невскому проспекту, можно устроить себе настоящую культурную прогулку вдоль канала Грибоедова, мимо банковского моста и величавого Казанского собора, прямиком в сердце города. Только перед этим обязательно насладитесь малиновым сидром бара Grand Bianco!</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">&nbsp;</span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">&nbsp;<span>Спэшл:&nbsp;</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- понедельник скидка 20% на десерты</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- вторник скидка 20% для женских компаний от 2-х человек</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- среда три кружки пенного по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- четверг винный безлимит за 900 рублей</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- пятница 3 шота по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- суббота 3 коктейля лонг-дринк по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- воскресенье скидка 20% на бутылку вина</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><strong id=\"docs-internal-guid-e4bc13cb-7fff-fbaa-1a23-528219620953\"><span style=\"font-size: 11pt; font-weight: normal;\">&nbsp;</span></strong></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-weight: 400; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Здесь можно посмотреть меню: <a style=\"color: #954f72;\" href=\"https://grandbianco.ru/#!/tab/422553998-2\">https://grandbianco.ru/#!/tab/422553998-2</a></span></span></p>",
"option": "Если вы желаете определенный стол, укажите это при регистрации в комментариях или позвоните нам.",
"menu": null,
"photos": [
"files/2022%2F09%2F6324b6588628a.jpeg",
"files/2022%2F09%2F6324b65896e20.jpeg",
"files/2022%2F09%2F6324b658b54fc.jpeg",
"files/2022%2F09%2F6324b658d478f.jpeg"
],
"time": "19:30",
"price": "400₽",
"text": "с человека, наличные или карта",
"payment_icon": 2,
"at": "в",
"description": "Тематическая игра 18+ для тех, кто не стесняется. \r\n39+ вопросов, 2,5 часа времени и непередаваемая атмосфера.",
"cityName": "Санкт-Петербург",
"status": 2,
"count": -17,
"is_teens": false,
"is_past": false,
"text_block": "<p class=\"p1\">Квиз, плиз! &mdash; это командная игра, отличный способ провести вечер с друзьями в уютном баре и получить новые позитивные впечатления.</p>\r\n<p class=\"p1\">Эта игра 18+. Мы не будем стесняться задавать неуместные и дерзкие вопросы (хотя мы всегда задаем дерзкие вопросы). Возможно, будем немного ругаться.</p>\r\n<p class=\"p1\">Более 15 тысяч команд играет в Квиз, плиз! по всему миру, попробуйте и вы.</p>\r\n<p>&nbsp;</p>",
"custom_fields": false,
"imageData": "/files/2019/07/5d31b6c8ac7ad.png",
"free_status": 2,
"success": true,
"game_type": 0,
"covid_free": 0,
"no_covid": 1,
"map_type": "yandex",
"latitude": "59.931450000000000",
"longitude": "30.321463000000000",
"special_mobile_banner": "/storage/source/1/P6ffuM7MxAMQQbGkkS8J6wKxV-GncP9g.jpg",
"datetime": "18.01.23 19:30",
"is_little_place": 0,
"link_to_bar": null,
"no_covid_general": ""
}`
	json6 = `{
"gameId": 50486,
"nameGame": "[назад в 2010-е] SPB",
"max_players": 9,
"titleGame": "#2",
"numberGame": "#2",
"game_type_id": 2,
"blockData": "19 января",
"blockNumberIs": 89,
"blockOf": 400,
"place": "ЦИНЬ",
"special": "",
"address": "16 линия В.О дом 83",
"place_description": "<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Метро: Василеостровская</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Сайт заведения:&nbsp;</span><u><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif; color: #1155cc;\"><a href=\"https://topevent-tchin.ru/\"><span style=\"color: blue;\">https://topevent-tchin.ru</span></a></span></u></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><u><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif; color: #1155cc;\"><br /><br /></span></u></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Цинь &mdash; это уникальная концертная площадка с собственным рестораном, местечко, в котором прошло уже не одно Открытие сезона Квиз, плиз!<br />Лучшие китайские повара были приглашены в Санкт-Петербург, чтобы передать все тонкости вкусов и неповторимости сочетания ингредиентов китайской кухни. В ресторане можно насладиться традиционными китайскими блюдами, попробовать ту самую утку по-пекински, курицу в кисло-сладком соусе или необычный корень лотоса.&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 10.5pt; font-family: Verdana, sans-serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Спэшл:</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\">&nbsp;</p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- скидка для именинников 20%</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">- скидка 20% всем жителям Васильевского острова&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 10.5pt; font-family: Verdana, sans-serif;\">&nbsp;</span></p>\r\n<p class=\"MsoNormal\" style=\"margin: 0cm; font-size: medium; font-family: Calibri, sans-serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman', serif;\">Здесь можно посмотреть меню: <a href=\"https://drive.google.com/file/d/1f-h_e8GQipF5ANn8Nj7Vn2ynhxrRUwNg/view\">https://drive.google.com/file/d/1f-h_e8GQipF5ANn8Nj7Vn2ynhxrRUwNg/view</a></span></p>",
"option": "Если вы желаете определенный стол, укажите это при регистрации в комментариях или позвоните нам.",
"menu": null,
"photos": [
"files/2022%2F09%2F632dc779e6a36.jpeg",
"files/2022%2F09%2F632dc77a07d00.jpeg",
"files/2022%2F09%2F632dc77a1de3e.jpeg",
"files/2022%2F09%2F632dc77a2c79f.jpeg",
"files/2022%2F09%2F632dc77a3930c.jpeg",
"files/2022%2F09%2F632dc77a480ea.jpeg"
],
"time": "19:30",
"price": "400₽",
"text": "с человека, наличные или карта",
"payment_icon": 2,
"at": "в",
"description": "В новой тематической игре от Квиз, плиз! попробуем восстановить хронологию десятых годов этого века – века социальных сетей, Тик-Ток челленджей и мини-сериалов.",
"cityName": "Санкт-Петербург",
"status": 1,
"count": 311,
"is_teens": false,
"is_past": false,
"text_block": "<p><span style=\"font-size: 11pt; font-family: Calibri, sans-serif;\">Квиз, плиз! &mdash; это не только интеллектуальная игра, но и отличный повод для встречи с друзьями. Вы приходите командой, сидите за столом в баре, общаетесь, вкусно пьёте и едите, при этом отвечаете на интересные вопросы, серьёзные и не очень. Полюбившийся формат тематических игр, в которых вспоминаем самое яркое, что произошло в отдельно взятом десятилетии.</span></p>",
"custom_fields": false,
"imageData": "/files/2022%2F02%2F61f933d21156e.png",
"free_status": 2,
"success": true,
"game_type": 0,
"covid_free": 0,
"no_covid": 0,
"map_type": "yandex",
"latitude": "59.930793000000000",
"longitude": "30.231295000000000",
"special_mobile_banner": "/storage/source/1/4flW0fyxQkq6IcrzVOVHSirbE4pecF0Q.png",
"datetime": "19.01.23 19:30",
"is_little_place": 0,
"link_to_bar": null
}`
	json7 = `{
"gameId": 50495,
"nameGame": "[кино и музыка] SPB",
"max_players": 9,
"titleGame": "#103",
"numberGame": "#103",
"game_type_id": 5,
"blockData": "22 января",
"blockNumberIs": 240,
"blockOf": 230,
"place": "Grand Bianco",
"special": "",
"address": "Набережная канала Грибоедова, д. 36",
"place_description": "<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">Метро: Сенная площадь, Невский проспект</span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Сайт заведения: </span></span><span style=\"font-size: 11pt;\"><a style=\"text-decoration: none;\" href=\"https://grandbianco.ru/\"><span style=\"color: #1155cc;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #1155cc; background-color: #ffffff; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: underline; text-decoration-skip: none; vertical-align: baseline; white-space: pre-wrap;\">https://grandbianco.ru</span></span></a></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-weight: 400; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Есть в Питере одно местечко, которое является для нас самым родным. Это бар, в котором 2 июня 2017 года состоялась самая первая игра Квиз, плиз! в Санкт-Петербурге. Помимо этого Grand Bianco представляет из себя стильное мультиформатное пространство, расположенное в самом излюбленном (и, следовательно, самом красивом) месте для туристов. После игры, если идете к Невскому проспекту, можно устроить себе настоящую культурную прогулку вдоль канала Грибоедова, мимо банковского моста и величавого Казанского собора, прямиком в сердце города. Только перед этим обязательно насладитесь малиновым сидром бара Grand Bianco!</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">&nbsp;</span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">&nbsp;<span>Спэшл:&nbsp;</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- понедельник скидка 20% на десерты</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- вторник скидка 20% для женских компаний от 2-х человек</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- среда три кружки пенного по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- четверг винный безлимит за 900 рублей</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- пятница 3 шота по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- суббота 3 коктейля лонг-дринк по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- воскресенье скидка 20% на бутылку вина</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><strong id=\"docs-internal-guid-e4bc13cb-7fff-fbaa-1a23-528219620953\"><span style=\"font-size: 11pt; font-weight: normal;\">&nbsp;</span></strong></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-weight: 400; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Здесь можно посмотреть меню: <a style=\"color: #954f72;\" href=\"https://grandbianco.ru/#!/tab/422553998-2\">https://grandbianco.ru/#!/tab/422553998-2</a></span></span></p>",
"option": "Если вы желаете определенный стол, укажите это при регистрации в комментариях или позвоните нам.",
"menu": null,
"photos": [
"files/2022%2F09%2F6324b6588628a.jpeg",
"files/2022%2F09%2F6324b65896e20.jpeg",
"files/2022%2F09%2F6324b658b54fc.jpeg",
"files/2022%2F09%2F6324b658d478f.jpeg"
],
"time": "15:30",
"price": "400₽",
"text": "с человека, наличные или карта",
"payment_icon": 2,
"at": "в",
"description": "Игра нашего необычного формата: вопросы только про кино и музыку. Меломанам и кинолюбителям посвящается!",
"cityName": "Санкт-Петербург",
"status": 2,
"count": -10,
"is_teens": false,
"is_past": false,
"text_block": "<p class=\"p1\">Квиз, плиз! &mdash; это командная игра, отличный способ провести вечер с друзьями в уютном баре и получить новые позитивные впечатления.</p>\r\n<p class=\"p1\">Вопросы только про кино и музыку, могут пригодиться знания как оскаровских лауреатов, так и авторских фильмов.&nbsp;</p>\r\n<p class=\"p1\">Более 15 тысяч команд играет в Квиз, плиз! по всему миру, попробуйте и вы.</p>",
"custom_fields": false,
"imageData": "/files/media_old/qp_files_30.png",
"free_status": 2,
"success": true,
"game_type": 0,
"covid_free": 0,
"no_covid": 1,
"map_type": "yandex",
"latitude": "59.931450000000000",
"longitude": "30.321463000000000",
"special_mobile_banner": "/storage/source/1/tJ0Mc8AyU8K0Bj2v7y_PaFXg2TbH0MVR.jpg",
"datetime": "22.01.23 15:30",
"is_little_place": 0,
"link_to_bar": null,
"no_covid_general": ""
}`
	json8 = `{
"gameId": -1,
"nameGame": "[кино и музыка] SPB",
"max_players": 9,
"titleGame": "#103",
"numberGame": "#103",
"game_type_id": 5,
"blockData": "22 января",
"blockNumberIs": 240,
"blockOf": 230,
"place": "Grand Bianco",
"special": "",
"address": "Набережная канала Грибоедова, д. 36",
"place_description": "<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">Метро: Сенная площадь, Невский проспект</span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Сайт заведения: </span></span><span style=\"font-size: 11pt;\"><a style=\"text-decoration: none;\" href=\"https://grandbianco.ru/\"><span style=\"color: #1155cc;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #1155cc; background-color: #ffffff; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: underline; text-decoration-skip: none; vertical-align: baseline; white-space: pre-wrap;\">https://grandbianco.ru</span></span></a></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-weight: 400; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Есть в Питере одно местечко, которое является для нас самым родным. Это бар, в котором 2 июня 2017 года состоялась самая первая игра Квиз, плиз! в Санкт-Петербурге. Помимо этого Grand Bianco представляет из себя стильное мультиформатное пространство, расположенное в самом излюбленном (и, следовательно, самом красивом) месте для туристов. После игры, если идете к Невскому проспекту, можно устроить себе настоящую культурную прогулку вдоль канала Грибоедова, мимо банковского моста и величавого Казанского собора, прямиком в сердце города. Только перед этим обязательно насладитесь малиновым сидром бара Grand Bianco!</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">&nbsp;</span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt;\">&nbsp;<span>Спэшл:&nbsp;</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- понедельник скидка 20% на десерты</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- вторник скидка 20% для женских компаний от 2-х человек</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- среда три кружки пенного по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- четверг винный безлимит за 900 рублей</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- пятница 3 шота по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- суббота 3 коктейля лонг-дринк по цене 2-х</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-family: 'Times New Roman'; font-size: 11pt; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">- воскресенье скидка 20% на бутылку вина</span></span></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><strong id=\"docs-internal-guid-e4bc13cb-7fff-fbaa-1a23-528219620953\"><span style=\"font-size: 11pt; font-weight: normal;\">&nbsp;</span></strong></p>\r\n<p style=\"margin-right: 0cm; margin-left: 0cm; font-size: medium; font-family: 'Times New Roman', serif;\"><span style=\"font-size: 11pt; font-family: 'Times New Roman'; color: #000000; background-color: #ffffff; font-weight: 400; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-variant-east-asian: normal; font-variant-position: normal; text-decoration: none; vertical-align: baseline; white-space: pre-wrap;\"><span style=\"font-size: 11pt;\">Здесь можно посмотреть меню: <a style=\"color: #954f72;\" href=\"https://grandbianco.ru/#!/tab/422553998-2\">https://grandbianco.ru/#!/tab/422553998-2</a></span></span></p>",
"option": "Если вы желаете определенный стол, укажите это при регистрации в комментариях или позвоните нам.",
"menu": null,
"photos": [
"files/2022%2F09%2F6324b6588628a.jpeg",
"files/2022%2F09%2F6324b65896e20.jpeg",
"files/2022%2F09%2F6324b658b54fc.jpeg",
"files/2022%2F09%2F6324b658d478f.jpeg"
],
"time": "15:30",
"price": "invalid price",
"text": "с человека, наличные или карта",
"payment_icon": 2,
"at": "в",
"description": "Игра нашего необычного формата: вопросы только про кино и музыку. Меломанам и кинолюбителям посвящается!",
"cityName": "Санкт-Петербург",
"status": 2,
"count": -10,
"is_teens": false,
"is_past": false,
"text_block": "<p class=\"p1\">Квиз, плиз! &mdash; это командная игра, отличный способ провести вечер с друзьями в уютном баре и получить новые позитивные впечатления.</p>\r\n<p class=\"p1\">Вопросы только про кино и музыку, могут пригодиться знания как оскаровских лауреатов, так и авторских фильмов.&nbsp;</p>\r\n<p class=\"p1\">Более 15 тысяч команд играет в Квиз, плиз! по всему миру, попробуйте и вы.</p>",
"custom_fields": false,
"imageData": "/files/media_old/qp_files_30.png",
"free_status": 2,
"success": true,
"game_type": 0,
"covid_free": 0,
"no_covid": 1,
"map_type": "yandex",
"latitude": "59.931450000000000",
"longitude": "30.321463000000000",
"special_mobile_banner": "/storage/source/1/tJ0Mc8AyU8K0Bj2v7y_PaFXg2TbH0MVR.jpg",
"datetime": "22.01.23 15:30",
"is_little_place": 0,
"link_to_bar": null,
"no_covid_general": ""
}`
)

func TestGamesFetcher_getGames(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r *strings.Reader
			switch req.URL.Query()["id"][0] {
			case "50069":
				r = strings.NewReader(json1)
			case "50071":
				r = strings.NewReader(json2)
			case "50068":
				r = strings.NewReader(json3)
			case "50502":
				r = strings.NewReader(json4)
			case "50484":
				r = strings.NewReader(json5)
			case "50486":
				r = strings.NewReader(json6)
			case "50495":
				r = strings.NewReader(json7)
			case "-1":
				r = strings.NewReader(json8)
			case "-2":
				r = strings.NewReader("invalid json")
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fetcher := Fetcher{
			client:             *http.DefaultClient,
			gameInfoPathFormat: GameInfoPathFormat,
			placeStorage:       mockPlaceStorage,
			url:                svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Grand Bianco", "Набережная канала Грибоедова, д. 36").Times(3).Return(database.Place{
			ID:         7,
			ExternalID: 5,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Дворец Олимпия", "Литейный пр., д. 14").Once().Return(database.Place{
			ID:         5,
			ExternalID: 3,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Puberty", "наб. Выборгская, д.47").Once().Return(database.Place{
			ID:         6,
			ExternalID: 4,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Везде, где есть интернет", "Невский проспект").Once().Return(database.Place{
			ID:         8,
			ExternalID: 6,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "ЦИНЬ", "16 линия В.О дом 83").Once().Return(database.Place{}, errors.New("some error"))

		got := fetcher.getGames(ctx, []int64{
			50069,
			50071,
			50068,
			50502,
			50484,
			50486, // get place ID error
			50495,
			-1, // invalid price
			-2, // invalid json
		})

		loc, err := time.LoadLocation(time_utils.TimeZoneMoscow)
		assert.NoError(t, err)

		dt1, err := time.ParseInLocation(timeFormatString, "15.01.23 19:30", loc)
		assert.NoError(t, err)

		dt2, err := time.ParseInLocation(timeFormatString, "15.01.23 19:30", loc)
		assert.NoError(t, err)

		dt3, err := time.ParseInLocation(timeFormatString, "15.01.23 16:00", loc)
		assert.NoError(t, err)

		dt4, err := time.ParseInLocation(timeFormatString, "15.01.23 20:00", loc)
		assert.NoError(t, err)

		dt5, err := time.ParseInLocation(timeFormatString, "18.01.23 19:30", loc)
		assert.NoError(t, err)

		dt7, err := time.ParseInLocation(timeFormatString, "22.01.23 15:30", loc)
		assert.NoError(t, err)

		expect := []model.Game{
			{
				ExternalID:  50069,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#1",
				Name:        "Квиз, плиз!.jpeg",
				PlaceID:     5,
				DateTime:    dt1,
				Price:       400,
				PaymentType: "cash,card",
				MaxPlayers:  9,
			},
			{
				ExternalID:  50071,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#608",
				Name:        "Квиз, плиз!",
				PlaceID:     3,
				DateTime:    dt2,
				Price:       400,
				PaymentType: "cash,card",
				MaxPlayers:  9,
			},
			{
				ExternalID:  50068,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_THEMATIC_MOVIES_AND_MUSIC),
				Number:      "#1",
				Name:        "[music party] кринж эдишн",
				PlaceID:     4,
				DateTime:    dt3,
				Price:       400,
				PaymentType: "cash,card",
				MaxPlayers:  9,
			},
			{
				ExternalID:  50502,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_MOVIES_AND_MUSIC),
				Number:      "#70",
				Name:        "Кино и музыка [стрим]",
				PlaceID:     6,
				DateTime:    dt4,
				Price:       1000,
				PaymentType: "",
				MaxPlayers:  9,
			},
			{
				ExternalID:  50484,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_THEMATIC),
				Number:      "#6",
				Name:        "[18+]",
				PlaceID:     5,
				DateTime:    dt5,
				Price:       400,
				PaymentType: "cash,card",
				MaxPlayers:  9,
			},
			{
				ExternalID:  50495,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_MOVIES_AND_MUSIC),
				Number:      "#103",
				Name:        "[кино и музыка]",
				PlaceID:     5,
				DateTime:    dt7,
				Price:       400,
				PaymentType: "cash,card",
				MaxPlayers:  9,
			},
		}

		assert.ElementsMatch(t, expect, got)
	})
}

func TestGamesFetcher_getGameIDs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			r := strings.NewReader(html1)
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		fetcher := Fetcher{
			client:             *http.DefaultClient,
			gameInfoPathFormat: GameInfoPathFormat,
			url:                svr.URL,
		}
		got, err := fetcher.getGameIDs(ctx)
		assert.ElementsMatch(t, []int64{
			50482, 50483, 50484, 50485,
			50486, 50487, 50488, 50490,
			50852, 50491, 50492, 50494,
			50493, 50853, 50495, 50496,
			50497, 50498, 50887, 50499,
			50833, 50834, 50835, 50836,
			50837, 50841, 50838, 50839,
			50840, 50844, 50843, 50842,
			50845, 50846, 50847, 50849,
			50848,
		}, got)
		assert.NoError(t, err)

	})
}
