package sixty_seconds

const (
	html1 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Клуб «60 секунд»</title><link href="/static/apple-touch-icon.png?v=1.4.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.4.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.4.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.4.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.4.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.4.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/style.css?v=1.4.1" rel="stylesheet"></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light sticky-top bg-white"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1 logo-img" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Открытая лига</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/league/117/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/league/117/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига-19</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12" style="margin-top: -20px;"><div class="d-flex flex-row-reverse"><h4><a class="blue" href="None" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="blue" href="None" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-md-6 col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига</h2><div class="white-razdel"></div><div class="l-blue pt-2">Клуб «60 Секунд» / Санкт-Петербург</div></div></div></div><div class="col-md-6 col-12"><h5 class="yellow text-center">Ближайшая игра</h5><h6 class="l-grey text-center"><a href="/game/17630/">Открытая лига | Игра #8</a></h6><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="l-blue ff-ibc vertical-text align-middle"><span>Играем</span></div><div><table><tbody><tr><td class="l-grey"><i class="fas fa-calendar-alt fa-fw l-blue"></i>&nbsp;06 февраля, Понедельник</td></tr><tr><td class="l-grey"><i class="fas fa-clock fa-fw l-blue"></i>&nbsp;19:30</td></tr><tr><td class="l-grey"><i class="fas fa-map-marker-alt fa-fw l-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="l-grey"><i class="fas fa-credit-card l-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2302336/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0 mt-md-n4"><ul class="nav nav-pills justify-content-md-start justify-content-center" id="myTab" role="tablist"><li class="nav-item nav-pills-yellow"><a aria-controls="schedule" aria-selected="true" class="nav-link tab-link active" data-toggle="tab" href="#schedule" id="schedule-tab" role="tab"><span>Лига</span></a></li><li class="nav-item nav-pills-yellow"><a aria-controls="seasons" aria-selected="true" class="nav-link tab-link" data-toggle="tab" href="#seasons" id="seasons-tab" role="tab"><span>Все сезоны</span></a></li><li class="nav-item nav-pills-yellow"><a class="nav-link tab-link" href="/schedule/71/"><span>Расписание клуба</span></a></li></ul></div></div></div></div><div class="container"><div class="row pt-4"><div class="col-12"><div class="tab-content" id="myTabContent"><div aria-labelledby="schedule-tab" class="tab-pane fade show active" id="schedule" role="tabpanel"><div class="row"><div class="col-12"><h4 class="text-center dark-blue">Ближайшие игры лиги</h4></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/17630/">Открытая лига | Игра #8</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;06 февраля, Понедельник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2302336/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/17631/">Открытая лига | Игра #9</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;13 февраля, Понедельник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2302338/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/17632/">Открытая лига | Игра #10</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;20 февраля, Понедельник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2302339/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div></div><div class="col-12"><h5 class="dark-blue text-center">Рейтинг сезона <a href="/season/1916/">Зима-22/23</a></h5></div><div class=""><div class="container-fluid"><div class="row"><div class="col-12"><nav class="tabbable"><ul class="nav nav-pills justify-content-sm-start justify-content-center" id="myTab" role="tablist"><li class="nav-item text-center nav-pills-dark"><a aria-controls="1" aria-selected="true" class="nav-link active" data-toggle="tab" href="#h1" id="h1-tab" role="tab"><span>Открытая лига</span></a></li></ul></nav></div></div></div><div class="tab-content" id="myTabContent"><div aria-labelledby="h1-tab" class="tab-pane fade show active" id="h1" role="tabpanel"><div class="table-responsive pt-2 pb-3"><table class="table table-hover table-sm table-nonfluid borderless-top" id="rate_table"><thead><tr><th class="text-center fit-strickt">&nbsp;&nbsp;#&nbsp;&nbsp;</th><th class="text-left fit-strickt">Название</th><th class="text-center fit-strickt">&nbsp;&nbsp;&nbsp;Сумма&nbsp;&nbsp;&nbsp;</th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17042/">28.11.22</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17170/">05.12.22</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17171/">12.12.22</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17626/">09.01.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17627/">16.01.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17628/">23.01.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17629/">30.01.23</a></th></tr></thead><tbody id="myTable"><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>1</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/32076/">Улица плохих снов</a></td><td class="bg-blue-dark text-white text-center fit-strickt">214.50</td><td class="text-center text-secondary"><span>26.00</span></td><td class="text-center text-secondary"><span>27.50</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>28.00</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>28.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>2</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/30782/">О, Спорт!</a></td><td class="bg-blue-dark text-white text-center fit-strickt">208.50</td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>30.00</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>27.00</span></td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center text-secondary"><span>24.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>3</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/19461/">«П» в кубе</a></td><td class="bg-blue-dark text-white text-center fit-strickt">163.00</td><td class="text-center text-secondary"><span>18.00</span></td><td class="text-center text-secondary"><span>27.50</span></td><td class="text-center text-secondary"><span>28.00</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>15.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>4</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/32077/">ПИКсики</a> <sup class="dark-blue">+1</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">146.00</td><td class="text-center text-secondary"><span>27.50</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>31.00</span></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>27.50</span></td><td class="text-center text-secondary"><span>35.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>5</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/30529/">Тигры Разума</a> <sup class="dark-blue">+1</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">132.00</td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>20.50</span></td><td class="text-center text-secondary"><span>26.50</span></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>19.50</span></td><td class="text-center text-secondary"><span>18.50</span></td><td class="text-center text-secondary"><span>22.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>6</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/7099/">Почему бы и нет</a> <sup class="yellow">-2</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">123.50</td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>23.50</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>22.00</span></td><td class="text-center text-secondary"><span>21.00</span></td><td class="text-center text-secondary"><span>20.00</span></td><td class="text-center text-secondary"><span>12.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>7</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/32079/">Мычание ягнят</a></td><td class="bg-blue-dark text-white text-center fit-strickt">118.50</td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>30.00</span></td><td class="text-center text-secondary"><span>32.00</span></td></tr></tbody></table></div></div></div></div><div class="row"><div class="col-12 pt-2"><h4 class="dark-blue text-center">Статистика</h4><p><small>Ниже представлены топ-5 команд в некоторых номинациях. Обратите внимание, что это общая статистика по всем играм. Кроме общей статистики, есть расширенная статистика для каждого из сезонов. Например, вот статистика <a class="table_a" href="/season/1916/#stats">текущего</a> и <a class="table_a" href="/season/740/#stats">предыдущего</a> сезонов. Кликнув на любую из команд лиги - вы попадете на страницу этой команды, где сможете ознакомится с её подробной статистикой, как за всё время, так и за каждый сезон в отдельности. Команды можно искать с помощью поиска, а также из списка всех <a class="table_a" href="/teams/117/">команд лиги</a>.</small></p></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Больше всех сыграли</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5099/">Газ до отказа</a></td><td class="text-center">87</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/10164/">Мамонт на Невском</a></td><td class="text-center">81</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Побед в сезонах</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td><td class="text-center">1</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">1</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6380/">Статистически незначимые</a></td><td class="text-center">25</td><td class="text-center">1</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/32076/">Улица плохих снов</a></td><td class="text-center">7</td><td class="text-center">1</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td><td class="text-center">0</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Победители</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td><td class="text-center">12</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">7</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/32076/">Улица плохих снов</a></td><td class="text-center">7</td><td class="text-center">3</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">3</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">3</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Призёры</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Пьедесталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">23</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td><td class="text-center">21</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">10</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5089/">Стыд и скрам</a></td><td class="text-center">61</td><td class="text-center">9</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td><td class="text-center">9</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучший процент взятия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Процент взятия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">72.5 %</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">68.7 %</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5086/">Дискриминант добра</a></td><td class="text-center">23</td><td class="text-center">66.4 %</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5074/">Квантовый мост</a></td><td class="text-center">43</td><td class="text-center">66.2 %</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5089/">Стыд и скрам</a></td><td class="text-center">61</td><td class="text-center">64.9 %</td></tr></tbody></table></div><small><sup>*</sup> - среди команд сыгравших больше 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучшая серия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">19</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td><td class="text-center">18</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6392/">Бред Стайлс</a></td><td class="text-center">22</td><td class="text-center">16</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/8040/">Борьба за безнравственность</a></td><td class="text-center">24</td><td class="text-center">16</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">16</td></tr></tbody></table></div><small><sup>*</sup> - количество правильных ответов подряд</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">лучшая худшая серия<sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">5</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/6375/">Хунвейбины за 40</a></td><td class="text-center">38</td><td class="text-center">6</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6380/">Статистически незначимые</a></td><td class="text-center">25</td><td class="text-center">6</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/6377/">Команда номер 8</a></td><td class="text-center">19</td><td class="text-center">6</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5074/">Квантовый мост</a></td><td class="text-center">43</td><td class="text-center">7</td></tr></tbody></table></div><small><sup>*</sup> - количество неправильных ответов подряд, среди команд сыгравших минимум 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Тоталов в турах <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Тоталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5089/">Стыд и скрам</a></td><td class="text-center">61</td><td class="text-center">6</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">5</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">3</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td><td class="text-center">2</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5074/">Квантовый мост</a></td><td class="text-center">43</td><td class="text-center">2</td></tr></tbody></table></div><small><sup>*</sup> - тотал - это тур в котором, команда ответила верно на все вопросы</small></div></div></div><div aria-labelledby="seasons-tab" class="tab-pane fade" id="seasons" role="tabpanel"><div class=""><h4 class="text-center dark-blue">Сезоны</h4><div class="table-responsive pb-4"><table class="table table-hover table-sm"><thead><tr class="active"><th class="text-nowrap">Сезон</th><th class="text-nowrap text-center">Игр</th><th class="text-nowrap">Начало</th><th class="text-nowrap">Окончание</th></tr></thead><tbody><tr><td class="text-nowrap"><a class="table_a" href="/season/1916/">Зима-22/23</a></td><td class="text-nowrap text-center">11</td><td class="text-nowrap">28 Ноябрь 2022</td><td class="text-nowrap">27 Февраль 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/740/">Весна 2020</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">09 Март 2020</td><td class="text-nowrap">25 Май 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/613/">Зима 2020</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">02 Декабрь 2019</td><td class="text-nowrap">25 Март 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/379/">Синхроны</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">09 Апрель 2019</td><td class="text-nowrap">23 Декабрь 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/532/">Осень 2019</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">09 Сентябрь 2019</td><td class="text-nowrap">25 Ноябрь 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/363/">Весна-лето 2019</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">06 Май 2019</td><td class="text-nowrap">26 Август 2019</td></tr></tbody></table></div></div></div></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12 my-n3"><div class="d-flex flex-row-reverse pt-1"><h4><a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Захарова Л. А. | <a class="foot-link" href="https://elba.kontur.ru/card/2b1oxlljoa" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.4.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.4.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.4.1"></script><script src="/static/js/popper.min.js?v=1.4.1"></script><script src="/static/js/bootstrap.min.js?v=1.4.1"></script><script src="/static/js/corner-popup.js?v=1.4.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.4.1"></script><script src="/static/js/alertify.min.js?v=1.4.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode('Санкт-Петербург'));
      $('#league_name').text(htmlDecode('Открытая лига'));
      $('.modal-city').text(htmlDecode('Санкт-Петербург'));
      
      set_city(71, 'Санкт-Петербург');
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'U05XpSrOvpk8wyROZnPJ3SrdL3LZqjQQ8YtCtjd52M4DRSURcP1ZVR7BCVg6FKe8',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><script type="text/javascript">
	
	$('a[data-toggle="tab"]').on('shown.bs.tab', function (e) {
		var hash = $(e.target).attr('href');
		if (history.pushState) {
			history.pushState(null, null, hash);
		} else {
			location.hash = hash;
		}
	});

	jQuery(document).ready(function ($) {
		let selectedTab = window.location.hash;
		$('.nav-link[href="' + selectedTab + '"]' ).trigger('click');
	})
</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html2 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Клуб «60 секунд»</title><link href="/static/apple-touch-icon.png?v=1.4.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.4.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.4.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.4.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.4.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.4.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/style.css?v=1.4.1" rel="stylesheet"></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light sticky-top bg-white"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1 logo-img" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Открытая лига</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/league/117/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/league/117/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига-19</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12" style="margin-top: -20px;"><div class="d-flex flex-row-reverse"><h4><a class="blue" href="None" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="blue" href="None" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-md-6 col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига</h2><div class="white-razdel"></div><div class="l-blue pt-2">Клуб «60 Секунд» / Санкт-Петербург</div></div></div></div><div class="col-md-6 col-12"><h5 class="yellow text-center">Ближайшая игра</h5><h6 class="l-grey text-center"><a href="/game/17631/">Открытая лига | Игра #9</a></h6><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="l-blue ff-ibc vertical-text align-middle"><span>Играем</span></div><div><table><tbody><tr><td class="l-grey"><i class="fas fa-calendar-alt fa-fw l-blue"></i>&nbsp;13 февраля, Понедельник</td></tr><tr><td class="l-grey"><i class="fas fa-clock fa-fw l-blue"></i>&nbsp;19:30</td></tr><tr><td class="l-grey"><i class="fas fa-map-marker-alt fa-fw l-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="l-grey"><i class="fas fa-credit-card l-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2302338/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0 mt-md-n4"><ul class="nav nav-pills justify-content-md-start justify-content-center" id="myTab" role="tablist"><li class="nav-item nav-pills-yellow"><a aria-controls="schedule" aria-selected="true" class="nav-link tab-link active" data-toggle="tab" href="#schedule" id="schedule-tab" role="tab"><span>Лига</span></a></li><li class="nav-item nav-pills-yellow"><a aria-controls="seasons" aria-selected="true" class="nav-link tab-link" data-toggle="tab" href="#seasons" id="seasons-tab" role="tab"><span>Все сезоны</span></a></li><li class="nav-item nav-pills-yellow"><a class="nav-link tab-link" href="/schedule/71/"><span>Расписание клуба</span></a></li></ul></div></div></div></div><div class="container"><div class="row pt-4"><div class="col-12"><div class="tab-content" id="myTabContent"><div aria-labelledby="schedule-tab" class="tab-pane fade show active" id="schedule" role="tabpanel"><div class="row"><div class="col-12"><h4 class="text-center dark-blue">Ближайшие игры лиги</h4></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/17631/">Открытая лига | Игра #9</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;13 февраля, Понедельник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2302338/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/17632/">Открытая лига | Игра #10</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;20 февраля, Понедельник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2302339/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/17633/">Открытая лига | Финал</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;27 февраля, Понедельник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2302340/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div></div><div class="col-12"><h5 class="dark-blue text-center">Рейтинг сезона <a href="/season/1916/">Зима-22/23</a></h5></div><div class=""><div class="container-fluid"><div class="row"><div class="col-12"><nav class="tabbable"><ul class="nav nav-pills justify-content-sm-start justify-content-center" id="myTab" role="tablist"><li class="nav-item text-center nav-pills-dark"><a aria-controls="1" aria-selected="true" class="nav-link active" data-toggle="tab" href="#h1" id="h1-tab" role="tab"><span>Открытая лига</span></a></li></ul></nav></div></div></div><div class="tab-content" id="myTabContent"><div aria-labelledby="h1-tab" class="tab-pane fade show active" id="h1" role="tabpanel"><div class="table-responsive pt-2 pb-3"><table class="table table-hover table-sm table-nonfluid borderless-top" id="rate_table"><thead><tr><th class="text-center fit-strickt">&nbsp;&nbsp;#&nbsp;&nbsp;</th><th class="text-left fit-strickt">Название</th><th class="text-center fit-strickt">&nbsp;&nbsp;&nbsp;Сумма&nbsp;&nbsp;&nbsp;</th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17042/">28.11.22</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17170/">05.12.22</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17171/">12.12.22</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17626/">09.01.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17627/">16.01.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17628/">23.01.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17629/">30.01.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/17630/">06.02.23</a></th></tr></thead><tbody id="myTable"><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>1</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/32076/">Улица плохих снов</a></td><td class="bg-blue-dark text-white text-center fit-strickt">248.00</td><td class="text-center text-secondary"><span>26.00</span></td><td class="text-center text-secondary"><span>27.50</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>28.00</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>28.00</span></td><td class="text-center text-secondary"><span>33.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>2</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/30782/">О, Спорт!</a></td><td class="bg-blue-dark text-white text-center fit-strickt">230.50</td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>30.00</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>27.00</span></td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>22.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>3</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/19461/">«П» в кубе</a></td><td class="bg-blue-dark text-white text-center fit-strickt">189.50</td><td class="text-center text-secondary"><span>18.00</span></td><td class="text-center text-secondary"><span>27.50</span></td><td class="text-center text-secondary"><span>28.00</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>15.00</span></td><td class="text-center text-secondary"><span>26.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>4</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/30529/">Тигры Разума</a></td><td class="bg-blue-dark text-white text-center fit-strickt">155.00</td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>20.50</span></td><td class="text-center text-secondary"><span>26.50</span></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>19.50</span></td><td class="text-center text-secondary"><span>18.50</span></td><td class="text-center text-secondary"><span>22.50</span></td><td class="text-center text-secondary"><span>23.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>5</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/33145/">Трудно плыть боком</a> <sup class="dark-blue">+2</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">150.50</td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>27.50</span></td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center text-secondary"><span>27.50</span></td><td class="text-center text-secondary"><span>30.00</span></td><td class="text-center text-secondary"><span>33.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>6</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/32079/">Мычание ягнят</a></td><td class="bg-blue-dark text-white text-center fit-strickt">148.50</td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>30.00</span></td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center text-secondary"><span>30.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>7</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/7099/">Почему бы и нет</a> <sup class="yellow">-2</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">148.00</td><td class="text-center"><i aria-hidden="true" class="fas fa-minus minus"></i></td><td class="text-center text-secondary"><span>23.50</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>22.00</span></td><td class="text-center text-secondary"><span>21.00</span></td><td class="text-center text-secondary"><span>20.00</span></td><td class="text-center text-secondary"><span>12.00</span></td><td class="text-center text-secondary"><span>24.50</span></td></tr></tbody></table></div></div></div></div><div class="row"><div class="col-12 pt-2"><h4 class="dark-blue text-center">Статистика</h4><p><small>Ниже представлены топ-5 команд в некоторых номинациях. Обратите внимание, что это общая статистика по всем играм. Кроме общей статистики, есть расширенная статистика для каждого из сезонов. Например, вот статистика <a class="table_a" href="/season/1916/#stats">текущего</a> и <a class="table_a" href="/season/740/#stats">предыдущего</a> сезонов. Кликнув на любую из команд лиги - вы попадете на страницу этой команды, где сможете ознакомится с её подробной статистикой, как за всё время, так и за каждый сезон в отдельности. Команды можно искать с помощью поиска, а также из списка всех <a class="table_a" href="/teams/117/">команд лиги</a>.</small></p></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Больше всех сыграли</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5099/">Газ до отказа</a></td><td class="text-center">87</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/10164/">Мамонт на Невском</a></td><td class="text-center">81</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Побед в сезонах</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td><td class="text-center">1</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">1</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6380/">Статистически незначимые</a></td><td class="text-center">25</td><td class="text-center">1</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/32076/">Улица плохих снов</a></td><td class="text-center">8</td><td class="text-center">1</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td><td class="text-center">0</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Победители</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td><td class="text-center">12</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">7</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/32076/">Улица плохих снов</a></td><td class="text-center">8</td><td class="text-center">4</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">3</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">3</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Призёры</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Пьедесталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">23</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td><td class="text-center">21</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">10</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5089/">Стыд и скрам</a></td><td class="text-center">61</td><td class="text-center">9</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td><td class="text-center">9</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучший процент взятия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Процент взятия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">72.5 %</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">68.7 %</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5086/">Дискриминант добра</a></td><td class="text-center">23</td><td class="text-center">66.4 %</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5074/">Квантовый мост</a></td><td class="text-center">43</td><td class="text-center">66.2 %</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5089/">Стыд и скрам</a></td><td class="text-center">61</td><td class="text-center">64.9 %</td></tr></tbody></table></div><small><sup>*</sup> - среди команд сыгравших больше 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучшая серия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">19</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td><td class="text-center">18</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6392/">Бред Стайлс</a></td><td class="text-center">22</td><td class="text-center">16</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/8040/">Борьба за безнравственность</a></td><td class="text-center">24</td><td class="text-center">16</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">16</td></tr></tbody></table></div><small><sup>*</sup> - количество правильных ответов подряд</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">лучшая худшая серия<sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">5</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/6375/">Хунвейбины за 40</a></td><td class="text-center">38</td><td class="text-center">6</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6380/">Статистически незначимые</a></td><td class="text-center">25</td><td class="text-center">6</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/6377/">Команда номер 8</a></td><td class="text-center">19</td><td class="text-center">6</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5074/">Квантовый мост</a></td><td class="text-center">43</td><td class="text-center">7</td></tr></tbody></table></div><small><sup>*</sup> - количество неправильных ответов подряд, среди команд сыгравших минимум 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Тоталов в турах <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Тоталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5089/">Стыд и скрам</a></td><td class="text-center">61</td><td class="text-center">6</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">5</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">3</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">107</td><td class="text-center">2</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5074/">Квантовый мост</a></td><td class="text-center">43</td><td class="text-center">2</td></tr></tbody></table></div><small><sup>*</sup> - тотал - это тур в котором, команда ответила верно на все вопросы</small></div></div></div><div aria-labelledby="seasons-tab" class="tab-pane fade" id="seasons" role="tabpanel"><div class=""><h4 class="text-center dark-blue">Сезоны</h4><div class="table-responsive pb-4"><table class="table table-hover table-sm"><thead><tr class="active"><th class="text-nowrap">Сезон</th><th class="text-nowrap text-center">Игр</th><th class="text-nowrap">Начало</th><th class="text-nowrap">Окончание</th></tr></thead><tbody><tr><td class="text-nowrap"><a class="table_a" href="/season/1916/">Зима-22/23</a></td><td class="text-nowrap text-center">11</td><td class="text-nowrap">28 Ноябрь 2022</td><td class="text-nowrap">27 Февраль 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/740/">Весна 2020</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">09 Март 2020</td><td class="text-nowrap">25 Май 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/613/">Зима 2020</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">02 Декабрь 2019</td><td class="text-nowrap">25 Март 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/379/">Синхроны</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">09 Апрель 2019</td><td class="text-nowrap">23 Декабрь 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/532/">Осень 2019</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">09 Сентябрь 2019</td><td class="text-nowrap">25 Ноябрь 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/363/">Весна-лето 2019</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">06 Май 2019</td><td class="text-nowrap">26 Август 2019</td></tr></tbody></table></div></div></div></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12 my-n3"><div class="d-flex flex-row-reverse pt-1"><h4><a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Захарова Л. А. | <a class="foot-link" href="https://elba.kontur.ru/card/2b1oxlljoa" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.4.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.4.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.4.1"></script><script src="/static/js/popper.min.js?v=1.4.1"></script><script src="/static/js/bootstrap.min.js?v=1.4.1"></script><script src="/static/js/corner-popup.js?v=1.4.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.4.1"></script><script src="/static/js/alertify.min.js?v=1.4.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode('Санкт-Петербург'));
      $('#league_name').text(htmlDecode('Открытая лига'));
      $('.modal-city').text(htmlDecode('Санкт-Петербург'));
      
      set_city(71, 'Санкт-Петербург');
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'FGNypSfLDyNbciZDOby48IHJiclpym6xTEbdtj12aVxGxC2G1DKk0Hn794QwNNuP',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><script type="text/javascript">
	
	$('a[data-toggle="tab"]').on('shown.bs.tab', function (e) {
		var hash = $(e.target).attr('href');
		if (history.pushState) {
			history.pushState(null, null, hash);
		} else {
			location.hash = hash;
		}
	});

	jQuery(document).ready(function ($) {
		let selectedTab = window.location.hash;
		$('.nav-link[href="' + selectedTab + '"]' ).trigger('click');
	})
</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"></body></html>`
	html17630 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига | Игра #8 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-22/23 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Открытая лига | Игра #8 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-22/23 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.4.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.4.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.4.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.4.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.4.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.4.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/style.css?v=1.4.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Открытая лига</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/17630/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/17630/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига-19</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига | Игра #8</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/117/">Открытая лига</a> / <a href="/season/1916/">Зима-22/23</a> / 6 февраля 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/17630/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/17630/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"><div id="table_1000"><h3 class="dark-blue text-center">Итоговая таблица</h3><div class="table-responsive"><table class="table table-hover table-sm text-center table-nonfluid mb-3 borderless-top tablesorter tablesorter-default tablesorter0146d4bb2e7088" id="jsonTable_1000" role="grid"><thead id="jsonTableHead_1000"><tr class="text-blue tablesorter-headerRow" role="row"><th class="text-center fit-half bg-white maingame tablesorter-header tablesorter-headerAsc" data-column="0" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="ascending" aria-label=": Ascending sort applied, activate to apply a descending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i class="fa-solid fa-medal"></i></div></th><th class="text-left tablesorter-header tablesorter-headerUnSorted" data-column="1" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="Название: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Название</div></th><th class="text-center fit maingame tablesorter-header tablesorter-headerAsc" data-column="2" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="ascending" aria-label="Сумма: Ascending sort applied, activate to apply a descending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Сумма</div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="3" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="1: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>1</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="4" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="2: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>2</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="5" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="3: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>3</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="6" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="4: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>4</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="7" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="5: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>5</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="8" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="6: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>6</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="9" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="7: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>7</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="10" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="8: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>8</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="11" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="9: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>9</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="12" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="10: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>10</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="13" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="11: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>11</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="14" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="12: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>12</small></div></th><th class="text-center fit-half main_a tablesorter-header tablesorter-headerUnSorted" data-column="15" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="1: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i onclick="toogletur(1)" class="table_blue_a far fa-plus-square tablesorter-noSort 1_but"></i>1</div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="16" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="13: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>13</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="17" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="14: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>14</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="18" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="15: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>15</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="19" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="16: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>16</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="20" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="17: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>17</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="21" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="18: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>18</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="22" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="19: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>19</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="23" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="20: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>20</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="24" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="21: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>21</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="25" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="22: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>22</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="26" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="23: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>23</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="27" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="24: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>24</small></div></th><th class="text-center fit-half main_a tablesorter-header tablesorter-headerUnSorted" data-column="28" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="2: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i onclick="toogletur(2)" class="table_blue_a far fa-plus-square tablesorter-noSort 2_but"></i>2</div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="29" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="25: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>25</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="30" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="26: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>26</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="31" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="27: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>27</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="32" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="28: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>28</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="33" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="29: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>29</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="34" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="30: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>30</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="35" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="31: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>31</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="36" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="32: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>32</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="37" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="33: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>33</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="38" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="34: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>34</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="39" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="35: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>35</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="40" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="36: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>36</small></div></th><th class="text-center fit-half main_a tablesorter-header tablesorter-headerUnSorted" data-column="41" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="3: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i onclick="toogletur(3)" class="table_blue_a far fa-plus-square tablesorter-noSort 3_but"></i>3</div></th><th class="text-center fit-half matrixgame tablesorter-header tablesorter-headerUnSorted" data-column="42" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label=": No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i class="fa-solid fa-medal"></i></div></th><th class="text-center fit matrixgame tablesorter-header tablesorter-headerUnSorted" data-column="43" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="Матрица: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Матрица</div></th></tr></thead><tbody class="tBody" id="jsonTableBody_1000" aria-live="polite" aria-relevant="all"><tr class="div_0 true" id="253264" role="row"><td class="bg-blue-dark text-white fit-half maingame">1</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/32076/" data-original-title="Улица плохих снов">Улица плохих снов</a></td><td class="bg-blue-dark text-white total maingame">29<sub class="yellow qrs"> 120</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>9</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>10</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>10</strong></td><td class="bg-blue-light text-white fit-half matrixgame">1</td><td class="bg-blue-light text-white matrix-total fit matrixgame">560</td></tr><tr class="div_0 true" id="252879" role="row"><td class="bg-blue-dark text-white fit-half maingame">2</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/33145/" data-original-title="Трудно плыть боком">Трудно плыть боком</a></td><td class="bg-blue-dark text-white total maingame">29<sub class="yellow qrs"> 117</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>10</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>9</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>10</strong></td><td class="bg-blue-light text-white fit-half matrixgame">2</td><td class="bg-blue-light text-white matrix-total fit matrixgame">470</td></tr><tr class="div_0 true" id="253305" role="row"><td class="bg-blue-dark text-white fit-half maingame">3</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/32079/" data-original-title="Мычание ягнят">Мычание ягнят</a></td><td class="bg-blue-dark text-white total maingame">28<sub class="yellow qrs"> 122</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>9</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>8</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>11</strong></td><td class="bg-blue-light text-white fit-half matrixgame">3</td><td class="bg-blue-light text-white matrix-total fit matrixgame">460</td></tr><tr class="div_0 true" id="253263" role="row"><td class="bg-blue-dark text-white fit-half maingame">4</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/31568/" data-original-title="Хроносинкластический Инфундибулум">Хроносинкластический Инфундибулум</a></td><td class="bg-blue-dark text-white total maingame">27<sub class="yellow qrs"> 100</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>9</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>8</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>10</strong></td><td class="bg-blue-light text-white fit-half matrixgame">5</td><td class="bg-blue-light text-white matrix-total fit matrixgame">370</td></tr><tr class="div_0 true" id="253260" role="row"><td class="bg-blue-dark text-white fit-half maingame">5</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/19461/" data-original-title="«П» в кубе">«П» в кубе</a></td><td class="bg-blue-dark text-white total maingame">23<sub class="yellow qrs"> 83</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>8</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>6</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>9</strong></td><td class="bg-blue-light text-white fit-half matrixgame">5</td><td class="bg-blue-light text-white matrix-total fit matrixgame">370</td></tr><tr class="div_0 true" id="253339" role="row"><td class="bg-blue-dark text-white fit-half maingame">6</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/21017/" data-original-title="Квантовый кот">Квантовый кот</a></td><td class="bg-blue-dark text-white total maingame">23<sub class="yellow qrs"> 84</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>9</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>6</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>8</strong></td><td class="bg-blue-light text-white fit-half matrixgame">9</td><td class="bg-blue-light text-white matrix-total fit matrixgame">250</td></tr><tr class="div_0 true" id="253338" role="row"><td class="bg-blue-dark text-white fit-half maingame">7</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/32214/" data-original-title="Коленки пчелы">Коленки пчелы</a></td><td class="bg-blue-dark text-white total maingame">21<sub class="yellow qrs"> 75</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>9</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>6</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>6</strong></td><td class="bg-blue-light text-white fit-half matrixgame">7</td><td class="bg-blue-light text-white matrix-total fit matrixgame">260</td></tr><tr class="div_0 true" id="253337" role="row"><td class="bg-blue-dark text-white fit-half maingame">8</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/7099/" data-original-title="Почему бы и нет">Почему бы и нет</a></td><td class="bg-blue-dark text-white total maingame">21<sub class="yellow qrs"> 61</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>7</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>3</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>11</strong></td><td class="bg-blue-light text-white fit-half matrixgame">12</td><td class="bg-blue-light text-white matrix-total fit matrixgame">150</td></tr><tr class="div_0 true" id="253261" role="row"><td class="bg-blue-dark text-white fit-half maingame">9</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/30529/" data-original-title="Тигры Разума">Тигры Разума</a></td><td class="bg-blue-dark text-white total maingame">20<sub class="yellow qrs"> 54</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>7</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>4</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>9</strong></td><td class="bg-blue-light text-white fit-half matrixgame">10</td><td class="bg-blue-light text-white matrix-total fit matrixgame">220</td></tr><tr class="div_0 true" id="253262" role="row"><td class="bg-blue-dark text-white fit-half maingame">10</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/30782/" data-original-title="О, Спорт!">О, Спорт!</a></td><td class="bg-blue-dark text-white total maingame">19<sub class="yellow qrs"> 69</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>6</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>4</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>9</strong></td><td class="bg-blue-light text-white fit-half matrixgame">4</td><td class="bg-blue-light text-white matrix-total fit matrixgame">450</td></tr><tr class="div_0 true" id="252878" role="row"><td class="bg-blue-dark text-white fit-half maingame">11</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/6379/" data-original-title="Байкодром Космодур">Байкодром Космодур</a></td><td class="bg-blue-dark text-white total maingame">18<sub class="yellow qrs"> 52</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>5</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>4</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>9</strong></td><td class="bg-blue-light text-white fit-half matrixgame">11</td><td class="bg-blue-light text-white matrix-total fit matrixgame">210</td></tr><tr class="div_0 true" id="253265" role="row"><td class="bg-blue-dark text-white fit-half maingame">12</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/33144/" data-original-title="Кусь Кусь Клан">Кусь Кусь Клан</a></td><td class="bg-blue-dark text-white total maingame">15<sub class="yellow qrs"> 37</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="text-blue"><strong>4</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>3</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-check text-light-blue" data-toggle="tooltip" data-placement="right" title="" data-original-title="+1"></i><span style="display:none;">1</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>8</strong></td><td class="bg-blue-light text-white fit-half matrixgame">7</td><td class="bg-blue-light text-white matrix-total fit matrixgame">260</td></tr></tbody><tfoot id="jsonTableFoot_1000" class="qrs"><tr role="row"><th class="maingame tablesorter-headerAsc" data-column="0"></th><th data-column="1"></th><th class="maingame tablesorter-headerAsc" data-column="2"></th><td class="d-none 1_tur" id="r1" data-column="3">3</td><td class="d-none 1_tur" id="r2" data-column="4">7</td><td class="d-none 1_tur" id="r3" data-column="5">3</td><td class="d-none 1_tur" id="r4" data-column="6">7</td><td class="d-none 1_tur" id="r5" data-column="7">2</td><td class="d-none 1_tur" id="r6" data-column="8">12</td><td class="d-none 1_tur" id="r7" data-column="9">2</td><td class="d-none 1_tur" id="r8" data-column="10">12</td><td class="d-none 1_tur" id="r9" data-column="11">11</td><td class="d-none 1_tur" id="r10" data-column="12">1</td><td class="d-none 1_tur" id="r11" data-column="13">2</td><td class="d-none 1_tur" id="r12" data-column="14">2</td><td data-column="15"></td><td class="d-none 2_tur" id="r13" data-column="16">7</td><td class="d-none 2_tur" id="r14" data-column="17">11</td><td class="d-none 2_tur" id="r15" data-column="18">3</td><td class="d-none 2_tur" id="r16" data-column="19">3</td><td class="d-none 2_tur" id="r17" data-column="20">8</td><td class="d-none 2_tur" id="r18" data-column="21">11</td><td class="d-none 2_tur" id="r19" data-column="22">7</td><td class="d-none 2_tur" id="r20" data-column="23">6</td><td class="d-none 2_tur" id="r21" data-column="24">4</td><td class="d-none 2_tur" id="r22" data-column="25">9</td><td class="d-none 2_tur" id="r23" data-column="26">9</td><td class="d-none 2_tur" id="r24" data-column="27">7</td><td data-column="28"></td><td class="d-none 3_tur" id="r25" data-column="29">3</td><td class="d-none 3_tur" id="r26" data-column="30">4</td><td class="d-none 3_tur" id="r27" data-column="31">4</td><td class="d-none 3_tur" id="r28" data-column="32">12</td><td class="d-none 3_tur" id="r29" data-column="33">1</td><td class="d-none 3_tur" id="r30" data-column="34">2</td><td class="d-none 3_tur" id="r31" data-column="35">1</td><td class="d-none 3_tur" id="r32" data-column="36">1</td><td class="d-none 3_tur" id="r33" data-column="37">1</td><td class="d-none 3_tur" id="r34" data-column="38">4</td><td class="d-none 3_tur" id="r35" data-column="39">2</td><td class="d-none 3_tur" id="r36" data-column="40">11</td><td data-column="41"></td><th class="matrixgame" data-column="42"></th><th class="matrixgame" data-column="43"></th></tr></tfoot></table></div></div></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Захарова Л. А. | <a class="foot-link" href="https://elba.kontur.ru/card/2b1oxlljoa" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.4.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.4.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.4.1"></script><script src="/static/js/popper.min.js?v=1.4.1"></script><script src="/static/js/bootstrap.min.js?v=1.4.1"></script><script src="/static/js/corner-popup.js?v=1.4.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.4.1"></script><script src="/static/js/alertify.min.js?v=1.4.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode('Санкт-Петербург'));
      $('#league_name').text(htmlDecode('Открытая лига'));
      $('.modal-city').text(htmlDecode('Санкт-Петербург'));
      
      set_city(71, 'Санкт-Петербург');
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'RuWJazV85E9suNublrmmTi4XXlCPoGo45skoe0HpC1TXP7xeyTyCLhKlOd7WD7Mm',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.4.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.4.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.4.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.4.1"></script><script src="/static/js/FileSaver.min.js?v=1.4.1"></script><script src="/static/js/tableexport.min.js?v=1.4.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [3, 7, 3, 7, 2, 12, 2, 12, 11, 1, 2, 2, 7, 11, 3, 3, 8, 11, 7, 6, 4, 9, 9, 7, 3, 4, 4, 12, 1, 2, 1, 1, 1, 4, 2, 11], "mediatours": [], "media_game": false, "matrix_game": true, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'RuWJazV85E9suNublrmmTi4XXlCPoGo45skoe0HpC1TXP7xeyTyCLhKlOd7WD7Mm',
            game_id: '17630',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [10, 6, 10, 6, 11, 1, 11, 1, 2, 12, 11, 11, 6, 2, 10, 10, 5, 2, 6, 7, 9, 4, 4, 6, 10, 9, 9, 1, 12, 11, 12, 12, 12, 9, 11, 2];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 12 | взятие: 63.19% | гробов:  0 | спасённых:  3 | кнопок:  5 | дыр: 6."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Открытая лига | Игра #8',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html17631 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига | Игра #9 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-22/23 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Открытая лига | Игра #9 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-22/23 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.4.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.4.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.4.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.4.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.4.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.4.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/style.css?v=1.4.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Открытая лига</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/17631/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/17631/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига-19</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига | Игра #9</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/117/">Открытая лига</a> / <a href="/season/1916/">Зима-22/23</a> / 13 февраля 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/17631/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/17631/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"><div id="table_1000"><h3 class="dark-blue text-center">Итоговая таблица</h3><div class="table-responsive"><table class="table table-hover table-sm text-center table-nonfluid mb-3 borderless-top tablesorter tablesorter-default tablesorter8fe90465f2a4a" id="jsonTable_1000" role="grid"><thead id="jsonTableHead_1000"><tr class="text-blue tablesorter-headerRow" role="row"><th class="text-center fit-half bg-white maingame tablesorter-header tablesorter-headerAsc" data-column="0" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="ascending" aria-label=": Ascending sort applied, activate to apply a descending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i class="fa-solid fa-medal"></i></div></th><th class="text-left tablesorter-header tablesorter-headerUnSorted" data-column="1" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="Название: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Название</div></th><th class="text-center fit maingame tablesorter-header tablesorter-headerAsc" data-column="2" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="ascending" aria-label="Сумма: Ascending sort applied, activate to apply a descending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Сумма</div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="3" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="1: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>1</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="4" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="2: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>2</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="5" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="3: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>3</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="6" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="4: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>4</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="7" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="5: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>5</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="8" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="6: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>6</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="9" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="7: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>7</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="10" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="8: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>8</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="11" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="9: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>9</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="12" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="10: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>10</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="13" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="11: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>11</small></div></th><th class="text-center d-none 1_tur tablesorter-header tablesorter-headerUnSorted" data-column="14" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="12: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>12</small></div></th><th class="text-center fit-half main_a tablesorter-header tablesorter-headerUnSorted" data-column="15" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="1: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i onclick="toogletur(1)" class="table_blue_a far fa-plus-square tablesorter-noSort 1_but"></i>1</div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="16" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="13: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>13</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="17" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="14: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>14</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="18" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="15: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>15</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="19" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="16: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>16</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="20" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="17: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>17</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="21" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="18: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>18</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="22" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="19: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>19</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="23" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="20: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>20</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="24" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="21: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>21</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="25" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="22: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>22</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="26" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="23: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>23</small></div></th><th class="text-center d-none 2_tur tablesorter-header tablesorter-headerUnSorted" data-column="27" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="24: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>24</small></div></th><th class="text-center fit-half main_a tablesorter-header tablesorter-headerUnSorted" data-column="28" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="2: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i onclick="toogletur(2)" class="table_blue_a far fa-plus-square tablesorter-noSort 2_but"></i>2</div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="29" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="25: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>25</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="30" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="26: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>26</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="31" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="27: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>27</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="32" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="28: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>28</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="33" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="29: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>29</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="34" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="30: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>30</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="35" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="31: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>31</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="36" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="32: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>32</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="37" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="33: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>33</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="38" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="34: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>34</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="39" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="35: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>35</small></div></th><th class="text-center d-none 3_tur tablesorter-header tablesorter-headerUnSorted" data-column="40" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="36: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><small>36</small></div></th><th class="text-center fit-half main_a tablesorter-header tablesorter-headerUnSorted" data-column="41" tabindex="0" scope="col" role="columnheader" aria-disabled="false" aria-controls="jsonTable_1000" unselectable="on" aria-sort="none" aria-label="3: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner"><i onclick="toogletur(3)" class="table_blue_a far fa-plus-square tablesorter-noSort 3_but"></i>3</div></th></tr></thead><tbody class="tBody" id="jsonTableBody_1000" aria-live="polite" aria-relevant="all"><tr class="div_0 true" id="252880" role="row"><td class="bg-blue-dark text-white fit-half maingame">1</td><td class="text-nowrap text-left main_a elipsis px-2 max-title-width"><a class="table_a " data-toggle="tooltip" title="" href="/team/33812/" data-original-title="Геном чебурека">Геном чебурека</a></td><td class="bg-blue-dark text-white total maingame">0<sub class="yellow qrs"> 0</sub></td><td class="d-none count 1_tur" id="1"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="2"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="3"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="4"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="5"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="6"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="7"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="8"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="9"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="10"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="11"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 1_tur" id="12"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>0</strong></td><td class="d-none count 2_tur" id="13"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="14"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="15"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="16"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="17"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="18"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="19"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="20"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="21"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="22"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="23"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 2_tur" id="24"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>0</strong></td><td class="d-none count 3_tur" id="25"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="26"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="27"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="28"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="29"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="30"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="31"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="32"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="33"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="34"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="35"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="d-none count 3_tur" id="36"><i class="fas fa-minus minus"></i><span style="display:none;">0</span></td><td class="text-blue"><strong>0</strong></td></tr></tbody><tfoot id="jsonTableFoot_1000" class="qrs"><tr role="row"><th class="maingame tablesorter-headerAsc" data-column="0"></th><th data-column="1"></th><th class="maingame tablesorter-headerAsc" data-column="2"></th><td class="d-none 1_tur" id="r1" data-column="3">2</td><td class="d-none 1_tur" id="r2" data-column="4">2</td><td class="d-none 1_tur" id="r3" data-column="5">2</td><td class="d-none 1_tur" id="r4" data-column="6">2</td><td class="d-none 1_tur" id="r5" data-column="7">2</td><td class="d-none 1_tur" id="r6" data-column="8">2</td><td class="d-none 1_tur" id="r7" data-column="9">2</td><td class="d-none 1_tur" id="r8" data-column="10">2</td><td class="d-none 1_tur" id="r9" data-column="11">2</td><td class="d-none 1_tur" id="r10" data-column="12">2</td><td class="d-none 1_tur" id="r11" data-column="13">2</td><td class="d-none 1_tur" id="r12" data-column="14">2</td><td data-column="15"></td><td class="d-none 2_tur" id="r13" data-column="16">2</td><td class="d-none 2_tur" id="r14" data-column="17">2</td><td class="d-none 2_tur" id="r15" data-column="18">2</td><td class="d-none 2_tur" id="r16" data-column="19">2</td><td class="d-none 2_tur" id="r17" data-column="20">2</td><td class="d-none 2_tur" id="r18" data-column="21">2</td><td class="d-none 2_tur" id="r19" data-column="22">2</td><td class="d-none 2_tur" id="r20" data-column="23">2</td><td class="d-none 2_tur" id="r21" data-column="24">2</td><td class="d-none 2_tur" id="r22" data-column="25">2</td><td class="d-none 2_tur" id="r23" data-column="26">2</td><td class="d-none 2_tur" id="r24" data-column="27">2</td><td data-column="28"></td><td class="d-none 3_tur" id="r25" data-column="29">2</td><td class="d-none 3_tur" id="r26" data-column="30">2</td><td class="d-none 3_tur" id="r27" data-column="31">2</td><td class="d-none 3_tur" id="r28" data-column="32">2</td><td class="d-none 3_tur" id="r29" data-column="33">2</td><td class="d-none 3_tur" id="r30" data-column="34">2</td><td class="d-none 3_tur" id="r31" data-column="35">2</td><td class="d-none 3_tur" id="r32" data-column="36">2</td><td class="d-none 3_tur" id="r33" data-column="37">2</td><td class="d-none 3_tur" id="r34" data-column="38">2</td><td class="d-none 3_tur" id="r35" data-column="39">2</td><td class="d-none 3_tur" id="r36" data-column="40">2</td><td data-column="41"></td></tr></tfoot></table></div></div></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Захарова Л. А. | <a class="foot-link" href="https://elba.kontur.ru/card/2b1oxlljoa" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.4.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.4.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.4.1"></script><script src="/static/js/popper.min.js?v=1.4.1"></script><script src="/static/js/bootstrap.min.js?v=1.4.1"></script><script src="/static/js/corner-popup.js?v=1.4.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.4.1"></script><script src="/static/js/alertify.min.js?v=1.4.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode('Санкт-Петербург'));
      $('#league_name').text(htmlDecode('Открытая лига'));
      $('.modal-city').text(htmlDecode('Санкт-Петербург'));
      
      set_city(71, 'Санкт-Петербург');
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'ipZxtI6Zl1s45j0G0uZ9mgz4fCk3Zj9hwnncx9SgSoczqD3JdWbpeffs6uPaeKxz',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.4.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.4.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.4.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.4.1"></script><script src="/static/js/FileSaver.min.js?v=1.4.1"></script><script src="/static/js/tableexport.min.js?v=1.4.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'ipZxtI6Zl1s45j0G0uZ9mgz4fCk3Zj9hwnncx9SgSoczqD3JdWbpeffs6uPaeKxz',
            game_id: '17631',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 1 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 36."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Открытая лига | Игра #9',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html17632 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига | Игра #10 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-22/23 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Открытая лига | Игра #10 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-22/23 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.4.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.4.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.4.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.4.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.4.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.4.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/style.css?v=1.4.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Открытая лига</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/17632/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/17632/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига-19</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига | Игра #10</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/117/">Открытая лига</a> / <a href="/season/1916/">Зима-22/23</a> / 20 февраля 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/17632/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/17632/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Захарова Л. А. | <a class="foot-link" href="https://elba.kontur.ru/card/2b1oxlljoa" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.4.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.4.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.4.1"></script><script src="/static/js/popper.min.js?v=1.4.1"></script><script src="/static/js/bootstrap.min.js?v=1.4.1"></script><script src="/static/js/corner-popup.js?v=1.4.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.4.1"></script><script src="/static/js/alertify.min.js?v=1.4.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode('Санкт-Петербург'));
      $('#league_name').text(htmlDecode('Открытая лига'));
      $('.modal-city').text(htmlDecode('Санкт-Петербург'));
      
      set_city(71, 'Санкт-Петербург');
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'sBLjZvnkoDiBPJ8sjqCJ2mDniPJOC0w9Gz9Y3W9BV026a3bvwSOZUljL9HeVRrUr',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.4.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.4.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.4.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.4.1"></script><script src="/static/js/FileSaver.min.js?v=1.4.1"></script><script src="/static/js/tableexport.min.js?v=1.4.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'sBLjZvnkoDiBPJ8sjqCJ2mDniPJOC0w9Gz9Y3W9BV026a3bvwSOZUljL9HeVRrUr',
            game_id: '17632',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Открытая лига | Игра #10',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html17633 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига | Финал | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-22/23 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Открытая лига | Финал | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-22/23 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.4.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.4.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.4.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.4.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.4.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.4.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/style.css?v=1.4.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Открытая лига</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/17633/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/17633/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига-19</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига | Финал</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/117/">Открытая лига</a> / <a href="/season/1916/">Зима-22/23</a> / 27 февраля 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/17633/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/17633/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Захарова Л. А. | <a class="foot-link" href="https://elba.kontur.ru/card/2b1oxlljoa" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.4.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.4.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.4.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.4.1"></script><script src="/static/js/popper.min.js?v=1.4.1"></script><script src="/static/js/bootstrap.min.js?v=1.4.1"></script><script src="/static/js/corner-popup.js?v=1.4.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.4.1"></script><script src="/static/js/alertify.min.js?v=1.4.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode('Санкт-Петербург'));
      $('#league_name').text(htmlDecode('Открытая лига'));
      $('.modal-city').text(htmlDecode('Санкт-Петербург'));
      
      set_city(71, 'Санкт-Петербург');
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'bPYyOaYwtvMLGdrJl8hr0yvTbwiDzEsgpNmdSBKN0Swg1xuMyAtHSxbh2oNKO5Qy',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.4.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.4.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.4.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.4.1"></script><script src="/static/js/FileSaver.min.js?v=1.4.1"></script><script src="/static/js/tableexport.min.js?v=1.4.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 4, "qlist": [10, 20, 30, 40], "questions": [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'bPYyOaYwtvMLGdrJl8hr0yvTbwiDzEsgpNmdSBKN0Swg1xuMyAtHSxbh2oNKO5Qy',
            game_id: '17633',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 0 | взятие: 0.00% | гробов:  40 | спасённых:  0 | кнопок:  40 | дыр: 0."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Открытая лига | Финал',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html3 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Клуб «60 секунд»</title><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/league/118/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/league/118/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/118/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/118/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12" style="margin-top: -20px;"><div class="d-flex flex-row-reverse"><h4><a class="blue" href="None" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="blue" href="None" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-md-6 col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Первая лига</h2><div class="white-razdel"></div><div class="l-blue pt-2">Клуб «60 Секунд» / Санкт-Петербург</div></div></div></div><div class="col-md-6 col-12"><h5 class="yellow text-center">Ближайшая игра</h5><h6 class="l-grey text-center"><a href="/game/18615/">Первая лига | Игра #4</a></h6><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="l-blue ff-ibc vertical-text align-middle"><span>Играем</span></div><div><table><tbody><tr><td class="l-grey"><i class="fas fa-calendar-alt fa-fw l-blue"></i>&nbsp;04 апреля, Вторник</td></tr><tr><td class="l-grey"><i class="fas fa-clock fa-fw l-blue"></i>&nbsp;19:30</td></tr><tr><td class="l-grey"><i class="fas fa-map-marker-alt fa-fw l-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="l-grey"><i class="fas fa-credit-card l-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2361915/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0 mt-md-n4"><ul class="nav nav-pills justify-content-md-start justify-content-center" id="myTab" role="tablist"><li class="nav-item nav-pills-yellow"><a aria-controls="schedule" aria-selected="true" class="nav-link tab-link active" data-toggle="tab" href="#schedule" id="schedule-tab" role="tab"><span>Лига</span></a></li><li class="nav-item nav-pills-yellow"><a aria-controls="seasons" aria-selected="true" class="nav-link tab-link" data-toggle="tab" href="#seasons" id="seasons-tab" role="tab"><span>Все сезоны</span></a></li><li class="nav-item nav-pills-yellow"><a class="nav-link tab-link" href="/schedule/71/"><span>Расписание клуба</span></a></li></ul></div></div></div></div><div class="container"><div class="row pt-4"><div class="col-12"><div class="tab-content" id="myTabContent"><div aria-labelledby="schedule-tab" class="tab-pane fade show active" id="schedule" role="tabpanel"><div class="row"><div class="col-12"><h4 class="text-center dark-blue">Ближайшие игры лиги</h4></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/18615/">Первая лига | Игра #4</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;04 апреля, Вторник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2361915/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/18616/">Первая лига | Игра #5</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;11 апреля, Вторник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/18617/">Первая лига | Игра #6</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;18 апреля, Вторник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1200 руб. с команды</td></tr></tbody></table></div></div></div></div></div><div class="col-12"><h5 class="dark-blue text-center">Рейтинг сезона <a href="/season/2091/">Весна 2023</a></h5></div><div class=""><div class="container-fluid"><div class="row"><div class="col-12"><nav class="tabbable"><ul class="nav nav-pills justify-content-sm-start justify-content-center" id="myTab" role="tablist"><li class="nav-item text-center nav-pills-dark"><a aria-controls="1" aria-selected="true" class="nav-link active" data-toggle="tab" href="#h1" id="h1-tab" role="tab"><span>Первая лига</span></a></li></ul></nav></div></div></div><div class="tab-content" id="myTabContent"><div aria-labelledby="h1-tab" class="tab-pane fade show active" id="h1" role="tabpanel"><div class="table-responsive pt-2 pb-3"><table class="table table-hover table-sm table-nonfluid borderless-top" id="rate_table"><thead><tr><th class="text-center fit-strickt">&nbsp;&nbsp;#&nbsp;&nbsp;</th><th class="text-left fit-strickt">Название</th><th class="text-center fit-strickt">&nbsp;&nbsp;&nbsp;Сумма&nbsp;&nbsp;&nbsp;</th><th class="text-center text-nowrap small"><a class="table_a" href="/game/18612/">14.03.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/18613/">21.03.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/18614/">28.03.23</a></th></tr></thead><tbody id="myTable"><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>1</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/8999/">Так получилось</a></td><td class="bg-blue-dark text-white text-center fit-strickt">89.00</td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center text-secondary"><span>33.50</span></td><td class="text-center text-secondary"><span>23.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>2</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/14921/">Картель Укупника</a></td><td class="bg-blue-dark text-white text-center fit-strickt">81.25</td><td class="text-center text-secondary"><span>30.00</span></td><td class="text-center text-secondary"><span>27.75</span></td><td class="text-center text-secondary"><span>23.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>3</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/30850/">Ангелы Кренделя</a> <sup class="dark-blue">+2</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">68.50</td><td class="text-center text-secondary"><span>23.50</span></td><td class="text-center text-secondary"><span>19.00</span></td><td class="text-center text-secondary"><span>26.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>4</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/5093/">Напутствие Широкова</a> <sup class="yellow">-1</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">67.75</td><td class="text-center text-secondary"><span>23.50</span></td><td class="text-center text-secondary"><span>27.75</span></td><td class="text-center text-secondary"><span>16.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>5</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/30230/">6x7</a> <sup class="yellow">-1</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">67.50</td><td class="text-center text-secondary"><span>25.50</span></td><td class="text-center text-secondary"><span>21.50</span></td><td class="text-center text-secondary"><span>20.50</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>6</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/8178/">The Hateful Seven</a> <sup class="dark-blue">+1</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">66.00</td><td class="text-center text-secondary"><span>20.50</span></td><td class="text-center text-secondary"><span>13.50</span></td><td class="text-center text-secondary"><span>32.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>7</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/5112/">Мама, я больше не Будда</a> <sup class="yellow">-1</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">65.75</td><td class="text-center text-secondary"><span>12.00</span></td><td class="text-center text-secondary"><span>27.75</span></td><td class="text-center text-secondary"><span>26.00</span></td></tr></tbody></table></div></div></div></div><div class="row"><div class="col-12 pt-2"><h4 class="dark-blue text-center">Статистика</h4><p><small>Ниже представлены топ-5 команд в некоторых номинациях. Обратите внимание, что это общая статистика по всем играм. Кроме общей статистики, есть расширенная статистика для каждого из сезонов. Например, вот статистика <a class="table_a" href="/season/2091/#stats">текущего</a> и <a class="table_a" href="/season/1927/#stats">предыдущего</a> сезонов. Кликнув на любую из команд лиги - вы попадете на страницу этой команды, где сможете ознакомится с её подробной статистикой, как за всё время, так и за каждый сезон в отдельности. Команды можно искать с помощью поиска, а также из списка всех <a class="table_a" href="/teams/118/">команд лиги</a>.</small></p></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Больше всех сыграли</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/7114/">Развал-снисхождение</a></td><td class="text-center">257</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">222</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5144/">Второе предельное состояние</a></td><td class="text-center">196</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">195</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5093/">Напутствие Широкова</a></td><td class="text-center">184</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Побед в сезонах</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">222</td><td class="text-center">4</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/24679/">Файервольные каменщики</a></td><td class="text-center">54</td><td class="text-center">3</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5142/">Power Of Sport</a></td><td class="text-center">99</td><td class="text-center">2</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">2</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/7114/">Развал-снисхождение</a></td><td class="text-center">257</td><td class="text-center">1</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Победители</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">37</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">222</td><td class="text-center">25</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/8603/">Витамин 404</a></td><td class="text-center">151</td><td class="text-center">22</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/7114/">Развал-снисхождение</a></td><td class="text-center">257</td><td class="text-center">21</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5142/">Power Of Sport</a></td><td class="text-center">99</td><td class="text-center">17</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Призёры</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Пьедесталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">222</td><td class="text-center">89</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/7114/">Развал-снисхождение</a></td><td class="text-center">257</td><td class="text-center">63</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/8603/">Витамин 404</a></td><td class="text-center">151</td><td class="text-center">62</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">58</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">195</td><td class="text-center">42</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучший процент взятия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Процент взятия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">80.4 %</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/24679/">Файервольные каменщики</a></td><td class="text-center">54</td><td class="text-center">76.0 %</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">222</td><td class="text-center">72.2 %</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5080/">Незнайки на Луне</a></td><td class="text-center">43</td><td class="text-center">71.2 %</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/8603/">Витамин 404</a></td><td class="text-center">151</td><td class="text-center">70.3 %</td></tr></tbody></table></div><small><sup>*</sup> - среди команд сыгравших больше 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучшая серия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">195</td><td class="text-center">35</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/7114/">Развал-снисхождение</a></td><td class="text-center">257</td><td class="text-center">31</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/29121/">Во всём виноват Тиндер</a></td><td class="text-center">31</td><td class="text-center">28</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5080/">Незнайки на Луне</a></td><td class="text-center">43</td><td class="text-center">28</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">27</td></tr></tbody></table></div><small><sup>*</sup> - количество правильных ответов подряд</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">лучшая худшая серия<sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/8999/">Так получилось</a></td><td class="text-center">30</td><td class="text-center">4</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/24679/">Файервольные каменщики</a></td><td class="text-center">54</td><td class="text-center">5</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/30230/">6x7</a></td><td class="text-center">28</td><td class="text-center">5</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">6</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5080/">Незнайки на Луне</a></td><td class="text-center">43</td><td class="text-center">6</td></tr></tbody></table></div><small><sup>*</sup> - количество неправильных ответов подряд, среди команд сыгравших минимум 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Тоталов в турах <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Тоталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">222</td><td class="text-center">26</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">23</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5142/">Power Of Sport</a></td><td class="text-center">99</td><td class="text-center">16</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">195</td><td class="text-center">10</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/7115/">Тритий лишний</a></td><td class="text-center">169</td><td class="text-center">9</td></tr></tbody></table></div><small><sup>*</sup> - тотал - это тур в котором, команда ответила верно на все вопросы</small></div></div></div><div aria-labelledby="seasons-tab" class="tab-pane fade" id="seasons" role="tabpanel"><div class=""><h4 class="text-center dark-blue">Сезоны</h4><div class="table-responsive pb-4"><table class="table table-hover table-sm"><thead><tr class="active"><th class="text-nowrap">Сезон</th><th class="text-nowrap text-center">Игр</th><th class="text-nowrap">Начало</th><th class="text-nowrap">Окончание</th></tr></thead><tbody><tr><td class="text-nowrap"><a class="table_a" href="/season/2091/">Весна 2023</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">14 Март 2023</td><td class="text-nowrap">30 Май 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1927/">Зима-22/23</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">29 Ноябрь 2022</td><td class="text-nowrap">28 Февраль 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1814/">Осень 2022</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">29 Август 2022</td><td class="text-nowrap">21 Ноябрь 2022</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1739/">Лето 2022</a></td><td class="text-nowrap text-center">10</td><td class="text-nowrap">06 Июнь 2022</td><td class="text-nowrap">22 Август 2022</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1624/">Весна 2022</a></td><td class="text-nowrap text-center">10</td><td class="text-nowrap">07 Март 2022</td><td class="text-nowrap">30 Май 2022</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1506/">Зима 2022</a></td><td class="text-nowrap text-center">11</td><td class="text-nowrap">06 Декабрь 2021</td><td class="text-nowrap">28 Февраль 2022</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1393/">Осень 2021</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">06 Сентябрь 2021</td><td class="text-nowrap">29 Ноябрь 2021</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1308/">Лето 2021</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">07 Июнь 2021</td><td class="text-nowrap">30 Август 2021</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1198/">Весна 2021</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">15 Март 2021</td><td class="text-nowrap">31 Май 2021</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1131/">Зима 2021</a></td><td class="text-nowrap text-center">6</td><td class="text-nowrap">18 Январь 2021</td><td class="text-nowrap">01 Март 2021</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1031/">Поздняя осень 2020</a></td><td class="text-nowrap text-center">7</td><td class="text-nowrap">12 Октябрь 2020</td><td class="text-nowrap">23 Ноябрь 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/969/">Блиц-сезон:Евротур</a></td><td class="text-nowrap text-center">4</td><td class="text-nowrap">14 Сентябрь 2020</td><td class="text-nowrap">05 Октябрь 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/949/">Блиц-сезон:Космическая Одиссея</a></td><td class="text-nowrap text-center">4</td><td class="text-nowrap">10 Август 2020</td><td class="text-nowrap">31 Август 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/741/">Весна 2020</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">10 Март 2020</td><td class="text-nowrap">26 Май 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/614/">Зима 2020</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">03 Декабрь 2019</td><td class="text-nowrap">03 Март 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/536/">Осень 2019</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">10 Сентябрь 2019</td><td class="text-nowrap">26 Ноябрь 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/420/">Лето 2019</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">04 Июнь 2019</td><td class="text-nowrap">27 Август 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/364/">Весна 2019</a></td><td class="text-nowrap text-center">8</td><td class="text-nowrap">15 Апрель 2019</td><td class="text-nowrap">28 Май 2019</td></tr></tbody></table></div></div></div></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12 my-n3"><div class="d-flex flex-row-reverse pt-1"><h4><a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: '9fyn1JKdkUnJYNR52qHH7fhdGuS9gpCKndW25awuRh7ej7U8fSTXZeXBxmngvQ02',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><script type="text/javascript">
	
	$('a[data-toggle="tab"]').on('shown.bs.tab', function (e) {
		var hash = $(e.target).attr('href');
		if (history.pushState) {
			history.pushState(null, null, hash);
		} else {
			location.hash = hash;
		}
	});

	jQuery(document).ready(function ($) {
		let selectedTab = window.location.hash;
		$('.nav-link[href="' + selectedTab + '"]' ).trigger('click');
	})
</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html4 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Клуб «60 секунд»</title><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><meta http-equiv="origin-trial" content="AymqwRC7u88Y4JPvfIF2F37QKylC04248hLCdJAsh8xgOfe/dVJPV3XS3wLFca1ZMVOtnBfVjaCMTVudWM//5g4AAAB7eyJvcmlnaW4iOiJodHRwczovL3d3dy5nb29nbGV0YWdtYW5hZ2VyLmNvbTo0NDMiLCJmZWF0dXJlIjoiUHJpdmFjeVNhbmRib3hBZHNBUElzIiwiZXhwaXJ5IjoxNjk1MTY3OTk5LCJpc1RoaXJkUGFydHkiOnRydWV9"></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/league/118/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/league/118/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/118/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/118/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12" style="margin-top: -20px;"><div class="d-flex flex-row-reverse"><h4><a class="blue" href="None" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="blue" href="None" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-md-6 col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Первая лига</h2><div class="white-razdel"></div><div class="l-blue pt-2">Клуб «60 Секунд» / Санкт-Петербург</div></div></div></div><div class="col-md-6 col-12"><h5 class="yellow text-center">Ближайшая игра</h5><h6 class="l-grey text-center"><a href="/game/21281/">Первая лига | Игра #8</a></h6><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="l-blue ff-ibc vertical-text align-middle"><span>Играем</span></div><div><table><tbody><tr><td class="l-grey"><i class="fas fa-calendar-alt fa-fw l-blue"></i>&nbsp;31 октября, Вторник</td></tr><tr><td class="l-grey"><i class="fas fa-clock fa-fw l-blue"></i>&nbsp;19:30</td></tr><tr><td class="l-grey"><i class="fas fa-map-marker-alt fa-fw l-blue"></i>&nbsp;Rossi's Club - ул. Зодчего Росси, 1-3</td></tr><tr><td class="l-grey"><i class="fas fa-credit-card l-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2638754/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0 mt-md-n4"><ul class="nav nav-pills justify-content-md-start justify-content-center" id="myTab" role="tablist"><li class="nav-item nav-pills-yellow"><a aria-controls="schedule" aria-selected="true" class="nav-link tab-link active" data-toggle="tab" href="#schedule" id="schedule-tab" role="tab"><span>Лига</span></a></li><li class="nav-item nav-pills-yellow"><a aria-controls="seasons" aria-selected="true" class="nav-link tab-link" data-toggle="tab" href="#seasons" id="seasons-tab" role="tab"><span>Все сезоны</span></a></li><li class="nav-item nav-pills-yellow"><a class="nav-link tab-link" href="/schedule/71/"><span>Расписание клуба</span></a></li></ul></div></div></div></div><div class="container"><div class="row pt-4"><div class="col-12"><div class="tab-content" id="myTabContent"><div aria-labelledby="schedule-tab" class="tab-pane fade show active" id="schedule" role="tabpanel"><div class="row"><div class="col-12"><h4 class="text-center dark-blue">Ближайшие игры лиги</h4></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/21281/">Первая лига | Игра #8</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;31 октября, Вторник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Rossi's Club - ул. Зодчего Росси, 1-3</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1500 руб. с команды</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2638754/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/21282/">Первая лига | Игра #9</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;07 ноября, Вторник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1500 руб. с команды</td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/21283/">Первая лига | Игра #10</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;14 ноября, Вторник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp; Rossi's Club - ул. Зодчего Росси, 1-3</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;1500 руб. с команды</td></tr></tbody></table></div></div></div></div></div><div class="col-12"><h5 class="dark-blue text-center">Рейтинг сезона <a href="/season/2333/">Осень 2023</a></h5></div><div class=""><div class="container-fluid"><div class="row"><div class="col-12"><nav class="tabbable"><ul class="nav nav-pills justify-content-sm-start justify-content-center" id="myTab" role="tablist"><li class="nav-item text-center nav-pills-dark"><a aria-controls="1" aria-selected="true" class="nav-link active" data-toggle="tab" href="#h1" id="h1-tab" role="tab"><span>Первая лига</span></a></li></ul></nav></div></div></div><div class="tab-content" id="myTabContent"><div aria-labelledby="h1-tab" class="tab-pane fade show active" id="h1" role="tabpanel"><div class="table-responsive pt-2 pb-3"><table class="table table-hover table-sm table-nonfluid borderless-top" id="rate_table"><thead><tr><th class="text-center fit-strickt">&nbsp;&nbsp;#&nbsp;&nbsp;</th><th class="text-left fit-strickt">Название</th><th class="text-center fit-strickt">&nbsp;&nbsp;&nbsp;Сумма&nbsp;&nbsp;&nbsp;</th><th class="text-center text-nowrap small"><a class="table_a" href="/game/21273/">05.09.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/21274/">12.09.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/21275/">19.09.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/21276/">26.09.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/21277/">10.10.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/21278/">17.10.23</a></th><th class="text-center text-nowrap small"><a class="table_a" href="/game/21280/">24.10.23</a></th></tr></thead><tbody id="myTable"><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>1</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/26664/">Капитолийская горчица</a></td><td class="bg-blue-dark text-white text-center fit-strickt">220.25</td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>29.25</span></td><td class="text-center text-secondary"><span>21.50</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>33.50</span></td><td class="text-center text-secondary"><span>31.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>2</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/38784/">Бегущие</a></td><td class="bg-blue-dark text-white text-center fit-strickt">194.33</td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>35.00</span></td><td class="text-center text-secondary"><span>28.33</span></td><td class="text-center text-secondary"><span>25.50</span></td><td class="text-center text-secondary"><span>30.00</span></td><td class="text-center text-secondary"><span>19.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>3</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/8178/">The Hateful Seven</a></td><td class="bg-blue-dark text-white text-center fit-strickt">175.50</td><td class="text-center text-secondary"><span>20.00</span></td><td class="text-center text-secondary"><span>20.50</span></td><td class="text-center text-secondary"><span>21.50</span></td><td class="text-center text-secondary"><span>32.00</span></td><td class="text-center text-secondary"><span>25.50</span></td><td class="text-center text-secondary"><span>25.00</span></td><td class="text-center text-secondary"><span>31.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>4</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/28237/">Маги Бумаги</a> <sup class="dark-blue">+1</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">167.00</td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>16.00</span></td><td class="text-center text-secondary"><span>27.00</span></td><td class="text-center text-secondary"><span>23.00</span></td><td class="text-center text-secondary"><span>29.00</span></td><td class="text-center text-secondary"><span>22.50</span></td><td class="text-center text-secondary"><span>25.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>5</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/5090/">Почеши Чешира</a> <sup class="yellow">-1</sup></td><td class="bg-blue-dark text-white text-center fit-strickt">161.75</td><td class="text-center text-secondary"><span>29.00</span></td><td class="text-center text-secondary"><span>29.25</span></td><td class="text-center text-secondary"><span>16.50</span></td><td class="text-center text-secondary"><span>25.50</span></td><td class="text-center text-secondary"><span>20.00</span></td><td class="text-center text-secondary"><span>22.50</span></td><td class="text-center text-secondary"><span>19.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>6</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/5112/">Мама, я больше не Будда</a></td><td class="bg-blue-dark text-white text-center fit-strickt">160.50</td><td class="text-center text-secondary"><span>20.00</span></td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>30.00</span></td><td class="text-center text-secondary"><span>15.50</span></td><td class="text-center text-secondary"><span>25.50</span></td><td class="text-center text-secondary"><span>26.00</span></td><td class="text-center text-secondary"><span>19.00</span></td></tr><tr><td class="bg-blue-dark text-white text-center fit-strickt"><strong>7</strong></td><td class="text-nowrap elipsis fit-strickt"><a class="table_a" href="/team/5093/">Напутствие Широкова</a></td><td class="bg-blue-dark text-white text-center fit-strickt">153.00</td><td class="text-center text-secondary"><span>16.50</span></td><td class="text-center text-secondary"><span>22.00</span></td><td class="text-center text-secondary"><span>24.50</span></td><td class="text-center text-secondary"><span>25.50</span></td><td class="text-center text-secondary"><span>22.00</span></td><td class="text-center text-secondary"><span>15.00</span></td><td class="text-center text-secondary"><span>27.50</span></td></tr></tbody></table></div></div></div></div><div class="row"><div class="col-12 pt-2"><h4 class="dark-blue text-center">Статистика</h4><p><small>Ниже представлены топ-5 команд в некоторых номинациях. Обратите внимание, что это общая статистика по всем играм. Кроме общей статистики, есть расширенная статистика для каждого из сезонов. Например, вот статистика <a class="table_a" href="/season/2333/#stats">текущего</a> и <a class="table_a" href="/season/2223/#stats">предыдущего</a> сезонов. Кликнув на любую из команд лиги - вы попадете на страницу этой команды, где сможете ознакомится с её подробной статистикой, как за всё время, так и за каждый сезон в отдельности. Команды можно искать с помощью поиска, а также из списка всех <a class="table_a" href="/teams/118/">команд лиги</a>.</small></p></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Больше всех сыграли</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">252</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">225</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5144/">Второе предельное состояние</a></td><td class="text-center">224</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5077/">Сайлоны</a></td><td class="text-center">218</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5093/">Напутствие Широкова</a></td><td class="text-center">211</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Побед в сезонах</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">252</td><td class="text-center">5</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/24679/">Файервольные каменщики</a></td><td class="text-center">61</td><td class="text-center">3</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5142/">Power Of Sport</a></td><td class="text-center">99</td><td class="text-center">2</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">2</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/26664/">Капитолийская горчица</a></td><td class="text-center">22</td><td class="text-center">2</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Победители</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">37</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">252</td><td class="text-center">31</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/8603/">Витамин 404</a></td><td class="text-center">177</td><td class="text-center">22</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">225</td><td class="text-center">19</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5142/">Power Of Sport</a></td><td class="text-center">99</td><td class="text-center">17</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Призёры</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Пьедесталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">252</td><td class="text-center">107</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/8603/">Витамин 404</a></td><td class="text-center">177</td><td class="text-center">66</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">58</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">225</td><td class="text-center">51</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5118/">Опять двойка</a></td><td class="text-center">152</td><td class="text-center">43</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучший процент взятия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Процент взятия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">80.4 %</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/24679/">Файервольные каменщики</a></td><td class="text-center">61</td><td class="text-center">76.0 %</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">252</td><td class="text-center">73.4 %</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5080/">Незнайки на Луне</a></td><td class="text-center">43</td><td class="text-center">71.2 %</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/32079/">Мычание ягнят</a></td><td class="text-center">18</td><td class="text-center">70.9 %</td></tr></tbody></table></div><small><sup>*</sup> - среди команд сыгравших больше 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучшая серия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">225</td><td class="text-center">35</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/5080/">Незнайки на Луне</a></td><td class="text-center">43</td><td class="text-center">28</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/29121/">Во всём виноват Тиндер</a></td><td class="text-center">56</td><td class="text-center">28</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">27</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5088/">НЛО</a></td><td class="text-center">86</td><td class="text-center">26</td></tr></tbody></table></div><small><sup>*</sup> - количество правильных ответов подряд</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">лучшая худшая серия<sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/24679/">Файервольные каменщики</a></td><td class="text-center">61</td><td class="text-center">5</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/8999/">Так получилось</a></td><td class="text-center">51</td><td class="text-center">5</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">6</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/8178/">The Hateful Seven</a></td><td class="text-center">46</td><td class="text-center">6</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5080/">Незнайки на Луне</a></td><td class="text-center">43</td><td class="text-center">6</td></tr></tbody></table></div><small><sup>*</sup> - количество неправильных ответов подряд, среди команд сыгравших минимум 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Тоталов в турах <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Тоталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/7101/">iSteria</a></td><td class="text-center">252</td><td class="text-center">34</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/5122/">Повторники</a></td><td class="text-center">86</td><td class="text-center">23</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">225</td><td class="text-center">17</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5142/">Power Of Sport</a></td><td class="text-center">99</td><td class="text-center">16</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/7115/">Тритий лишний</a></td><td class="text-center">196</td><td class="text-center">13</td></tr></tbody></table></div><small><sup>*</sup> - тотал - это тур в котором, команда ответила верно на все вопросы</small></div></div></div><div aria-labelledby="seasons-tab" class="tab-pane fade" id="seasons" role="tabpanel"><div class=""><h4 class="text-center dark-blue">Сезоны</h4><div class="table-responsive pb-4"><table class="table table-hover table-sm"><thead><tr class="active"><th class="text-nowrap">Сезон</th><th class="text-nowrap text-center">Игр</th><th class="text-nowrap">Начало</th><th class="text-nowrap">Окончание</th></tr></thead><tbody><tr><td class="text-nowrap"><a class="table_a" href="/season/2333/">Осень 2023</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">05 Сентябрь 2023</td><td class="text-nowrap">28 Ноябрь 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/2223/">Лето 2023</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">06 Июнь 2023</td><td class="text-nowrap">29 Август 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/2091/">Весна 2023</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">14 Март 2023</td><td class="text-nowrap">30 Май 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1927/">Зима-22/23</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">29 Ноябрь 2022</td><td class="text-nowrap">28 Февраль 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1814/">Осень 2022</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">29 Август 2022</td><td class="text-nowrap">21 Ноябрь 2022</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1739/">Лето 2022</a></td><td class="text-nowrap text-center">10</td><td class="text-nowrap">06 Июнь 2022</td><td class="text-nowrap">22 Август 2022</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1624/">Весна 2022</a></td><td class="text-nowrap text-center">10</td><td class="text-nowrap">07 Март 2022</td><td class="text-nowrap">30 Май 2022</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1506/">Зима 2022</a></td><td class="text-nowrap text-center">11</td><td class="text-nowrap">06 Декабрь 2021</td><td class="text-nowrap">28 Февраль 2022</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1393/">Осень 2021</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">06 Сентябрь 2021</td><td class="text-nowrap">29 Ноябрь 2021</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1308/">Лето 2021</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">07 Июнь 2021</td><td class="text-nowrap">30 Август 2021</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1198/">Весна 2021</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">15 Март 2021</td><td class="text-nowrap">31 Май 2021</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1131/">Зима 2021</a></td><td class="text-nowrap text-center">6</td><td class="text-nowrap">18 Январь 2021</td><td class="text-nowrap">01 Март 2021</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1031/">Поздняя осень 2020</a></td><td class="text-nowrap text-center">7</td><td class="text-nowrap">12 Октябрь 2020</td><td class="text-nowrap">23 Ноябрь 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/969/">Блиц-сезон:Евротур</a></td><td class="text-nowrap text-center">4</td><td class="text-nowrap">14 Сентябрь 2020</td><td class="text-nowrap">05 Октябрь 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/949/">Блиц-сезон:Космическая Одиссея</a></td><td class="text-nowrap text-center">4</td><td class="text-nowrap">10 Август 2020</td><td class="text-nowrap">31 Август 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/741/">Весна 2020</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">10 Март 2020</td><td class="text-nowrap">26 Май 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/614/">Зима 2020</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">03 Декабрь 2019</td><td class="text-nowrap">03 Март 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/536/">Осень 2019</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">10 Сентябрь 2019</td><td class="text-nowrap">26 Ноябрь 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/420/">Лето 2019</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">04 Июнь 2019</td><td class="text-nowrap">27 Август 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/364/">Весна 2019</a></td><td class="text-nowrap text-center">8</td><td class="text-nowrap">15 Апрель 2019</td><td class="text-nowrap">28 Май 2019</td></tr></tbody></table></div></div></div></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12 my-n3"><div class="d-flex flex-row-reverse pt-1"><h4><a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="mailto:ab@madrobit.com">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: '3wOOZ1gn6SyFhdTuIRr6eRxHDxwROF4Ghuct3s2EDfiaCxWxVjDm6Qd5up1Y36sY',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><script type="text/javascript">
	
	$('a[data-toggle="tab"]').on('shown.bs.tab', function (e) {
		var hash = $(e.target).attr('href');
		if (history.pushState) {
			history.pushState(null, null, hash);
		} else {
			location.hash = hash;
		}
	});

	jQuery(document).ready(function ($) {
		let selectedTab = window.location.hash;
		$('.nav-link[href="' + selectedTab + '"]' ).trigger('click');
	})
</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"></body></html>`
	html18615 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Первая лига | Игра #4 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Весна 2023 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Первая лига | Игра #4 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Весна 2023 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/18615/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/18615/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/118/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/118/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Первая лига | Игра #4</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/118/">Первая лига</a> / <a href="/season/2091/">Весна 2023</a> / 4 апреля 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/18615/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/18615/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'H4v1xvPw9mGv17Lv3cTlFjpv7c9kLCG7V2TGBWBNGJq0mrOygE5Bxi5TY4Er034p',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.5.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.5.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.5.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.5.1"></script><script src="/static/js/FileSaver.min.js?v=1.5.1"></script><script src="/static/js/tableexport.min.js?v=1.5.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'H4v1xvPw9mGv17Lv3cTlFjpv7c9kLCG7V2TGBWBNGJq0mrOygE5Bxi5TY4Er034p',
            game_id: '18615',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Первая лига | Игра #4',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html18616 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Первая лига | Игра #5 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Весна 2023 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Первая лига | Игра #5 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Весна 2023 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/18616/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/18616/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/118/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/118/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Первая лига | Игра #5</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/118/">Первая лига</a> / <a href="/season/2091/">Весна 2023</a> / 11 апреля 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/18616/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/18616/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'JVWMJHKC6i1NqHLlrurrvw59tKdHGATUXTkrN8wTDFLiL1OoEWDHnvLxkCIOV1hc',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.5.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.5.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.5.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.5.1"></script><script src="/static/js/FileSaver.min.js?v=1.5.1"></script><script src="/static/js/tableexport.min.js?v=1.5.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'JVWMJHKC6i1NqHLlrurrvw59tKdHGATUXTkrN8wTDFLiL1OoEWDHnvLxkCIOV1hc',
            game_id: '18616',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Первая лига | Игра #5',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html18617 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Первая лига | Игра #6 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Весна 2023 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Первая лига | Игра #6 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Весна 2023 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/18617/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/18617/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/118/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/118/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Нур-Султан')">Нур-Султан</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Первая лига | Игра #6</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/118/">Первая лига</a> / <a href="/season/2091/">Весна 2023</a> / 18 апреля 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/18617/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/18617/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"><a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="https://madrobit.com" rel="noopener" target="_blank">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'bs5X3Vw0YDiNttD89mxnCbxoGz4YWuP4pqtC7mihv02iONGbmOJDuadMxrz5bVdm',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.5.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.5.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.5.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.5.1"></script><script src="/static/js/FileSaver.min.js?v=1.5.1"></script><script src="/static/js/tableexport.min.js?v=1.5.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'bs5X3Vw0YDiNttD89mxnCbxoGz4YWuP4pqtC7mihv02iONGbmOJDuadMxrz5bVdm',
            game_id: '18617',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Первая лига | Игра #6',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html21281 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Первая лига | Игра #8 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Осень 2023 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Первая лига | Игра #8 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Осень 2023 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><meta http-equiv="origin-trial" content="AymqwRC7u88Y4JPvfIF2F37QKylC04248hLCdJAsh8xgOfe/dVJPV3XS3wLFca1ZMVOtnBfVjaCMTVudWM//5g4AAAB7eyJvcmlnaW4iOiJodHRwczovL3d3dy5nb29nbGV0YWdtYW5hZ2VyLmNvbTo0NDMiLCJmZWF0dXJlIjoiUHJpdmFjeVNhbmRib3hBZHNBUElzIiwiZXhwaXJ5IjoxNjk1MTY3OTk5LCJpc1RoaXJkUGFydHkiOnRydWV9"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/21281/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/21281/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/118/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/118/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Первая лига | Игра #8</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/118/">Первая лига</a> / <a href="/season/2333/">Осень 2023</a> / 31 октября 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/21281/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/21281/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="mailto:ab@madrobit.com">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'BHH5zdxWsmcmq9i7ELkTDFBKtp67afGJPF5KDEjdZJWRLtlaRdw9vEh8khBepG41',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.5.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.5.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.5.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.5.1"></script><script src="/static/js/FileSaver.min.js?v=1.5.1"></script><script src="/static/js/tableexport.min.js?v=1.5.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'BHH5zdxWsmcmq9i7ELkTDFBKtp67afGJPF5KDEjdZJWRLtlaRdw9vEh8khBepG41',
            game_id: '21281',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Первая лига | Игра #8',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html21282 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Первая лига | Игра #9 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Осень 2023 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Первая лига | Игра #9 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Осень 2023 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style><meta http-equiv="origin-trial" content="AymqwRC7u88Y4JPvfIF2F37QKylC04248hLCdJAsh8xgOfe/dVJPV3XS3wLFca1ZMVOtnBfVjaCMTVudWM//5g4AAAB7eyJvcmlnaW4iOiJodHRwczovL3d3dy5nb29nbGV0YWdtYW5hZ2VyLmNvbTo0NDMiLCJmZWF0dXJlIjoiUHJpdmFjeVNhbmRib3hBZHNBUElzIiwiZXhwaXJ5IjoxNjk1MTY3OTk5LCJpc1RoaXJkUGFydHkiOnRydWV9"></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/21282/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/21282/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/118/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/118/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Первая лига | Игра #9</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/118/">Первая лига</a> / <a href="/season/2333/">Осень 2023</a> / 7 ноября 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/21282/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/21282/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="mailto:ab@madrobit.com">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: '4rc9tRqI269poSdEf2nrxPaMFZTKMeAVipAOxicZztTUJcgHsuzHpOQawRoR1FYd',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.5.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.5.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.5.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.5.1"></script><script src="/static/js/FileSaver.min.js?v=1.5.1"></script><script src="/static/js/tableexport.min.js?v=1.5.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: '4rc9tRqI269poSdEf2nrxPaMFZTKMeAVipAOxicZztTUJcgHsuzHpOQawRoR1FYd',
            game_id: '21282',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Первая лига | Игра #9',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html21283 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Первая лига | Игра #10 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Осень 2023 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Первая лига | Игра #10 | Первая лига (Санкт-Петербург, Клуб «60 Секунд») | Осень 2023 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style><meta http-equiv="origin-trial" content="AymqwRC7u88Y4JPvfIF2F37QKylC04248hLCdJAsh8xgOfe/dVJPV3XS3wLFca1ZMVOtnBfVjaCMTVudWM//5g4AAAB7eyJvcmlnaW4iOiJodHRwczovL3d3dy5nb29nbGV0YWdtYW5hZ2VyLmNvbTo0NDMiLCJmZWF0dXJlIjoiUHJpdmFjeVNhbmRib3hBZHNBUElzIiwiZXhwaXJ5IjoxNjk1MTY3OTk5LCJpc1RoaXJkUGFydHkiOnRydWV9"></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/21283/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/21283/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/118/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/118/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Первая лига | Игра #10</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/118/">Первая лига</a> / <a href="/season/2333/">Осень 2023</a> / 14 ноября 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/21283/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/21283/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"></div></div></div><div class="col-md-6 small"><p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a></p></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="mailto:ab@madrobit.com">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'YGINM6iL2UxT5wt1zCZXkPYamZAHzvmMcE6sQx42zhhoqQw4M4bdcOEydR5OOWK4',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.5.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.5.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.5.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.5.1"></script><script src="/static/js/FileSaver.min.js?v=1.5.1"></script><script src="/static/js/tableexport.min.js?v=1.5.1"></script><script type="text/javascript">

const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};

function get_game_results(key) {
    $.ajax({
          url: "/get_game_results/", // the endpoint
          type: "GET", // http method
          data: {
            csrfmiddlewaretoken: 'YGINM6iL2UxT5wt1zCZXkPYamZAHzvmMcE6sQx42zhhoqQw4M4bdcOEydR5OOWK4',
            game_id: '21283',
            key: key,
          },
          async: false,
          success: function (json) {
                if(json.teams > 0){
                    createtable(json, game_settings, key);
                    $("#jsonTable_"+key).tablesorter({
                        sortList : [[0,0], [2,0]]
                    });
                };
          },
          error: function (xhr, errmsg, err) {
            alertify.set('notifier','position', 'top-right');
            alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
          }
    });
};


function toogletur(tur_id) {
    event.stopPropagation();
    $(".table ."+tur_id+"_tur").toggleClass("d-none");
    $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
};

$(function () {
  $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
});


$(document).ready(function () {
    if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
    {
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    } else {
        get_game_results(1000);
    };

    var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
    var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      var ctx = document.getElementById('bar-chart').getContext('2d');
      var barChart = new Chart(ctx, {
          type: 'bar',
          data: {
              datasets: [{
                  label: 'Правильных ответов',
                  yAxisID: 'A',
                  backgroundColor: '#0096E4',
                  data: graph,
              }],
              labels: labels,
          },
          options: {
            aspectRatio: 2.5,
            title: {
                display: true,
                fontColor: '#090438',
                text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
            },
            legend: {
              labels: {
                  fontColor: '#090438'
                }
              },
            scales: {
              xAxes: [{
                 ticks: {
                    fontColor: "#090438",
                 }
              }],
              yAxes: [{
                id: 'A',
                type: 'linear',
                position: 'left',
                ticks: {
                  //max: 40,
                  min: 0,
                  fontColor: "#090438",
                }
              }]
            }
          }
      });
      $(function () {
  $('[data-toggle="tooltip"]').tooltip()
})
});




function QRS() {
    //iterate through each row in the table
    $('.qrs').each(function(){
        $(this).fadeToggle();
    });
};

function divshow() {
    if (game_settings.divisions_count < 2){return false};
    $("#main_content").empty();
    if (Cookies.get('divs') == 'True')
    {
        Cookies.set('divs', 'False', { expires: 10 });
        get_game_results(1000); 
    } else {
        Cookies.set('divs', 'True', { expires: 10 });
        $.each(game_settings.divisions, function(key, value) {
            get_game_results(key);  
        });
    }
};

function xlsdownload() {
    $("table").each(function(){
        TableExport.prototype.charset = "charset=utf-8";

        if (this.id.substring(0,4) == 'json'){
            var tableid = this.id
            var exporttable = $(this).tableExport({
                formats: ['csv'],
                exportButtons: false,
                filename: 'Первая лига | Игра #10',
                separator: ',',
                mimeType: 'text/csv',
                fileExtension: '.csv',
                enforceStrictRFC4180: true
            });
            
        var exportData = exporttable.getExportData();
        var xlsxData = exportData[tableid].csv; 
        exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
        }
    });
};

</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html5 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Клуб «60 секунд»</title><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><meta http-equiv="origin-trial" content="AymqwRC7u88Y4JPvfIF2F37QKylC04248hLCdJAsh8xgOfe/dVJPV3XS3wLFca1ZMVOtnBfVjaCMTVudWM//5g4AAAB7eyJvcmlnaW4iOiJodHRwczovL3d3dy5nb29nbGV0YWdtYW5hZ2VyLmNvbTo0NDMiLCJmZWF0dXJlIjoiUHJpdmFjeVNhbmRib3hBZHNBUElzIiwiZXhwaXJ5IjoxNjk1MTY3OTk5LCJpc1RoaXJkUGFydHkiOnRydWV9"></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/league/117/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/league/117/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12" style="margin-top: -20px;"><div class="d-flex flex-row-reverse"><h4><a class="blue" href="None" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="blue" href="None" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-md-6 col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига</h2><div class="white-razdel"></div><div class="l-blue pt-2">Клуб «60 Секунд» / Санкт-Петербург</div></div></div></div><div class="col-md-6 col-12"><h5 class="yellow text-center">Ближайшая игра</h5><h6 class="l-grey text-center"><a href="/game/22678/">Открытая лига | Игра #1</a></h6><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="l-blue ff-ibc vertical-text align-middle"><span>Играем</span></div><div><table><tbody><tr><td class="l-grey"><i class="fas fa-calendar-alt fa-fw l-blue"></i>&nbsp;04 декабря, Понедельник</td></tr><tr><td class="l-grey"><i class="fas fa-clock fa-fw l-blue"></i>&nbsp;19:30</td></tr><tr><td class="l-grey"><i class="fas fa-map-marker-alt fa-fw l-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="l-grey"><i class="fas fa-credit-card l-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2695369/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0 mt-md-n4"><ul class="nav nav-pills justify-content-md-start justify-content-center" id="myTab" role="tablist"><li class="nav-item nav-pills-yellow"><a aria-controls="schedule" aria-selected="true" class="nav-link tab-link active" data-toggle="tab" href="#schedule" id="schedule-tab" role="tab"><span>Лига</span></a></li><li class="nav-item nav-pills-yellow"><a aria-controls="seasons" aria-selected="true" class="nav-link tab-link" data-toggle="tab" href="#seasons" id="seasons-tab" role="tab"><span>Все сезоны</span></a></li><li class="nav-item nav-pills-yellow"><a class="nav-link tab-link" href="/schedule/71/"><span>Расписание клуба</span></a></li></ul></div></div></div></div><div class="container"><div class="row pt-4"><div class="col-12"><div class="tab-content" id="myTabContent"><div aria-labelledby="schedule-tab" class="tab-pane fade show active" id="schedule" role="tabpanel"><div class="row"><div class="col-12"><h4 class="text-center dark-blue">Ближайшие игры лиги</h4></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/22678/">Открытая лига | Игра #1</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;04 декабря, Понедельник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;400 руб. с человека</td></tr><tr><td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp; <a class="" href="https://club60sec.timepad.ru/event/2695369/" target="_blank">Купить билет</a></td></tr></tbody></table></div></div></div></div><div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3"><div class="container py-3" style="border: 1px solid #E4E4E4;"><h5 class="dark-blue text-center"><a class="" href="/game/22679/">Открытая лига | Игра #2</a></h5><div class="col-12 d-flex flex-row justify-content-center align-items-center"><div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div><div><table><tbody><tr><td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;18 декабря, Понедельник</td></tr><tr><td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td></tr><tr><td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td></tr><tr><td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>&nbsp;400 руб. с человека</td></tr></tbody></table></div></div></div></div></div><div class="row"><div class="col-12 pt-2"><h4 class="dark-blue text-center">Статистика</h4><p><small>Ниже представлены топ-5 команд в некоторых номинациях. Обратите внимание, что это общая статистика по всем играм. Кроме общей статистики, есть расширенная статистика для каждого из сезонов. Например, вот статистика <a class="table_a" href="/season/2462/#stats">текущего</a> и <a class="table_a" href="/season/2207/#stats">предыдущего</a> сезонов. Кликнув на любую из команд лиги - вы попадете на страницу этой команды, где сможете ознакомится с её подробной статистикой, как за всё время, так и за каждый сезон в отдельности. Команды можно искать с помощью поиска, а также из списка всех <a class="table_a" href="/teams/117/">команд лиги</a>.</small></p></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Больше всех сыграли</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">134</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/10164/">Мамонт на Невском</a></td><td class="text-center">108</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5099/">Газ до отказа</a></td><td class="text-center">87</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Побед в сезонах</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">134</td><td class="text-center">1</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/6380/">Статистически незначимые</a></td><td class="text-center">25</td><td class="text-center">1</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/36839/">Per aspera ad aspera!</a></td><td class="text-center">21</td><td class="text-center">1</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/10164/">Мамонт на Невском</a></td><td class="text-center">108</td><td class="text-center">0</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td><td class="text-center">0</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Победители</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Побед</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">134</td><td class="text-center">12</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/36839/">Per aspera ad aspera!</a></td><td class="text-center">21</td><td class="text-center">10</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">7</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/32077/">ПИКсики</a></td><td class="text-center">17</td><td class="text-center">4</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/30782/">О, Спорт!</a></td><td class="text-center">41</td><td class="text-center">4</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Призёры</h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Пьедесталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">23</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">134</td><td class="text-center">21</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/36839/">Per aspera ad aspera!</a></td><td class="text-center">21</td><td class="text-center">13</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/30782/">О, Спорт!</a></td><td class="text-center">41</td><td class="text-center">12</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/32077/">ПИКсики</a></td><td class="text-center">17</td><td class="text-center">10</td></tr></tbody></table></div></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучший процент взятия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Процент взятия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/32077/">ПИКсики</a></td><td class="text-center">17</td><td class="text-center">76.1 %</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/36839/">Per aspera ad aspera!</a></td><td class="text-center">21</td><td class="text-center">74.9 %</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6429/">ИТ-ГРАД</a></td><td class="text-center">26</td><td class="text-center">72.5 %</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">68.7 %</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/5086/">Дискриминант добра</a></td><td class="text-center">23</td><td class="text-center">66.4 %</td></tr></tbody></table></div><small><sup>*</sup> - среди команд сыгравших больше 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Лучшая серия <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/30782/">О, Спорт!</a></td><td class="text-center">41</td><td class="text-center">26</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/36839/">Per aspera ad aspera!</a></td><td class="text-center">21</td><td class="text-center">19</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">19</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/9807/">КЛЕВЕР</a></td><td class="text-center">100</td><td class="text-center">18</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/6392/">Бред Стайлс</a></td><td class="text-center">22</td><td class="text-center">16</td></tr></tbody></table></div><small><sup>*</sup> - количество правильных ответов подряд</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">лучшая худшая серия<sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Серия</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">5</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/36839/">Per aspera ad aspera!</a></td><td class="text-center">21</td><td class="text-center">5</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/6375/">Хунвейбины за 40</a></td><td class="text-center">38</td><td class="text-center">6</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/6380/">Статистически незначимые</a></td><td class="text-center">25</td><td class="text-center">6</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/6377/">Команда номер 8</a></td><td class="text-center">19</td><td class="text-center">6</td></tr></tbody></table></div><small><sup>*</sup> - количество неправильных ответов подряд, среди команд сыгравших минимум 15 игр</small></div><div class="col-sm-6 col-lg-4 col-12 pt-2"><h6 class="dark-blue text-center">Тоталов в турах <sup>*</sup></h6><div class="table-responsive"><table class="table table-hover table-sm small"><thead><tr class="active"><th class="text-right">#</th><th>Команда</th><th class="text-center">Игр</th><th class="text-center">Тоталов</th></tr></thead><tbody><tr><th class="text-right">1</th><td><a class="table_a" href="/team/5089/">Стыд и скрам</a></td><td class="text-center">61</td><td class="text-center">6</td></tr><tr><th class="text-right">2</th><td><a class="table_a" href="/team/8941/">МеГаМозгоВиты</a></td><td class="text-center">88</td><td class="text-center">5</td></tr><tr><th class="text-right">3</th><td><a class="table_a" href="/team/16961/">Планы на ветер</a></td><td class="text-center">134</td><td class="text-center">3</td></tr><tr><th class="text-right">4</th><td><a class="table_a" href="/team/5073/">Невский берег - 2</a></td><td class="text-center">38</td><td class="text-center">3</td></tr><tr><th class="text-right">5</th><td><a class="table_a" href="/team/32077/">ПИКсики</a></td><td class="text-center">17</td><td class="text-center">3</td></tr></tbody></table></div><small><sup>*</sup> - тотал - это тур в котором, команда ответила верно на все вопросы</small></div></div></div><div aria-labelledby="seasons-tab" class="tab-pane fade" id="seasons" role="tabpanel"><div class=""><h4 class="text-center dark-blue">Сезоны</h4><div class="table-responsive pb-4"><table class="table table-hover table-sm"><thead><tr class="active"><th class="text-nowrap">Сезон</th><th class="text-nowrap text-center">Игр</th><th class="text-nowrap">Начало</th><th class="text-nowrap">Окончание</th></tr></thead><tbody><tr><td class="text-nowrap"><a class="table_a" href="/season/2462/">Зима-Весна 23/24</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">04 Декабрь 2023</td><td class="text-nowrap">18 Декабрь 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/2207/">Лето-Осень 2023</a></td><td class="text-nowrap text-center">21</td><td class="text-nowrap">05 Июнь 2023</td><td class="text-nowrap">20 Ноябрь 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/2090/">Весна 2023</a></td><td class="text-nowrap text-center">11</td><td class="text-nowrap">13 Март 2023</td><td class="text-nowrap">29 Май 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/1916/">Зима-22/23</a></td><td class="text-nowrap text-center">11</td><td class="text-nowrap">28 Ноябрь 2022</td><td class="text-nowrap">27 Февраль 2023</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/740/">Весна 2020</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">09 Март 2020</td><td class="text-nowrap">25 Май 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/613/">Зима 2020</a></td><td class="text-nowrap text-center">12</td><td class="text-nowrap">02 Декабрь 2019</td><td class="text-nowrap">25 Март 2020</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/379/">Синхроны</a></td><td class="text-nowrap text-center">2</td><td class="text-nowrap">09 Апрель 2019</td><td class="text-nowrap">23 Декабрь 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/532/">Осень 2019</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">09 Сентябрь 2019</td><td class="text-nowrap">25 Ноябрь 2019</td></tr><tr><td class="text-nowrap"><a class="table_a" href="/season/363/">Весна-лето 2019</a></td><td class="text-nowrap text-center">13</td><td class="text-nowrap">06 Май 2019</td><td class="text-nowrap">26 Август 2019</td></tr></tbody></table></div></div></div></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12 my-n3"><div class="d-flex flex-row-reverse pt-1"><h4><a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-telegram-plane"></i></a>&nbsp;&nbsp;&nbsp; <a class="no-bb foot-link" href="None" rel="noopener" target="_blank"><i class="fab fa-vk"></i></a>&nbsp;&nbsp;&nbsp;</h4></div></div><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"></div></div></div><div class="col-md-6 small"></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="mailto:ab@madrobit.com">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
    setTimeout(function(){
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    }, 5000);

    $(window).on("load", function (e) {
    // Animate loader off screen
    
    $(".single").hide();
    

    if (Cookies.get('league') != undefined)
    {
      $('#city').text(htmlDecode(''));
      $('#league_name').text(htmlDecode(''));
      $('.modal-city').text(htmlDecode(''));
      
    }
    else{
      set_city(71, 'Санкт-Петербург');
    };

    $.fn.cornerpopup({
      variant: 2
    });
    $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
  });
    function toggleOn() {
      $(".online .modal-leagues").toggle();
      $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleOff() {
      $(".offline .modal-leagues").toggle();
      $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    function toggleArchive() {
      $(".archive .modal-leagues").toggle();
      $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
    };
    
    function htmlDecode(value) {
      return $("<textarea/>").html(value).text();
    };
    function set_city(club_id, city) {

      $(".modal-leagues").empty();
      $(".archive").hide();
      $.ajax({
        url: "/getclubleagues/", // the endpoint
        type: "POST", // http method
        data: {
          club_id: club_id,
          csrfmiddlewaretoken: 'iROAgP4HbI27u3lLxWRsxbLcQo6lYSWmMJLnUKErxLIoytmNTMrmq05XOzCsNNN5',
        }, // data sent with the post request
        // handle a successful response
        success: function (json) {
         if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
          $(".single").hide();
          $(".multi").show();

          if (Object.keys(json.off).length > 0){
            $.each(json.off, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".offline .modal-leagues").append(div);
            });
            $(".offline").show();
            if (club_id == 25) {
              $(".offline .modal-leagues").hide();
            }
            
          } else{
            $(".offline").hide();
          }
          if (Object.keys(json.on).length > 0){
            $.each(json.on, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".online .modal-leagues").append(div);
            });
            $(".online").show();
            if (club_id == 25) {
              $(".online .modal-leagues").hide();
            }
            $(".offline .off-title").show();
              //$(".offline .modal-leagues").hide();
            } else {
              $(".online").hide();
              //$(".offline .modal-leagues").show();
              $(".offline .off-title").hide();
            }
          }
          else {
            $(".modal-city-yes").removeAttr('onclick');
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            } else {
              $.each(json.on, function(index, element) {
                $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
              });
            };
            $(".multi").hide();
            $(".single").show();
            $(".online").hide();
            $(".offline").hide();
          };
          if (Object.keys(json.archive).length > 0){
            $.each(json.archive, function(index, element) {
              var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
              $(".archive .modal-leagues").append(div);
            });
            $(".archive").show();
            $(".archive .modal-leagues").hide();
          } else {
            $(".archive").hide();
          };

        },
        // handle a non-successful response
        error: function (xhr, errmsg, err) {
          $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
            console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
          }
        });

      $('.modal-city').text(city);
      $('#leagues').slideDown(200);
      $('#cityes').fadeOut(200);

    };

    function set_league(league_id) {
      Cookies.set('league', league_id, { expires: 90 });
      document.location.href="/";
    };

    function select_city() {
      $('#leagues').fadeOut(200);
      $('#cityes').slideDown(200);
    };

    function cityFilter() {
  // Declare variables
  var input, filter, table, tr, td, i;
  input = document.getElementById("cityInput");
  filter = input.value.toUpperCase();

  $('.modal-a-div').each(function() {
    if ($(this).text().toUpperCase().indexOf(filter) > -1) {
      $(this).show();
    } else {
      $(this).css("display","none");
    }
  });

};
$(function () {
  $(document).scroll(function () {
    var $nav = $(".navbar");
    var logo_img = $("#logo-img");
    //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
    $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
    logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
  });
});
</script><script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'UA-104995998-1');
</script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><script type="text/javascript">
	
	$('a[data-toggle="tab"]').on('shown.bs.tab', function (e) {
		var hash = $(e.target).attr('href');
		if (history.pushState) {
			history.pushState(null, null, hash);
		} else {
			location.hash = hash;
		}
	});

	jQuery(document).ready(function ($) {
		let selectedTab = window.location.hash;
		$('.nav-link[href="' + selectedTab + '"]' ).trigger('click');
	})
</script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html22678 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига | Игра #1 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-Весна 23/24 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Открытая лига | Игра #1 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-Весна 23/24 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><style type="text/css">/* Chart.js */
  @keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style><meta http-equiv="origin-trial" content="AymqwRC7u88Y4JPvfIF2F37QKylC04248hLCdJAsh8xgOfe/dVJPV3XS3wLFca1ZMVOtnBfVjaCMTVudWM//5g4AAAB7eyJvcmlnaW4iOiJodHRwczovL3d3dy5nb29nbGV0YWdtYW5hZ2VyLmNvbTo0NDMiLCJmZWF0dXJlIjoiUHJpdmFjeVNhbmRib3hBZHNBUElzIiwiZXhwaXJ5IjoxNjk1MTY3OTk5LCJpc1RoaXJkUGFydHkiOnRydWV9"></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/22678/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/22678/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига | Игра #1</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/117/">Открытая лига</a> / <a href="/season/2462/">Зима-Весна 23/24</a> / 4 декабря 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/22678/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/22678/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"></div></div></div><div class="col-md-6 small"></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="mailto:ab@madrobit.com">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
      setTimeout(function(){
        $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
      }, 5000);
  
      $(window).on("load", function (e) {
      // Animate loader off screen
      
      $(".single").hide();
      
  
      if (Cookies.get('league') != undefined)
      {
        $('#city').text(htmlDecode(''));
        $('#league_name').text(htmlDecode(''));
        $('.modal-city').text(htmlDecode(''));
        
      }
      else{
        set_city(71, 'Санкт-Петербург');
      };
  
      $.fn.cornerpopup({
        variant: 2
      });
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    });
      function toggleOn() {
        $(".online .modal-leagues").toggle();
        $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
      };
      function toggleOff() {
        $(".offline .modal-leagues").toggle();
        $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
      };
      function toggleArchive() {
        $(".archive .modal-leagues").toggle();
        $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
      };
      
      function htmlDecode(value) {
        return $("<textarea/>").html(value).text();
      };
      function set_city(club_id, city) {
  
        $(".modal-leagues").empty();
        $(".archive").hide();
        $.ajax({
          url: "/getclubleagues/", // the endpoint
          type: "POST", // http method
          data: {
            club_id: club_id,
            csrfmiddlewaretoken: 'Pc9kEaF2N1ayqeVRa3IYjPMgIdjBEa7uj467i5fM94QPuEWTwTiScE61GoPIt5Yd',
          }, // data sent with the post request
          // handle a successful response
          success: function (json) {
           if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
            $(".single").hide();
            $(".multi").show();
  
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
                $(".offline .modal-leagues").append(div);
              });
              $(".offline").show();
              if (club_id == 25) {
                $(".offline .modal-leagues").hide();
              }
              
            } else{
              $(".offline").hide();
            }
            if (Object.keys(json.on).length > 0){
              $.each(json.on, function(index, element) {
                var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
                $(".online .modal-leagues").append(div);
              });
              $(".online").show();
              if (club_id == 25) {
                $(".online .modal-leagues").hide();
              }
              $(".offline .off-title").show();
                //$(".offline .modal-leagues").hide();
              } else {
                $(".online").hide();
                //$(".offline .modal-leagues").show();
                $(".offline .off-title").hide();
              }
            }
            else {
              $(".modal-city-yes").removeAttr('onclick');
              if (Object.keys(json.off).length > 0){
                $.each(json.off, function(index, element) {
                  $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
                });
              } else {
                $.each(json.on, function(index, element) {
                  $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
                });
              };
              $(".multi").hide();
              $(".single").show();
              $(".online").hide();
              $(".offline").hide();
            };
            if (Object.keys(json.archive).length > 0){
              $.each(json.archive, function(index, element) {
                var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
                $(".archive .modal-leagues").append(div);
              });
              $(".archive").show();
              $(".archive .modal-leagues").hide();
            } else {
              $(".archive").hide();
            };
  
          },
          // handle a non-successful response
          error: function (xhr, errmsg, err) {
            $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                  " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
              console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
            }
          });
  
        $('.modal-city').text(city);
        $('#leagues').slideDown(200);
        $('#cityes').fadeOut(200);
  
      };
  
      function set_league(league_id) {
        Cookies.set('league', league_id, { expires: 90 });
        document.location.href="/";
      };
  
      function select_city() {
        $('#leagues').fadeOut(200);
        $('#cityes').slideDown(200);
      };
  
      function cityFilter() {
    // Declare variables
    var input, filter, table, tr, td, i;
    input = document.getElementById("cityInput");
    filter = input.value.toUpperCase();
  
    $('.modal-a-div').each(function() {
      if ($(this).text().toUpperCase().indexOf(filter) > -1) {
        $(this).show();
      } else {
        $(this).css("display","none");
      }
    });
  
  };
  $(function () {
    $(document).scroll(function () {
      var $nav = $(".navbar");
      var logo_img = $("#logo-img");
      //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
      $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
      logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
    });
  });
  </script><script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());
    gtag('config', 'UA-104995998-1');
  </script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.5.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.5.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.5.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.5.1"></script><script src="/static/js/FileSaver.min.js?v=1.5.1"></script><script src="/static/js/tableexport.min.js?v=1.5.1"></script><script type="text/javascript">
  
  const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};
  
  function get_game_results(key) {
      $.ajax({
            url: "/get_game_results/", // the endpoint
            type: "GET", // http method
            data: {
              csrfmiddlewaretoken: 'Pc9kEaF2N1ayqeVRa3IYjPMgIdjBEa7uj467i5fM94QPuEWTwTiScE61GoPIt5Yd',
              game_id: '22678',
              key: key,
            },
            async: false,
            success: function (json) {
                  if(json.teams > 0){
                      createtable(json, game_settings, key);
                      $("#jsonTable_"+key).tablesorter({
                          sortList : [[0,0], [2,0]]
                      });
                  };
            },
            error: function (xhr, errmsg, err) {
              alertify.set('notifier','position', 'top-right');
              alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
            }
      });
  };
  
  
  function toogletur(tur_id) {
      event.stopPropagation();
      $(".table ."+tur_id+"_tur").toggleClass("d-none");
      $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
  };
  
  $(function () {
    $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
  });
  
  
  $(document).ready(function () {
      if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
      {
          $.each(game_settings.divisions, function(key, value) {
              get_game_results(key);  
          });
      } else {
          get_game_results(1000);
      };
  
      var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
      var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
        var ctx = document.getElementById('bar-chart').getContext('2d');
        var barChart = new Chart(ctx, {
            type: 'bar',
            data: {
                datasets: [{
                    label: 'Правильных ответов',
                    yAxisID: 'A',
                    backgroundColor: '#0096E4',
                    data: graph,
                }],
                labels: labels,
            },
            options: {
              aspectRatio: 2.5,
              title: {
                  display: true,
                  fontColor: '#090438',
                  text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
              },
              legend: {
                labels: {
                    fontColor: '#090438'
                  }
                },
              scales: {
                xAxes: [{
                   ticks: {
                      fontColor: "#090438",
                   }
                }],
                yAxes: [{
                  id: 'A',
                  type: 'linear',
                  position: 'left',
                  ticks: {
                    //max: 40,
                    min: 0,
                    fontColor: "#090438",
                  }
                }]
              }
            }
        });
        $(function () {
    $('[data-toggle="tooltip"]').tooltip()
  })
  });
  
  
  
  
  function QRS() {
      //iterate through each row in the table
      $('.qrs').each(function(){
          $(this).fadeToggle();
      });
  };
  
  function divshow() {
      if (game_settings.divisions_count < 2){return false};
      $("#main_content").empty();
      if (Cookies.get('divs') == 'True')
      {
          Cookies.set('divs', 'False', { expires: 10 });
          get_game_results(1000); 
      } else {
          Cookies.set('divs', 'True', { expires: 10 });
          $.each(game_settings.divisions, function(key, value) {
              get_game_results(key);  
          });
      }
  };
  
  function xlsdownload() {
      $("table").each(function(){
          TableExport.prototype.charset = "charset=utf-8";
  
          if (this.id.substring(0,4) == 'json'){
              var tableid = this.id
              var exporttable = $(this).tableExport({
                  formats: ['csv'],
                  exportButtons: false,
                  filename: 'Открытая лига | Игра #1',
                  separator: ',',
                  mimeType: 'text/csv',
                  fileExtension: '.csv',
                  enforceStrictRFC4180: true
              });
              
          var exportData = exporttable.getExportData();
          var xlsxData = exportData[tableid].csv; 
          exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
          }
      });
  };
  
  </script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
	html22679 = `<html lang="en"><head><meta charset="utf-8"><meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport"><meta content="" name="description"><meta content="madrobit" name="author"><title>Открытая лига | Игра #2 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-Весна 23/24 | Клуб «60 секунд»</title><meta content="Таблица игры сезона: Открытая лига | Игра #2 | Открытая лига (Санкт-Петербург, Клуб «60 Секунд») | Зима-Весна 23/24 | Клуб «60 секунд»" name="Description"><link href="/static/apple-touch-icon.png?v=1.5.1" rel="apple-touch-icon" sizes="180x180"><link href="/static/favicon-32x32.png?v=1.5.1" rel="icon" sizes="32x32" type="image/png"><link href="/static/favicon-16x16.png?v=1.5.1" rel="icon" sizes="16x16" type="image/png"><link href="/static/site.webmanifest" rel="manifest"><link color="#5bbad5" href="/static/safari-pinned-tab.svg?v=1.5.1" rel="mask-icon"><link href="/static/favicon.ico?v=1.5.1" rel="shortcut icon"><meta content="#2b5797" name="msapplication-TileColor"><meta content="/static/browserconfig.xml?v=1.5.1" name="msapplication-config"><meta content="#ffffff" name="theme-color"><meta content="https://60sec.online/static/img/snippet.jpg" property="og:image"><link href="/static/css/bootstrap.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/style.css?v=1.5.1" rel="stylesheet"><meta http-equiv="origin-trial" content="AymqwRC7u88Y4JPvfIF2F37QKylC04248hLCdJAsh8xgOfe/dVJPV3XS3wLFca1ZMVOtnBfVjaCMTVudWM//5g4AAAB7eyJvcmlnaW4iOiJodHRwczovL3d3dy5nb29nbGV0YWdtYW5hZ2VyLmNvbTo0NDMiLCJmZWF0dXJlIjoiUHJpdmFjeVNhbmRib3hBZHNBUElzIiwiZXhwaXJ5IjoxNjk1MTY3OTk5LCJpc1RoaXJkUGFydHkiOnRydWV9"><style type="text/css">/* Chart.js */
  @keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head><body><div class="loadingio-spinner-bean-eater-i7x2im0jb9" style="display: none;"><div class="int-packman-wrapper"><div class="ldio-yj1mc95c7go"><div><div></div><div></div><div></div></div><div><div></div><div></div><div></div></div></div></div></div><nav class="navbar navbar-expand-lg navbar-light"><div class="container"><div class="navbar-brand d-flex flex-row align-items-center"><div class="pr-2"><a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a></div><div class="d-flex flex-column" id="club-selector"><div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="league_name" id="league_name">Выбрать Лигу</span></a></div><div style="margin-top: -6px;"><a data-target="#cityModal" data-toggle="modal" href="#"><span class="city" id="city">Санкт-Петербург</span></a></div></div></div><div class="d-flex flex-row order-2 order-lg-3"><ul class="navbar-nav flex-row"><li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a><div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout"><form action="/search/" class="form-inline dropdown-item" id="form-id" method="get"><div class="input-group" style="min-width: 200px;"><input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="search" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск"><span class="input-group-append"><div class="input-group-text bg-transparent"><i class="fa fa-search" onclick="document.getElementById('form-id').submit();" style="cursor: pointer;"></i></div></span></div></form></div></li><li class="dropdown nav-item"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button"><img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24"></a><div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout"><a class="dropdown-item" href="/en/game/22679/"><img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English</a><a class="dropdown-item" href="/game/22679/"><img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский</a></div></li></ul><button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button></div><div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent"><ul class="navbar-nav mr-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/igroki/" style="color:#0096E4 !important;">ИГРОКИ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/faq/">FAQ</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/syncro/">Синхроны</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/teams/117/">Команды</a></li><li class="nav-item"><a class="nav-link nav-link-nav" href="/reglament/117/">Регламент</a></li></ul><ul class="navbar-nav ml-auto uppercase"><li class="nav-item"><a class="nav-link nav-link-nav" href="/login/">Вход</a></li></ul></div></div></nav><div aria-hidden="true" aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1"><div class="modal-dialog modal-dialog-centered modal-lg" role="document"><div class="modal-content"><div class="modal-body"><button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button><div class="container" id="leagues"><div class="row"><div class="mx-auto text-center dark-blue" id="modal_city"><p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? <span class="modal-city-inv multi" onclick="select_city()">Нет</span> <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p><div class="mx-auto text-center pb-2 single" style="display: none;"><span class="modal-city-yes" onclick="set_league(215)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div><p></p></div><div class="col-lg-6 col-12 text-center container pb-2 offline"><h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div></div></div><div class="col-lg-6 col-12 text-center container pb-1 online"><h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5><div class="container modal-leagues pb-3"><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div><div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div></div></div><div class="mx-auto text-center container pb-2 archive" style="display: none;"><h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-angle-down fa-xs" id="archive-icon"></i></h5><div class="container modal-leagues pb-3"></div></div></div></div><div class="container" id="cityes" style="display:none;"><h2 class="black pt-4">Наша география</h2><input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text"><br><div class="container pb-3 list-group d-flex flex-row flex-wrap"><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div><div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div></div></div></div></div></div></div><div id="header"><div class="container unskew"><div class="row"><div class="col-12 pb-2"><div class="card transparent"><div class="card-body" style="padding:0;"><h2 class="white">Открытая лига | Игра #2</h2><div class="white-razdel"></div><div class="l-blue pt-2">Санкт-Петербург / <a href="/league/117/">Открытая лига</a> / <a href="/season/2462/">Зима-Весна 23/24</a> / 18 декабря 2023 19:30</div></div></div></div><div class="col-md-6 col-12 py-3 pt-md-0"><a class="btn btn-yellow" href="/enterthegame/22679/" style="min-width: 250px;">Войти в игру</a></div></div></div></div><div class="container-fluid pt-3"><div class="row justify-content-center"><div class="col-12" style="margin-bottom: -38px; z-index: 10"><div class="d-flex flex-row-reverse"><div class="dropdown show text-right"><a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button"><h2><i class="fas fa-bars"></i></h2></a><div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right"><a class="dropdown-item" href="#" onclick="divshow()">По дивизионам</a> <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a> <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a> <a class="dropdown-item" href="/game/22679/show">Режим демонстрации</a></div></div></div></div><div class="collapse container-fluid" id="graph"><div class="row justify-content-center"><div class="col-12 col-lg-10 col-xl-8"><canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas></div></div></div><div class="col-12"><div class="pb-4" id="main_content"></div></div><div class="col-12"><div class="pb-3"></div></div></div></div><footer class="footer"><div class="footer_top mt-n3"></div><div class="container"><div class="row"><div class="col-12"><div class="d-flex justify-content-between"><div class="pb-3"><a class="no-bb" href="/"><img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" width="80"></a></div><div class="mt-4"></div></div></div><div class="col-md-6 small"></div><div class="col-md-6 text-right small"><p><a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a><br>powered by <a class="foot-link" href="mailto:ab@madrobit.com">madrobit.com</a></p></div></div></div></footer><link href="/static/fontawesome/css/all.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/alertify.min.css?v=1.5.1" rel="stylesheet"><link href="/static/css/corner-popup.min.css?v=1.5.1" rel="stylesheet"><link href="https://fonts.googleapis.com/icon?family=Material+Icons&amp;display=swap" rel="stylesheet"><script type="text/javascript" async="" src="https://www.google-analytics.com/analytics.js"></script><script type="text/javascript" async="" src="https://www.googletagmanager.com/gtag/js?id=G-B97GDESXDV&amp;l=dataLayer&amp;cx=c"></script><script src="/static/js/js.cookie.min.js?v=1.5.1"></script><script src="/static/js/jquery-3.4.1.min.js?v=1.5.1"></script><script src="/static/js/popper.min.js?v=1.5.1"></script><script src="/static/js/bootstrap.min.js?v=1.5.1"></script><script src="/static/js/corner-popup.js?v=1.5.1"></script><script src="/static/js/reconnecting-websocket.min.js?v=1.5.1"></script><script src="/static/js/alertify.min.js?v=1.5.1"></script><script type="text/javascript">
      setTimeout(function(){
        $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
      }, 5000);
  
      $(window).on("load", function (e) {
      // Animate loader off screen
      
      $(".single").hide();
      
  
      if (Cookies.get('league') != undefined)
      {
        $('#city').text(htmlDecode(''));
        $('#league_name').text(htmlDecode(''));
        $('.modal-city').text(htmlDecode(''));
        
      }
      else{
        set_city(71, 'Санкт-Петербург');
      };
  
      $.fn.cornerpopup({
        variant: 2
      });
      $(".loadingio-spinner-bean-eater-i7x2im0jb9").fadeOut("slow");
    });
      function toggleOn() {
        $(".online .modal-leagues").toggle();
        $("#on-icon").toggleClass("fa-angle-up fa-angle-down");
      };
      function toggleOff() {
        $(".offline .modal-leagues").toggle();
        $("#off-icon").toggleClass("fa-angle-up fa-angle-down");
      };
      function toggleArchive() {
        $(".archive .modal-leagues").toggle();
        $("#archive-icon").toggleClass("fa-angle-up fa-angle-down");
      };
      
      function htmlDecode(value) {
        return $("<textarea/>").html(value).text();
      };
      function set_city(club_id, city) {
  
        $(".modal-leagues").empty();
        $(".archive").hide();
        $.ajax({
          url: "/getclubleagues/", // the endpoint
          type: "POST", // http method
          data: {
            club_id: club_id,
            csrfmiddlewaretoken: 'FIGyz2BappuJP8FylNfuQo3cPPyZCsvs9ADldXbULsa0TyGAHDPoJdnXN046rnmb',
          }, // data sent with the post request
          // handle a successful response
          success: function (json) {
           if (Object.keys(json.on).length + Object.keys(json.off).length > 1){
            $(".single").hide();
            $(".multi").show();
  
            if (Object.keys(json.off).length > 0){
              $.each(json.off, function(index, element) {
                var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
                $(".offline .modal-leagues").append(div);
              });
              $(".offline").show();
              if (club_id == 25) {
                $(".offline .modal-leagues").hide();
              }
              
            } else{
              $(".offline").hide();
            }
            if (Object.keys(json.on).length > 0){
              $.each(json.on, function(index, element) {
                var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
                $(".online .modal-leagues").append(div);
              });
              $(".online").show();
              if (club_id == 25) {
                $(".online .modal-leagues").hide();
              }
              $(".offline .off-title").show();
                //$(".offline .modal-leagues").hide();
              } else {
                $(".online").hide();
                //$(".offline .modal-leagues").show();
                $(".offline .off-title").hide();
              }
            }
            else {
              $(".modal-city-yes").removeAttr('onclick');
              if (Object.keys(json.off).length > 0){
                $.each(json.off, function(index, element) {
                  $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
                });
              } else {
                $.each(json.on, function(index, element) {
                  $(".modal-city-yes").attr('onclick', 'set_league('+element+');');
                });
              };
              $(".multi").hide();
              $(".single").show();
              $(".online").hide();
              $(".offline").hide();
            };
            if (Object.keys(json.archive).length > 0){
              $.each(json.archive, function(index, element) {
                var div = '<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(' + element + ')">'+ index +'</a></div>'
                $(".archive .modal-leagues").append(div);
              });
              $(".archive").show();
              $(".archive .modal-leagues").hide();
            } else {
              $(".archive").hide();
            };
  
          },
          // handle a non-successful response
          error: function (xhr, errmsg, err) {
            $('#results').html("<div class='alert-box alert radius' data-alert>Oops! We have encountered an error: " + errmsg +
                  " <a href='#' class='close'>&times;</a></div>"); // add the error to the dom
              console.log(xhr.status + ": " + xhr.responseText); // provide a bit more info about the error to the console
            }
          });
  
        $('.modal-city').text(city);
        $('#leagues').slideDown(200);
        $('#cityes').fadeOut(200);
  
      };
  
      function set_league(league_id) {
        Cookies.set('league', league_id, { expires: 90 });
        document.location.href="/";
      };
  
      function select_city() {
        $('#leagues').fadeOut(200);
        $('#cityes').slideDown(200);
      };
  
      function cityFilter() {
    // Declare variables
    var input, filter, table, tr, td, i;
    input = document.getElementById("cityInput");
    filter = input.value.toUpperCase();
  
    $('.modal-a-div').each(function() {
      if ($(this).text().toUpperCase().indexOf(filter) > -1) {
        $(this).show();
      } else {
        $(this).css("display","none");
      }
    });
  
  };
  $(function () {
    $(document).scroll(function () {
      var $nav = $(".navbar");
      var logo_img = $("#logo-img");
      //$nav.toggleClass('', $(this).scrollTop() > $nav.height());
      $nav.toggleClass('sticky-top bg-white', $(this).scrollTop() > $nav.height());
      logo_img.toggleClass('logo-img', $(this).scrollTop() > $nav.height());
    });
  });
  </script><script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());
    gtag('config', 'UA-104995998-1');
  </script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-104995998-1"></script><link href="/static/css/chart.min.css?v=1.5.1" rel="stylesheet"><script src="/static/js/chart.bundle.min.js?v=1.5.1"></script><script src="/static/js/jquery.tablesorter.min.js?v=1.5.1" type="text/javascript"></script><script src="/static/js/gametable.js?v=1.5.1"></script><script src="/static/js/FileSaver.min.js?v=1.5.1"></script><script src="/static/js/tableexport.min.js?v=1.5.1"></script><script type="text/javascript">
  
  const game_settings = {"divisions": {"100": "\u0412\u043d\u0435 \u0437\u0430\u0447\u0451\u0442\u0430"}, "divisions_count": 1, "countmatrix": true, "countrating": false, "dontskipplaces": false, "tours": 3, "qlist": [12, 24, 36], "questions": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], "mediatours": [], "media_game": false, "matrix_game": false, "show_city": false, "language": "ru"};
  
  function get_game_results(key) {
      $.ajax({
            url: "/get_game_results/", // the endpoint
            type: "GET", // http method
            data: {
              csrfmiddlewaretoken: 'FIGyz2BappuJP8FylNfuQo3cPPyZCsvs9ADldXbULsa0TyGAHDPoJdnXN046rnmb',
              game_id: '22679',
              key: key,
            },
            async: false,
            success: function (json) {
                  if(json.teams > 0){
                      createtable(json, game_settings, key);
                      $("#jsonTable_"+key).tablesorter({
                          sortList : [[0,0], [2,0]]
                      });
                  };
            },
            error: function (xhr, errmsg, err) {
              alertify.set('notifier','position', 'top-right');
              alertify.error('Проблема с получением информации с сервера. Попробуйте обновить страницу позже.');
            }
      });
  };
  
  
  function toogletur(tur_id) {
      event.stopPropagation();
      $(".table ."+tur_id+"_tur").toggleClass("d-none");
      $("."+tur_id+"_but").toggleClass("fa-plus-square").toggleClass("fa-minus-square");
  };
  
  $(function () {
    $('[data-toggle="tooltip"]').tooltip({delay: { "show": 500, "hide": 200 }});
  });
  
  
  $(document).ready(function () {
      if ((Cookies.get('divs') == 'True') && (game_settings.divisions_count > 1))
      {
          $.each(game_settings.divisions, function(key, value) {
              get_game_results(key);  
          });
      } else {
          get_game_results(1000);
      };
  
      var labels = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36];
      var graph = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
        var ctx = document.getElementById('bar-chart').getContext('2d');
        var barChart = new Chart(ctx, {
            type: 'bar',
            data: {
                datasets: [{
                    label: 'Правильных ответов',
                    yAxisID: 'A',
                    backgroundColor: '#0096E4',
                    data: graph,
                }],
                labels: labels,
            },
            options: {
              aspectRatio: 2.5,
              title: {
                  display: true,
                  fontColor: '#090438',
                  text: "Команд в игре: 0 | взятие: 0.00% | гробов:  36 | спасённых:  0 | кнопок:  0 | дыр: 0."
              },
              legend: {
                labels: {
                    fontColor: '#090438'
                  }
                },
              scales: {
                xAxes: [{
                   ticks: {
                      fontColor: "#090438",
                   }
                }],
                yAxes: [{
                  id: 'A',
                  type: 'linear',
                  position: 'left',
                  ticks: {
                    //max: 40,
                    min: 0,
                    fontColor: "#090438",
                  }
                }]
              }
            }
        });
        $(function () {
    $('[data-toggle="tooltip"]').tooltip()
  })
  });
  
  
  
  
  function QRS() {
      //iterate through each row in the table
      $('.qrs').each(function(){
          $(this).fadeToggle();
      });
  };
  
  function divshow() {
      if (game_settings.divisions_count < 2){return false};
      $("#main_content").empty();
      if (Cookies.get('divs') == 'True')
      {
          Cookies.set('divs', 'False', { expires: 10 });
          get_game_results(1000); 
      } else {
          Cookies.set('divs', 'True', { expires: 10 });
          $.each(game_settings.divisions, function(key, value) {
              get_game_results(key);  
          });
      }
  };
  
  function xlsdownload() {
      $("table").each(function(){
          TableExport.prototype.charset = "charset=utf-8";
  
          if (this.id.substring(0,4) == 'json'){
              var tableid = this.id
              var exporttable = $(this).tableExport({
                  formats: ['csv'],
                  exportButtons: false,
                  filename: 'Открытая лига | Игра #2',
                  separator: ',',
                  mimeType: 'text/csv',
                  fileExtension: '.csv',
                  enforceStrictRFC4180: true
              });
              
          var exportData = exporttable.getExportData();
          var xlsxData = exportData[tableid].csv; 
          exporttable.export2file(xlsxData.data, xlsxData.mimeType, xlsxData.filename, xlsxData.fileExtension, xlsxData.merges);
          }
      });
  };
  
  </script><link href="https://www.google-analytics.com" rel="preconnect"><link href="https://stats.g.doubleclick.net" rel="preconnect"><div id="corner-popup" class="popup-xs" style="display: flex;"><div class="hide-mobile col sm-4"><img src="/static/img/cookie.png" class="corner-img-cookie responsive"></div><div class="col xs-12 sm-8"><div class="corner-close close-change"></div><div class="corner-container"><p class="corner-text">This website uses cookies to ensure you get the best experience on our website. <a href="https://en.wikipedia.org/wiki/HTTP_cookie" target="_blank" class="cookie-more">More information.</a></p><a class="corner-btn-cookie">Got it</a></div></div></div></body></html>`
)
