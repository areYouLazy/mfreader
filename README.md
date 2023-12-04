# mfreader

A tool inspired by [mfdread](https://github.com/zhovner/mfdread) to print MIFARE products dumps in a human readable format

## Index

- [mfreader](#mfreader)
  - [Index](#index)
  - [Usage](#usage)
  - [Examples](#examples)
  - [References](#references)

## Usage

```console
root@localhost$ mfreader -h
mfreader, a tool inspired by mfdread ( https://github.com/zhovner/mfdread )
to print MIFARE dumps in a human readable format

Usage:
  mfreader [flags]

Flags:
  -f, --file string    dump file to analyze
  -h, --help           help for mfreader
  -j, --json           output in JSON format
  -m, --manufacturer   print manufacturer data
```

## Examples

```console
root@localhost$ # Get only manufacturer data
root@localhost$ mfreader -f dump.mfd -m
+----------+-----+-----+------+
| UID      | BCC | SAK | ATQA |
+----------+-----+-----+------+
| 33bd9d3f | 2c  | 98  | 02   |
+----------+-----+-----+------+
```

```console
root@localhost$ # Get only manufacturer data in json format
root@localhost$ mfreader -f dump.mfd -m -j
{
  "UID": "33bd9d3f",
  "BCC": "2c",
  "SAK": "98",
  "ATQA": "02"
}
```

```console
root@localhost$ # Default output (without manufacturer data by default)
root@localhost$ mfreader -f dump.mfd
+--------+-------+----------------------------------+-------------+-------------------------------------+
| SECTOR | BLOCK |               DATA               | ACCESS BITS |                RIGHTS               |
|        |       |                                  |             |                                     |
|        |       |      KEY A ACCESS BITS KEY B     |             |                                     |
+--------+-------+----------------------------------+-------------+-------------------------------------+
|    0   |   0   | 33bd9d3f2c980200648f841441502212 |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    0   |   1   | 090f180800000000000003010000400b |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    0   |   2   | 00000000400c400c400c000400040005 |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    0   |   3   | a0a1a2a3a4a5787788c17de02a7f6025 |     011     | - B | A/B B | - B                   |
+--------+-------+----------------------------------+-------------+-------------------------------------+
|    1   |   4   | 418d50c98d7f962462004c800000ffcc |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    1   |   5   | 1fa1014100d101c060000000049a2a9f |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    1   |   6   | 1fa1014100d101c060000000049a2a9f |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    1   |   7   | 2735fc18180778778800bf23a53c1f63 |     011     | - B | A/B B | - B                   |
+--------+-------+----------------------------------+-------------+-------------------------------------+
|    2   |   8   | 3065061730077220296012505b74c05d |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    2   |   9   | 68c701da24c027ece0ee9a99c0caadb1 |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    2   |   10  | c82591842f0b8304a2a068d1f4e016e7 |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    2   |   11  | 2aba9519f574787788ffcb9a1f2d7368 |     011     | - B | A/B B | - B                   |
+--------+-------+----------------------------------+-------------+-------------------------------------+
|    3   |   12  | 6c135ade77c0f7a11f09ad059d45720c |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    3   |   13  | 3c0dc85010e3ef723bfad584c4ad509d |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    3   |   14  | 040e821625f14168040ed8ee61a8f635 |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    3   |   15  | 84fd7f7a12b6787788ffc7c0adb3284f |     011     | - B | A/B B | - B                   |
+--------+-------+----------------------------------+-------------+-------------------------------------+
|    4   |   16  | 420d53f9dbd3362461004c800000bc18 |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    4   |   17  | 1f51014100d101c0900004240280bdce |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    4   |   18  | 1f51014100d101c0900004240280bdce |     100     | A/B |   B   |  -  |  -  [r/w]       |
|    4   |   19  | 73068f118c13787788002b7f3253fac5 |     011     | - B | A/B B | - B                   |
+--------+-------+----------------------------------+-------------+-------------------------------------+
|    5   |   20  | 00000000000000000000000000000000 |     110     | A/B |   B   |   B | A/B [value]     |
|    5   |   21  | 01770000907222029653352020202020 |     110     | A/B |   B   |   B | A/B [value]     |
|    5   |   22  | 00000000000000000000000000000000 |     110     | A/B |   B   |   B | A/B [value]     |
|    5   |   23  | 186d8c4b93f908778f029f131d8c2057 |     011     | - B | A/B B | - B                   |
+--------+-------+----------------------------------+-------------+-------------------------------------+
|    6   |   24  | 00000000000000000000000000000000 |     110     | A/B |   B   |   B | A/B [value]     |
|    6   |   25  | 00000000000000000000000000000000 |     110     | A/B |   B   |   B | A/B [value]     |
|    6   |   26  | 00000000000000000000000000000000 |     110     | A/B |   B   |   B | A/B [value]     |
|    6   |   27  | 3a4bba8adaf008778f0267362d90f973 |     011     | - B | A/B B | - B                   |
+--------+-------+----------------------------------+-------------+-------------------------------------+
[...]
```

```console
root@localhost$ # JSON output (includes manufacturer data by default)
root@localhost$ mfreader -f dump.mfd -j
{
  "Sectors": [
    {
      "Number": 0,
      "Blocks": [
        {
          "Number": 0,
          "Position": 0,
          "Rights": "100",
          "Data": "33bd9d3f2c980200648f841441502212",
          "IsTrailerSector": false
        },
        {
          "Number": 1,
          "Position": 1,
          "Rights": "100",
          "Data": "090f180800000000000003010000400b",
          "IsTrailerSector": false
        },
        {
          "Number": 2,
          "Position": 2,
          "Rights": "100",
          "Data": "00000000400c400c400c000400040005",
          "IsTrailerSector": false
        },
        {
          "Number": 3,
          "Position": 3,
          "Rights": "011",
          "Data": "a0a1a2a3a4a5787788c17de02a7f6025",
          "IsTrailerSector": true
        }
      ],
      "AccessBytes": "787788",
      "Data": "33bd9d3f2c980200648f841441502212090f180800000000000003010000400b00000000400c400c400c000400040005a0a1a2a3a4a5787788c17de02a7f6025"
    },
    {
      "Number": 1,
      "Blocks": [
        {
          "Number": 4,
          "Position": 0,
          "Rights": "100",
          "Data": "418d50c98d7f962462004c800000ffcc",
          "IsTrailerSector": false
        },
        {
          "Number": 5,
          "Position": 1,
          "Rights": "100",
          "Data": "1fa1014100d101c060000000049a2a9f",
          "IsTrailerSector": false
        },
        {
          "Number": 6,
          "Position": 2,
          "Rights": "100",
          "Data": "1fa1014100d101c060000000049a2a9f",
          "IsTrailerSector": false
        },
        {
          "Number": 7,
          "Position": 3,
          "Rights": "011",
          "Data": "2735fc18180778778800bf23a53c1f63",
          "IsTrailerSector": true
        }
      ],
      "AccessBytes": "787788",
      "Data": "418d50c98d7f962462004c800000ffcc1fa1014100d101c060000000049a2a9f1fa1014100d101c060000000049a2a9f2735fc18180778778800bf23a53c1f63"
    },
    [...]
  ],
  "Size": 4096,
  "UID": "33bd9d3f",
  "BCC": "2c",
  "SAK": "98",
  "ATQA": "02"
}
```

## References

- https://www.nxp.com/docs/en/data-sheet/MF1S50YYX_V1.pdf
- https://www.nxp.com/docs/en/application-note/AN10833.pdf
