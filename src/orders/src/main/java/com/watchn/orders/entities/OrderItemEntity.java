package com.watchn.orders.entities;

import lombok.Data;

import java.io.Serializable;

@Data
public class OrderItemEntity {

    public static class Key implements Serializable {

        private String orderId;
        private String productId;

        public String getOrderId() {
            return orderId;
        }

        public void setOrderId(String orderId) {
            this.orderId = orderId;
        }

        public String getProductId() {
            return productId;
        }

        public void setProductId(String productId) {
            this.productId = productId;
        }
    }

    private Key id;

    private String productId;

    private int quantity;

    private int price;

    private OrderEntity order;
}