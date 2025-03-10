# golang-kafka

包含kafka & log & teams notify & mysql gorm

# log用法

有Infof & Debugf & Warningf & Errorf可以使用

	import "golang-kafka/util/log"
    
    log.Errorf("message: %v ,message: %v. end", "messageA", "messageB")

# notify用法

    import 	notifier "golang-kafka/util/notify"

    notifier.GetNotify().Send("title", "message")

# kafka producer用法

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


# kafka consumer用法

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