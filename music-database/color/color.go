package color

func Cyan(s string) string {
   return "\x1b[1;36m" + s + "\x1b[m"
}

func Green(s string) string {
   return "\x1b[1;32m" + s + "\x1b[m"
}

func Red(s string) string {
   return "\x1b[1;31m" + s + "\x1b[m"
}

func Yellow(s string) string {
   return "\x1b[30;43m" + s + "\x1b[m"
}
