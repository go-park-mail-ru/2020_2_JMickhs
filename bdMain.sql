ALTER SYSTEM SET shared_buffers = '128MB';
CREATE EXTENSION citext;
CREATE EXTENSION POSTGIS;
CREATE EXTENSION pg_trgm;
CREATE EXTENSION btree_gist;

create table hotels (
    hotel_id serial PRIMARY KEY NOT NULL,
    name citext,
    location citext,
    country citext,
    city   citext,
    coordinates geometry(POINT,4326),
    description text,
    img text,
    photos text[],
    curr_rating float DEFAULT 0 CHECK (curr_rating >= 0  AND curr_rating <=5),
    comm_count int DEFAULT 0 CHECK(comm_count >= 0),
    comm_count_for_each  int[5] DEFAULT '{0,0,0,0,0,0}'
);

create table recommendations(
    user_id int UNIQUE ,
    hotel_id int[],
    time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create table chats(
    chat_id text,
    user_id int
);

CREATE TABLE wishlists(
    wishlist_id serial PRIMARY KEY NOT NULL,
    name citext,
    user_id int
);

CREATE TABLE wishlistshotels(
    wishlist_id int,
    CONSTRAINT fk_wishlists
        FOREIGN KEY(wishlist_id)
            REFERENCES wishlists(wishlist_id)
                ON DELETE CASCADE,
    hotel_id int,
    UNIQUE (wishlist_id, hotel_id)
);

CREATE INDEX if not exists hotels_gist_idx ON hotels USING gist (coordinates);
CREATE INDEX hotels_trgm_idx ON hotels
    USING gist (name);
CREATE INDEX hotels_trgm_idx2 ON hotels
    USING gist (location);
CREATE INDEX id_idx ON hotels USING btree (hotel_id);

create table comments (
    comm_id serial not null PRIMARY KEY,
    user_id int not null,
    hotel_id int not null,
    message text,
    rating int DEFAULT 0 CHECK (rating  >= 0  AND rating <=5),
    photos text[],
    time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (hotel_id,user_id),
        CONSTRAINT fk_comments
           FOREIGN KEY(hotel_id)
           REFERENCES hotels(hotel_id)
         ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION public.trproc_upd_comments()
    RETURNS trigger AS
    $body$
    BEGIN
	    NEW.time:=CURRENT_TIMESTAMP;
	    RETURN NEW;
    END;
    $body$
LANGUAGE 'plpgsql';

CREATE TRIGGER tr_comments_on_change
    BEFORE UPDATE
    ON comments FOR EACH ROW
EXECUTE PROCEDURE public.trproc_upd_comments();

CREATE OR REPLACE FUNCTION public.trproc_upd_hotels()
    RETURNS trigger AS
$body$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        UPDATE hotels
        SET comm_count_for_each[NEW.rating+1] = comm_count_for_each[NEW.rating+1]::int + 1,
            comm_count = comm_count + 1
            WHERE hotel_id = NEW.hotel_id;
            RETURN OLD;
    ELSEIF (TG_OP = 'UPDATE') THEN
        UPDATE hotels
        SET comm_count_for_each[OLD.rating+1] = comm_count_for_each[OLD.rating+1]::int - 1,
            comm_count_for_each[NEW.rating+1] = comm_count_for_each[NEW.rating+1]::int + 1
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

CREATE TRIGGER tr_comments_on_add
    AFTER INSERT OR DELETE OR UPDATE
    ON comments FOR EACH ROW
EXECUTE PROCEDURE public.trproc_upd_hotels();