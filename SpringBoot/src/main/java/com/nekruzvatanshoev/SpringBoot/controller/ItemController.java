package com.nekruzvatanshoev.SpringBoot.controller;

import com.nekruzvatanshoev.SpringBoot.domain.Item;
import com.nekruzvatanshoev.SpringBoot.kafka.KafkaProducer;
import com.nekruzvatanshoev.SpringBoot.service.OrderItemService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ItemController {
    public final OrderItemService orderItemService;

    public ItemController(OrderItemService orderItemService) {
        this.orderItemService = orderItemService;
    }

    @PostMapping("/items")
    public ResponseEntity<Item> addToShoppingCart(@RequestBody Item item) {
        orderItemService.addItemToCart(item);
        System.out.println(item);
        return new ResponseEntity<>(HttpStatus.ACCEPTED);
    }
}
