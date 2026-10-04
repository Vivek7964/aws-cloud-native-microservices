package com.watchn.orders.repositories;

import com.watchn.orders.entities.OrderEntity;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;

@Repository
public class OrderRepository {

    private final Map<String, OrderEntity> orders = new ConcurrentHashMap<>();

    public OrderEntity save(OrderEntity order) {
        if (order.getId() == null || order.getId().isEmpty()) {
            order.setId(UUID.randomUUID().toString());
        }

        orders.put(order.getId(), order);
        return order;
    }

    public List<OrderEntity> findAll() {
        return new ArrayList<>(orders.values());
    }
}