# dump1090-simulator

In order to ease the _What these filghts?_ development, this program is simulating the `dump1090` in,ayr output when queried via API `/data.json`.

## Getting started

Start the web server, with or withou prior setup
```bash
$ go run main.go
NO port defintion found in environment variable. Will use the default: 8080
Starting server on port 8080
NOTE: only ONE endpoint is valid: /data.json
TIP : curl -s http://localhost:8080/data.json

^C
$ export DUMP1090_SIMU_PORT=82
$ go run main.go
Port definition found in environment variable. Will use it: 82
Starting server on port 82
NOTE: only ONE endpoint is valid: /data.json
TIP : curl -s http://localhost:82/data.json
```

## Endpoint

There is only one endpoint: `/data.json`


```bash

$ curl -s http://localhost:82/data.json
[
{"hex":"34644c", "squawk":"0000", "flight":"", "lat":0.000000, "lon":0.000000, "validposition":0, "altitude":35975,  "vert_rate":0,"track":44, "validtrack":1,"speed":421, "messages":6, "seen":9},
{"hex":"40688a", "squawk":"2053", "flight":"BAW734  ", "lat":48.692198, "lon":2.342618, "validposition":1, "altitude":31000,  "vert_rate":0,"track":142, "validtrack":1,"speed":432, "messages":254, "seen":3},
{"hex":"4404a4", "squawk":"1000", "flight":"", "lat":0.000000, "lon":0.000000, "validposition":0, "altitude":15000,  "vert_rate":0,"track":0, "validtrack":0,"speed":0, "messages":11, "seen":11},
{"hex":"34734f", "squawk":"1000", "flight":"VLG30EY ", "lat":48.742940, "lon":2.548684, "validposition":1, "altitude":5125,  "vert_rate":2048,"track":87, "validtrack":1,"speed":239, "messages":258, "seen":0},
{"hex":"392ae2", "squawk":"1000", "flight":"AFR63PX ", "lat":48.697877, "lon":2.689121, "validposition":1, "altitude":15975,  "vert_rate":2432,"track":225, "validtrack":1,"speed":356, "messages":1040, "seen":0},
{"hex":"40653e", "squawk":"5727", "flight":"EZY29ZV ", "lat":49.079636, "lon":3.201740, "validposition":1, "altitude":38000,  "vert_rate":0,"track":316, "validtrack":1,"speed":417, "messages":85, "seen":14},
{"hex":"406b5b", "squawk":"5766", "flight":"BAW2741 ", "lat":48.929947, "lon":2.984830, "validposition":1, "altitude":36000,  "vert_rate":0,"track":324, "validtrack":1,"speed":435, "messages":249, "seen":0},
{"hex":"39cea6", "squawk":"1000", "flight":"TVF13QN ", "lat":48.754440, "lon":2.694186, "validposition":1, "altitude":9000,  "vert_rate":64,"track":82, "validtrack":1,"speed":258, "messages":1069, "seen":0},
{"hex":"3944f8", "squawk":"1000", "flight":"AFR52DA ", "lat":48.917650, "lon":2.892079, "validposition":1, "altitude":14500,  "vert_rate":0,"track":313, "validtrack":1,"speed":356, "messages":104, "seen":64},
{"hex":"3946e1", "squawk":"1000", "flight":"AFR86BE ", "lat":48.581869, "lon":2.506762, "validposition":1, "altitude":12875,  "vert_rate":3072,"track":0, "validtrack":0,"speed":360, "messages":2028, "seen":0},
{"hex":"3c0699", "squawk":"1000", "flight":"", "lat":0.000000, "lon":0.000000, "validposition":0, "altitude":4725,  "vert_rate":0,"track":266, "validtrack":1,"speed":264, "messages":133, "seen":53},
{"hex":"02012b", "squawk":"1020", "flight":"MAC821D ", "lat":48.931616, "lon":2.869889, "validposition":1, "altitude":14875,  "vert_rate":0,"track":313, "validtrack":1,"speed":359, "messages":56, "seen":154},
{"hex":"3c0ac7", "squawk":"1000", "flight":"CFG8LT  ", "lat":48.722270, "lon":3.109637, "validposition":1, "altitude":36000,  "vert_rate":0,"track":62, "validtrack":1,"speed":419, "messages":2525, "seen":0},
{"hex":"398a63", "squawk":"1000", "flight":"NAK444  ", "lat":0.000000, "lon":0.000000, "validposition":0, "altitude":2725,  "vert_rate":0,"track":0, "validtrack":0,"speed":0, "messages":278, "seen":216},
{"hex":"4d236f", "squawk":"1000", "flight":"LYX8101 ", "lat":0.000000, "lon":0.000000, "validposition":0, "altitude":27975,  "vert_rate":0,"track":0, "validtrack":0,"speed":0, "messages":1343, "seen":119},
{"hex":"3c49b6", "squawk":"1000", "flight":"TUI6DL  ", "lat":48.930130, "lon":3.484920, "validposition":1, "altitude":35125,  "vert_rate":-1024,"track":70, "validtrack":1,"speed":412, "messages":3904, "seen":23},
{"hex":"3444c3", "squawk":"7672", "flight":"AEA41HG ", "lat":48.501755, "lon":2.417344, "validposition":1, "altitude":17300,  "vert_rate":3520,"track":214, "validtrack":1,"speed":351, "messages":2395, "seen":239}
]


$ curl -s http://localhost:82/data.json
[
{"hex":"03e015", "squawk":"0000", "flight":"", "lat":0.000000, "lon":0.000000, "validposition":0, "altitude":40000,  "vert_rate":0,"track":0, "validtrack":0,"speed":0, "messages":4, "seen":6},
{"hex":"484164", "squawk":"2104", "flight":"KLM11S  ", "lat":48.841736, "lon":2.601154, "validposition":1, "altitude":39000,  "vert_rate":64,"track":201, "validtrack":1,"speed":477, "messages":404, "seen":0},
{"hex":"3420ca", "squawk":"7673", "flight":"IBE34AL ", "lat":48.674698, "lon":2.557702, "validposition":1, "altitude":7225,  "vert_rate":448,"track":224, "validtrack":1,"speed":264, "messages":1338, "seen":0},
{"hex":"4d20d8", "squawk":"5724", "flight":"VJT836  ", "lat":49.009460, "lon":2.817500, "validposition":1, "altitude":38000,  "vert_rate":0,"track":320, "validtrack":1,"speed":462, "messages":533, "seen":16},
{"hex":"4d02dd", "squawk":"5777", "flight":"JFA77E  ", "lat":48.619436, "lon":2.641995, "validposition":1, "altitude":22000,  "vert_rate":0,"track":312, "validtrack":1,"speed":288, "messages":661, "seen":0},
{"hex":"a026a5", "squawk":"1000", "flight":"", "lat":48.931915, "lon":2.868770, "validposition":1, "altitude":15000,  "vert_rate":0,"track":313, "validtrack":1,"speed":353, "messages":125, "seen":139},
{"hex":"3cd636", "squawk":"5732", "flight":"", "lat":0.000000, "lon":0.000000, "validposition":0, "altitude":38000,  "vert_rate":0,"track":330, "validtrack":1,"speed":437, "messages":30, "seen":74},
{"hex":"4d2265", "squawk":"1000", "flight":"RYR6024 ", "lat":48.912064, "lon":2.955178, "validposition":1, "altitude":31950,  "vert_rate":0,"track":331, "validtrack":1,"speed":407, "messages":553, "seen":0},
{"hex":"39ce41", "squawk":"1000", "flight":"VLJ994N ", "lat":48.537277, "lon":2.633972, "validposition":1, "altitude":21000,  "vert_rate":0,"track":191, "validtrack":1,"speed":366, "messages":1670, "seen":6},
{"hex":"484b6e", "squawk":"2111", "flight":"TRA5679 ", "lat":48.409332, "lon":2.571012, "validposition":1, "altitude":39000,  "vert_rate":0,"track":189, "validtrack":1,"speed":476, "messages":2098, "seen":0},
{"hex":"0a004a", "squawk":"7601", "flight":"DAH1123 ", "lat":48.536041, "lon":2.438965, "validposition":1, "altitude":11950,  "vert_rate":2688,"track":228, "validtrack":1,"speed":367, "messages":2040, "seen":6},
{"hex":"34644c", "squawk":"1000", "flight":"IBS3672 ", "lat":48.931274, "lon":2.955674, "validposition":1, "altitude":36000,  "vert_rate":0,"track":71, "validtrack":1,"speed":423, "messages":990, "seen":2},
{"hex":"40688a", "squawk":"2053", "flight":"BAW734  ", "lat":48.417160, "lon":2.663058, "validposition":1, "altitude":31025,  "vert_rate":0,"track":142, "validtrack":1,"speed":430, "messages":1635, "seen":137},
{"hex":"4404a4", "squawk":"1000", "flight":"EJU53RN ", "lat":0.000000, "lon":0.000000, "validposition":0, "altitude":14000,  "vert_rate":-832,"track":314, "validtrack":1,"speed":355, "messages":68, "seen":227},
{"hex":"34734f", "squawk":"1000", "flight":"VLG30EY ", "lat":48.766123, "lon":2.846615, "validposition":1, "altitude":9000,  "vert_rate":0,"track":83, "validtrack":1,"speed":271, "messages":1433, "seen":107},
{"hex":"392ae2", "squawk":"1000", "flight":"AFR63PX ", "lat":48.347432, "lon":2.461516, "validposition":1, "altitude":23750,  "vert_rate":448,"track":163, "validtrack":1,"speed":418, "messages":2902, "seen":58},
{"hex":"406b5b", "squawk":"5766", "flight":"BAW2741 ", "lat":48.965012, "lon":2.945815, "validposition":1, "altitude":36000,  "vert_rate":0,"track":324, "validtrack":1,"speed":435, "messages":321, "seen":236},
{"hex":"39cea6", "squawk":"1000", "flight":"TVF13QN ", "lat":48.762817, "lon":2.804401, "validposition":1, "altitude":9000,  "vert_rate":0,"track":83, "validtrack":1,"speed":251, "messages":1497, "seen":233},
{"hex":"3946e1", "squawk":"1000", "flight":"AFR86BE ", "lat":48.387187, "lon":2.306768, "validposition":1, "altitude":20375,  "vert_rate":2240,"track":0, "validtrack":0,"speed":375, "messages":2440, "seen":172},
{"hex":"3c0ac7", "squawk":"1000", "flight":"CFG8LT  ", "lat":48.806992, "lon":3.355126, "validposition":1, "altitude":35975,  "vert_rate":64,"track":62, "validtrack":1,"speed":420, "messages":2714, "seen":173}
]

```


## Data

Data are fake. There are 29 different data set, retrieved from real data.

If you request more than 29 times the API, it will loop back to first result.