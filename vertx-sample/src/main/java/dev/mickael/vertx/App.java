package dev.mickael.vertx;

import io.vertx.core.Vertx;
import io.vertx.ext.web.Router;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Random;

public class App {

  private final String[] quotes;
  private final Random rand = new Random();

  public App(String[] quotes) {
    this.quotes = quotes;
  }

  public void startServer() {
    var vertx = Vertx.vertx();
    var server = vertx.createHttpServer();
    var router = Router.router(vertx);

    router
        .route("/franklin-says")
        .handler(
            routingContext -> {
              var response = routingContext.response();
              response.putHeader("content-type", "text/plain");

              response.end(quotes[rand.nextInt(quotes.length)]);
            });

    server.requestHandler(router).listen(8080);
  }

  public static void main(String[] args) {
    BufferedReader file = null;
    var quotes = new ArrayList<String>();

    try {
      file = new BufferedReader(new FileReader("../quotes.txt"));
      String quote;
      while ((quote = file.readLine()) != null) {
        quotes.add(quote);
      }
    } catch (IOException e) {
      e.printStackTrace();
    } finally {
      try {
        if (file != null) file.close();
      } catch (IOException e) {
        e.printStackTrace();
      }
    }

    new App(quotes.stream().toArray(String[]::new)).startServer();
  }
}
