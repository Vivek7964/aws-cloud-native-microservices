package com.watchn.carts.configuration;

import com.watchn.carts.services.CartService;
import com.watchn.carts.services.InMemoryCartService;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class InMemoryConfiguration {

    @Bean
    public CartService cartService() {
        return new InMemoryCartService();
    }
}