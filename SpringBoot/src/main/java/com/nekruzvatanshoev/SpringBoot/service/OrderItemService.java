package com.nekruzvatanshoev.SpringBoot.service;

import com.nekruzvatanshoev.SpringBoot.amqp.Producer;
import com.nekruzvatanshoev.SpringBoot.domain.Item;
import com.nekruzvatanshoev.SpringBoot.kafka.KafkaProducer;
import org.springframework.stereotype.Service;

@Service
public class OrderItemService {

    private final Producer producer;
    private final KafkaProducer kafkaProducer;

    public OrderItemService(Producer producer, KafkaProducer kafkaProducer) {

        this.producer = producer;
        this.kafkaProducer = kafkaProducer;
    }

    public String addItemToCart(Item item) {
        //producer.sendMessage("Added " + id + " to shopping cart!");
        producer.sendMessage(item);
        kafkaProducer.sendItem(item);
        return "Added " + item.getId() + " shopping to cart!";
    }
}
