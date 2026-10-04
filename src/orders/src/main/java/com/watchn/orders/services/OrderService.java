package com.watchn.orders.services;

import com.watchn.orders.entities.OrderEntity;
import com.watchn.orders.entities.OrderItemEntity;
import com.watchn.orders.messaging.OrdersEventHandler;
import com.watchn.orders.repositories.OrderRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

@Service
@Slf4j
public class OrderService {

    @Autowired
    private OrderRepository repository;

    @Autowired
    private OrdersEventHandler eventHandler;

    public OrderEntity create(OrderEntity order) {

        if (order.getId() == null || order.getId().isEmpty()) {
            order.setId(UUID.randomUUID().toString());
        }

        for (OrderItemEntity item : order.getItems()) {
            item.setOrder(order);

            OrderItemEntity.Key key = new OrderItemEntity.Key();
            key.setOrderId(order.getId());
            key.setProductId(item.getProductId());

            item.setId(key);
        }

        OrderEntity entity = repository.save(order);

        eventHandler.postCreatedEvent(entity);

        return entity;
    }

    public List<OrderEntity> list() {
        return repository.findAll();
    }
}