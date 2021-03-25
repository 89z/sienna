func (m manager) sync(tar string) error {
   open, e := os.Open(tar)
   if e != nil {
      return e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      values, e := m.getValue(
         scan.Text(), "%FILENAME%",
      )
      if e != nil {
         return e
      }
      file := values[0]
      archive := path.Join(m.Cache, file)
      _, e = x.Copy(
         getRepo(file) + file, archive,
      )
      if e != nil {
         return e
      }
      e = unarchive(archive, m.Dest)
      if e != nil {
         return e
      }
   }
   return nil
}
