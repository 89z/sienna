<?php
declare(strict_types = 1);

class Radix64 {
   function __construct() {
      $s = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';
      $this->s_dig = $s;
   }

   function encode(int $n_in): string {
      $s_out = '';
      do {
         $s_out = $this->s_dig[$n_in % 64] . $s_out;
         $n_in = intdiv($n_in, 64);
      } while ($n_in > 0);
      return $s_out;
   }
}
