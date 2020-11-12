package com.nekruzvatanshoev.SpringBoot.amqp;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageListener;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;

@Component
public class Consumer {
    @RabbitListener(queues="${nekruzvatanshoev.amqp.queue}")
    public void onMessage(Message message) {
        System.out.println(message);
    }
}
