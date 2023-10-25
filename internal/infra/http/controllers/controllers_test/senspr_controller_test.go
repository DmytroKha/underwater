package controllers_test

import "net/http"

var sensorControllerTests = []*requestTest{
	{
		"Get the average temperature detected by a particular sensor 'StatusOK'",
		"/api/v1/sensor/gamma 1/temperature/average?from=1634764800&till=1698058383",
		"GET",
		``,
		http.StatusOK,
		`(\d+\.\d+)`,
		"wrong sensor average temperature request",
	},
	{
		"Get the average temperature detected by a particular sensor 'StatusInternalServerError'",
		"/api/v1/sensor/dfgjkdllk/temperature/average?from=1634764800&till=1698058383",
		"GET",
		``,
		http.StatusInternalServerError,
		`{"error":"upper: no more rows in this result set"}`,
		"wrong sensor average temperature request",
	},
	{
		"Get the average temperature detected by a particular sensor 'StatusBadRequest'",
		"/api/v1/sensor/gamma 1/temperature/average?from=kj&till=ljl",
		"GET",
		``,
		http.StatusBadRequest,
		`{"error":"from parameter is not a valid UNIX timestamp"}`,
		"wrong sensor average temperature request",
	},
	{
		"Get the minimum temperature detected by sensors in the region 'StatusOK'",
		"/api/v1/region/temperature/min?xMin=10&xMax=20&yMin=20&yMax=25&zMin=30&zMax=32",
		"GET",
		``,
		http.StatusOK,
		`(\d+\.\d+)`,
		"wrong region minimum temperature request",
	},
	{
		"Get the minimum temperature detected by sensors in the region 'StatusBadRequest'",
		"/api/v1/region/temperature/min?xMin=fg&xMax=20&yMin=20&yMax=25&zMin=30&zMax=32",
		"GET",
		``,
		http.StatusBadRequest,
		`{"error":"xMin parameter is not valid"}`,
		"wrong region minimum temperature request",
	},
	{
		"Get the maximum temperature detected by sensors in the region 'StatusOK'",
		"/api/v1/region/temperature/max?xMin=10&xMax=20&yMin=20&yMax=25&zMin=30&zMax=32",
		"GET",
		``,
		http.StatusOK,
		`(\d+\.\d+)`,
		"wrong region maximum temperature request",
	},
	{
		"Get the maximum temperature detected by sensors in the region 'StatusBadRequest'",
		"/api/v1/region/temperature/max?xMin=fg&xMax=20&yMin=20&yMax=25&zMin=30&zMax=32",
		"GET",
		``,
		http.StatusBadRequest,
		`{"error":"xMin parameter is not valid"}`,
		"wrong region maximum temperature request",
	},
}
