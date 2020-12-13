package release

func (m Map) Date() (string, bool) {
   s, b := m.S("date")
   if s == "" {
      return "", false
   }
   return s + "-12-31", true
}

func (m Map) IsOfficial() bool {
   return m["status"] == "Official"
}

func (m Map) TrackLen() int {
   n := 0
   for _, item_m := range m.A("media") {
      n += item_m["track-count"]
   }
   return n
}

func Reduce(old_n int, new_m Map, new_n int, old_a Slice) int {
   if new_n == 0 {
      return 0
   }
   old_m := old_a[old_n]
   date_old_s, b := old_m.Date()
   if ! b {
      return new_n
   }
   date_new_s, b := new_m.Date()
   if ! b {
      return old_n
   }
   if ! new_m.IsOfficial() {
      return old_n
   }
   if date_new_s > date_old_s {
      return old_n
   }
   if date_new_s < date_old_s {
      return new_n
   }
   if new_m.TrackLen() >= old_m.TrackLen() {
      return old_n
   }
   return new_n
}
