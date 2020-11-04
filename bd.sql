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
	        ON DELETE CASCADE,
	 CONSTRAINT fk_users
        FOREIGN KEY(user_id)
	        REFERENCES users(user_id)
);

create table hotels (
    hotel_id serial PRIMARY KEY NOT NULL,
    name text ,
    location text,
    description text,
    img text,
    photos text[],
    curr_rating float DEFAULT 0 CHECK (curr_rating >= 0  AND curr_rating <=10),
    comm_count int DEFAULT 0 CHECK(comm_count >= 0)
);

CREATE EXTENSION pg_trgm;

CREATE INDEX hotels_trgm_name_idx ON hotels
  USING gin (name gin_trgm_ops);

CREATE INDEX hotels_trgm_loc_idx ON hotels
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

CREATE TRIGGER tr_comments_on_add
    AFTER INSERT OR DELETE
    ON comments FOR EACH ROW
EXECUTE PROCEDURE public.trproc_upd_hotels();

CREATE OR REPLACE FUNCTION public.trproc_upd_hotels()
    RETURNS trigger AS
    $body$
    BEGIN
        IF (TG_OP = 'INSERT') THEN
	        UPDATE hotels
            SET comm_count = comm_count + 1
            WHERE hotel_id = NEW.hotel_id;
            RETURN OLD;
        ELSEIF (TG_OP = 'DELETE') THEN
            UPDATE hotels
            SET comm_count = comm_count - 1
            WHERE hotel_id = NEW.hotel_id;
            RETURN OLD;
        END IF;
    END;
    $body$
LANGUAGE 'plpgsql';

INSERT INTO hotels(hotel_id,name,location,description,img,photos) VALUES
(1,'Villa Domina','Россия г.Москва','Вилла Domina находится в городе Сплит, всего в 5 минутах ходьбы от дворца Диоклетиана, находящегося под охраной ЮНЕСКО.','static/img/hotelImg1.jpg','{"static/img/hotelImg2.jpg","static/img/hotelImg3.jpg"}'),
(2,'Apartments Tudor','Италия г.Рим','В апартаментах Tudor, расположенных на побережье, всего в 200 метрах от пляжа Фируле, к услугам гостей номера с кондиционером, бесплатным WI-Fi, бесплатной парковкой и спутниковым телевидением.','static/img/hotelImg2.jpg','{"static/img/hotelImg3.jpg","static/img/hotelImg3.jpg"}'),
(3,'Villa Muller Apartments','Греция г.Афины','Комплекс апартаментов Villa Muller расположен в Сплите в жупании Сплитско-Далматинска, недалеко от пляжа Баквиче и дворца Диоклетиана.','static/img/hotelImg3.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg4.jpg"}'),
(4,'Split Inn Apartments','Мексика  г.Мехико','Эти красивые и стильные апартаменты идеально расположены в самом центре Сплита и отлично подходят для гостей, желающих ознакомиться с потрясающими достопримечательностями этого старинного города.','static/img/hotelImg4.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg3.jpg"}'),
(5,'Villa Domina','Россия г.Москва','Вилла Domina находится в городе Сплит, всего в 5 минутах ходьбы от дворца Диоклетиана, находящегося под охраной ЮНЕСКО.','static/img/hotelImg1.jpg','{"static/img/hotelImg2.jpg","static/img/hotelImg3.jpg"}'),
(6,'Apartments Tudor','Италия г.Рим','В апартаментах Tudor, расположенных на побережье, всего в 200 метрах от пляжа Фируле, к услугам гостей номера с кондиционером, бесплатным WI-Fi, бесплатной парковкой и спутниковым телевидением.','static/img/hotelImg2.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg3.jpg"}'),
(7,'Villa Muller Apartments','Греция г.Афины','Комплекс апартаментов Villa Muller расположен в Сплите в жупании Сплитско-Далматинска, недалеко от пляжа Баквиче и дворца Диоклетиана.','static/img/hotelImg3.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg4.jpg"}'),
(8,'Split Inn Apartments','Мексика  г.Мехико','Эти красивые и стильные апартаменты идеально расположены в самом центре Сплита и отлично подходят для гостей, желающих ознакомиться с потрясающими достопримечательностями этого старинного города.','static/img/hotelImg4.jpg','{"static/img/hotelImg1.jpg","static/img/hotelImg2.jpg"}');