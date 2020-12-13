# CSV sql

artist_id | artist        | check
----------|---------------|-----------
0         | Cocteau Twins | 2020-06-21
1         | Chicane       |

album_id | album           | date
---------|-----------------|-----------
0        | Blue Bell Knoll | 1988-09-19
1        | Treasure        | 1984-10-01

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

- <https://github.com/mattn/go-sqlite3>
- <https://github.com/mithrandie/csvq-driver>
- <https://github.com/mithrandie/csvq>

## SQLite

- <https://sqlite.org/cli.html#export_to_csv>
- <https://sqlite.org/cli.html#importing_csv_files>
