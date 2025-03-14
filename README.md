# golang-kafka

包含kafka & log & teams notify & mysql gorm

# log用法

有Infof & Debugf & Warningf & Errorf可以使用

	import "golang-kafka/util/log"
    
    log.Errorf("message: %v ,message: %v. end", "messageA", "messageB")

# notify用法

    import 	notifier "golang-kafka/util/notify"

    notifier.GetNotify().Send("title", "message")

# kafka producer用法(參考worker)

    wg := &sync.WaitGroup{} // 使用 WaitGroup 等待 goroutine 完成
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 7; i++ {
			str := strconv.Itoa(int(time.Now().UnixNano()))
			if err := kafka.ProduceMessage(KakfaTopic, str); err != nil {
				log.Errorf("Failed to send message from worker: %v\n", err)
			}
		}
	}()

	wg.Wait()


# kafka consumer用法(參考worker2)

    ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < KafkaNumConsumers; i++ {
		wg.Add(1)
		go startConsumer(ctx, &wg, i)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-shutdown:
		log.Errorf("%v", sig)
		notifier.GetNotify().Send("worker consumer shutdown", "worker consumer shutdown")

		cancel()

		wg.Wait()
	}


# redis用法(參考worker3)

    ctx, cancel := context.WithCancel(context.Background())
	setting.InitConfig(ctx)

	defer cancel()
	defer kafka.CloseProducer()
	defer redis.CloseRedisClient()

	redisTest := redis.GetRedisClient()
	if redisTest == nil {
		log.Errorf("Redis client is nil, exiting program")
	}

	err := redisTest.Set(ctx, "testKey", "testValue", 1*time.Hour).Err()
	if err != nil {
		fmt.Println("Failed to add testKey key-value pair")
		log.Errorf("%v", err)
	}

	val, _ := redisTest.Get(ctx, "testKey").Result()
	fmt.Println("redis:", val)


# mysql gorm用法(參考worker4)

    ctx, cancel := context.WithCancel(context.Background())
	setting.InitConfig(ctx)

	defer cancel()
	defer kafka.CloseProducer()
	defer redis.CloseRedisClient()
	defer database.CloseDB()

	db := database.GetDB()

	//不存在則建立table
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Errorf("table建立失敗: %v", err)
	}

	//insert
	newUser := models.User{Name: "Alice", Email: "alice@example.com", Age: 25}
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Errorf("insert失敗: %v", result.Error)
	}

	//select
	var users []models.User
	db.Find(&users)
	fmt.Println("select結果:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", user.ID, user.Name, user.Email, user.Age)
	}