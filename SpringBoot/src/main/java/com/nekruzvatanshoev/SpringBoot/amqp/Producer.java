package com.nekruzvatanshoev.SpringBoot.amqp;

import com.nekruzvatanshoev.SpringBoot.domain.Item;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class Producer {
    private RabbitTemplate tempalte;
    private String exchange;
    private String routingKey;

    Producer(final RabbitTemplate template,
             @Value("${nekruzvatanshoev.amqp.exchange:}") final String exchange,
             @Value("${nekruzvatanshoev.amqp.queue}") final String routingKey) {
        this.tempalte = template;
        this.exchange = exchange;
        this.routingKey = routingKey;
    }

    public void sendMessage(Item item) {
        this.tempalte.convertAndSend(exchange, routingKey, item);
    }
}
