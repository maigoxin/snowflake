package id

import (
  "fmt"
  "sync"
  "time"
  "errors"
)

const (
  workerIdBits       = uint(5)
  datacenterIdBits   = uint(5)
  sequenceBits       = uint(12)
  workerIdShift      = sequenceBits
  datacenterIdShift  = sequenceBits + workerIdBits
  timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits
  sequenceMask       = -1 ^ (-1 << sequenceBits)
  maxNextIdsNum      = 1024
)

type Id struct {
  sequence      int64
  lastTimestamp int64
  workerId      int64
  twepoch       int64
  datacenterId  int64
  mutex         sync.Mutex
}

func NewId(workerId, datacenterId, twepoch int64)(*Id, error){
  id := &Id{}
  id.sequence = 0
  id.lastTimestamp = -1
  id.twepoch = twepoch
  id.mutex = sync.Mutex{}

  return id, nil
}

func timeGen() (int64) {
  return time.Now().UnixNano()/int64(time.Millisecond)
}

func utillNextMillis(lastTimestamp int64) (int64){
  timestamp := timeGen()
  for timestamp <= lastTimestamp {
    timestamp = timeGen()
  }
  return timestamp
}

func (id *Id) NextId() (int64, error) {
  id.mutex.Lock()
  defer id.mutex.Unlock()
  timestamp := timeGen()

  if id.lastTimestamp == timestamp {
    id.sequence = (id.sequence + 1) & sequenceMask
    if id.sequence == 0 {
      timestamp = utillNextMillis(id.lastTimestamp)
    }
  } else {
    id.sequence = 0
  }
  id.lastTimestamp = timestamp
  return ((timestamp - id.twepoch) << timestampLeftShift) | (id.datacenterId << datacenterIdShift) | (id.workerId << workerIdShift) | id.sequence, nil
}

func (id *Id) NextIds(num int) ([]int64, error) {
  if num > maxNextIdsNum || num < 0 {
    return nil, errors.New(fmt.Sprintf("NextIds num: %d error", num))
  }
  ids := make([]int64, num)
  id.mutex.Lock()
  defer id.mutex.Unlock()

  for i := 0; i < num; i++ {
    timestamp := timeGen()
    if id.lastTimestamp == timestamp {
      id.sequence = (id.sequence + 1) & sequenceMask
      if id.sequence == 0 {
        timestamp = utillNextMillis(id.lastTimestamp)
      }
    } else {
      id.sequence = 0
    }
    id.lastTimestamp = timestamp
    ids[i] = ((timestamp - id.twepoch) << timestampLeftShift) | (id.datacenterId << datacenterIdShift) | (id.workerId << workerIdShift) | id.sequence
  }
  return ids, nil
}
