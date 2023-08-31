# ORCReader 1337

Simple CLI tool that uses the ORC-library https://github.com/scritchley/orc/ to read orc files and provide useful information about the given file. 

## Run
`go run main.go`

## Build
`go build main.go`

## How to use

**Help**
```
./orcreader
-- O R C R e a d e r 1337 -- 
Usage: ./orcreader <action> <inputfile>
Available actions: read, schema, schemajson, schemasql, readjson, metadata, count
Example: ./orcreader read myfile.orc
```

**Read**
```
./orcreader read examples/test.orc
[1 M M Primary 500 Good 0 0 0]
[2 F M Primary 500 Good 0 0 0]
[3 M S Primary 500 Good 0 0 0]
[4 F S Primary 500 Good 0 0 0]
[5 M D Primary 500 Good 0 0 0]
[6 F D Primary 500 Good 0 0 0]
[7 M W Primary 500 Good 0 0 0]
[8 F W Primary 500 Good 0 0 0]
[9 M U Primary 500 Good 0 0 0]
[10 F U Primary 500 Good 0 0 0]
[11 M M Secondary 500 Good 0 0 0]
[12 F M Secondary 500 Good 0 0 0]
```

**Readjson**
```
./orcreader readjson examples/test.orc
{"_col0": 638, "_col1": "F", "_col2": "W", "_col3": "Primary", "_col4": 5000, "_col5": "Good", "_col6": 0, "_col7": 0, "_col8": 0}
{"_col0": 639, "_col1": "M", "_col2": "U", "_col3": "Primary", "_col4": 5000, "_col5": "Good", "_col6": 0, "_col7": 0, "_col8": 0}
{"_col0": 640, "_col1": "F", "_col2": "U", "_col3": "Primary", "_col4": 5000, "_col5": "Good", "_col6": 0, "_col7": 0, "_col8": 0}
{"_col0": 641, "_col1": "M", "_col2": "M", "_col3": "Secondary", "_col4": 5000, "_col5": "Good", "_col6": 0, "_col7": 0, "_col8": 0}
{"_col0": 642, "_col1": "F", "_col2": "M", "_col3": "Secondary", "_col4": 5000, "_col5": "Good", "_col6": 0, "_col7": 0, "_col8": 0}
{"_col0": 643, "_col1": "M", "_col2": "S", "_col3": "Secondary", "_col4": 5000, "_col5": "Good", "_col6": 0, "_col7": 0, "_col8": 0}
{"_col0": 644, "_col1": "F", "_col2": "S", "_col3": "Secondary", "_col4": 5000, "_col5": "Good", "_col6": 0, "_col7": 0, "_col8": 0}
```

**Schema**
```
./orcreader schema examples/test.orc
struct<_col0:int,_col1:string,_col2:string,_col3:string,_col4:int,_col5:string,_col6:int,_col7:int,_col8:int>
```

**SchemaJSON**
```
./orcreader schemajson examples/test.orc                                                                                                           20:35:02
{"category": "struct", "id": 0, "max": 9, "fields": {"_col0": {"category": "int", "id": 1, "max": 1},"_col1": {"category": "string", "id": 2, "max": 2},"_col2": {"category": "string", "id": 3, "max": 3},"_col3": {"category": "string", "id": 4, "max": 4},"_col4": {"category": "int", "id": 5, "max": 5},"_col5": {"category": "string", "id": 6, "max": 6},"_col6": {"category": "int", "id": 7, "max": 7},"_col7": {"category": "int", "id": 8, "max": 8},"_col8": {"category": "int", "id": 9, "max": 9}}}
```

**SchemaSQL**
```
./orcreader schemasql examples/test.orc
CREATE TABLE mytable (
    _col0 INTEGER, 
    _col1 VARCHAR, 
    _col2 VARCHAR, 
    _col3 VARCHAR, 
    _col4 INTEGER, 
    _col5 VARCHAR, 
    _col6 INTEGER, 
    _col7 INTEGER, 
    _col8 INTEGER
)
WITH (format = 'ORC')
```

**Metadata**
```
./orcreader metadata examples/test.orc
stripeStats:
<colStats:
<numberOfValues:1920800 >
 colStats:
<numberOfValues:1920800 intStatistics:
<minimum:1 maximum:1920800 sum:1844737280400 >
 >
 colStats:
<numberOfValues:1920800 stringStatistics:
<minimum:"F" maximum:"M" sum:1920800 >
 >
 colStats:
<numberOfValues:1920800 stringStatistics:
<minimum:"D" maximum:"W" sum:1920800 >
 >
 colStats:
<numberOfValues:1920800 stringStatistics:
<minimum:"2 yr Degree" maximum:"Unknown" sum:18384800 >
 >
 colStats:
<numberOfValues:1920800 intStatistics:
<minimum:500 maximum:10000 sum:10084200000 >
 >
 colStats:
<numberOfValues:1920800 stringStatistics:
<minimum:"Good" maximum:"Unknown" sum:13445600 >
 >
 colStats:
<numberOfValues:1920800 intStatistics:
<minimum:0 maximum:6 sum:5762400 >
 >
 colStats:
<numberOfValues:1920800 intStatistics:
<minimum:0 maximum:6 sum:5762400 >
 >
 colStats:
<numberOfValues:1920800 intStatistics:
<minimum:0 maximum:6 sum:5762400 >
 >
 >
```

**Count**
```
./orcreader count examples/test.orc
1920800
```