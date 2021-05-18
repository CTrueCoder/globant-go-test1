<!DOCTYPE html>

<html>
<head>
  <title>Globant go api</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="/static/css/base.css">
</head>

<body>
  <div>
    <h1 class="charcoal rounded box">api</h1>

    <div class="">
      <h3>
        get all books
      </h3>
      <div class="charcoal rounded-box">
        GET: /books/get
      </div>
    </div>

    <div class="">
      <h3>
        get all books using filter by name, minprice, maxprice, genre
      </h3>
      <div class="charcoal rounded-box">
        GET: /books/get?name=Hero
        <br>
        GET: /books/get?name=Hero&genre=2
        <br>
        POST: /books/get {json}
      </div>
    </div>

    <div class="">
      <h3>
        get book by id
      </h3>
      <div class="charcoal rounded-box">
        /books/get/1
      </div>
    </div>

    <div class="">
      <h3>
        create book using json{Name, Price, Genre, Amount}
      </h3>
      <div class="charcoal rounded-box">
        POST: /books/create
      </div>
    </div>
    
    <div class="">
      <h3>
        Update book data {Name, Price, Genre, Amount} by id
        <br>
        Do update can use one-more field
      </h3>
      <div class="charcoal rounded-box">
        PUT: /books/update/1 {json}
      </div>
    </div>

    <div class="">
      <h3>
        Delete book by id
      </h3>
      <div class="charcoal rounded-box">
        DELETE: /books/delete/1
      </div>
    </div>

  </div>
</body>
</html>
