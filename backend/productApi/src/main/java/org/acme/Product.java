package org.acme;

import io.quarkus.hibernate.orm.panache.PanacheEntity;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Table;

@Entity
@Table(name = "products")
public class Product extends PanacheEntity {
    @Column(length = 100)
    public String name;

    @Column(length = 100)
    public String brand;

    public int portionSizeInGrams;
    public int caloriesPer100Grams;
    public double proteinsInGrams;
    public double fatsInGrams;
    public double carbsInGrams;
}