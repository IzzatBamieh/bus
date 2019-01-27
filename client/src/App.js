import React, { Component } from 'react';
import styled from 'styled-components';

const BusHome = styled.div`
`;

const Dashbourd = styled.div`
  display: flex;
  flex-direction: row;
`;

const Subscribers = styled.div`
  display: flex;
  flex-direction: column;
`;

const Publishers = styled.div`
  display: flex;
  flex-direction: column;
`;

const SubscriberCard = styled.button`
  text-align: left;
  width: 100%;
  border-radius: 239px;
  background-color: aqua;
  border: solid;
  margin: 30px;
`;

const SubscriberName = styled.h1`
  font-size: 1.5em;
  color: palevioletred;
`;

const PublisherCard = styled.button`
  text-align: left;
  width: 100%;
  border-radius: 239px;
  background-color: aqua;
  border: solid;
  margin: 30px;
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
      data: { Topics: [{ name: "test_1", publisher: "publisher_service_1", subscribers: [{ name: "subscriber_service_1" }, { name: "subscriber_service_2" }, { name: "subscriber_service_3" }] }, { name: "test_2", publisher: "publisher_service_2", subscribers: [{name: "service_1"}]}]},
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
      <BusHome>
        {this.state.data.Topics.map((item, i) => {
          return (<TopicsDashbourd
          topic={item} />)
        })}
      </BusHome>
    );
  }
}

class TopicsDashbourd extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
      return (
        <Dashbourd>
          <Publishers>
            <Publisher
              topic={this.props.topic} />
          </Publishers>
          <Subscribers>
            {this.props.topic.subscribers.map((subscriber, i) => {
              return (<Subscriber
                subscriber={subscriber} />)
            })}
          </Subscribers>
        </Dashbourd>
      );
    }
  }

function Subscriber(props){
  return (
    <SubscriberCard>
      <SubscriberName>
        {props.subscriber.name}
      </SubscriberName>
    </SubscriberCard>
  );
}

function Publisher(props) {
  return (
    <PublisherCard>
      <PublisherName>
        {props.topic.publisher}
      </PublisherName>
      <PublisherTopic>
        {props.topic.name}
      </PublisherTopic>
    </PublisherCard>
  );
}

export default App;
