package org.acme;

import jakarta.inject.Inject;
import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;
import org.eclipse.microprofile.rest.client.inject.RegisterRestClient;
import org.eclipse.microprofile.rest.client.inject.RestClient;

import java.util.List;

@Path("/tracker")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class TrackerResource {

    @Inject
    @RestClient
    ProductService productService;

    @GET
    @Path("/products")
    public Response getAllProducts() {
        List<ProductDTO> products = productService.getAllProducts();
        return Response.ok(products).build();
    }

    @GET
    @Path("/products/{name}")
    public Response getProduct(@PathParam("name") String name) {
        ProductDTO product = productService.getProductByName(name);
        return Response.ok(product).build();
    }

    @POST
    @Path("/products")
    public Response addProduct(ProductDTO product) {
        productService.createProduct(product);
        return Response.status(Response.Status.CREATED).build();
    }

    @DELETE
    @Path("/products/{name}")
    public Response deleteProduct(@PathParam("name") String name) {
        productService.deleteProduct(name);
        return Response.noContent().build();
    }
}