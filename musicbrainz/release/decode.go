package decode

import (
   "encoding/json"
   "log"
   "net/http"
   "net/url"
)

type Map map[string]interface{}
type Slice []interface{}

type Decode struct {
   API string
   MBID string
   Query url.Values
}

func NewDecode(mbid_s string) Decode {
   return Decode{
      "https://musicbrainz.org/ws/2/release",
      mbid_s,
      url.Values{
         "fmt": []string{"json"},
         "inc": []string{"artist-credits", "recordings"},
      },
   }
}

func (dec_o Decode) Group() (Slice, error) {
   dec_o.Query["release-group"] = []string{dec_o.MBID}
   url_s := dec_o.API + "?" + dec_o.Query.Encode()
   println(url_s)
   get_o, e := http.Get(url_s)
   if e != nil {
      return nil, e
   }
   m := Map{}
   e = json.NewDecoder(get_o.Body).Decode(&m)
   if e != nil {
      return nil, e
   }
   return m["releases"], nil
}

func (dec_o Decode) Release() (Map, error) {
   url_s := dec_o.API + "/" + dec_o.MBID + "?" + dec_o.Query.Encode()
   println(url_s)
   get_o, e := http.Get(url_s)
   if e != nil {
      return nil, e
   }
   m := Map{}
   e = json.NewDecoder(get_o.Body).Decode(&m)
   if e != nil {
      return nil, e
   }
   return m, nil
}
