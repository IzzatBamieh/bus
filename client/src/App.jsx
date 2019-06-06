import React, { Component, Fragment, createContext} from 'react';
import styled from 'styled-components';

const ServiceContext =  createContext();

const Home = styled.div``;

const Button = styled.button``;

const Input = styled.input ``;


class App extends Component {

  constructor(props) {
    super(props);

    this.state ={
      service: 'function',
      options: {
        serviceName: '',
        serviceRoute: '',
        serviceModel: '',
      },
      updateInputValue: (event) => this.updateInputValue(event),
      createService: this.createService

    }
  }

  updateInputValue(event) {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  createService() {
    const headers = {'Content-Type': 'application/json',
                      'Accept': 'application/json'};
    const body = JSON.stringify(this.state.options);
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

  render() {
    return (
      <Home>
        <ServiceContext.Provider value={this.state}>
           <CreateService  />
        </ServiceContext.Provider>
      </Home>
    );
  }
}

const CreateService = (ServiceContext) => (
  <ServiceContext.Consumer>
    {(this.context)}=> (
      <Fragment>
        <form>
          <label>Service Name
            <Input type='text' name='serviceName' value={this.context.options.serviceName} onChange={event => this.context.updateInputValue(event)}/>
          </label>
          <br/>
          <label>
            Service Route 
            <Input type='text' name='serviceRoute' value={this.context.options.serviceRoute} onChange={event => this.context.updateInputValue(event)}/>
          </label>
          <br/>
          <label>
            Service Model 
            <Input type='text' name='serviceModel' value={this.context.options.serviceModel} onChange={event => this.context.updateInputValue(event)}/>
          </label>
          <br/>
            <Button onClick={this.context.createFunction}>Submit</Button>
        </form>
      </Fragment>
    )}

  </ServiceContext.Consumer>
)

export default App;
