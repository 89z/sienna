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

func main() {
   packSet := map[string]bool{}
   for packs := []string{target}; len(packs) > 0; packs = packs[1:] {
      target := packs[0]
      deps, e := man.getValue(target, "%DEPENDS%")
      if e != nil {
         log.Fatal(e)
      }
      packs = append(packs, deps...)
      if packSet[target] {
         continue
      }
      println(target)
      packSet[target] = true
   }
}

type description struct {
   name string
   filename string
   provides []string
   depends []string
}

func newDescription(file, repo, variant string) (description, error) {
   open, e := os.Open(file)
   if e != nil {
      return description{}, e
   }
   scan := bufio.NewScanner(open)
   var desc description
   for scan.Scan() {
      switch scan.Text() {
      case "%FILENAME%":
         scan.Scan()
         mirror.Path = path.Join(repo, variant, scan.Text())
         desc.filename = mirror.String()
      case "%NAME%":
         scan.Scan()
         desc.name = scan.Text()
      case "%DEPENDS%":
         for scan.Scan() {
            line := scan.Text()
            if line == "" {
               break
            }
            desc.depends = append(
               desc.depends, baseName(line, "=>"),
            )
         }
      case "%PROVIDES%":
         for scan.Scan() {
            line := scan.Text()
            if line == "" {
               break
            }
            desc.provides = append(desc.provides, line)
         }
      }
   }
   return desc, nil
}
