defmodule Bus.TopicController do
  use Bus.Web, :controller

  alias Bus.Topic

  def index(conn, _params) do
    topics = {"topics": [{"id": "1", "name": "userUpdate"}, {"id": "2", "name": "userUpdatePassword"}]}
    render(conn, "index.json", topics: topics)
  end
end
