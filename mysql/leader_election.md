CreateOrUpdate SQL:
```
INSERT IGNORE INTO leader_election (service_id, leader_id, last_seen_active)   VALUES (?, ?, now()) ON DUPLICATE KEY UPDATE leader_id = IF(last_seen_active                            <DATE_SUB(VALUES(last_seen_active), INTERVAL 30 SECOND), VALUES(leader_id), leader_id),last_seen_active = IF(leader_id =VALUES(leader_id), VALUES(last_seen_active),last_seen_active);
```

```
CREATE TABLE `leader_election` (
  `service_id` varchar(128) NOT NULL COMMENT 'service_id',
  `leader_id` varchar(128) NOT NULL COMMENT 'leader_id',
  `last_seen_active` varchar(128) NOT NULL COMMENT 'last_seen_active',
  PRIMARY KEY (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='leader_election'
```


ref:
https://gist.github.com/ljjjustin/f2213ac9b9b8c31df746f8b56095ea32
