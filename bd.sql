create table users (
    user_id  serial not null PRIMARY KEY,
    username VARCHAR (50) UNIQUE,
    email VARCHAR (50) UNIQUE,
    password text,
    avatar text
);

create table comments (
    comm_id serial not null PRIMARY KEY,
    user_id int not null,
    hotel_id int not null,
    message text,
    rating int
);

create table hotels (
    hotel_id serial PRIMARY KEY NOT NULL,
    name text ,
    location text,
    description text,
    img text,
    curr_rating int DEFAULT 0 CHECK (curr_rating >= 0  AND curr_rating <=10)
);

create table rating (
    rate_id serial PRIMARY KEY NOT NULL,
    hotel_id int not null,
    user_id int not null,
    rate int CHECK (rate >= 0  AND rate <=10)
);


CREATE EXTENSION pg_trgm;

CREATE INDEX hotels_trgm_idx ON hotels
  USING gin (name gin_trgm_ops);

CREATE INDEX hotels_trgm_idx ON hotels
  USING gin (location gin_trgm_ops);

CREATE INDEX id_idx ON hotels USING btree (hotel_id);

INSERT INTO hotels(hotel_id,name,location,description,img) VALUES
(5,'Villa Domina','Россия г.Москва','Вилла Domina находится в городе Сплит, всего в 5 минутах ходьбы от дворца Диоклетиана, находящегося под охраной ЮНЕСКО.','static/img/hotelImg1.jpg'),
(6,'Apartments Tudor','Италия г.Рим','В апартаментах Tudor, расположенных на побережье, всего в 200 метрах от пляжа Фируле, к услугам гостей номера с кондиционером, бесплатным WI-Fi, бесплатной парковкой и спутниковым телевидением.','static/img/hotelImg2.jpg'),
(7,'Villa Muller Apartments','Греция г.Афины','Комплекс апартаментов Villa Muller расположен в Сплите в жупании Сплитско-Далматинска, недалеко от пляжа Баквиче и дворца Диоклетиана.','static/img/hotelImg3.jpg'),
(8,'Split Inn Apartments','Мексика  г.Мехико','Эти красивые и стильные апартаменты идеально расположены в самом центре Сплита и отлично подходят для гостей, желающих ознакомиться с потрясающими достопримечательностями этого старинного города.','static/img/hotelImg4.jpg');