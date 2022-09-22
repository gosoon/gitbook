CreateOrUpdate SQL:
```
INSERT IGNORE INTO leader_election (service_id, leader_id, last_seen_active)   VALUES (?, ?, now()) ON DUPLICATE KEY UPDATE leader_id = IF(last_seen_active                            <DATE_SUB(VALUES(last_seen_active), INTERVAL 30 SECOND), VALUES(leader_id), leader_id),last_seen_active = IF(leader_id =VALUES(leader_id), VALUES(last_seen_active),last_seen_active);
```

ref:
https://gist.github.com/ljjjustin/f2213ac9b9b8c31df746f8b56095ea32
