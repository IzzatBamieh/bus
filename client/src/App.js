import React, { Component } from 'react';
import styled from 'styled-components';

const TopicHome = styled.div`
  display: flex;
  flex-direction: row;
`;

const TopicCard = styled.div`
  text-align: left;
  border: solid;
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
      data: [],
    };
  }


  componentDidMount() {

    const headers = {
      "Content-Type": "application/json",
    };

    fetch('127.0.0.1:9000/topics', { method: "get", headers })
      .then(response => response.json())
      .then(data => {
        this.setState({ data: prepareFeed(data) })
      })
  }
  
  render() {
    return (
      <TopicHome>
        {this.state.data.map((item, i) => {
          return (<Topic
          value={item}
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
            this.props.value
            </TopicName>
          <TopicCount>
            50
            </TopicCount>
          <a
            className="Topic Details"
            target="_blk"
            rel="noopener noreferrer"
          >
            More details
            </a>
        </TopicCard>
      );
    }
  }

export default App;
