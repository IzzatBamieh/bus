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
    case 'DB':
      return <CreateDB setService={props.setService}/>
    case 'External Service':
      return <CreateExternalService setService={props.setService}/>
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
    this.createDB = this.createDB.bind(this);
    this.createExternalService = this.createExternalService.bind(this);
  }

  createFunction() {
    this.props.setService('function');
  }

  createDB() {
    this.props.setService('DB');
  }

  createExternalService() {
    this.props.setService('External Service');
  }

  render() {
    return (
      <Fragment>
        <Button onClick={this.createFunction}>Function</Button>
        <Button onClick={this.createDB}>DB</Button>
        <Button onClick={this.createExternalService}>Service External Service</Button>
      </Fragment>
    );
  }
}


class CreateFunction extends Component  {

  constructor(props) {
    super(props);

    this.setService = this.setService.bind(this);
    this.createFunction = this.createFunction.bind(this);
  }

  setService() {
    this.props.setService('create');
  }

  createFunction(){
    // make api request to server
    this.setService();
  }

  render() {
    return (
      <Home>
        <Input type='text'/>
        <Button onClick={this.createFunction}>Submit</Button>
      </Home>
    );
  }
}


class CreateDB extends Component  {

  constructor(props) {
    super(props);

    this.setService = this.setService.bind(this);
    this.createDB = this.createDB.bind(this);
  }

  setService() {
    this.props.setService('create');
  }

  createDB(){
    // make api request to server
    this.setService();
  }

  render() {
    return (
      <Home>
        <Input type='text'/>
        <Button onClick={this.createDB}>Submit</Button>
      </Home>
    );
  }
}


class CreateExternalService extends Component  {

  constructor(props) {
    super(props);

    this.setService = this.setService.bind(this);
    this.createExternalService = this.createExternalService.bind(this);
  }

  setService() {
    this.props.setService('create');
  }

  createExternalService(){
    // make api request to server
    this.setService();
  }

  render() {
    return (
      <Home>
        <Input type='text'/>
        <Button onClick={this.createExternalService}>Submit</Button>
      </Home>
    );
  }
}

export default App;
