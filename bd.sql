create table users (
    id  serial not null PRIMARY KEY,
    username text ,
    email text UNIQUE,
    password text,
    avatar text
);

select password from userss where username='kostikan';


create table hotels (
    id  serial not null PRIMARY KEY,
    name text ,
    description text,
    img text
);

INSERT INTO hotels(id,name,description,img) VALUES 
(1,'Villa Domina','Вилла Domina находится в городе Сплит, всего в 5 минутах ходьбы от дворца Диоклетиана, находящегося под охраной ЮНЕСКО.','static/img/hotelImg1.jpg'),
(2,'Apartments Tudor','В апартаментах Tudor, расположенных на побережье, всего в 200 метрах от пляжа Фируле, к услугам гостей номера с кондиционером, бесплатным WI-Fi, бесплатной парковкой и спутниковым телевидением.','static/img/hotelImg2.jpg'),
(3,'Villa Muller Apartments','Комплекс апартаментов Villa Muller расположен в Сплите в жупании Сплитско-Далматинска, недалеко от пляжа Баквиче и дворца Диоклетиана.','static/img/hotelImg3.jpg'),
(4,'Split Inn Apartments','Эти красивые и стильные апартаменты идеально расположены в самом центре Сплита и отлично подходят для гостей, желающих ознакомиться с потрясающими достопримечательностями этого старинного города.','static/img/hotelImg4.jpg');