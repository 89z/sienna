func (p Player) Views() error {
   view, err := strconv.ParseFloat(p.ViewCount, 64)
   if err != nil { return err }
   hour, err := sinceHours(p.PublishDate)
   if err != nil { return err }
   view /= hour / 24 / 365
   format := numberFormat(view)
   if view > 10_000_000 {
      x.LogFail("Fail", format)
   } else {
      x.LogPass("Pass", format)
   }
   return nil
}
