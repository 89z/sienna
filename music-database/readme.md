# CSV sql

artist_n | artist_s
---------|--------------
0        | Cocteau Twins
1        | Chicane

album_n | album_s         | date_d     | url_s
--------|-----------------|------------|------
0       | Blue Bell Knoll | 1988-09-19 |
1       | Treasure        | 1984-10-01 |

song_n | song_s     | note_s
-------|------------|-------
0      | Ivo        | good
1      | Persephone | good

song_n | album_n
-------|--------
0      | 1
1      | 1

song_n | artist_n
-------|---------
0      | 0
1      | 0

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

- <https://github.com/mithrandie/csvq-driver>
- <https://github.com/mithrandie/csvq>
