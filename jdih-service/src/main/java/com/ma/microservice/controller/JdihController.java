package com.ma.microservice.controller;

import jakarta.annotation.PostConstruct;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/jdih")
public class JdihController {

    @PostConstruct
    public void init() {
        System.out.println(">>> JdihController aktif");
    }
    @GetMapping("/dokumen")
    public String getDokumen() {
        return "Dokumen hukum tersedia";
    }
}
