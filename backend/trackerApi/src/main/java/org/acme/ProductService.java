package org.acme;

import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import org.eclipse.microprofile.rest.client.inject.RegisterRestClient;

import java.util.List;

@RegisterRestClient(configKey = "product-api")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public interface ProductService {

    @GET
    List<ProductDTO> getAllProducts();

    @GET
    @Path("/{name}")
    ProductDTO getProductByName(@PathParam("name") String name);

    @POST
    void createProduct(ProductDTO productDTO);

    @DELETE
    @Path("/{name}")
    void deleteProduct(@PathParam("name") String name);
}