package youtube

import (
   "math"
   "net/url"
)

func NumberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n3 := n / math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n3) + []string{"", " k", " M", " B"}[n2]
}

func GetContents(s string) (string, error) {
   o, e := http.Get(s)
   if e != nil {
      return "", e
   }
   y, e := ioutil.ReadAll(o.Body)
   if e != nil {
      return "", e
   }
   return string(y), nil
}

function Info(id_s string) (Map, error) {
   info_s := "https://www.youtube.com/get_video_info?video_id=" + id_s
   println(info_s)
   query_s, e := GetContents(info_s)
   if e != nil {
      return nil, e
   }
   m, e := url.ParseQuery(query_s)
   if e != nil {
      return nil, e
   }
   resp_s := m["player_response"]
   return json_decode($resp_s)->microformat->playerMicroformatRenderer;
}

function youtube_views(object $info_o): string {
   $views_n = (int)($info_o->viewCount);
   $date_o = DateTime::createFromFormat('!Y-m-d', $info_o->publishDate);
   $sec_n = seconds(new DateTime, $date_o);
   $rate_n = $views_n / ($sec_n / 60 / 60 / 24 / 365);
   $rate_s = format_number($rate_n);
   if ($rate_n > 8_000_000) {
      return 'RED ' . color_red($rate_s);
   }
   return 'GREEN ' . color_green($rate_s);
}
