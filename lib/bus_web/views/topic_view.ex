defmodule BusWeb.TopicView do
  use BusWeb, :view

  def render("index.json", %{topics: topics}) do
    %{data: render_many(topics, BusWeb.TopicView, "topic.json")}
  end

  def render("topic.json", %{topic: topic}) do
    %{topics: topic}
  end
end
