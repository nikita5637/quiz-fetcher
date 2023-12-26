package sixty_seconds

const (
	html1 = `<html lang="ru"><head>
	<meta charset="UTF-8">
	<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
	<link rel="icon" href="/fav.png" type="image/png">
	<link rel="shortcut icon" href="/fav.png" type="image/png">
	<title>
	    Первая лига
	    | Клуб «60 секунд»</title>
	<meta name="title" content="Первая лига | Клуб «60 секунд»">
	<meta property="og:title" content="Первая лига | Клуб «60 секунд»">
    
	<meta name="description" content="Первая лига">
	<meta property="og:description" content="Первая лига">
    
	<meta content="/static/img/snippet.jpg" property="og:image">
    
	<link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
	<link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
	<link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
	<link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
	<link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
	<link href="/static/main/libs/chart/chart.css" rel="stylesheet">
	<link href="/static/main/css/style.css" rel="stylesheet">
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
				    <a href="/quizgames/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
			    </div>
		
			    <div class="d-flex flex-column" id="club-selector">
				    <div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
					    <a data-target="#cityModal" data-toggle="modal" href="#">
						    <span class="league_name" id="league_name">
							    Первая лига</span>
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
						    <a class="dropdown-item" href="/en/league/119/">
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
			    <div class="col-12">
				    <div class="d-flex flex-row-reverse">
					    <h4 class="mb-0">
						    
						    
					    </h4>
				    </div>
			    </div>
			    <div class="col-md-6 col-12 pb-4">
				    <div class="card transparent">
					    <div class="card-body" style="padding:0;">
    
						    <h2 class="white">
							    
							    Первая лига
						    </h2>
						    <div class="white-razdel"></div>
						    <div class="l-blue pt-2">Клуб «60 Секунд» / Санкт-Петербург</div>
					    </div>
				    </div>
			    </div>
			    <div class="col-md-6 col-12">
    
		    
		  <h5 class="yellow text-center">Ближайшая игра</h5>
		  <h6 class="l-grey text-center"><a href="/quizgames/game/23284/">Первая лига | Игра #4</a></h6>
		  <div class="col-12 d-flex flex-row justify-content-center align-items-center">
		    <div class="l-blue ff-ibc vertical-text align-middle"><span>Играем</span></div>
		    <div>
		      <table>
			<tbody>
			  <tr>
			    <td class="l-grey"><i class="fas fa-calendar-alt fa-fw l-blue"></i>&nbsp;9 January, Tuesday</td>
			  </tr>
			  <tr>
			    <td class="l-grey"><i class="fas fa-clock fa-fw l-blue"></i>&nbsp;19:30</td>
			  </tr>
			  
			  <tr>
			    <td class="l-grey"><i class="fas fa-map-marker-alt fa-fw l-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td>
			  </tr>
			  
			  
			</tbody>
		      </table>
		    </div>
		  </div>
		    
		</div>
			    <div class="col-md-6 col-12 py-3 pt-md-0 mt-md-n4">
				    <ul class="nav nav-pills justify-content-md-start justify-content-center" id="myTab" role="tablist">
					    <li class="nav-item nav-pills-yellow">
						    <a aria-controls="schedule" aria-selected="false" class="nav-link tab-link active" data-toggle="tab" href="#schedule" id="schedule-tab" role="tab">
							    <span>Лига</span>
						    </a>
					    </li>
					    <li class="nav-item nav-pills-yellow">
						    <a aria-controls="seasons" aria-selected="true" class="nav-link tab-link" data-toggle="tab" href="#seasons" id="seasons-tab" role="tab">
							    <span>Все сезоны</span>
						    </a>
					    </li>
					    <li class="nav-item nav-pills-yellow">
						    <a class="nav-link tab-link" href="/quizgames/schedule/71/">
							    <span>Расписание клуба</span>
						    </a>
					    </li>
				    </ul>
			    </div>
		    </div>
	    </div>
    </div>
    
    <div class="container">
	    <div class="row pt-4">
		    <div class="col-12">
			    <div class="tab-content" id="myTabContent">
				    <div aria-labelledby="schedule-tab" class="tab-pane fade show active" id="schedule" role="tabpanel">
    
	<div class="row">
	    <div class="col-12"><h4 class="text-center dark-blue">Ближайшие игры лиги</h4></div>
	    
	    <div class="col-lg-4 col-md-6 col-12 d-flex align-items-stretch pb-3">
		<div class="container py-3" style="border: 1px solid #E4E4E4;">
		    <h5 class="dark-blue text-center"><a href="/quizgames/game/23284/">Первая лига | Игра #4</a></h5>
		    <div class="col-12 d-flex flex-row justify-content-center align-items-center">
			<div class="text-light-blue ff-ibc vertical-text align-middle igraem"><span>Играем</span></div>
			<div>
			    <table>
				<tbody>
				    <tr>
					<td class="dark-blue"><i class="fas fa-calendar-alt fa-fw text-light-blue"></i>&nbsp;9 January, Tuesday</td>
				    </tr>
				    <tr>
					<td class="dark-blue"><i class="fas fa-clock fa-fw text-light-blue"></i>&nbsp;19:30</td>
				    </tr>
				    
				    <tr>
					<td class="dark-blue"><i class="fas fa-map-marker-alt fa-fw text-light-blue"></i>&nbsp;Дворец «Олимпия» - Литейный пр., д. 14</td>
				    </tr>
				    
				    
				    <tr>
					<td class="dark-blue">
					    <i class="fas fa-ruble-sign fa-fw text-light-blue"></i>
					    &nbsp;1500 руб. с команды</td>
				    </tr>
				    
				    
				</tbody>
			    </table>
			</div>
		    </div>
		</div>
	    </div>
    
	
	</div>
    
    
    
    <div class="row">
	<div class="col-12 pt-2">
	    <h4 class="dark-blue text-center">Статистика</h4>
	    <p><small>Ниже представлены топ-5 команд в некоторых номинациях. Обратите внимание, что это
		общая статистика по всем играм. Кроме общей статистики, есть расширенная статистика для
		каждого из сезонов. Например, вот статистика
		    <a class="table_a" href="/quizgames/season/2463/#stats">текущего</a>
			и
			<a class="table_a" href="/quizgames/season/2333/#stats">предыдущего</a> сезонов.
			
		
		Кликнув на любую из команд лиги - вы попадете на страницу этой команды, где сможете
		ознакомится с её подробной статистикой, как за всё время, так и за каждый сезон в
		отдельности. Команды можно искать с помощью поиска, а также из списка всех
		<a class="table_a" href="/quizgames/teams_league/118/">команд лиги</a>.</small>
	    </p>
	</div>
    
	<div class="col-sm-6 col-lg-4 col-12 pt-2">
	    <h6 class="dark-blue text-center">Больше всех сыграли</h6>
    
	    <div class="table-responsive">
	      <table class="table table-hover table-sm small">
		<thead>
		  <tr class="active">
		    <th class="text-right">#</th>
		    <th>Команда</th>
		    <th class="text-center">Игр</th>
		  </tr>
		</thead>
		<tbody id="stat_more_games">
    
    
		<tr><td class="text-right">1</td><td class="text-left"><a class="table_a" href="/quizgames/team/7114/">Развал-снисхождение</a></td><td class="text-center">259</td></tr><tr><td class="text-right">2</td><td class="text-left"><a class="table_a" href="/quizgames/team/7101/">iSteria</a></td><td class="text-center">228</td></tr><tr><td class="text-right">3</td><td class="text-left"><a class="table_a" href="/quizgames/team/7213/">С.А.Н.Д.А.л.И</a></td><td class="text-center">205</td></tr><tr><td class="text-right">4</td><td class="text-left"><a class="table_a" href="/quizgames/team/5144/">Второе предельное состояние</a></td><td class="text-center">199</td></tr><tr><td class="text-right">5</td><td class="text-left"><a class="table_a" href="/quizgames/team/7259/">Два капитана</a></td><td class="text-center">191</td></tr></tbody>
	      </table>
	    </div>
	</div>
    
	<div class="col-sm-6 col-lg-4 col-12 pt-2">
	    <h6 class="dark-blue text-center">Побед в сезонах</h6>
    
	    <div class="table-responsive">
	      <table class="table table-hover table-sm small">
		<thead>
		  <tr class="active">
		    <th class="text-right">#</th>
		    <th>Команда</th>
		    <th class="text-center">Игр</th>
		    <th class="text-center">Побед</th>
		  </tr>
		</thead>
		<tbody id="stat_win_season">
    
    
		<tr><td class="text-right">1</td><td class="text-left"><a class="table_a" href="/quizgames/team/9615/">АлисА в городах</a></td><td class="text-center">127</td><td class="text-center">6</td></tr><tr><td class="text-right">2</td><td class="text-left"><a class="table_a" href="/quizgames/team/20142/">Шевелись, Плотва!</a></td><td class="text-center">120</td><td class="text-center">6</td></tr><tr><td class="text-right">3</td><td class="text-left"><a class="table_a" href="/quizgames/team/24679/">Файервольные каменщики</a></td><td class="text-center">55</td><td class="text-center">5</td></tr><tr><td class="text-right">4</td><td class="text-left"><a class="table_a" href="/quizgames/team/5070/">АйКогнито</a></td><td class="text-center">179</td><td class="text-center">2</td></tr><tr><td class="text-right">5</td><td class="text-left"><a class="table_a" href="/quizgames/team/12403/">Зайчатки разума</a></td><td class="text-center">185</td><td class="text-center">2</td></tr></tbody>
	      </table>
	    </div>
	</div>
    
	<div class="col-sm-6 col-lg-4 col-12 pt-2">
	    <h6 class="dark-blue text-center">Победители</h6>
    
	    <div class="table-responsive">
	      <table class="table table-hover table-sm small">
		<thead>
		  <tr class="active">
		    <th class="text-right">#</th>
		    <th>Команда</th>
		    <th class="text-center">Игр</th>
		    <th class="text-center">Побед</th>
		  </tr>
		</thead>
		<tbody id="stat_top_one">
    
    
		<tr><td class="text-right">1</td><td class="text-left"><a class="table_a" href="/quizgames/team/20142/">Шевелись, Плотва!</a></td><td class="text-center">120</td><td class="text-center">49</td></tr><tr><td class="text-right">2</td><td class="text-left"><a class="table_a" href="/quizgames/team/9615/">АлисА в городах</a></td><td class="text-center">127</td><td class="text-center">28</td></tr><tr><td class="text-right">3</td><td class="text-left"><a class="table_a" href="/quizgames/team/7101/">iSteria</a></td><td class="text-center">228</td><td class="text-center">17</td></tr><tr><td class="text-right">4</td><td class="text-left"><a class="table_a" href="/quizgames/team/24679/">Файервольные каменщики</a></td><td class="text-center">55</td><td class="text-center">17</td></tr><tr><td class="text-right">5</td><td class="text-left"><a class="table_a" href="/quizgames/team/12403/">Зайчатки разума</a></td><td class="text-center">185</td><td class="text-center">13</td></tr></tbody>
	      </table>
	    </div>
	</div>
    
	<div class="col-sm-6 col-lg-4 col-12 pt-2">
	    <h6 class="dark-blue text-center">Призёры</h6>
    
	    <div class="table-responsive">
	      <table class="table table-hover table-sm small">
		<thead>
		  <tr class="active">
		    <th class="text-right">#</th>
		    <th>Команда</th>
		    <th class="text-center">Игр</th>
		    <th class="text-center">Пъедесталов</th>
		  </tr>
		</thead>
		<tbody id="stat_top_three">
    
    
		<tr><td class="text-right">1</td><td class="text-left"><a class="table_a" href="/quizgames/team/9615/">АлисА в городах</a></td><td class="text-center">127</td><td class="text-center">82</td></tr><tr><td class="text-right">2</td><td class="text-left"><a class="table_a" href="/quizgames/team/20142/">Шевелись, Плотва!</a></td><td class="text-center">120</td><td class="text-center">76</td></tr><tr><td class="text-right">3</td><td class="text-left"><a class="table_a" href="/quizgames/team/7101/">iSteria</a></td><td class="text-center">228</td><td class="text-center">65</td></tr><tr><td class="text-right">4</td><td class="text-left"><a class="table_a" href="/quizgames/team/12403/">Зайчатки разума</a></td><td class="text-center">185</td><td class="text-center">54</td></tr><tr><td class="text-right">5</td><td class="text-left"><a class="table_a" href="/quizgames/team/24679/">Файервольные каменщики</a></td><td class="text-center">55</td><td class="text-center">33</td></tr></tbody>
	      </table>
	    </div>
	</div>
    
	<div class="col-sm-6 col-lg-4 col-12 pt-2">
	    <h6 class="dark-blue text-center">Лучший процент взятия*</h6>
    
	    <div class="table-responsive">
	      <table class="table table-hover table-sm small">
		<thead>
		  <tr class="active">
		    <th class="text-right">#</th>
		    <th>Команда</th>
		    <th class="text-center">Игр</th>
		    <th class="text-center">Процент взятия</th>
		  </tr>
		</thead>
		<tbody id="stat_correct_percent">
    
    
		<tr><td class="text-right">1</td><td class="text-left"><a class="table_a" href="/quizgames/team/24679/">Файервольные каменщики</a></td><td class="text-center">55</td><td class="text-center">85.23%</td></tr><tr><td class="text-right">2</td><td class="text-left"><a class="table_a" href="/quizgames/team/8178/">The Hateful Seven</a></td><td class="text-center">43</td><td class="text-center">80.52%</td></tr><tr><td class="text-right">3</td><td class="text-left"><a class="table_a" href="/quizgames/team/32079/">Мычание ягнят</a></td><td class="text-center">16</td><td class="text-center">79.56%</td></tr><tr><td class="text-right">4</td><td class="text-left"><a class="table_a" href="/quizgames/team/30105/">Мы Вам перезвоним</a></td><td class="text-center">21</td><td class="text-center">79.25%</td></tr><tr><td class="text-right">5</td><td class="text-left"><a class="table_a" href="/quizgames/team/20142/">Шевелись, Плотва!</a></td><td class="text-center">120</td><td class="text-center">77.81%</td></tr></tbody>
	      </table>
	    </div>
	    <small><sup>*</sup> - среди команд сыгравших больше 15 игр</small>
	</div>
    
	<div class="col-sm-6 col-lg-4 col-12 pt-2">
	    <h6 class="dark-blue text-center">Лучшая серия*</h6>
    
	    <div class="table-responsive">
	      <table class="table table-hover table-sm small">
		<thead>
		  <tr class="active">
		    <th class="text-right">#</th>
		    <th>Команда</th>
		    <th class="text-center">Игр</th>
		    <th class="text-center">Серия</th>
		  </tr>
		</thead>
		<tbody id="stat_best_series">
    
    
		<tr><td class="text-right">1</td><td class="text-left"><a class="table_a" href="/quizgames/team/7114/">Развал-снисхождение</a></td><td class="text-center">259</td><td class="text-center">37</td></tr><tr><td class="text-right">2</td><td class="text-left"><a class="table_a" href="/quizgames/team/12403/">Зайчатки разума</a></td><td class="text-center">185</td><td class="text-center">36</td></tr><tr><td class="text-right">3</td><td class="text-left"><a class="table_a" href="/quizgames/team/5094/">Повстанцы с Будуина</a></td><td class="text-center">188</td><td class="text-center">35</td></tr><tr><td class="text-right">4</td><td class="text-left"><a class="table_a" href="/quizgames/team/6512/">Гильдия Старой Гвардии</a></td><td class="text-center">39</td><td class="text-center">35</td></tr><tr><td class="text-right">5</td><td class="text-left"><a class="table_a" href="/quizgames/team/3585/">НИИЧАВО</a></td><td class="text-center">53</td><td class="text-center">34</td></tr></tbody>
	      </table>
	    </div>
	    <small><sup>*</sup> - количество правильных ответов подряд</small>
	</div>
    
	<div class="col-sm-6 col-lg-4 col-12 pt-2">
	    <h6 class="dark-blue text-center">Лучшая худшая серия*</h6>
    
	    <div class="table-responsive">
	      <table class="table table-hover table-sm small">
		<thead>
		  <tr class="active">
		    <th class="text-right">#</th>
		    <th>Команда</th>
		    <th class="text-center">Игр</th>
		    <th class="text-center">Серия</th>
		  </tr>
		</thead>
		<tbody id="stat_bad_series">
    
    
		<tr><td class="text-right">1</td><td class="text-left"><a class="table_a" href="/quizgames/team/35056/">Фру-фру</a></td><td class="text-center">28</td><td class="text-center">39</td></tr><tr><td class="text-right">2</td><td class="text-left"><a class="table_a" href="/quizgames/team/30329/">СКИДИЗМ</a></td><td class="text-center">1</td><td class="text-center">35</td></tr><tr><td class="text-right">3</td><td class="text-left"><a class="table_a" href="/quizgames/team/7099/">Почему бы и нет</a></td><td class="text-center">46</td><td class="text-center">27</td></tr><tr><td class="text-right">4</td><td class="text-left"><a class="table_a" href="/quizgames/team/28774/">Мы верим в вино</a></td><td class="text-center">9</td><td class="text-center">25</td></tr><tr><td class="text-right">5</td><td class="text-left"><a class="table_a" href="/quizgames/team/7116/">Харизма атавизма</a></td><td class="text-center">117</td><td class="text-center">24</td></tr></tbody>
	      </table>
	    </div>
	    <small><sup>*</sup> - количество неправильных ответов подряд, среди команд сыгравших минимум 15 игр</small>
	</div>
    
	<div class="col-sm-6 col-lg-4 col-12 pt-2">
	    <h6 class="dark-blue text-center">Тоталов в турах*</h6>
    
	    <div class="table-responsive">
	      <table class="table table-hover table-sm small">
		<thead>
		  <tr class="active">
		    <th class="text-right">#</th>
		    <th>Команда</th>
		    <th class="text-center">Игр</th>
		    <th class="text-center">Тоталов</th>
		  </tr>
		</thead>
		<tbody id="stat_total">
    
    
		<tr><td class="text-right">1</td><td class="text-left"><a class="table_a" href="/quizgames/team/7101/">iSteria</a></td><td class="text-center">228</td><td class="text-center">16</td></tr><tr><td class="text-right">2</td><td class="text-left"><a class="table_a" href="/quizgames/team/12403/">Зайчатки разума</a></td><td class="text-center">185</td><td class="text-center">12</td></tr><tr><td class="text-right">3</td><td class="text-left"><a class="table_a" href="/quizgames/team/20142/">Шевелись, Плотва!</a></td><td class="text-center">120</td><td class="text-center">11</td></tr><tr><td class="text-right">4</td><td class="text-left"><a class="table_a" href="/quizgames/team/9615/">АлисА в городах</a></td><td class="text-center">127</td><td class="text-center">8</td></tr><tr><td class="text-right">5</td><td class="text-left"><a class="table_a" href="/quizgames/team/5142/">Power Of Sport</a></td><td class="text-center">70</td><td class="text-center">7</td></tr></tbody>
	      </table>
	    </div>
	    <small><sup>*</sup> - тотал - это тур в котором, команда ответила верно на все вопросы</small>
	</div>
    
    </div>
    
    </div>
				    
    <div aria-labelledby="seasons-tab" class="tab-pane fade" id="seasons" role="tabpanel">
	<h4 class="text-center dark-blue">Сезоны
	    
	</h4>
	<div class="table-responsive pb-4">
	    <table class="table table-hover table-sm tablesorter tablesorter-default tablesorter10b862ff76c91" role="grid">
		<thead>
		    <tr role="row" class="tablesorter-headerRow">
			<th class="text-nowrap tablesorter-header tablesorter-headerUnSorted" data-column="0" tabindex="0" scope="col" role="columnheader" aria-disabled="false" unselectable="on" aria-sort="none" aria-label="Сезон: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Сезон</div></th>
			<th class="text-nowrap text-center tablesorter-header tablesorter-headerUnSorted" data-column="1" tabindex="0" scope="col" role="columnheader" aria-disabled="false" unselectable="on" aria-sort="none" aria-label="Игр: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Игр</div></th>
			<th class="text-nowrap tablesorter-header tablesorter-headerUnSorted" data-column="2" tabindex="0" scope="col" role="columnheader" aria-disabled="false" unselectable="on" aria-sort="none" aria-label="Начало: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Начало</div></th>
			<th class="text-nowrap tablesorter-header tablesorter-headerUnSorted" data-column="3" tabindex="0" scope="col" role="columnheader" aria-disabled="false" unselectable="on" aria-sort="none" aria-label="Окончание: No sort applied, activate to apply an ascending sort" style="-webkit-user-select: none;"><div class="tablesorter-header-inner">Окончание</div></th>
		    </tr>
		</thead>
		<tbody aria-live="polite" aria-relevant="all">
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/2463/">Зима-23/24</a>
			</td>
			<td class="text-nowrap text-center">4</td>
			<td class="text-nowrap">5 Dec 2023</td>
			<td class="text-nowrap">9 Jan 2024</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/2333/">Осень 2023</a>
			</td>
			<td class="text-nowrap text-center">12</td>
			<td class="text-nowrap">5 Sep 2023</td>
			<td class="text-nowrap">28 Nov 2023</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/2223/">Лето 2023</a>
			</td>
			<td class="text-nowrap text-center">12</td>
			<td class="text-nowrap">6 Jun 2023</td>
			<td class="text-nowrap">29 Aug 2023</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/2091/">Весна 2023</a>
			</td>
			<td class="text-nowrap text-center">12</td>
			<td class="text-nowrap">14 Mar 2023</td>
			<td class="text-nowrap">30 May 2023</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1927/">Зима-22/23</a>
			</td>
			<td class="text-nowrap text-center">12</td>
			<td class="text-nowrap">29 Nov 2022</td>
			<td class="text-nowrap">28 Feb 2023</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1814/">Осень 2022</a>
			</td>
			<td class="text-nowrap text-center">13</td>
			<td class="text-nowrap">29 Aug 2022</td>
			<td class="text-nowrap">21 Nov 2022</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1739/">Лето 2022</a>
			</td>
			<td class="text-nowrap text-center">10</td>
			<td class="text-nowrap">6 Jun 2022</td>
			<td class="text-nowrap">22 Aug 2022</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1624/">Весна 2022</a>
			</td>
			<td class="text-nowrap text-center">10</td>
			<td class="text-nowrap">7 Mar 2022</td>
			<td class="text-nowrap">30 May 2022</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1506/">Зима 2022</a>
			</td>
			<td class="text-nowrap text-center">11</td>
			<td class="text-nowrap">6 Dec 2021</td>
			<td class="text-nowrap">28 Feb 2022</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1393/">Осень 2021</a>
			</td>
			<td class="text-nowrap text-center">12</td>
			<td class="text-nowrap">6 Sep 2021</td>
			<td class="text-nowrap">29 Nov 2021</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1308/">Лето 2021</a>
			</td>
			<td class="text-nowrap text-center">13</td>
			<td class="text-nowrap">7 Jun 2021</td>
			<td class="text-nowrap">30 Aug 2021</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1198/">Весна 2021</a>
			</td>
			<td class="text-nowrap text-center">12</td>
			<td class="text-nowrap">15 Mar 2021</td>
			<td class="text-nowrap">31 May 2021</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1131/">Зима 2021</a>
			</td>
			<td class="text-nowrap text-center">6</td>
			<td class="text-nowrap">18 Jan 2021</td>
			<td class="text-nowrap">1 Mar 2021</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/1031/">Поздняя осень 2020</a>
			</td>
			<td class="text-nowrap text-center">7</td>
			<td class="text-nowrap">12 Oct 2020</td>
			<td class="text-nowrap">23 Nov 2020</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/969/">Блиц-сезон:Евротур</a>
			</td>
			<td class="text-nowrap text-center">4</td>
			<td class="text-nowrap">14 Sep 2020</td>
			<td class="text-nowrap">5 Oct 2020</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/949/">Блиц-сезон:Космическая Одиссея</a>
			</td>
			<td class="text-nowrap text-center">4</td>
			<td class="text-nowrap">10 Aug 2020</td>
			<td class="text-nowrap">31 Aug 2020</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/741/">Весна 2020</a>
			</td>
			<td class="text-nowrap text-center">2</td>
			<td class="text-nowrap">10 Mar 2020</td>
			<td class="text-nowrap">17 Mar 2020</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/614/">Зима 2020</a>
			</td>
			<td class="text-nowrap text-center">12</td>
			<td class="text-nowrap">3 Dec 2019</td>
			<td class="text-nowrap">3 Mar 2020</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/536/">Осень 2019</a>
			</td>
			<td class="text-nowrap text-center">12</td>
			<td class="text-nowrap">10 Sep 2019</td>
			<td class="text-nowrap">26 Nov 2019</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/420/">Лето 2019</a>
			</td>
			<td class="text-nowrap text-center">13</td>
			<td class="text-nowrap">4 Jun 2019</td>
			<td class="text-nowrap">27 Aug 2019</td>
		    </tr>
		
		    <tr role="row">
			<td class="text-nowrap">
			    
			    <a class="table_a" href="/quizgames/season/364/">Весна 2019</a>
			</td>
			<td class="text-nowrap text-center">8</td>
			<td class="text-nowrap">15 Apr 2019</td>
			<td class="text-nowrap">28 May 2019</td>
		    </tr>
		
		</tbody>
	    </table>
	</div>
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
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div>
			    
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
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div>
			    
				<div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div>
			    
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
						    <a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
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
    <script src="/static/CACHE/js/output.2d352d1a7b33.js"></script>
    
    
    
    
    
	<script src="/static/CACHE/js/output.e9aea6c5d367.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>
    
    
    
    </body></html>`
	html23284 = `<html lang="ru"><head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
    <link rel="icon" href="/fav.png" type="image/png">
    <link rel="shortcut icon" href="/fav.png" type="image/png">
    <title>
        Первая лига | Игра #4 | Первая лига (Санкт-Петербург) | Зима-23/24
        | Клуб «60 секунд»</title>
    <meta name="title" content="Первая лига | Игра #4 | Первая лига (Санкт-Петербург) | Зима-23/24 | Клуб «60 секунд»">
    <meta property="og:title" content="Первая лига | Игра #4 | Первая лига (Санкт-Петербург) | Зима-23/24 | Клуб «60 секунд»">

    <meta name="description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">
    <meta property="og:description" content="Интеллектуальные игры Online. Клуб 60 секунд. Выбери  свою лигу.">

    <meta content="/static/img/snippet.jpg" property="og:image">

    <link href="/static/main/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/main/libs/fontawesome/css/all.css" rel="stylesheet">
    <link href="/static/main/css/bootstrap-datetimepicker.min.css" rel="stylesheet">
    <link href="/static/main/libs/jquery/jquery-ui.min.css" rel="stylesheet">
    <link href="/static/main/libs/alertify/alertify.min.css" rel="stylesheet">
    <link href="/static/main/libs/chart/chart.css" rel="stylesheet">
    <link href="/static/main/css/style.css" rel="stylesheet">
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
				<a href="/quizgames/"><img alt="Клуб 60 Секунд" class="logo-img1" id="logo-img" src="/static/img/logo.svg"></a>
			</div>
            
			<div class="d-flex flex-column" id="club-selector">
				<div class="text-truncate" style="max-width: 250px; margin-bottom: -6px;">
					<a data-target="#cityModal" data-toggle="modal" href="#">
						<span class="league_name" id="league_name">
							Первая лига</span>
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
						<a class="dropdown-item" href="/en/league/119/">
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
                            <a href="/quizgames/season/2463/">Зима-23/24</a> /
                            9 Jan 2024, 7:30 p.m.
                        </div>
					</div>
				</div>

                
                    <div class="pt-2">
                        <a class="btn btn-yellow" href="/quizgames/game_enter/23284/" style="min-width: 250px;">Войти в игру</a>
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
                  
                  <a class="dropdown-item" href="/quizgames/game/23284/show">Режим демонстрации</a>
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
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(57, 'Актау')">Актау</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(121, 'Аланья')">Аланья</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(104, 'Алматы')">Алматы</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(116, 'Анталья')">Анталья</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(2, 'Антананариву')">Антананариву</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(45, 'Архангельск и Северодвинск')">Архангельск и Северодвинск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(97, 'Астана')">Астана</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(34, 'Астрахань')">Астрахань</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(43, 'Баку')">Баку</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(81, 'Барнаул')">Барнаул</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(133, 'Барселона')">Барселона</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(126, 'Батуми')">Батуми</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(24, 'Берлин')">Берлин</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(102, 'Бишкек')">Бишкек</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(96, 'Бремен')">Бремен</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(12, 'Брест')">Брест</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(11, 'Великий Новгород')">Великий Новгород</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(122, 'Вильнюс')">Вильнюс</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(6, 'Владивосток')">Владивосток</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(22, 'Владимир')">Владимир</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(53, 'Вологда')">Вологда</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(61, 'Воркута')">Воркута</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(30, 'Воронеж')">Воронеж</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(28, 'Воткинск')">Воткинск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(73, 'Гаага')">Гаага</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(76, 'Гамбург')">Гамбург</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(51, 'Гатчина')">Гатчина</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(40, 'Глазов')">Глазов</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(106, 'Гуанчжоу')">Гуанчжоу</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(77, 'Детройт')">Детройт</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(136, 'Долгопрудный')">Долгопрудный</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(39, 'Дубай')">Дубай</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(47, 'Екатеринбург')">Екатеринбург</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(113, 'Ереван')">Ереван</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(109, 'Жуковский')">Жуковский</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(87, 'Зарайск')">Зарайск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(44, 'Иваново')">Иваново</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(94, 'Игра Головой Online')">Игра Головой Online</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(68, 'Ижевск')">Ижевск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(90, 'Иматра')">Иматра</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(128, 'Йошкар-Ола')">Йошкар-Ола</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(108, 'Казань')">Казань</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(99, 'Калгари')">Калгари</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(13, 'Калининград')">Калининград</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(63, 'Калуга')">Калуга</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(112, 'Караганда')">Караганда</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(101, 'Кёльн')">Кёльн</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(37, 'Кемерово')">Кемерово</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(75, 'Киль')">Киль</a></div>
                        
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
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(91, 'Лаппеенранта')">Лаппеенранта</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(1, 'Лимасол')">Лимасол</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(111, 'Лиссабон')">Лиссабон</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(21, 'Лондон')">Лондон</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(80, 'Лос-Анджелес')">Лос-Анджелес</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(105, 'Луховицы')">Луховицы</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(124, 'Манавгат')">Манавгат</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(67, 'Миасс, Златоуст, Сатка')">Миасс, Златоуст, Сатка</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(78, 'Минск')">Минск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(131, 'Молодечно')">Молодечно</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(25, 'Москва')">Москва</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(62, 'Мурманск')">Мурманск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(110, 'Набережные Челны')">Набережные Челны</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(64, 'Надым')">Надым</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(8, 'Находка')">Находка</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(23, 'Нижнекамск')">Нижнекамск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(16, 'Нижний Новгород')">Нижний Новгород</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(88, 'Никосия')">Никосия</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(123, 'Нови Сад')">Нови Сад</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(46, 'Новокузнецк')">Новокузнецк</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(50, 'Новомосковск')">Новомосковск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(17, 'Новосибирск')">Новосибирск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(119, 'Ноябрьск')">Ноябрьск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(125, 'Обнинск')">Обнинск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(82, 'Омск')">Омск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(92, 'Орск')">Орск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(84, 'Островцы и Октябрьский')">Островцы и Октябрьский</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(89, 'Павлодар')">Павлодар</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(127, 'Пафос')">Пафос</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(107, 'Пекин')">Пекин</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(58, 'Пенза')">Пенза</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(120, 'Петропавловск-Камчатский')">Петропавловск-Камчатский</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(114, 'Пинск')">Пинск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(140, 'Порту')">Порту</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(33, 'Прага')">Прага</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(59, 'Псков')">Псков</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(139, 'Пушкин')">Пушкин</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(132, 'Пхукет')">Пхукет</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(129, 'Раменское')">Раменское</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(49, 'Ростов-на-Дону')">Ростов-на-Дону</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(20, 'Рязань')">Рязань</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(26, 'Самара')">Самара</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(71, 'Санкт-Петербург')">Санкт-Петербург</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(95, 'Саратов')">Саратов</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(85, 'Североморск')">Североморск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(66, 'Симферополь')">Симферополь</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(134, 'Сочи')">Сочи</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(93, 'Ставрополь')">Ставрополь</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(31, 'Стокгольм')">Стокгольм</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(135, 'Сыктывкар')">Сыктывкар</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(48, 'Таганрог')">Таганрог</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(15, 'Таллинн')">Таллинн</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(55, 'Ташкент')">Ташкент</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(118, 'Тбилиси')">Тбилиси</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(29, 'Тверь')">Тверь</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(56, 'Тенерифе')">Тенерифе</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(70, 'Томск')">Томск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(137, 'Тосно')">Тосно</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(69, 'Тула')">Тула</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(130, 'Тюмень')">Тюмень</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(83, 'Удомля')">Удомля</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(54, 'Уссурийск')">Уссурийск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(60, 'Ухта')">Ухта</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(9, 'Хабаровск')">Хабаровск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(115, 'Чайковский')">Чайковский</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(38, 'Чебоксары')">Чебоксары</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(35, 'Челябинск')">Челябинск</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(18, 'Чита')">Чита</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(100, 'Шанхай')">Шанхай</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(72, 'Шахты')">Шахты</a></div>
                        
                            <div class="modal-a-div" style="padding: 10px; padding-right: 20px;"><a class="modal-a" href="#" onclick="set_city(103, 'Шымкент')">Шымкент</a></div>
                        
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
						<a class="btn btn_w_ol btn-sm" href="http://vopros.club60sec.ru/" rel="noopener" target="_blank">Прислать вопрос</a>
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
<script src="/static/CACHE/js/output.d470fde14561.js"></script>





    <script src="/static/CACHE/js/output.55c3fe27ba56.js"></script><div class="alertify-notifier ajs-top ajs-right"></div>



</body></html>`
)
