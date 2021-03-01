
# 普通SQL
1. `UNIX_TIMESTAMP(ss.createdate) > UNIX_TIMESTAMP('2018-12-04 11:39:05')`


# Hive



# solr

- checkstatus_ld:(-1)
- 包含，分词有tsd 比如： `category_tsd:收纳整理` —— 会切分成字，一个个字包含。
  - 如果需要词： `category_tsd:"收纳整理"`
- 或：`category_tsd:"收纳" OR category_tsd:"整理"`