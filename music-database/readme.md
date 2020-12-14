# Music database

## CSV import

~~~
sqlite3 music.db
.read music.sql
.mode csv
.import --skip 1 artist_t.csv artist_t
.import --skip 1 album_t.csv album_t
.import --skip 1 song_album_t.csv song_album_t
.import --skip 1 song_artist_t.csv song_artist_t
.import --skip 1 song_t.csv song_t
.quit
~~~

<https://sqlite.org/cli.html#importing_csv_files>

## CSV export

~~~
sqlite3 music.db
.mode csv
.once artist_t.csv
select * from artist_t;
.once album_t.csv
select * from album_t;
.once song_t.csv
select * from song_t;
.once song_album_t.csv
select * from song_album_t;
.once song_artist_t.csv
select * from song_artist_t;
~~~

<https://sqlite.org/cli.html#export_to_csv>

## Primary key

~~~
create table artist_t(artist_n integer primary key, artist_s);
insert into artist_t(artist_s) values('Cocteau Twins');
select * from artist_t;
~~~

~~~
GREEN 7,506,058 Sade Lovers Rock
GREEN 7,558,395 Pet Shop Boys Actually
GREEN 7,763,194 Pet Shop Boys Please
GREEN 7,772,017 Modest Mouse Good News for People Who Love Bad News
GREEN 7,937,060 Pearl Jam Jeremy
undefined property: viewCount Chris Isaak Heart Shaped World
undefined property: viewCount OutKast Stankonia
~~~

- <https://github.com/mattn/go-sqlite3/tree/master/_example>
- <https://golang.org/pkg/database/sql#DB.Query>
