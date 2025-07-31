package com.ma.microservice;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Import;

@SpringBootApplication(scanBasePackages = {
		"com.ma.jdih", // hanya scan package jdih
		"com.ma.common.jwt", // kalau kamu ingin pakai JwtFilter spesifik saja
		"com.ma.microservice.config",
		"com.ma.microservice.controller"
})
public class JdihServiceApplication {

	public static void main(String[] args) {
		SpringApplication.run(JdihServiceApplication.class, args);
	}

}
