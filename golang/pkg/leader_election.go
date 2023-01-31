func startLeaderElection(ctx context.Context) {
	glog.Info("start leader election")
	var latestIsLeader bool
	leaderElectionTicker := time.NewTicker(20 * time.Second)
	defer leaderElectionTicker.Stop()

	endpointID := leaderElection.GetEndpointIdentity()

	// first start
	isLeader, err := leaderElectionModel.CreateOrUpdate(leaderElectionModel.LeaderElectionServiceID, endpointID)
	if err != nil {
		glog.Fatalf("leader election failed with:%v", err)
	}

	if isLeader {
		fmt.Printf("节点: %s 切换为 leader", endpointID)
		latestIsLeader = true
	}

	// leader和worker都需要执行的task
	// 定期刷新缓存
	err = cron.StartCacheCronTask()
	if err != nil {
		glog.Errorf("failed to StartCacheCronTask :%v", err)
	}

	for {
	LOOP:
		glog.Infof("in loop")
		select {
		case <-leaderElectionTicker.C:
			glog.Infof("instance:%v start leader election", endpointID)
			isLeader, err := leaderElectionModel.CreateOrUpdate(leaderElectionModel.LeaderElectionServiceID, endpointID)
			if err != nil {
				glog.Errorf("leader election failed with:%v", err)
			}

			reloadCronTask()

			// start taskManager
			if err := taskManager.Run(endpointID, index, count); err != nil {
				glog.Errorf(err.Error())
			}

			if isLeader {
				if !latestIsLeader {
					latestIsLeader = true
					fmt.Printf("节点: %s 切换为 leader", endpointID)
				}

				err = startDistributeTask()
				if err != nil {
					glog.Errorf("distribute task failed with:%v", err)
					goto LOOP
				}
			} else {
				if latestIsLeader {
					latestIsLeader = false
				}
			}
		case <-ctx.Done():
			taskManager.Stop()
			return
		}
	}
}

type LeaderElection struct {
	ServiceID      string `orm:"column(service_id);pk" json:"service_id" description:"ID"`
	LeaderID       string `orm:"column(leader_id)" json:"leader_id" description:"ID"`
	LastSeenActive string `orm:"column(last_seen_active)" json:"last_seen_active" description:"last_seen_active"`
	TaskList       string `orm:"column(task_list)" json:"task_list" description:"task_list"`
	TaskType       string `orm:"column(task_type)" json:"task_type" description:"task_type"`
}

func CreateOrUpdate(serviceID string, leaderID string) (bool, error) {
	orm := models.GetOrmer()
	res, err := orm.Raw("INSERT IGNORE INTO leader_election (service_id, leader_id, last_seen_active) VALUES (?, ?, now()) ON DUPLICATE KEY UPDATE       leader_id = IF(last_seen_active <DATE_SUB(VALUES(last_seen_active), INTERVAL 30 SECOND), VALUES(leader_id), leader_id),last_seen_active = IF(leader_id   =VALUES(leader_id), VALUES(last_seen_active),last_seen_active);", serviceID, leaderID).Exec()
	if err != nil {
		return false, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if affected >= 1 {
		return true, nil
	}
	return false, nil
}
