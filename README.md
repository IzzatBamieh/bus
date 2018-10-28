# Bus

To start your Phoenix server:
  * Make sure you have a postgres container or local instance running. You can find DB config:
    `/Users/izzatbamieh/Desktop/code/bus/config/dev.exs`
  * Install dependencies with `mix deps.get`
  * Create and migrate your database with `mix ecto.create && mix ecto.migrate`
  * Install Node.js dependencies with `cd assets && npm install`
  * Start Phoenix endpoint with `mix phx.server`

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.
Or `curl localhost:4000/topics` to get all NATS topics
