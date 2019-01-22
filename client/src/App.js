import React, { Component } from 'react';
import styled from 'styled-components';

const TopicHome = styled.div`
  display: flex;
  flex-direction: column;
`;

const TopicCard = styled.button`
  text-align: left;
  width: 20%;
  border-radius: 239px;
  background-color: aqua;
  border: solid;
  margin: 10px;
`;

const TopicName = styled.h1`
  font-size: 1.5em;
  color: palevioletred;
`;

const TopicCount = styled.h1`
  font-size: 1em; 
  color: palevioletred;
`;

class App extends Component {


  constructor(props) {
    super(props);

    this.state = {
      data: {Topics: [{name: "test_1"}, {name: "test_2"}]},
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
          return (<Topic
          topic={item}
          topics_count={this.topics_count}
          key={i}
          id={i} />)
        })}
      </TopicHome>
    );
  }
}

class Topic extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
      return (
        <TopicCard>
          <TopicName>
            {this.props.topic.name}
          </TopicName>
          <TopicCount>
            {this.props.topics_count}
          </TopicCount>
        </TopicCard>
      );
    }
  }

export default App;
