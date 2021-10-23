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

keys = ["NTE"]

// Load in batch unique sensors for each airport specified by keys
keys = ["NTE"]

from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> keep(columns: ["airportId", "sensorId"])
    |> filter(fn: (r) => exists r.sensorId and exists r.airportId and contains(value: r.airportId, set: keys))
    |> distinct(column: "sensorId")