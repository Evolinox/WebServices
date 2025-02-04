package org.acme;

import io.quarkus.hibernate.orm.panache.PanacheRepository;
import jakarta.enterprise.context.ApplicationScoped;

import java.util.List;
import java.util.Optional;

@ApplicationScoped
public class ProductRepository implements PanacheRepository<Product> {
    public Optional<Product> findByName(String name) {
        return find("name", name).firstResultOptional();
    }

    public List<Product> findAllProducts() {
        return listAll();
    }
}