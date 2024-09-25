This program processes a MaxMind `.mmdb` database file and removes all information except for the `latitude` and `longitude` fields inside of the `location`.

The purpose is to reduce the filesize of a database when you only need a subset of the information.

### Example

```
$ ./mmdb-latlongonly GeoLite2-City.mmdb GeoLite2-City-LatLongOnly.mmdb
NodeCount: 5,329,799
Processed  5,035,446 records
Reduced size of the database by 38.86%

$ mmdbctl metadata --data-types GeoLite2-City.mmdb
- Binary Format 2.0
- Database Type GeoLite2-City
- IP Version    6
- Record Size   28
- Node Count    5329799 (5.08 MB)
- Tree Size     37308593 (35.58 MB)
- Data Section Size 23839475 (22.74 MB)
    - Pointer Size  5571095 (5.31 MB)
    - UTF-8 String Size 354418210353 (330.08 GB)
    - Double Size   18382760 (17.53 MB)
    - Bytes Size    149670234548 (139.39 GB)
    - Unsigned 16-bit Integer Size 215685842178 (200.87 GB)
    - Unsigned 32-bit Integer Size 98483459588 (91.72 GB)
    - Signed 32-bit Integer Size 75618973 (72.12 MB)
    - Unsigned 64-bit Integer Size 198983857 (189.77 MB)
    - Unsigned 128-bit Integer Size 208099056 (198.46 MB)
    - Map Key-Value Pair Count 88056163579 (82.01 GB)
    - Array Length  166810431 (159.08 MB)
    - Float Size    10592 (10.34 KB)
- Data Section Start Offset 37308609 (35.58 MB)
- Data Section End Offset 61148084 (58.32 MB)
- Metadata Section Start Offset 61148098 (58.32 MB)
- Description
    en GeoLite2City database
- Languages     de, en, es, fr, ja, pt-BR, ru, zh-CN
- Build Epoch   1727177173

$ mmdbctl metadata --data-types GeoLite2-City-LatLongOnly.mmdb
- Binary Format 2.0
- Database Type GeoLite2-City-LatLongOnly
- IP Version    6
- Record Size   28
- Node Count    4970883 (4.74 MB)
- Tree Size     34796181 (33.18 MB)
- Data Section Size 2590454 (2.47 MB)
    - Pointer Size  243363 (237.66 KB)
    - UTF-8 String Size 37503822934 (34.93 GB)
    - Double Size   1029632 (1005.50 KB)
    - Bytes Size    24123392101 (22.47 GB)
    - Unsigned 16-bit Integer Size 76637906657 (71.37 GB)
    - Unsigned 32-bit Integer Size 24433239117 (22.76 GB)
    - Signed 32-bit Integer Size 33624 (32.84 KB)
    - Unsigned 64-bit Integer Size 51702308 (49.31 MB)
    - Unsigned 128-bit Integer Size 114231741 (108.94 MB)
    - Map Key-Value Pair Count 25877854273 (24.10 GB)
    - Array Length  10341493 (9.86 MB)
    - Float Size    1656 (1.62 KB)
- Data Section Start Offset 34796197 (33.18 MB)
- Data Section End Offset 37386651 (35.65 MB)
- Metadata Section Start Offset 37386665 (35.65 MB)
- Description
- Languages
- Build Epoch   1727177173
```
