defmodule BusWeb.TopicController do
  use BusWeb, :controller

  alias BusWeb.Topic

  def index(conn, _params) do
    topics = [ %{id: 1, name: "userUpdate"}, %{id: 2, name: "userUpdatePassword"}]
    render(conn, "index.json", topics: topics)
  end
end
