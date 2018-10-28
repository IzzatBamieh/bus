defmodule BusWeb.Topic do
  use Ecto.Schema
  import Ecto.Query, warn: false

  schema "topics" do
    field :name, :string

    timestamps()
  end
end
