package main

import (
  "flag"
  "fmt"
  "net/http"
  "strconv"
  "github.com/maigoxin/snowflake/colorize"
  "github.com/maigoxin/snowflake/id"
)

var generator *id.Id

func RequestHanler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  count_num := 1
  if count, ok := r.Form[`count`]; ok == true {
    count_num, _ = strconv.Atoi(count[0])
  }

  w.Header().Set("Content-Type", "text/plain; charset=utf-8")

  if ids, err := generator.NextIds(count_num); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, err)
    colorize.Err(`err:%s`, err)
  }else {
    for _, value := range ids {
      w.Write([]byte(fmt.Sprintln(value)))
    }
    colorize.Info(`generate:%x`, ids)
  }
}

func main() {
  var port = flag.Int("port", 80, `port, default is 80`)
  var isDebug = flag.Bool("debug", false, `debug, default is false`)
  var centerId = flag.Int64("center", 0, `centerId, default is 0`)
  var workerId = flag.Int64("worker", 0, `workerId, default is 0`)
  var twepoch = flag.Int64("twepoch", 0, `twepoch, default is 0`)
  flag.Parse()

  generator, _  = id.NewId(*workerId, *centerId, *twepoch)
  colorize.IsDebug = *isDebug
  colorize.Info(`going to run :%d`, *port)

  http.HandleFunc("/", RequestHanler)
  http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
