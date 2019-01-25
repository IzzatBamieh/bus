import React, { Component } from 'react';
import styled from 'styled-components';

const TopicHome = styled.div`
  display: flex;
  flex-direction: column;
`;

const Subscribers = styled.div`
  display: flex;
  flex-direction: column;
`;

const Publishers = styled.div`
`;

const SubscriberCard = styled.button`
  text-align: left;
  width: 20%;
  border-radius: 239px;
  background-color: aqua;
  border: solid;
  margin: 10px;
`;

const SubscriberName = styled.h1`
  font-size: 1.5em;
  color: palevioletred;
`;

const PublisherCard = styled.button`
  text-align: left;
  width: 20%;
  border-radius: 239px;
  background-color: aqua;
  border: solid;
  margin: 10px;
`;

const PublisherName = styled.h1`
  font-size: 1.5em;
  color: palevioletred;
`;

const PublisherTopic = styled.h1`
  font-size: 1.5em;
  color: palevioletred;
`;

class App extends Component {


  constructor(props) {
    super(props);

    this.state = {
      data: { Topics: [{ name: "test_1", publisher: "service_p_1", subscribers: [{ name: "service_1" }] }, { name: "test_2", publisher: "service_p_2", subscribers: [{name: "service_1"}]}]},
      topics_count: 0
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
      <TopicHome>
        {this.state.data.Topics.map((item, i) => {
          return (<Publisher
          topic={item} />)
        })}
      </TopicHome>
    );
  }
}

class Publisher extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
      return (
        <Publishers>
          <PublisherCard>
            <PublisherName>
              {this.props.topic.publisher}
            </PublisherName>
            <PublisherTopic>
              {this.props.topic.name}
            </PublisherTopic>
          </PublisherCard>
          <Subscribers>
            {this.props.topic.subscribers.map((subscriber, i) => {
              return (<Subscriber
                subscriber={subscriber} />)
            })}
          </Subscribers>
        </Publishers>
      );
    }
  }

class Subscriber extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <SubscriberCard>
        <SubscriberName>
          {this.props.subscriber.name}
        </SubscriberName>
      </SubscriberCard>
    );
  }
}

export default App;
