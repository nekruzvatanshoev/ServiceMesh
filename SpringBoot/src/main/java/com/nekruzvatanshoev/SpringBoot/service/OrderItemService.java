package com.nekruzvatanshoev.SpringBoot.service;

import com.nekruzvatanshoev.SpringBoot.amqp.Producer;
import org.springframework.stereotype.Service;

@Service
public class OrderItemService {

    private final Producer producer;

    public OrderItemService(Producer producer) {
        this.producer = producer;
    }

    public String addItemToCart(String id) {
        //producer.sendMessage("Added " + id + " to shopping cart!");
        producer.sendMessage(id);
        return "Added " + id + " shopping to cart!";
    }
}
