# CSV sql

artist_id | artist
----------|--------------
0         | Cocteau Twins
1         | Chicane

album_id | album           | date       | youtube | musicbrainz
---------|-----------------|------------|---------|------------
0        | Blue Bell Knoll | 1988-09-19 |         |
1        | Treasure        | 1984-10-01 |         |

song_id | song       | note
--------|------------|-----
0       | Ivo        | good
1       | Persephone | good

song_id | album_id
--------|---------
0       | 1
1       | 1

song_id | artist_id
--------|----------
0       | 0
1       | 0

~~~
create table artist_t(artist_n integer primary key, artist_s);
insert into artist_t(artist_s) values('Cocteau Twins');
select * from artist_t;
~~~

- <https://github.com/mithrandie/csvq-driver>
- <https://github.com/mithrandie/csvq>
- <https://sqlite.org/autoinc.html>
