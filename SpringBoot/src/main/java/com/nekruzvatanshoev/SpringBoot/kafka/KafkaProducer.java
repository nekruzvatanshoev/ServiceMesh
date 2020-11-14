package com.nekruzvatanshoev.SpringBoot.kafka;

import com.nekruzvatanshoev.SpringBoot.domain.Item;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
public class KafkaProducer {
    private KafkaTemplate<String, Item> kafkaTemplate;

    @Autowired
    public KafkaProducer(KafkaTemplate<String, Item> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }


    public void sendItem(Item item) {
        kafkaTemplate.send("test",item);
    }
}
