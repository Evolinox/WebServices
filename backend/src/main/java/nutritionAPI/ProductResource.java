package nutritionAPI;

import jakarta.inject.Inject;
import jakarta.transaction.Transactional;
import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;

import java.util.List;
import java.util.Optional;

@Path("/products")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class ProductResource {
    @Inject
    ProductRepository repository;

    @GET
    public Response getAll() {
        List<Product> products = repository.findAllProducts();
        return Response.ok(products).build();
    }

    @GET
    @Path("/{name}")
    public Response getByName(@PathParam("name") String name) {
        return  repository.find("name", name)
                .singleResultOptional()
                .map(product -> Response.ok(product).build())
                .orElse(Response.status(Response.Status.NOT_FOUND).build());
    }

    @POST
    @Transactional
    public Response createProduct(Product product) {
        product.persist();
        return Response.status(Response.Status.CREATED).entity(product).build();
    }

    @DELETE
    @Path("/{name}")
    @Transactional
    public Response delete(@PathParam("name") String name) {
        Optional<Product> food = repository.findByName(name);
        if (food.isPresent()) {
            repository.delete(food.get());
            return Response.noContent().build();
        } else {
            return Response.status(Response.Status.NOT_FOUND).entity("Food not found").build();
        }
    }
}
