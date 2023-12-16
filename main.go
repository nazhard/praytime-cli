package main

import (
  "encoding/json"
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
  "time"
)

const DDMMYYYY = "02-01-2006"

func main()  {
  address := "Bandung"
  if len(os.Args) >= 2 {
    address = os.Args[1]
  }

  now := time.Now().UTC()
  date := now.Format(DDMMYYYY)

  res, err := http.Get("https://api.aladhan.com/v1/timingsByAddress/" + date + "?address=" + address + "&method=20")
  HandleError(err)
  
  defer res.Body.Close()
  
  if res.StatusCode == 200 {
    body, err := io.ReadAll(res.Body)
    HandleError(err)

    var r Response
    err = json.Unmarshal(body, &r)
    HandleError(err)

    today := r.Data.Date.Gregorian
    tim := r.Data.Timings
    
    fmt.Printf(
      "Pray time for: %s\n\nCurrent date: %s, %s %s %s \n\n",
      address,
      today.Weekday.En,
      today.Day,
      today.Month.En,
      today.Year,
    )
    fmt.Printf("Imsak   %s\n", tim.Imsak)
    fmt.Printf("Fajr    %s\n", tim.Fajr)
    fmt.Printf("Sunrise %s\n", tim.Sunrise)
    fmt.Printf("Dhuhr   %s\n", tim.Dhuhr)
    fmt.Printf("Asr     %s\n", tim.Asr)
    fmt.Printf("Maghrib %s\n", tim.Maghrib)
    fmt.Printf("Isha    %s\n", tim.Isha)
  } else {
    fmt.Println("There was an error while trying to get data")
  }
}

func HandleError(err error) {
  if err != nil {
    log.Fatalln(err)
  }
}
