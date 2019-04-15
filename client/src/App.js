import React, { Component } from 'react';
import styled from 'styled-components';
import Modal from 'react-modal';
import { BrowserRouter as Router, Route, Link  } from 'react-router-dom';

const Home = styled.div``;

const Button = styled.button``;


class App extends Component {

  constructor(props) {
    super(props);

    this.state = {
      ifBus: false,
    };
  }

  //componentDidMount() {
//
  //  const headers = {
  //    "Content-Type": "application/json",
  //  };
//
  //  fetch('127.0.0.1:9000/topics', { method: "get", headers })
  //    .then(response => response.json())
  //    .then(data => {
  //      this.setState({ data: { Topics: [{ name: "test" }] } })
  //      this.setState({topics_count: {Topics:[{name: "test"}]}})
  //    })
  //}

  render() {
    return (
      <Home>
        <BusHome ifBus={this.state.ifBus}/>
      </Home>
    );
  }
}

function updateState(text){
  this.setState({text})
}

function BusHome(props) {
  const ifBus = props.ifBus;
  if (ifBus) {
    return <loadBus />;
  }
  return <CreateBus />;
}


function CreateBus(props) {
  return <Button onClick={this.}>add bus</Button>
}

function loadBus(props) {
  return <h1>Please sign up.</h1>;
}

export default App;
