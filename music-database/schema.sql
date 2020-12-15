create table album_t (album_n integer, album_s text, date_s text, url_s text);
create table artist_t (artist_n integer, artist_s text, check_s text);
create table song_album_t (song_n integer, album_n text);
create table song_artist_t (song_n integer, artist_n integer);
create table song_t (song_n integer, song_s text, note_s text);
