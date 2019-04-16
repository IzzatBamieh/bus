import React, { Component, Fragment } from 'react';
import styled from 'styled-components';
import Modal from 'react-modal';
import { BrowserRouter as Router, Route, Link  } from 'react-router-dom';

const Home = styled.div``;

const Button = styled.button``;


class App extends Component {

  constructor(props) {
    super(props);

    this.state ={
      services: 'create'
    }
    this.createService = this.createService.bind(this);
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

  createService(service) {
    this.setState({services : service});
  }

  render() {
    return (
      <Home>
        <CreateService  createService = {this.createService} service = {this.state.services}/>
      </Home>
    );
  }
}



function CreateService(props) {
  switch(props.service) {
    case 'create':
      const chooseAction = props.createService('choose');
      return <StartSerivce onClick={chooseAction}/>
    case 'choose':
      return < ChooseSerivce onClick={props.createService} />
    default:
      return 'action unknown'
  }
}

function StartSerivce(props) {
  //props.createService('')
  return (
    <Fragment>
       <Button onClick={props.createService}>Create a Service</Button>
    </Fragment>
  );
}

function ChooseSerivce(props) {
  return (
    <Fragment>
      <Button onClick={props.createService}> Function</Button>
      <Button onClick={props.createService}> DB </Button>
      <Button onClick={props.createService}> Service integration</Button>
    </Fragment>
  );
}

export default App;
