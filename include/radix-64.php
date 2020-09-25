<?php
declare(strict_types = 1);
error_reporting(E_ALL);

class Radix64 {
   var $s_dig = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';

   function encode(int $n_in): string {
      $s_out = '';
      do {
         $s_out = $this->s_dig[$n_in % 64] . $s_out;
         $n_in = intdiv($n_in, 64);
      } while ($n_in > 0);
      return $s_out;
   }
}
