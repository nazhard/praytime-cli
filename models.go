package main

type Response struct {
  Data struct {
    Timings struct {
      Fajr     string `json:"Fajr"`
      Sunrise  string `json:"Sunrise"`
      Dhuhr    string `json:"Dhuhr"`
      Asr      string `json:"Asr"`
      Maghrib  string `json:"Maghrib"`
      Isha     string `json:"Isha"`
      Imsak    string `json:"Imsak"`
      Midnight string `json:"Midnight"`
    } `json:"timings"`
    Date struct {
      Gregorian struct {
        Day string `json:"day"`
        Weekday struct {
          En string `json:"en"`
        } `json:"weekday"`
        Month struct {
          En string `json:"en"`
        }
        Year string `json:"year"`
      } `json:"gregorian"`
    } `json:"date"`
  } `json:"data"`
}
