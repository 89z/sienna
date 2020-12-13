package release

func (m Map) Date() string {
   s := m.S("date")
   switch len(s) {
   case 0:
      return "9999-12-31"
   case 4:
      return s + "-12-31"
   case 6:
      return s + "-31"
   default:
      return s
   }
}

func (m Map) IsOfficial() bool {
   return m["status"] == "Official"
}

func (m Map) TrackLen() float64 {
   track_n := float64(0)
   a := m.A("media")
   for n := range a {
      track_n += a.M(n).N("track-count")
   }
   return track_n
}

func Reduce(old_n int, new_m Map, new_n int, old_a Slice) int {
   if new_n == 0 {
      return 0
   }
   old_m := old_a.M(old_n)
   if ! new_m.IsOfficial() {
      return old_n
   }
   if new_m.Date() > old_m.Date() {
      return old_n
   }
   if new_m.Date() < old_m.Date() {
      return new_n
   }
   if new_m.TrackLen() >= old_m.TrackLen() {
      return old_n
   }
   return new_n
}
