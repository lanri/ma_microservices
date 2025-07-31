package com.ma.microservice.controller;

import com.ma.microservice.model.LoginRequest;
import com.ma.microservice.security.JwtService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/auth")
public class UserController {

    @Autowired
    private JwtService jwtService;

    @PostMapping("/login")
    public String login(@RequestBody LoginRequest request) {
        // Demo hardcoded user validation
        if ("admin".equals(request.getUsername()) && "password".equals(request.getPassword())) {
            return jwtService.generateToken(request.getUsername());
            //token can be saved on hazelcast or redis
        }
        throw new RuntimeException("Username/password salah");
    }
}
