package com.ma.microservice.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.oauth2.jwt.JwtClaimsSet;
import org.springframework.security.oauth2.server.authorization.token.JwtEncodingContext;
import org.springframework.security.oauth2.server.authorization.token.OAuth2TokenCustomizer;

@Configuration
public class TokenCustomizerConfig {

    @Bean
    public OAuth2TokenCustomizer<JwtEncodingContext> tokenCustomizer() {
        return context -> {
            if ("access_token".equals(context.getTokenType().getValue())) {
                JwtClaimsSet.Builder claims = context.getClaims();

                // Konversi scope ke string dipisahkan spasi
                String scope = String.join(" ", context.getAuthorizedScopes());

                // Simpan sebagai string, bukan array
                claims.claim("scope", scope);
            }
        };
    }
}


