import React, { Component, Fragment } from 'react';
import styled from 'styled-components';

const Home = styled.div``;

const Button = styled.button``;

const Input = styled.input ``;


class App extends Component {

  constructor(props) {
    super(props);

    this.state ={
      service: 'create'
    }
    this.setService = this.setService.bind(this);
  }

  setService(service) {
    this.setState({service : service});
  }

  render() {
    return (
      <Home>
        <CreateService  setService = {this.setService} service = {this.state.service}/>
      </Home>
    );
  }
}

function CreateService(props) {
  switch(props.service) {
    case 'create':
      return <StartService setService={props.setService}/>
    case 'choose':
      return <ChooseService setService={props.setService}/>
    case 'function':
      return <CreateFunction setService={props.setService}/>
    default:
      throw Error('action unknown')
  }
}

class StartService extends Component {

  constructor(props) {
    super(props);

    this.chooseService = this.chooseService.bind(this);
  }


  chooseService() {
    this.props.setService('choose');
  }

  render() {
    return (
      <Home>
         <Button onClick={this.chooseService}>Create a Service</Button>
      </Home>
    );
  }
}

class ChooseService extends Component  {

  constructor(props) {
    super(props);

    this.createFunction = this.createFunction.bind(this);
  }

  createFunction() {
    this.props.setService('function');
  }

  render() {
    return (
      <Fragment>
        <Button onClick={this.createFunction}>Function</Button>
      </Fragment>
    );
  }
}


class CreateFunction extends Component  {

  constructor(props) {
    super(props);

    this.state ={
      serviceName: null
    }

    this.setService = this.setService.bind(this);
    this.createFunction = this.createFunction.bind(this);
    this.createService = createService.bind(this);
  }

  setService() {
    this.props.setService('create');
  }

  updateInputValue(event) {
    this.setState({serviceName : event.target.value});
  }

  createFunction(){
    createService(this.state.serviceName)
    // make api request to server serviceName
    this.setService();
  }

  render() {
    return (
      <Home>
        <Input type='text' value={this.state.serviceName} onChange={event => this.updateInputValue(event)}/>
        <Button onClick={this.createFunction} disabled={!this.state.serviceName}>Submit</Button>
      </Home>
    );
  }
}

function createService(service) {
  const headers = {'Content-Type': 'application/json',
                    'Accept': 'application/json'};
  const body = JSON.stringify({'name': service})
  return fetch('http://127.0.0.1:9000/services',  {
      cache: "no-cache", 
      method: "POST",
      headers: headers,
      body: body
  })
  .then(response => response.json())
  .catch(err => console.log(err))
}

export default App;
