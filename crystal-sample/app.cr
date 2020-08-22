require "router"

class App
  include Router

  QUOTES = File.read_lines("../quotes.txt")

  def draw_routes
    get "/franklin-says" do |context, params|
      context.response.content_type = "text/plain"
      context.response.print "#{QUOTES.sample}\n"

      context
    end
  end

  def run
    server = HTTP::Server.new(route_handler)
    address = server.bind_tcp "0.0.0.0", 8002
    server.listen
  end
end

app = App.new
app.draw_routes
app.run
