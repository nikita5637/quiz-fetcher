package sixty_seconds

const (
	html1 = `<html lang="ru"><head>
	<meta charset="UTF-8">
	<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
	<link rel="icon" href="/fav.png" type="image/png">
	<link rel="shortcut icon" href="/fav.png" type="image/png">
	<title>
	    Расписание игр
	    | Клуб «60 секунд»</title>
	<meta name="title" content="Расписание игр | Клуб «60 секунд»">
	<meta property="og:title" content="Расписание игр | Клуб «60 секунд»">
    
	<meta name="description" content="Расписание игр">
	<meta property="og:description" content="Расписание игр">
    
	<meta content="/static/img/snippet.jpg" property="og:image">
    
	<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
	<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
	<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
	<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
	<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
	<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
	<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
	<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
	<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
	<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
    </head>
    <body class="">
    <div class="loaderArea" style="display: none;">
      <div class="loader" style="display: none;">
	<div class="dot dot1"></div>
	<div class="dot dot2"></div>
	<div class="dot dot3"></div>
	<div class="dot dot4"></div>
      </div>
    </div>
    
    
    <nav class="navbar navbar-expand-lg navbar-light">
	    <div class="container">
		    <div class="navbar-brand d-flex flex-row align-items-center">
			    <div class="pr-2">
				    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
			    </div>
		
			    <div class="d-flex flex-column" id="club-selector">
				    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
					    <a data-target="#cityModal" data-toggle="modal" href="#">
						    <span class="league_name" id="league_name">Выбрать лигу</span>
					    </a>
				    </div>
				    <div style="margin-top: -6px;">
					    <a data-target="#cityModal" data-toggle="modal" href="#">
						    <span class="city" id="city">Санкт-Петербург</span>
					    </a>
				    </div>
			    </div>
	    
		    </div>
		    <div class="d-flex flex-row order-2 order-lg-3">
			    <ul class="navbar-nav flex-row">
		    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
		      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
			<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
			  <div class="input-group" style="min-width: 200px;">
			    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
			      <span class="input-group-append">
				  <div class="input-group-text bg-transparent">
				      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
				  </div>
			      </span>
			  </div>
			</form>
		      </div>
		    </li>
				    <li class="dropdown nav-item">
					    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
						    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
					    </a>
					    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
						    <a class="dropdown-item d-none" href="/en/league/119/">
							    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
						    </a>
						    <a class="dropdown-item" href="/league/119/">
							    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
						    </a>
					    </div>
				    </li>
			    </ul>
			    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
		    </div>
    
		    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
		
			    <ul class="navbar-nav mr-auto uppercase">
				    <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
				    </li>
				    <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
				    </li>
				    
				    <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/117/">Команды</a>
				    </li>
				    
				    
				    <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/reglament/117/">Регламент</a>
				    </li>
				    
			    </ul>
    
			    <ul class="navbar-nav ml-auto uppercase">
				    
					      <li class="nav-item">
						    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
					      </li>
				    
    
			    </ul>
	    
		    </div>
    
	    </div>
    </nav>
    
    
    <div id="header">
	    <div class="container py-3">
		    <div class="row">
			    <div class="col-md-6 col-12 pb-1">
				    <div class="card transparent">
					    <div class="card-body" style="padding:0;">
						    <h2 class="white">Расписание игр</h2>
						    <div class="white-razdel"></div>
						    <div class="l-blue pt-2">Клуб «60 Секунд» / Санкт-Петербург</div>
					    </div>
				    </div>
		</div>
		    </div>
	    </div>
    </div>
    
    <div class="container">
	    <div class="row pt-4">
		    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25678">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25678/" class="game_application_name">Открытая лига | Игра #3</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">17 June, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
					<tr>
					    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
						<a class="" href="https://club60sec.timepad.ru/event/2904719/" target="_blank">Купить билет</a>
					    </td>
					</tr>
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25781">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25781/" class="game_application_name">Первая лига | Игра #2</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">18 June, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
					<tr>
					    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
						<a class="" href="https://club60sec.timepad.ru/event/2905304/" target="_blank">Купить билет</a>
					    </td>
					</tr>
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25794">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25794/" class="game_application_name">Высшая лига | Игра #3</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">19 June, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
					<tr>
					    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
						<a class="" href="https://club60sec.timepad.ru/event/2907578/" target="_blank">Купить билет</a>
					    </td>
					</tr>
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25622">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25622/" class="game_application_name">Корпоративная лига | Игра #6</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">20 June, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1800 руб. с команды</td>
								    </tr>
				    
				    
					
					<tr>
					    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
						<a class="" href="https://club60sec.timepad.ru/event/2904734/" target="_blank">Купить билет</a>
					    </td>
					</tr>
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25807">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25807/" class="game_application_name">Киберлига | Игра #3</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">20 June, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
					<tr>
					    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
						<a class="" href="https://club60sec.timepad.ru/event/2927359/" target="_blank">Купить билет</a>
					    </td>
					</tr>
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25679">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25679/" class="game_application_name">Открытая лига | Игра #4</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">24 June, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
					<tr>
					    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
						<a class="" href="https://club60sec.timepad.ru/event/2927450/" target="_blank">Купить билет</a>
					    </td>
					</tr>
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25782">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25782/" class="game_application_name">Первая лига | Игра #3</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">25 June, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Фрегат "Благодать", Петровская наб., 2А</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;2400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25808">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25808/" class="game_application_name">Киберлига | Игра #4</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">27 June, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25680">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25680/" class="game_application_name">Открытая лига | Игра #5</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">1 July, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25783">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25783/" class="game_application_name">Первая лига | Игра #4</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">2 July, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25796">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25796/" class="game_application_name">Высшая лига | Игра #4</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">3 July, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25809">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25809/" class="game_application_name">Киберлига | Игра #5</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">4 July, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25681">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25681/" class="game_application_name">Открытая лига | Игра #6</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">8 July, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25784">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25784/" class="game_application_name">Первая лига | Игра #5</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">9 July, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25797">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25797/" class="game_application_name">Высшая лига | Игра #5</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">10 July, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25810">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25810/" class="game_application_name">Киберлига | Игра #6</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">11 July, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25682">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25682/" class="game_application_name">Открытая лига | Игра #7</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">15 July, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25785">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25785/" class="game_application_name">Первая лига | Игра #6</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">16 July, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25798">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25798/" class="game_application_name">Высшая лига | Игра #6</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">17 July, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25623">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25623/" class="game_application_name">Корпоративная лига | Игра #7</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">18 July, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1800 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25811">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25811/" class="game_application_name">Киберлига | Игра #7</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">18 July, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25683">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25683/" class="game_application_name">Открытая лига | Игра #8</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">22 July, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25786">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25786/" class="game_application_name">Первая лига | Игра #7</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">23 July, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25799">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25799/" class="game_application_name">Высшая лига | Игра #7</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">24 July, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25812">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25812/" class="game_application_name">Киберлига | Игра #8</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">25 July, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25684">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25684/" class="game_application_name">Открытая лига | Игра #9</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">29 July, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25787">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25787/" class="game_application_name">Первая лига | Игра #8</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">30 July, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25800">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25800/" class="game_application_name">Высшая лига | Игра #8</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">31 July, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25813">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25813/" class="game_application_name">Киберлига | Игра #9</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">1 August, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25685">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25685/" class="game_application_name">Открытая лига | Игра #10</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">5 August, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25788">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25788/" class="game_application_name">Первая лига | Игра #9</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">6 August, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25801">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25801/" class="game_application_name">Высшая лига | Игра #9</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">7 August, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25814">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25814/" class="game_application_name">Киберлига | Игра #10</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">8 August, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25686">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25686/" class="game_application_name">Открытая лига | Игра #11</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">12 August, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25789">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25789/" class="game_application_name">Первая лига | Игра #10</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">13 August, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25802">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25802/" class="game_application_name">Высшая лига | Игра #10</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">14 August, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25815">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25815/" class="game_application_name">Киберлига | Игра #11</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">15 August, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25687">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25687/" class="game_application_name">Открытая лига | Игра #12</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">19 August, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25790">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25790/" class="game_application_name">Первая лига | Игра #11</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">20 August, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25803">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25803/" class="game_application_name">Высшая лига | Игра #11</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">21 August, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25624">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25624/" class="game_application_name">Корпоративная лига | Игра #8</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">22 August, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1800 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25816">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25816/" class="game_application_name">Киберлига | Игра #12</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">22 August, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25688">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25688/" class="game_application_name">Открытая лига Финал</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">26 August, Monday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;400 руб. с человека</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25791">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25791/" class="game_application_name">Первая лига | Финал</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">27 August, Tuesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25804">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25804/" class="game_application_name">Высшая лига | Финал</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">28 August, Wednesday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
		    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25817">
			    <div class="container py-3" style="border: 1px solid #E4E4E4;">
				    <h5 class="dark-blue text-center"><a href="/quizgames/game/25817/" class="game_application_name">Киберлига | Финал</a></h5>
				    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
					    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
					    <div>
						    <table>
							    <tbody>
								    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">29 August, Thursday</span></td>
								    </tr>
								    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
								    </tr>
				    
								    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
								    </tr>
				    
				    
								    <tr>
									    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1400 руб. с команды</td>
								    </tr>
				    
				    
					
				
							    </tbody>
						    </table>
					    </div>
				    </div>
			    </div>
		    </div>
			    
	    </div>
    </div>
	
    <div aria-hidden="true" aria-labelledby="gameApplicationModal" class="modal modal_application fade" id="gameApplicationModal" role="dialog" tabindex="-1" data-backdrop="static">
      <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	<div class="modal-content">
	    <div class="modal-header">
		<button type="button" class="close" data-dismiss="modal" aria-label="Close">
		  <span aria-hidden="true"><i class="fa-regular fa-circle-xmark"></i></span>
		</button>
	    </div>
	    <h4 class="game_application_title"></h4>
	    <p class="game_application_description">Оставьте свои контакты, и наш добрый менеджер сделает всё за Вас!</p>
	    <form id="form_application">
		<input type="hidden" name="csrfmiddlewaretoken" value="s5Sqaoe4uuse8VbdYP6lMI204s7K0dIJJjgVe8IwFZE0PfpVcC7AaDhti4uP3uhO">
		<input type="hidden" name="game_id" value="">
		<input class="inp input gsheets required" type="text" name="name" placeholder="Ваше имя">
		<input class="inp input gsheets required" type="text" name="phone" placeholder="Телефон">
		<input class="inp input gsheets required" type="text" name="email" id="input-email" placeholder="Email">
    
		<div class="btn-radio btn-i_have_team">
		    <input value="У меня есть команда" type="radio" id="rdo-1" name="i_have_team" class="">
		    <svg width="20px" height="20px" viewBox="0 0 20 20"><circle class="svg_cirk2" cx="10" cy="10" r="9"></circle><circle cx="10" cy="10" r="9"></circle><path d="M10,7 C8.34314575,7 7,8.34314575 7,10 C7,11.6568542 8.34314575,13 10,13 C11.6568542,13 13,11.6568542 13,10 C13,8.34314575 11.6568542,7 10,7 Z" class="inner"></path><path d="M10,1 L10,1 L10,1 C14.9705627,1 19,5.02943725 19,10 L19,10 L19,10 C19,14.9705627 14.9705627,19 10,19 L10,19 L10,19 C5.02943725,19 1,14.9705627 1,10 L1,10 L1,10 C1,5.02943725 5.02943725,1 10,1 L10,1 Z" class="outer"></path></svg>
		    <span>У меня есть команда (от 2 до <output id="output1">8</output> человек)</span>
		</div>
		<div class="hidden_input" style="display: none;">
		    <input class="inp input gsheets" type="text" id="team" name="team" placeholder="Название команды" value="Нет команды">
		    <label for="num_players" class="label">Количество игроков: <output id="output"></output></label>
    
		    <div class="form-range">
			<input type="range" id="num_players" name="num_players" min="2" max="8" class="inp" oninput="document.getElementById('output').innerHTML = this.value;">
			<div class="ticks-10 d-none">
			    <span class="tick">2</span>
			    <span class="tick">3</span>
			    <span class="tick">4</span>
			    <span class="tick">5</span>
			    <span class="tick">6</span>
			    <span class="tick">7</span>
			    <span class="tick">8</span>
			    <span class="tick">9</span>
			    <span class="tick">10</span>
			</div>
			<div class="ticks-8">
			    <span class="tick">2</span>
			    <span class="tick">3</span>
			    <span class="tick">4</span>
			    <span class="tick">5</span>
			    <span class="tick">6</span>
			    <span class="tick">7</span>
			    <span class="tick">8</span>
			</div>
			<div class="ticks-6 d-none">
			    <span class="tick">2</span>
			    <span class="tick">3</span>
			    <span class="tick">4</span>
			    <span class="tick">5</span>
			    <span class="tick">6</span>
			</div>
		    </div>
		</div>
    
		<div class="btn-radio btn-i_dont_have_team">
		    <input value="Я без команды, присоединюсь к существующей" type="radio" id="rdo-2" name="team_radio" class="checked">
		    <svg width="20px" height="20px" viewBox="0 0 20 20"><circle class="svg_cirk2" cx="10" cy="10" r="9"></circle><circle cx="10" cy="10" r="9"></circle><path d="M10,7 C8.34314575,7 7,8.34314575 7,10 C7,11.6568542 8.34314575,13 10,13 C11.6568542,13 13,11.6568542 13,10 C13,8.34314575 11.6568542,7 10,7 Z" class="inner"></path><path d="M10,1 L10,1 L10,1 C14.9705627,1 19,5.02943725 19,10 L19,10 L19,10 C19,14.9705627 14.9705627,19 10,19 L10,19 L10,19 C5.02943725,19 1,14.9705627 1,10 L1,10 L1,10 C1,5.02943725 5.02943725,1 10,1 L10,1 Z" class="outer"></path></svg>
		    <span>Я без команды, присоединюсь к существующей</span>
		</div>
    
		<textarea class="inp input" name="comment" rows="5" placeholder="Комментарий"></textarea>
    
		<div class="row justify-content-center">
		    <div class="col-auto md-auto block-center">
			<button class="btn_modal" type="submit">Записаться на игру!</button>
			<img class="loader-img d-none" src="/static/main/img/loader.gif" style="height: 20px;">
		    </div>
		</div>
	    </form>
    
	</div>
      </div>
    </div>
    
    <div class="modal fade" id="thanksModal" tabindex="-1" role="dialog" aria-labelledby="filters" aria-hidden="true">
      <div class="modal-dialog">
	<div class="modal-content">
	  <div class="modal-header">
		<button type="button" class="close" data-dismiss="modal" aria-label="Close">
		  <span aria-hidden="true"><i class="fa-regular fa-circle-xmark"></i></span>
		</button>
	    </div>
	  <div class="modal-body text-center">
		<h3>Спасибо за заявку</h3>
	      <p>Для подтверждения регистрации в день игры с вами свяжется менеджер.<br> А сейчас ловите SMS или email с подробностями вашей игры.</p>
	  </div>
	</div>
      </div>
    </div>
    
    
    
    <div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
	    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
		    <div class="modal-content">
			    <div class="modal-body">
				    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
				    <div class="container" id="leagues">
					    <div class="row">
						    <div class="mx-auto text-center dark-blue" id="modal_city">
							    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
				    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
				    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
							    <div class="mx-auto text-center pb-2 single" style="display: none;">
				    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
							    <p></p>
						    </div>
						    <div class="col-lg-6 col-12 text-center container pb-2 offline">
							    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
							    <div class="container modal-leagues pb-3">
				    
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
								
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
								
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
								
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
								
				</div>
						    </div>
						    <div class="col-lg-6 col-12 text-center container pb-1 online">
							    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
							    <div class="container modal-leagues pb-3">
								    
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
								
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
								
							    </div>
						    </div>
						    <div class="mx-auto text-center container pb-2 archive">
							    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
				<div class="container modal-leagues pb-3" style="display:none">
							    
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
				
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
				
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
				
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
				
								    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
				
				</div>
						    </div>
					    </div>
				    </div>
				    <div class="container" id="cityes" style="display:none;">
					    <h2 class="black pt-4">Наша география</h2>
					    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
					    <br>
					    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
			    
			</div>
				    </div>
			    </div>
		    </div>
	    </div>
    </div>
    
    <footer class="footer">
	    <div class="container">
		    <div class="row">
			    <div class="col-12">
				    <div class="d-flex justify-content-between">
					    <div class="pb-3">
						    <a class="no-bb" href="/">
							    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
						    </a>
					    </div>
					    <div class="mt-4">
						    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
					    </div>
				    </div>
			    </div>
			    <div class="col-md-6 small">
    
				    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
				    </p>
    
			    </div>
			    <div class="col-md-6 text-right small">
				    <p>
					    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
				    </p>
			    </div>
		    </div>
	    </div>
    </footer>
    
    
    
    <script>
	var BASE_DIR = "quizgames/";
    </script>
    <script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>
    
    
    
    
    
    
    
    
    </body></html>`
	html25781 = `<html lang="ru"><head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
    <link rel="icon" href="/fav.png" type="image/png">
    <link rel="shortcut icon" href="/fav.png" type="image/png">
    <title>
        Первая лига | Игра #2 | Первая лига (Санкт-Петербург) | Лето 2024
        | Клуб «60 секунд»</title>
    <meta name="title" content="Первая лига | Игра #2 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
    <meta property="og:title" content="Первая лига | Игра #2 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

    <meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
    <meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

    <meta content="/static/img/snippet.jpg" property="og:image">

    <link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
    <link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
    <link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
    <link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
    <link href="/static/main/libs/chart/chart.css" rel="stylesheet">
    <link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
    <link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
    <link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
    <link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
  <div class="loader" style="display: none;">
    <div class="dot dot1"></div>
    <div class="dot dot2"></div>
    <div class="dot dot3"></div>
    <div class="dot dot4"></div>
  </div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
	<div class="container">
		<div class="navbar-brand d-flex flex-row align-items-center">
			<div class="pr-2">
				<a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
			</div>
            
			<div class="d-flex flex-column" id="club-selector">
				<div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
					<a data-target="#cityModal" data-toggle="modal" href="#">
						<span class="league_name" id="league_name">Выбрать лигу</span>
					</a>
				</div>
				<div style="margin-top: -6px;">
					<a data-target="#cityModal" data-toggle="modal" href="#">
						<span class="city" id="city">Санкт-Петербург</span>
					</a>
				</div>
			</div>
        
		</div>
		<div class="d-flex flex-row order-2 order-lg-3">
			<ul class="navbar-nav flex-row">
                <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
                  <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
                    <form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
                      <div class="input-group" style="min-width: 200px;">
                        <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
                          <span class="input-group-append">
                              <div class="input-group-text bg-transparent">
                                  <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
                              </div>
                          </span>
                      </div>
                    </form>
                  </div>
                </li>
				<li class="dropdown nav-item">
					<a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
						<img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
					</a>
					<div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
						<a class="dropdown-item d-none" href="/en/league/119/">
							<img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
						</a>
						<a class="dropdown-item" href="/league/119/">
							<img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
						</a>
					</div>
				</li>
			</ul>
			<button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
		</div>

		<div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
            
			<ul class="navbar-nav mr-auto uppercase">
				<li class="nav-item">
					<a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
				</li>
				<li class="nav-item">
					<a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
				</li>
				
				<li class="nav-item">
					<a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
				</li>
				
				
				<li class="nav-item">
					<a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
				</li>
				
			</ul>

			<ul class="navbar-nav ml-auto uppercase">
				
				  	<li class="nav-item">
						<a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				 	 </li>
				

			</ul>
        
		</div>

	</div>
</nav>


<div id="header">
	<div class="container py-3">
		<div class="row">
			<div class="col-md-12 col-12 pb-2">
				<div class="card transparent">
					<div class="card-body" style="padding:0;">
						<h2 class="white">
                            
                            Первая лига | Игра #2</h2>
						<div class="white-razdel"></div>
						<div class="l-blue pt-2">Санкт-Петербург /
                            <a href="/quizgames/league/118/">Первая лига</a> /
                            <a href="/quizgames/season/2755/">Лето 2024</a> /
                            18 Jun 2024, 7:30 p.m.
                        </div>
					</div>
				</div>

                
                    <div class="pt-2">
                        <a class="btn btn-yellow" href="/quizgames/game_enter/25781/" style="min-width: 250px;">Войти в игру</a>
                    </div>
                
            </div>
		</div>
	</div>
</div>

<div class="container-fluid">
	<div class="row justify-content-center">

		<div class="col-12" style="margin-bottom: -38px; z-index: 10">
          <div class="d-flex flex-row-reverse">
            <div class="dropdown text-right">
                <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
                    <h2><i class="fas fa-bars"></i></h2>
                </a>
              <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
                  <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
                  <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
                  <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
                  
                  <a class="dropdown-item" href="/quizgames/game/25781/show/">Режим демонстрации</a>
              </div>
            </div>
          </div>
        </div>

        <div class="collapse container-fluid" id="graph">
            <div class="row justify-content-center">
                <div class="col-12 col-lg-10 col-xl-8">
                    <canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
                </div>
            </div>
        </div>
        <div class="col-12 mx-auto text-center py-4">
        

            <div class="pb-4" id="main_content">
                <div id="table_1000">
                    <h3 class="dark-blue text-center">Итоговая таблица</h3>
                    <div class="text-center spiner d-none">
                        <div class="load">
                            <hr><hr><hr><hr>
                        </div>
                    </div>
                    <div class="table-results"></div>

                    <p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
                    <p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

                </div>

        
        </div>
        </div>
	</div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
	<div class="modal-dialog modal-dialog-centered modal-lg" role="document">
		<div class="modal-content">
			<div class="modal-body">
				<button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
				<div class="container" id="leagues">
					<div class="row">
						<div class="mx-auto text-center dark-blue" id="modal_city">
							<p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
                                <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
                                <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
							<div class="mx-auto text-center pb-2 single" style="display: none;">
                                <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
							<p></p>
						</div>
						<div class="col-lg-6 col-12 text-center container pb-2 offline">
							<h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
							<div class="container modal-leagues pb-3">
                                
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							    
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							    
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							    
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							    
                            </div>
						</div>
						<div class="col-lg-6 col-12 text-center container pb-1 online">
							<h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
							<div class="container modal-leagues pb-3">
								
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							    
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							    
							</div>
						</div>
						<div class="mx-auto text-center container pb-2 archive">
							<h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
                            <div class="container modal-leagues pb-3" style="display:none">
							
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
                            
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
                            
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
                            
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
                            
								<div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
                            
                            </div>
						</div>
					</div>
				</div>
				<div class="container" id="cityes" style="display:none;">
					<h2 class="black pt-4">Наша география</h2>
					<input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
					<br>
					<div class="container pb-3 list-group d-flex flex-row flex-wrap">
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
                        
                    </div>
				</div>
			</div>
		</div>
	</div>
</div>

<footer class="footer">
	<div class="container">
		<div class="row">
			<div class="col-12">
				<div class="d-flex justify-content-between">
					<div class="pb-3">
						<a class="no-bb" href="/">
							<img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
						</a>
					</div>
					<div class="mt-4">
						<a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
					</div>
				</div>
			</div>
			<div class="col-md-6 small">

				<p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
				</p>

			</div>
			<div class="col-md-6 text-right small">
				<p>
					<a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
				</p>
			</div>
		</div>
	</div>
</footer>



<script>
    var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





    <script src="/static/CACHE/js/output.b0e61c65fec3.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25782 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #3 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #3 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #3 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #3</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			25 Jun 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25782/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25782/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.e942631c22c5.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25783 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #4 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #4 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #4 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #4</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			2 Jul 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25783/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25783/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.bb0c64a00b1f.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25784 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #5 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #5 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #5 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #5</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			9 Jul 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25784/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25784/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.a2934a32eddc.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25785 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #6 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #6 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #6 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #6</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			16 Jul 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25785/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25785/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.a907295354ed.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25786 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #7 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #7 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #7 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #7</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			23 Jul 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25786/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25786/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.355fd9cd4899.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25787 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #8 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #8 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #8 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #8</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			30 Jul 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25787/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25787/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.da89a298ca7a.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25788 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #9 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #9 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #9 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #9</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			6 Aug 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25788/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25788/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.f88d3add4715.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25789 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #10 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #10 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #10 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #10</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			13 Aug 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25789/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25789/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.903f0d30cd14.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25790 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Игра #11 | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Игра #11 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Игра #11 | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Игра #11</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			20 Aug 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25790/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25790/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>



<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>





<script src="/static/CACHE/js/output.57596e3e2e91.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
	html25791 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Первая лига | Финал | Первая лига (Санкт-Петербург) | Лето 2024
    | Клуб «60 секунд»</title>
<meta name="title" content="Первая лига | Финал | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">
<meta property="og:title" content="Первая лига | Финал | Первая лига (Санкт-Петербург) | Лето 2024 | Клуб «60 секунд»">

<meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
<meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
<style type="text/css">/* Chart.js */
@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>


<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/syncro/">Синхроны</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/118/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/118/">Регламент</a>
			    </li>
			    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-12 col-12 pb-2">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">
			
			Первая лига | Финал</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Санкт-Петербург /
			<a href="/quizgames/league/118/">Первая лига</a> /
			<a href="/quizgames/season/2755/">Лето 2024</a> /
			27 Aug 2024, 7:30 p.m.
		    </div>
				    </div>
			    </div>

	    
		<div class="pt-2">
		    <a class="btn btn-yellow" href="/quizgames/game_enter/25791/" style="min-width: 250px;">Войти в игру</a>
		</div>
	    
	</div>
	    </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row justify-content-center">

	    <div class="col-12" style="margin-bottom: -38px; z-index: 10">
      <div class="d-flex flex-row-reverse">
	<div class="dropdown text-right">
	    <a aria-label="Game menu button" class="table_a" data-toggle="dropdown" href="#" id="dropdownMenuLink" role="button" aria-expanded="false">
		<h2><i class="fas fa-bars"></i></h2>
	    </a>
	  <div aria-labelledby="dropdownMenuLink" class="dropdown-menu dropdown-menu-right" style="will-change: transform;">
	      <a class="dropdown-item" href="#" onclick="division_show()">По дивизионам</a>
	      <a class="dropdown-item" data-target="#graph" data-toggle="collapse" href="#">График взятия</a>
	      <a class="dropdown-item" href="#" onclick="QRS()">Рейтинг вопросов</a>
	      
	      <a class="dropdown-item" href="/quizgames/game/25791/show/">Режим демонстрации</a>
	  </div>
	</div>
      </div>
    </div>

    <div class="collapse container-fluid" id="graph">
	<div class="row justify-content-center">
	    <div class="col-12 col-lg-10 col-xl-8">
		<canvas id="bar-chart" height="0" width="0" class="chartjs-render-monitor" style="display: block; height: 0px; width: 0px;"></canvas>
	    </div>
	</div>
    </div>
    <div class="col-12 mx-auto text-center py-4">
    

	<div class="pb-4" id="main_content">
	    <div id="table_1000">
		<h3 class="dark-blue text-center">Итоговая таблица</h3>
		<div class="text-center spiner d-none">
		    <div class="load">
			<hr><hr><hr><hr>
		    </div>
		</div>
		<div class="table-results"></div>

		<p class="explanation d-none pt-3"><span class="yellow">*</span> - результаты команды рассчитаны с учетом понижающего коэффициента</p>
		<p class="explanation-shooting d-none">* - победитель по дополнительным показателям</p>

	    </div>

    
    </div>
    </div>
    </div>
</div>


<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2023. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>
<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>
<script src="/static/CACHE/js/output.2802d9dbdd35.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>
</body></html>`
	html2 = `<html lang="ru"><head>
<meta charset="UTF-8">
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<link rel="icon" href="/fav.png" type="image/png">
<link rel="shortcut icon" href="/fav.png" type="image/png">
<title>
    Расписание игр
    | Клуб «60 секунд»</title>
<meta name="title" content="Расписание игр | Клуб «60 секунд»">
<meta property="og:title" content="Расписание игр | Клуб «60 секунд»">

<meta name="description" content="Расписание игр">
<meta property="og:description" content="Расписание игр">

<meta content="/static/img/snippet.jpg" property="og:image">

<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
<link href="/static/main/libs/select2/select2.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.carousel.min.css" rel="stylesheet">
<link href="/static/main/libs/owl/owl.theme.default.min.css" rel="stylesheet">
<link href="/static/main/css/style.css?v=1.2" rel="stylesheet">
</head>
<body class="">
<div class="loaderArea" style="display: none;">
<div class="loader" style="display: none;">
<div class="dot dot1"></div>
<div class="dot dot2"></div>
<div class="dot dot3"></div>
<div class="dot dot4"></div>
</div>
</div>



<nav class="navbar navbar-expand-lg navbar-light">
    <div class="container">
	    <div class="navbar-brand d-flex flex-row align-items-center">
		    <div class="pr-2">
			    <a href="/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
		    </div>
	
		    <div class="d-flex flex-column" id="club-selector">
			    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="league_name" id="league_name">Выбрать лигу</span>
				    </a>
			    </div>
			    <div style="margin-top: -6px;">
				    <a data-target="#cityModal" data-toggle="modal" href="#">
					    <span class="city" id="city">Санкт-Петербург</span>
				    </a>
			    </div>
		    </div>
    
	    </div>
	    <div class="d-flex flex-row order-2 order-lg-3">
		    <ul class="navbar-nav flex-row">
	    <li class="nav-item dropdown"><a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="navbarDropdown" role="button"><i class="fa fa-search mt-1 pr-3 pr-lg-0"></i></a>
	      <div aria-labelledby="navbarDropdown" class="dropdown-menu dropdown-menu-right popout">
		<form action="/quizgames/search/" class="form-inline dropdown-item" id="form-search" method="get">
		  <div class="input-group" style="min-width: 200px;">
		    <input autocomplete="off" class="form-control py-2 border-right-0 border" id="example-search-input" name="s" placeholder="Поиск" style="outline: 0 !important; border-color: initial; box-shadow: none;" type="Поиск">
		      <span class="input-group-append">
			  <div class="input-group-text bg-transparent">
			      <i class="fa fa-search" onclick="document.getElementById('form-search').submit();" style="cursor: pointer;"></i>
			  </div>
		      </span>
		  </div>
		</form>
	      </div>
	    </li>
			    <li class="dropdown nav-item">
				    <a aria-expanded="false" aria-haspopup="true" class="nav-link nav-link-nav" data-toggle="dropdown" href="#" id="lang_dropdown" role="button">
					    <img alt="Русский" class="mb-1" id="flag-gb" src="/static/img/ru.svg" width="24">
				    </a>
				    <div aria-labelledby="lang_dropdown" class="dropdown-menu dropdown-menu-right popout">
					    <a class="dropdown-item d-none" href="/en/league/119/">
						    <img alt="English" class="mb-1 pr-2" id="flag-gb" src="/static/img/gb.svg" width="24">English
					    </a>
					    <a class="dropdown-item" href="/league/119/">
						    <img alt="Русский" class="mb-1 pr-2" id="flag-ru" src="/static/img/ru.svg" width="24">Русский
					    </a>
				    </div>
			    </li>
		    </ul>
		    <button aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation" class="navbar-toggler" data-target="#navbarSupportedContent" data-toggle="collapse" style="border: 0px;" type="button"><i class="material-icons dark-blue">more_vert</i></button>
	    </div>

	    <div class="collapse navbar-collapse order-3 order-lg-2" id="navbarSupportedContent">
	
		    <ul class="navbar-nav mr-auto uppercase">
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/faq/">FAQ</a>
			    </li>
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/teams_league/117/">Команды</a>
			    </li>
			    
			    
			    <li class="nav-item">
				    <a class="nav-link nav-link-nav" href="/quizgames/reglament/117/">Регламент</a>
			    </li>
			    
	    
		    </ul>

		    <ul class="navbar-nav ml-auto uppercase">
			    
				      <li class="nav-item">
					    <a class="nav-link nav-link-nav" href="/quizgames/account/login/">Вход</a>
				      </li>
			    

		    </ul>
    
	    </div>

    </div>
</nav>


<div id="header">
    <div class="container py-3">
	    <div class="row">
		    <div class="col-md-6 col-12 pb-1">
			    <div class="card transparent">
				    <div class="card-body" style="padding:0;">
					    <h2 class="white">Расписание игр</h2>
					    <div class="white-razdel"></div>
					    <div class="l-blue pt-2">Клуб «60 Секунд» / Санкт-Петербург</div>
				    </div>
			    </div>
	</div>
	    </div>
    </div>
</div>

<div class="container">
    <div class="row pt-4">
	    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25683">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25683/" class="game_application_name">Открытая лига | Игра #8</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">22 July, Monday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;400 руб. с человека</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2953538/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25786">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25786/" class="game_application_name">Первая лига | Игра #7</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">23 July, Tuesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2965859/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25799">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25799/" class="game_application_name">Высшая лига | Игра #7</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">24 July, Wednesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2965968/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25812">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25812/" class="game_application_name">Киберлига | Игра #8</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">25 July, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1400 руб. с команды</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2970182/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25684">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25684/" class="game_application_name">Открытая лига | Игра #9</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">29 July, Monday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;400 руб. с человека</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2970306/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25787">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25787/" class="game_application_name">Первая лига | Игра #8</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">30 July, Tuesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2970395/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25800">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25800/" class="game_application_name">Высшая лига | Игра #8</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">31 July, Wednesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2970405/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25813">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25813/" class="game_application_name">Киберлига | Игра #9</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">1 August, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1400 руб. с команды</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2970188/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25685">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25685/" class="game_application_name">Открытая лига | Игра #10</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">5 August, Monday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;400 руб. с человека</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2970313/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25788">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25788/" class="game_application_name">Первая лига | Игра #9</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">6 August, Tuesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2970397/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25801">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25801/" class="game_application_name">Высшая лига | Игра #9</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">7 August, Wednesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2970416/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25814">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25814/" class="game_application_name">Киберлига | Игра #10</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">8 August, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1400 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25686">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25686/" class="game_application_name">Открытая лига | Игра #11</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">12 August, Monday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;400 руб. с человека</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25789">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25789/" class="game_application_name">Первая лига | Игра #10</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">13 August, Tuesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25802">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25802/" class="game_application_name">Высшая лига | Игра #10</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">14 August, Wednesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25815">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25815/" class="game_application_name">Киберлига | Игра #11</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">15 August, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1400 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25687">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25687/" class="game_application_name">Открытая лига | Игра #12</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">19 August, Monday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;400 руб. с человека</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25790">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25790/" class="game_application_name">Первая лига | Игра #11</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">20 August, Tuesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25803">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25803/" class="game_application_name">Высшая лига | Игра #11</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">21 August, Wednesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25624">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25624/" class="game_application_name">Корпоративная лига | Игра #8</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">22 August, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1800 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
				<tr>
				    <td class="dark-blue"><i class="fas fa-credit-card text-light-blue"></i>&nbsp;
					<a class="" href="https://club60sec.timepad.ru/event/2969706/" target="_blank">Купить билет</a>
				    </td>
				</tr>
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25816">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25816/" class="game_application_name">Киберлига | Игра #12</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">22 August, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1400 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25688">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25688/" class="game_application_name">Открытая лига Финал</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">26 August, Monday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;400 руб. с человека</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25791">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25791/" class="game_application_name">Первая лига | Финал</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">27 August, Tuesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25804">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25804/" class="game_application_name">Высшая лига | Финал</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">28 August, Wednesday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1500 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_25817">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/25817/" class="game_application_name">Киберлига | Финал</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">29 August, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">20:00</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Онлайн</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1400 руб. с команды</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_26733">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/26733/" class="game_application_name">Корпоративная лига | Игра #9</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">19 September, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1800 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_26734">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/26734/" class="game_application_name">Корпоративная лига | Игра #10</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">24 October, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1800 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3" id="game_application_26735">
		    <div class="container py-3" style="border: 1px solid #E4E4E4;">
			    <h5 class="dark-blue text-center"><a href="/quizgames/game/26735/" class="game_application_name">Корпоративная лига | Игра #11</a></h5>
			    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
				    <div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
				    <div>
					    <table>
						    <tbody>
							    <tr>
				<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_date">21 November, Thursday</span></td>
							    </tr>
							    <tr>
				<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;<span class="game_application_time">19:30</span></td>
							    </tr>
			    
							    <tr>
				<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;<span class="game_application_place">Дворец «Олимпия» - Литейный пр., д. 14</span></td>
							    </tr>
			    
			    
							    <tr>
								    <td class="dark-blue"><i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
				    &nbsp;1800 руб. с команды + депозит</td>
							    </tr>
			    
			    
				
			
						    </tbody>
					    </table>
				    </div>
			    </div>
		    </div>
	    </div>
		    
    </div>
</div>

<div aria-hidden="true" aria-labelledby="gameApplicationModal" class="modal modal_application fade" id="gameApplicationModal" role="dialog" tabindex="-1" data-backdrop="static">
<div class="modal-dialog modal-dialog-centered modal-lg" role="document">
<div class="modal-content">
    <div class="modal-header">
	<button type="button" class="close" data-dismiss="modal" aria-label="Close">
	  <span aria-hidden="true"><i class="fa-regular fa-circle-xmark"></i></span>
	</button>
    </div>
    <h4 class="game_application_title"></h4>
    <p class="game_application_description">Оставьте свои контакты, и наш добрый менеджер сделает всё за Вас!</p>
    <form id="form_application">
	<input type="hidden" name="csrfmiddlewaretoken" value="VtUs3gSv0gBoUNSMiAQwTowjx1SYed7UcHiX70mXbLNaB76uwnRLhjLMLDf3huGZ">
	<input type="hidden" name="game_id" value="">
	<input class="inp input gsheets required" type="text" name="name" placeholder="Ваше имя">
	<input class="inp input gsheets required" type="text" name="phone" placeholder="Телефон">
	<input class="inp input gsheets required" type="text" name="email" id="input-email" placeholder="Email">

	<div class="btn-radio btn-i_have_team">
	    <input value="У меня есть команда" type="radio" id="rdo-1" name="i_have_team" class="">
	    <svg width="20px" height="20px" viewBox="0 0 20 20"><circle class="svg_cirk2" cx="10" cy="10" r="9"></circle><circle cx="10" cy="10" r="9"></circle><path d="M10,7 C8.34314575,7 7,8.34314575 7,10 C7,11.6568542 8.34314575,13 10,13 C11.6568542,13 13,11.6568542 13,10 C13,8.34314575 11.6568542,7 10,7 Z" class="inner"></path><path d="M10,1 L10,1 L10,1 C14.9705627,1 19,5.02943725 19,10 L19,10 L19,10 C19,14.9705627 14.9705627,19 10,19 L10,19 L10,19 C5.02943725,19 1,14.9705627 1,10 L1,10 L1,10 C1,5.02943725 5.02943725,1 10,1 L10,1 Z" class="outer"></path></svg>
	    <span>У меня есть команда (от 2 до <output id="output1">8</output> человек)</span>
	</div>
	<div class="hidden_input" style="display: none;">
	    <input class="inp input gsheets" type="text" id="team" name="team" placeholder="Название команды" value="Нет команды">
	    <label for="num_players" class="label">Количество игроков: <output id="output"></output></label>

	    <div class="form-range">
		<input type="range" id="num_players" name="num_players" min="2" max="8" class="inp" oninput="document.getElementById('output').innerHTML = this.value;">
		<div class="ticks-10 d-none">
		    <span class="tick">2</span>
		    <span class="tick">3</span>
		    <span class="tick">4</span>
		    <span class="tick">5</span>
		    <span class="tick">6</span>
		    <span class="tick">7</span>
		    <span class="tick">8</span>
		    <span class="tick">9</span>
		    <span class="tick">10</span>
		</div>
		<div class="ticks-8">
		    <span class="tick">2</span>
		    <span class="tick">3</span>
		    <span class="tick">4</span>
		    <span class="tick">5</span>
		    <span class="tick">6</span>
		    <span class="tick">7</span>
		    <span class="tick">8</span>
		</div>
		<div class="ticks-6 d-none">
		    <span class="tick">2</span>
		    <span class="tick">3</span>
		    <span class="tick">4</span>
		    <span class="tick">5</span>
		    <span class="tick">6</span>
		</div>
	    </div>
	</div>

	<div class="btn-radio btn-i_dont_have_team">
	    <input value="Я без команды, присоединюсь к существующей" type="radio" id="rdo-2" name="team_radio" class="checked">
	    <svg width="20px" height="20px" viewBox="0 0 20 20"><circle class="svg_cirk2" cx="10" cy="10" r="9"></circle><circle cx="10" cy="10" r="9"></circle><path d="M10,7 C8.34314575,7 7,8.34314575 7,10 C7,11.6568542 8.34314575,13 10,13 C11.6568542,13 13,11.6568542 13,10 C13,8.34314575 11.6568542,7 10,7 Z" class="inner"></path><path d="M10,1 L10,1 L10,1 C14.9705627,1 19,5.02943725 19,10 L19,10 L19,10 C19,14.9705627 14.9705627,19 10,19 L10,19 L10,19 C5.02943725,19 1,14.9705627 1,10 L1,10 L1,10 C1,5.02943725 5.02943725,1 10,1 L10,1 Z" class="outer"></path></svg>
	    <span>Я без команды, присоединюсь к существующей</span>
	</div>

	<textarea class="inp input" name="comment" rows="5" placeholder="Комментарий"></textarea>

	<div class="row justify-content-center">
	    <div class="col-auto md-auto block-center">
		<button class="btn_modal" type="submit">Записаться на игру!</button>
		<img class="loader-img d-none" src="/static/main/img/loader.gif" style="height: 20px;">
	    </div>
	</div>
    </form>

</div>
</div>
</div>

<div class="modal fade" id="thanksModal" tabindex="-1" role="dialog" aria-labelledby="filters" aria-hidden="true">
<div class="modal-dialog">
<div class="modal-content">
  <div class="modal-header">
	<button type="button" class="close" data-dismiss="modal" aria-label="Close">
	  <span aria-hidden="true"><i class="fa-regular fa-circle-xmark"></i></span>
	</button>
    </div>
  <div class="modal-body text-center">
	<h3>Спасибо за заявку</h3>
      <p>Для подтверждения регистрации в день игры с вами свяжется менеджер.<br> А сейчас ловите SMS или email с подробностями вашей игры.</p>
  </div>
</div>
</div>
</div>



<div aria-labelledby="cityModalTitle" class="modal fade bd-example-modal-lg" id="cityModal" role="document" tabindex="-1" aria-modal="true">
    <div class="modal-dialog modal-dialog-centered modal-lg" role="document">
	    <div class="modal-content">
		    <div class="modal-body">
			    <button aria-label="Close" class="close" data-dismiss="modal" type="button"><span aria-hidden="true">×</span></button>
			    <div class="container" id="leagues">
				    <div class="row">
					    <div class="mx-auto text-center dark-blue" id="modal_city">
						    <p>Ваш город <span class="modal-city">Санкт-Петербург</span> верно? 
			    <span class="modal-city-inv multi" onclick="select_city()">Нет</span>
			    <span class="multi"><br>Выберите лигу. Вы сможете поменять выбор в любой момент.</span></p>
						    <div class="mx-auto text-center pb-2 single" style="display: none;">
			    <span class="modal-city-yes" onclick="set_league(197)">&nbsp;Да&nbsp;</span> <span class="modal-city-inv" onclick="select_city()">Нет</span></div>
						    <p></p>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-2 offline">
						    <h5 class="dark-blue pb-1" onclick="toggleOff()" style="cursor: pointer;">Offline <i class="fas fa-angle-down fa-xs" id="off-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
			    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(121)">Корпоративная лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(119)">Высшая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(118)">Первая лига</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(117)">Открытая лига</a></div>
							
			</div>
					    </div>
					    <div class="col-lg-6 col-12 text-center container pb-1 online">
						    <h5 class="dark-blue pb-1" onclick="toggleOn()" style="cursor: pointer;">Online <i class="fas fa-angle-down fa-xs" id="on-icon"></i></h5>
						    <div class="container modal-leagues pb-3">
							    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(207)">Спецпроекты</a></div>
							
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(198)">Киберлига</a></div>
							
						    </div>
					    </div>
					    <div class="mx-auto text-center container pb-2 archive">
						    <h5 class="dark-blue" onclick="toggleArchive()" style="cursor: pointer;">Archive <i class="fas fa-xs fa-angle-up" id="archive-icon"></i></h5>
			<div class="container modal-leagues pb-3" style="display:none">
						    
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(197)">Киберлига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(120)">Лига Чемпионов (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(215)">Воскресная Киберлига (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(122)">Медиаигры (Временно не играем)</a></div>
			
							    <div class="mx-auto text-center pb-2"><a href="#" class="modal-a" onclick="set_league(130)">Воскресная Лига  (Временно не играем)</a></div>
			
			</div>
					    </div>
				    </div>
			    </div>
			    <div class="container" id="cityes" style="display:none;">
				    <h2 class="black pt-4">Наша география</h2>
				    <input class="form-control" id="cityInput" onkeyup="cityFilter()" placeholder="Поиск" type="text">
				    <br>
				    <div class="container pb-3 list-group d-flex flex-row flex-wrap">
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(147, 'Владикавказ')">Владикавказ</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(144, 'Иннополис')">Иннополис</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(3, 'Киров')">Киров</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(74, 'Кливленд')">Кливленд</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(41, 'Коломна')">Коломна</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(138, 'Колпино')">Колпино</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(117, 'Комсомольск-на-Амуре')">Комсомольск-на-Амуре</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(65, 'Кострома')">Кострома</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(52, 'Котлас')">Котлас</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(142, 'Краснодар')">Краснодар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(141, 'Красноярск')">Красноярск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(19, 'Курск')">Курск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимассол')">Лимассол</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(145, 'Пермь')">Пермь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(146, 'Северодвинск')">Северодвинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллин')">Таллин</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(27, 'Элиста')">Элиста</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(14, 'Югорск')">Югорск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(143, 'Якутск')">Якутск</a></div>
		    
			<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(42, 'Ярославль')">Ярославль</a></div>
		    
		</div>
			    </div>
		    </div>
	    </div>
    </div>
</div>

<footer class="footer">
    <div class="container">
	    <div class="row">
		    <div class="col-12">
			    <div class="d-flex justify-content-between">
				    <div class="pb-3">
					    <a class="no-bb" href="/">
						    <img alt="Клуб 60 Секунд" src="/static/img/logo_white.svg" data-toggle="tooltip" width="80" data-original-title="" title="">
					    </a>
				    </div>
				    <div class="mt-4">
					    <a class="btn btn_w_ol btn-sm d-none" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
				    </div>
			    </div>
		    </div>
		    <div class="col-md-6 small">

			    <p class="truncate">© 2016-2024. Интеллектуальный клуб «60 секунд».<br>ИП Жирнов Д. В. | <a class="foot-link" href="/static/requisites.pdf" rel="noopener" target="_blank">реквизиты</a>
			    </p>

		    </div>
		    <div class="col-md-6 text-right small">
			    <p>
				    <a class="link_privacy foot-link" href="https://club60sec.ru/privacy/" rel="noopener" target="_blank">Политика конфиденциальности</a>
			    </p>
		    </div>
	    </div>
    </div>
</footer>
<script>
var BASE_DIR = "quizgames/";
</script>
<script src="/static/CACHE/js/output.f75f2bfe86be.js"></script>
</body></html>`
)
