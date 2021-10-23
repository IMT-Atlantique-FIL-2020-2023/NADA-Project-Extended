import "influxdata/influxdb/sample"

sample.data(set: "airSensor")
    |> rename(columns: {sensor_id: "sensorId"})
    |> filter(fn: (r) => r._field == "temperature")
    |> set(key: "airportId", value: "NTE")
    |> set(key: "_measurement", value: "temperature")
    |> to(
        org: "org.nada",
        bucket: "nada-bucket",
    )

sample.data(set: "airSensor")
    |> rename(columns: {sensor_id: "sensorId"})
    |> filter(fn: (r) => r._field == "humidity")
    |> set(key: "airportId", value: "NTE")
    |> set(key: "_measurement", value: "humidity")
    |> to(
        org: "org.nada",
        bucket: "nada-bucket",
    )

sample.data(set: "airSensor")
    |> rename(columns: {sensor_id: "sensorId"})
    |> filter(fn: (r) => r._field == "co")
    |> set(key: "airportId", value: "NTE")
    |> set(key: "_measurement", value: "co")
    |> to(
        org: "org.nada",
        bucket: "nada-bucket",
    )