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
    rating int,
    time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (hotel_id,user_id),
     CONSTRAINT fk_comments
        FOREIGN KEY(hotel_id)
	        REFERENCES hotels(hotel_id)
	        ON DELETE CASCADE
);

create table hotels (
    hotel_id serial PRIMARY KEY NOT NULL,
    name text ,
    location text,
    description text,
    img text,
    photos text[],
    curr_rating float DEFAULT 0 CHECK (curr_rating >= 0  AND curr_rating <=10)
);


CREATE EXTENSION pg_trgm;

CREATE INDEX hotels_trgm_idx ON hotels
  USING gin (name gin_trgm_ops);

CREATE INDEX hotels_trgm_idx ON hotels
  USING gin (location gin_trgm_ops);

CREATE INDEX id_idx ON hotels USING btree (hotel_id);

CREATE TRIGGER tr_comments_on_change
    BEFORE UPDATE
    ON comments FOR EACH ROW
EXECUTE PROCEDURE public.trproc_upd_comments();

CREATE OR REPLACE FUNCTION public.trproc_upd_comments()
    RETURNS trigger AS
    $body$
    BEGIN
	    NEW.time:=CURRENT_TIMESTAMP;
	    RETURN NEW;
    END;
    $body$
LANGUAGE 'plpgsql';

INSERT INTO hotels(hotel_id,name,location,description,img,photos) VALUES
(1,'Villa Domina','Россия г.Москва','Вилла Domina находится в городе Сплит, всего в 5 минутах ходьбы от дворца Диоклетиана, находящегося под охраной ЮНЕСКО.','static/img/hotelImg1.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg3.jpg"}'),
(2,'Apartments Tudor','Италия г.Рим','В апартаментах Tudor, расположенных на побережье, всего в 200 метрах от пляжа Фируле, к услугам гостей номера с кондиционером, бесплатным WI-Fi, бесплатной парковкой и спутниковым телевидением.','static/img/hotelImg2.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg3.jpg"}'),
(3,'Villa Muller Apartments','Греция г.Афины','Комплекс апартаментов Villa Muller расположен в Сплите в жупании Сплитско-Далматинска, недалеко от пляжа Баквиче и дворца Диоклетиана.','static/img/hotelImg3.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg3.jpg"}'),
(4,'Split Inn Apartments','Мексика  г.Мехико','Эти красивые и стильные апартаменты идеально расположены в самом центре Сплита и отлично подходят для гостей, желающих ознакомиться с потрясающими достопримечательностями этого старинного города.','static/img/hotelImg4.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg3.jpg"}');