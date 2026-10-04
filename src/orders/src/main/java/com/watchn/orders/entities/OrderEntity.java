package com.watchn.orders.entities;

import lombok.Data;

import java.util.ArrayList;
import java.util.List;

@Data
public class OrderEntity {

    private String id;
    private String firstName;
    private String lastName;
    private String email;

    private List<OrderItemEntity> items = new ArrayList<>();
}