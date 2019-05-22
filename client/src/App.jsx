import React, { Component, Fragment} from 'react';
import styled from 'styled-components';

import { Dropdown, DropdownToggle, DropdownMenu, DropdownItem } from 'reactstrap';

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
      serviceName: '',
      serviceRoute: '',
      serviceModel: ''
    }

    this.setService = this.setService.bind(this);
    this.createFunction = this.createFunction.bind(this);
    this.createService = createService.bind(this);
    this.isEnabled = this.isEnabled();
  }

  setService() {
    this.props.setService('choose');
  }

  updateInputValue(event) {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  createFunction(){
    createService(this.state)

    this.setService();
  }

  isEnabled() {
    const isAllFieldsTrue= this.state.serviceName && this.state.serviceRoute && this.state.serviceModel;
    return isAllFieldsTrue;
   }

  render() {
    return (
      <Home>
        <form>
          <label>
            Service Name
            <Input type='text' name='serviceName' value={this.state.serviceName} onChange={event => this.updateInputValue(event)}/>
          </label>
          <br/>
          <label>
            Service Route 
            <Input type='text' name='serviceRoute' value={this.state.serviceRoute} onChange={event => this.updateInputValue(event)}/>
          </label>
          <br/>
          <label>
            Service Model 
            <Input type='text' name='serviceModel' value={this.state.serviceModel} onChange={event => this.updateInputValue(event)}/>
          </label>
          <br/>
          <Button onClick={this.createFunction}>Submit</Button>
        </form>
      </Home>
    );    
  }
}

function createService(options) {
  const headers = {'Content-Type': 'application/json',
                    'Accept': 'application/json'};
  const body = JSON.stringify(options);
  console.log(body);
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
