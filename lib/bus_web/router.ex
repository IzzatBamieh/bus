defmodule BusWeb.Router do
  use BusWeb, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_flash
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/", BusWeb do
    pipe_through :browser # Use the default browser stack

    get "/", PageController, :index
  end

  scope "/topics", Bus.Web do
    pipe_through :api # Use the default browser stack

    get "/", TopicController, :index
  end

  # Other scopes may use custom stacks.
  # scope "/api", BusWeb do
  #   pipe_through :api
  # end
end
