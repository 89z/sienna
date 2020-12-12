package decode

class MusicBrainzRelease {
   function __construct($release_o) {
      foreach ($release_o as $k => $v) {
         $this->$k = $v;
      }
   }
   function date_b(): bool {
      if (! property_exists($this, 'date')) {
         return false;
      }
      if ($this->date == '') {
         return false;
      }
      return true;
   }
   function date_s(): string {
      return $this->date . '-12-31';
   }
   function status(): bool {
      return $this->status == 'Official';
   }
   function tracks(): int {
      $ca_n = 0;
      foreach ($this->media as $it_o) {
         $ca_n += $it_o->{'track-count'};
      }
      return $ca_n;
   }
}

func Reduce(acc_n int, cur_m Map, idx_n int, src_a Slice) int {
   if idx_n == 0 {
      return 0
   }
   $old_o = new MusicBrainzRelease($src_a[$acc_n]);
   if (! $old_o->date_b()) {
      return $idx_n;
   }
   $new_o = new MusicBrainzRelease($cur_o);
   if (! $new_o->date_b()) {
      return $acc_n;
   }
   if (! $new_o->status()) {
      return $acc_n;
   }
   if ($new_o->date_s() > $old_o->date_s()) {
      return $acc_n;
   }
   if ($new_o->date_s() < $old_o->date_s()) {
      return $idx_n;
   }
   if ($new_o->tracks() >= $old_o->tracks()) {
      return $acc_n;
   }
   return $idx_n;
}
