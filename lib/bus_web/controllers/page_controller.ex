defmodule BusWeb.PageController do
  use BusWeb, :controller

  def index(conn, _params) do
    render conn, "index.html"
  end
end
