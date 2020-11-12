package com.nekruzvatanshoev.SpringBoot.service;

import com.nekruzvatanshoev.SpringBoot.amqp.Producer;
import com.nekruzvatanshoev.SpringBoot.domain.Item;
import org.springframework.stereotype.Service;

@Service
public class OrderItemService {

    private final Producer producer;

    public OrderItemService(Producer producer) {
        this.producer = producer;
    }

    public String addItemToCart(Item item) {
        //producer.sendMessage("Added " + id + " to shopping cart!");
        producer.sendMessage(item);
        return "Added " + item.getId() + " shopping to cart!";
    }
}
