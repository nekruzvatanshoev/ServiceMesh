package com.nekruzvatanshoev.SpringBoot.config;

import org.springframework.amqp.core.Queue;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Configuration
@EnableConfigurationProperties(AMQPProperties.class)
public class AMQPConfig {

    public Queue queue(@Value("${nekruzvatanshoev.amqp.queue}")String queueName) {
        return new Queue(queueName, false);
    }
}
