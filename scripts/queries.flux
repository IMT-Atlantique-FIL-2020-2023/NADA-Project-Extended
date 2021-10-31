from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> keep(columns: ["airportId"])
    |> filter(fn: (r) => exists r.airportId)
    |> rename(columns: {airportId: "_value"})

from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> keep(columns: ["airportId"])
    |> filter(fn: (r) => r.airportId == "NTE")
    |> rename(columns: {airportId: "_value"})
    |> first()

// Load in batch unique sensors for each airport specified by keys, filter by airport
import "dict"

keys = ["NTE",]
sensorsIds = ["":[],"NTE":[],]

filterFunc = (r) => {
    item = dict.get(dict: sensorsIds, key: r.airportId, default: [])
    return (contains(value: r.sensorId, set: item))
}

from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> filter(fn: (r) => exists r.sensorId and exists r.airportId)
    |> filter(fn: filterFunc)
    |> distinct(column: "sensorId")
    |> keep(columns: ["airportId", "_measurement", "_value"])
    |> rename(columns: {airportId: "_field"})

// Load in batch unique sensors for each airport specified by keys, filter by airport
import "dict"

keys = ["NTE",]
sensorsIds = ["":[],"NTE":["TLM0101"],]
timeRangeStart = 2018-05-22T23:30:00Z
timeRangeEnd = 2023-05-22T23:30:00Z
everyValue = 1h
discritzeInto = 10
discretizeMode = true
duration = if discretizeMode then duration(v: (int(v: timeRangeEnd) - int(v:timeRangeStart)) / discritzeInto) else everyValue

filterFunc = (r) => {
    item = dict.get(dict: sensorsIds, key: r.airportId, default: [])
    len = length(arr: item) 
    return (len == 0 or contains(value: r.sensorId, set: item))
}

from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> filter(fn: (r) => exists r.sensorId and exists r.airportId)
    |> filter(fn: filterFunc)
    |> window(every: duration,  createEmpty: true)
    |> mean()
    |> keep(columns: ["airportId", "_measurement", "_value", "_start", "_stop", "_time", "sensorId"])
    |> rename(columns: {airportId: "_field"})