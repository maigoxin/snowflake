package colorize

import (
  "time"
  "runtime"
  "fmt"
)

const (
  OFF = "\033[0m"
  RED = "\033[31m"
  GREEN = "\033[32m"
  YELLO = "\033[33m"
  BLUE = "\033[34m"
  PURPLE = "\033[35m"
  CYAN = "\033[36m"
  WHITE = "\033[37m"
)

var (
  IsDebug = false
)

func Color(color, message string)(string){
  return color  + message + OFF
}

func Red(message string)(string){
  return Color(RED, message)
}

func Green(message string)(string){
  return Color(GREEN, message)
}

func Yello(message string)(string){
  return Color(YELLO, message)
}

func Blue(message string)(string){
  return Color(BLUE, message)
}

func Purple(message string)(string){
  return Color(PURPLE, message)
}

func Cyan(message string)(string){
  return Color(CYAN, message)
}

func White(message string)(string){
  return Color(WHITE, message)
}

func Info(format string, a ...interface{}){
  if IsDebug {
    msg := fmt.Sprintf(format, a...)
    header := Header("Info")
    fmt.Println(White(header + msg))
  }
}

func Warn(format string, a ...interface{}){
  if IsDebug {
    msg := fmt.Sprintf(format, a...)
    header := Header("Warn")
    fmt.Println(Yello(header + msg))
  }
}

func Err(format string, a ...interface{}){
  if IsDebug {
    msg := fmt.Sprintf(format, a...)
    header := Header("Error")
    fmt.Println(Red(header + msg))
  }
}

func Header(h string)(string){
  pc, _, line, _:= runtime.Caller(2)
  f := runtime.FuncForPC(pc)
  return fmt.Sprintf("[%s][%s:%d][%v]", h, f.Name(), line, time.Now())
}
