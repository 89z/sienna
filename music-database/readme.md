# CSV sql

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

- <https://github.com/mattn/go-sqlite3>
- <https://github.com/mithrandie/csvq-driver>
- <https://github.com/mithrandie/csvq>
